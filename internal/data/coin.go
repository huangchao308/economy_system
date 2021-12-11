package data

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	pkgErrors "github.com/pkg/errors"
	"gorm.io/gorm"

	"economy_system/internal/biz"
)

type CoinRepoImpl struct {
	data *Data
	log  *log.Helper
}

func NewCoinRepo(data *Data, logger log.Logger) biz.CoinRepo {
	return &CoinRepoImpl{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *CoinRepoImpl) Recharge(ctx context.Context, user uint64, amount float32) (float32, error) {
	coin := &biz.Coin{UserID: user}
	now := time.Now()
	err := r.data.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		bill := &biz.Bill{
			BillType: biz.BillTypeRecharge,
			Amount:   amount,
			From:     user,
			To:       user,
			CreateAt: now,
		}
		res := tx.Create(bill)
		if res.Error != nil {
			return pkgErrors.Wrapf(res.Error, "Create err, bill: %v", bill)
		}
		res = tx.Where(&biz.Coin{UserID: user}).Attrs(&biz.Coin{CreateAt: now, UpdateAt: now}).FirstOrCreate(coin)
		if res.Error != nil {
			return pkgErrors.Wrapf(res.Error, "FirstOrCreate err, coin: %v", coin)
		}
		coin.Balance += amount
		coin.UpdateAt = now

		res = tx.Updates(coin)
		if res.Error != nil {
			return pkgErrors.Wrapf(res.Error, "Update coin err. coin: %v", coin)
		}
		return nil
	})
	if err != nil {
		return 0, pkgErrors.Wrap(err, "Recharge err")
	}
	return coin.Balance, nil
}

func (r *CoinRepoImpl) ExChange(ctx context.Context, from, to uint64, amount float32) (float32, float32, error) {
	fromCoin := &biz.Coin{UserID: from}
	toCoin := &biz.Coin{UserID: to}
	now := time.Now()
	err := r.data.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		bill := &biz.Bill{
			BillType: biz.BillTypeSendGift,
			Amount:   amount,
			From:     from,
			To:       to,
			CreateAt: now,
		}
		res := tx.Create(bill)
		if res.Error != nil {
			return res.Error
		}
		res = tx.First(fromCoin)
		if res.Error != nil {
			return res.Error
		}
		if fromCoin.Balance < amount {
			err := errors.Newf(400, "Insufficient_balance", "This user %v has not enough coins", from)
			return pkgErrors.Wrap(err, "")
		}
		fromCoin.Balance -= amount
		fromCoin.UpdateAt = now
		res = tx.Updates(fromCoin)
		if res.Error != nil {
			return pkgErrors.Wrapf(res.Error, "Updates fromCoin err. %v", fromCoin)
		}
		res = tx.Where(&biz.Coin{UserID: to}).Attrs(&biz.Coin{CreateAt: now, UpdateAt: now}).FirstOrCreate(toCoin)
		if res.Error != nil {
			return pkgErrors.Wrapf(res.Error, "FirstOrCreate err. %v", toCoin)
		}
		toCoin.Balance += amount
		toCoin.UpdateAt = now
		res = tx.Updates(toCoin)
		if res.Error != nil {
			return pkgErrors.Wrapf(res.Error, "Updates toCoin err. %v", toCoin)
		}
		return nil
	})
	if err != nil {
		return 0, 0, err
	}
	return fromCoin.Balance, toCoin.Balance, nil
}

func (r *CoinRepoImpl) GetCoin(ctx context.Context, user uint64) (*biz.Coin, error) {
	coin := &biz.Coin{UserID: user}
	tx := r.data.db.WithContext(ctx).First(coin)
	if tx.Error != nil {
		return nil, pkgErrors.Wrapf(tx.Error, "First err. %v", coin)
	}

	return coin, nil
}
