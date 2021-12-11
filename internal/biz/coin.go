package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type BillType int8

const (
	BillTypeRecharge BillType = 1
	BillTypeSendGift BillType = 2
)

type Bill struct {
	ID       uint64
	BillType BillType
	Amount   float32
	From     uint64
	To       uint64
	CreateAt time.Time
}

type Coin struct {
	UserID        uint64 `gorm:"column:user_id;primaryKey"`
	Balance       float32
	FreezeBalance float32
	CreateAt      time.Time `gorm:"column:create_at"`
	UpdateAt      time.Time `gorm:"column:update_at"`
}

type CoinRepo interface {
	Recharge(ctx context.Context, user uint64, amount float32) (float32, error)
	ExChange(ctx context.Context, from, to uint64, amount float32) (float32, float32, error)
	GetCoin(ctx context.Context, user uint64) (*Coin, error)
}

type CoinUseCase struct {
	repo CoinRepo
	log  *log.Helper
}

func NewCoinUseCase(repo CoinRepo, logger log.Logger) *CoinUseCase {
	return &CoinUseCase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (uc *CoinUseCase) Recharge(ctx context.Context, user uint64, amount float32) (float32, error) {
	uc.log.WithContext(ctx).Infof("input#user: %v, input#amount: %v", user, amount)
	return uc.repo.Recharge(ctx, user, amount)
}

//func (uc *CoinUseCase) Consume(ctx context.Context, user uint64, amount float32) (float32, error) {
//	coin, err := uc.giftRepo.GetCoin(ctx, user)
//	if err != nil {
//		return 0, err
//	}
//	if coin == nil {
//		return 0, err
//	}
//	if coin.Balance < amount {
//		return coin.Balance, errors.New(400001, "余额不足", "余额不足")
//	}
//	return uc.giftRepo.ExChange(ctx, user, amount)
//}

func (uc *CoinUseCase) GetCoin(ctx context.Context, user uint64) (*Coin, error) {
	return uc.repo.GetCoin(ctx, user)
}
