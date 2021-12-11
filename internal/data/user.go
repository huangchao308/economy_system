package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"economy_system/internal/biz"
)

type UserRepoImpl struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &UserRepoImpl{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *UserRepoImpl) CreateUser(ctx context.Context, u *biz.User) error {
	tx := r.data.db.WithContext(ctx).Create(u)
	return tx.Error
}

func (r *UserRepoImpl) UpdateUser(ctx context.Context, u *biz.User) error {
	tx := r.data.db.WithContext(ctx).Updates(*u)
	return tx.Error
}

func (r *UserRepoImpl) GetUser(ctx context.Context, id uint64) (*biz.User, error) {
	user := &biz.User{ID: id}
	tx := r.data.db.WithContext(ctx).First(user)
	return user, tx.Error
}
