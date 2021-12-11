package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	pb "economy_system/api/coin/v1"
	"economy_system/internal/biz"
)

type CoinService struct {
	pb.UnimplementedCoinServer

	uc  *biz.CoinUseCase
	log *log.Helper
}

func NewCoinService(uc *biz.CoinUseCase, logger log.Logger) *CoinService {
	return &CoinService{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

func (s *CoinService) Recharge(ctx context.Context, req *pb.RechargeCoinRequest) (*pb.RechargeCoinReply, error) {
	s.log.WithContext(ctx).Infof("req: %v", req)
	balance, err := s.uc.Recharge(ctx, req.GetUserId(), req.GetCoinNum())
	if err != nil {
		return nil, err
	}
	reply := &pb.RechargeCoinReply{
		Balance: balance,
	}
	return reply, nil
}

//func (s *CoinService) ConsumeCoin(ctx context.Context, req *pb.ConsumeCoinRequest) (*pb.ConsumeCoinReply, error) {
//	reply := &pb.ConsumeCoinReply{}
//	balance, err := s.uc.Consume(ctx, req.GetUserId(), req.GetCoinNum())
//	if err != nil {
//		return nil, err
//	}
//
//	reply.Balance = balance
//
//	return reply, nil
//}
func (s *CoinService) GetCoinBalance(ctx context.Context, req *pb.GetCoinBalanceRequest) (*pb.GetCoinBalanceReply, error) {
	coin, err := s.uc.GetCoin(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}
	reply := &pb.GetCoinBalanceReply{
		ErrCode: 0,
		Balance: coin.Balance,
	}
	return reply, nil
}
