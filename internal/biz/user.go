package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	ID       uint64
	UserName string    `gorm:"column:username"`
	PassMd5  string    `gorm:"column:pass_md5"`
	Nickname string    `gorm:"column:nickname"`
	HeadImg  string    `gorm:"column:head_img"`
	CreateAt time.Time `gorm:"column:create_at"`
	UpdateAt time.Time `gorm:"column:update_at"`
}

type UserRepo interface {
	CreateUser(ctx context.Context, u *User) error
	UpdateUser(ctx context.Context, u *User) error
	GetUser(ctx context.Context, id uint64) (*User, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (uc *UserUseCase) Create(ctx context.Context, u *User) error {
	return uc.repo.CreateUser(ctx, u)
}

func (uc *UserUseCase) Update(ctx context.Context, u *User) error {
	return uc.repo.UpdateUser(ctx, u)
}

func (uc *UserUseCase) Get(ctx context.Context, id uint64) (*User, error) {
	return uc.repo.GetUser(ctx, id)
}
