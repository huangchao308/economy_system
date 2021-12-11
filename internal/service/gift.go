package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	pb "economy_system/api/gift/v1"
	"economy_system/internal/biz"
)

type GiftService struct {
	pb.UnimplementedGiftServer

	uc  *biz.GiftUseCase
	log *log.Helper
}

func NewGiftService(uc *biz.GiftUseCase, logger log.Logger) *GiftService {
	return &GiftService{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

func (s *GiftService) GiveGift(ctx context.Context, req *pb.GiveGiftRequest) (*pb.GiveGiftReply, error) {
	return s.uc.GiveGift(ctx, req)
}
func (s *GiftService) GetSendingHistory(ctx context.Context, req *pb.GetSendingHistoryRequest) (*pb.GetSendingHistoryReply, error) {
	return s.uc.GetSendingHistory(ctx, req)
}
func (s *GiftService) GetReceivingHistory(ctx context.Context, req *pb.GetReceivingHistoryRequest) (*pb.GetReceivingHistoryReply, error) {
	return s.uc.GetReceivingHistory(ctx, req)
}
