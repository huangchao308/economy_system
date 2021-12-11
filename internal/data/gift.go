package data

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"

	"economy_system/internal/biz"
)

type GiftRepoImpl struct {
	data *Data
	log  *log.Helper
}

func (g *GiftRepoImpl) AddGiftBill(ctx context.Context, sender, receiver, giftID uint64, num uint64) (uint64, error) {
	bill := &biz.GiftBill{
		GiftID:   giftID,
		Num:      num,
		From:     sender,
		To:       receiver,
		Status:   biz.SendGiftStatusPending,
		CreateAt: time.Now(),
	}
	tx := g.data.db.WithContext(ctx).Create(bill)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return bill.ID, nil
}

func (g *GiftRepoImpl) UpdateGiftBillStatus(ctx context.Context, billID uint64, status biz.SendGiftStatus) error {
	bill := &biz.GiftBill{ID: billID}
	tx := g.data.db.WithContext(ctx).Model(bill).Update("status", status)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (g *GiftRepoImpl) GetGift(ctx context.Context, id uint64) (*biz.Gift, error) {
	gift := &biz.Gift{ID: id}
	tx := g.data.db.First(gift)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return gift, nil
}

func (g *GiftRepoImpl) CreateGift(ctx context.Context, gift *biz.Gift) error {
	tx := g.data.db.WithContext(ctx).Create(gift)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (g *GiftRepoImpl) UpdateGift(ctx context.Context, gift *biz.Gift) error {
	tx := g.data.db.WithContext(ctx).Updates(gift)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (g *GiftRepoImpl) ListGift(ctx context.Context, pageIndex int, pageSize int) ([]*biz.Gift, error) {
	giftList := make([]*biz.Gift, 0, pageSize)
	offset := (pageIndex - 1) * pageSize
	tx := g.data.db.WithContext(ctx).Offset(offset).Limit(pageSize).Find(giftList)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return giftList, nil
}

func (g *GiftRepoImpl) ListGiveGiftBill(ctx context.Context, sender uint64) ([]*biz.GiftBill, error) {
	billList := make([]*biz.GiftBill, 0)
	tx := g.data.db.WithContext(ctx).Where(&biz.GiftBill{From: sender, Status: biz.SendGiftStatusCompleted}).Find(&billList)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return billList, nil
}

func (g *GiftRepoImpl) ListReceiveGiftBill(ctx context.Context, receiver uint64) ([]*biz.GiftBill, error) {
	billList := make([]*biz.GiftBill, 0)
	tx := g.data.db.WithContext(ctx).Where(&biz.GiftBill{To: receiver, Status: biz.SendGiftStatusCompleted}).Find(&billList)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return billList, nil
}

func NewGiftRepo(data *Data, logger log.Logger) biz.GiftRepo {
	return &GiftRepoImpl{
		data: data,
		log:  log.NewHelper(logger),
	}
}
