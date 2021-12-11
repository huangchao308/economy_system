package biz

import (
	"context"
	"sync"
	"time"

	"github.com/go-kratos/kratos/v2/log"

	pb "economy_system/api/gift/v1"
	"economy_system/internal/rpc"
)

type Gift struct {
	ID    uint64
	Name  string
	Price float32
	Icon  string
}

type SendGiftStatus int8

const (
	SendGiftStatusPending   SendGiftStatus = 1
	SendGiftStatusCompleted SendGiftStatus = 2
	SendGiftStatusFailed    SendGiftStatus = 3
)

type GiftBill struct {
	ID       uint64
	GiftID   uint64
	Num      uint64
	From     uint64
	To       uint64
	Status   SendGiftStatus
	CreateAt time.Time
}

type GiftRepo interface {
	GetGift(ctx context.Context, id uint64) (*Gift, error)
	CreateGift(ctx context.Context, gift *Gift) error
	UpdateGift(ctx context.Context, gift *Gift) error
	ListGift(ctx context.Context, pageIndex int, pageSize int) ([]*Gift, error)
	ListGiveGiftBill(ctx context.Context, sender uint64) ([]*GiftBill, error)
	ListReceiveGiftBill(ctx context.Context, receiver uint64) ([]*GiftBill, error)
	AddGiftBill(ctx context.Context, sender, receiver, giftID, num uint64) (uint64, error)
	UpdateGiftBillStatus(ctx context.Context, billID uint64, status SendGiftStatus) error
}

type SendMsgRes struct {
	Partition int32
	Offset    int64
}

type KafkaProducer interface {
	SendMsg(content interface{}) error
}

type GiftUseCase struct {
	coinRepo CoinRepo
	giftRepo GiftRepo
	userRpc  rpc.UserRpc
	log      *log.Helper
	producer KafkaProducer
}

func NewGiftUseCase(coinRepo CoinRepo, giftRepo GiftRepo, userRpc rpc.UserRpc, producer KafkaProducer, logger log.Logger) *GiftUseCase {
	return &GiftUseCase{
		coinRepo: coinRepo,
		giftRepo: giftRepo,
		userRpc:  userRpc,
		log:      log.NewHelper(logger),
		producer: producer,
	}
}

// MGiveGift 用于在打榜时的送礼
func (s *GiftUseCase) MGiveGift(ctx context.Context, req []*pb.GiveGiftRequest) (*pb.GiveGiftReply, error) {
	wg := sync.WaitGroup{}
	for _, content := range req {
		wg.Add(1)
		go func(c *pb.GiveGiftRequest) {
			err := s.producer.SendMsg(c)
			if err != nil {
				s.log.WithContext(ctx).Errorf("SendMsg err: %+v", err)
				return
			}
		}(content)
	}
	wg.Wait()
	return &pb.GiveGiftReply{}, nil
}

func (s *GiftUseCase) GiveGift(ctx context.Context, req *pb.GiveGiftRequest) (*pb.GiveGiftReply, error) {
	gift, err := s.giftRepo.GetGift(ctx, req.GetGiftId())
	if err != nil {
		return nil, err
	}
	billID, err := s.giftRepo.AddGiftBill(ctx, req.GetSenderId(), req.GetReceiverId(), req.GetGiftId(), req.GetNum())
	if err != nil {
		return nil, err
	}
	_, _, err = s.coinRepo.ExChange(ctx, req.GetSenderId(), req.GetReceiverId(), gift.Price*float32(req.GetNum()))
	if err != nil {
		_ = s.giftRepo.UpdateGiftBillStatus(ctx, billID, SendGiftStatusFailed)
		return nil, err
	}
	err = s.giftRepo.UpdateGiftBillStatus(ctx, billID, SendGiftStatusCompleted)
	if err != nil {
		return nil, err
	}

	return &pb.GiveGiftReply{}, nil
}

func (s *GiftUseCase) GetSendingHistory(ctx context.Context, req *pb.GetSendingHistoryRequest) (*pb.GetSendingHistoryReply, error) {
	bills, err := s.giftRepo.ListGiveGiftBill(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}
	historyList := make([]*pb.GiftHistory, 0, len(bills))
	for _, bill := range bills {
		history := s.billToHistory(ctx, bill)
		historyList = append(historyList, history)
	}

	return &pb.GetSendingHistoryReply{History: historyList}, nil
}
func (s *GiftUseCase) GetReceivingHistory(ctx context.Context, req *pb.GetReceivingHistoryRequest) (*pb.GetReceivingHistoryReply, error) {
	bills, err := s.giftRepo.ListReceiveGiftBill(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}
	historyList := make([]*pb.GiftHistory, 0, len(bills))
	for _, bill := range bills {
		history := s.billToHistory(ctx, bill)
		historyList = append(historyList, history)
	}

	return &pb.GetReceivingHistoryReply{History: historyList}, nil
}

func (s *GiftUseCase) billToHistory(ctx context.Context, bill *GiftBill) *pb.GiftHistory {
	history := &pb.GiftHistory{
		GiftId:           bill.GiftID,
		GiftName:         "",
		GiftValue:        0,
		SenderId:         bill.From,
		ReceiverId:       bill.To,
		SenderNickname:   "",
		ReceiverNickname: "",
		SendTime:         uint64(bill.CreateAt.Unix()),
		Num:              bill.Num,
	}
	gift, err := s.giftRepo.GetGift(ctx, bill.GiftID)
	if err != nil {
		s.log.WithContext(ctx).Errorf("GetGift err: %+v", err)
	}
	if gift != nil {
		history.GiftName = gift.Name
		history.GiftValue = gift.Price
	}

	userInfos, err := s.userRpc.GetUser(ctx, []uint64{bill.From, bill.To})
	if err != nil {
		s.log.WithContext(ctx).Errorf("GetUser err: %+v", err)
	}
	if u, ok := userInfos[bill.From]; ok {
		history.SenderNickname = u.GetNickname()
	}
	if u, ok := userInfos[bill.To]; ok {
		history.ReceiverNickname = u.GetNickname()
	}

	return history
}
