package service

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"time"

	"github.com/go-kratos/kratos/v2/log"

	pb "economy_system/api/user/v1"
	"economy_system/internal/biz"
)

type UserService struct {
	pb.UnimplementedUserServer

	uc  *biz.UserUseCase
	log *log.Helper
}

func NewUserService(uc *biz.UserUseCase, logger log.Logger) *UserService {
	return &UserService{uc: uc, log: log.NewHelper(logger)}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	reply := &pb.CreateUserReply{}
	hash := md5.New()
	hash.Write([]byte(req.GetPassword()))
	now := time.Now()
	user := &biz.User{
		UserName: req.GetUsername(),
		PassMd5:  hex.EncodeToString(hash.Sum(nil)),
		Nickname: req.GetNickname(),
		HeadImg:  req.GetHeadImg(),
		CreateAt: now,
		UpdateAt: now,
	}
	err := s.uc.Create(ctx, user)
	if err != nil {
		reply.ErrCode = -1
	}
	return reply, err
}
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	return &pb.UpdateUserReply{}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	reply := &pb.GetUserReply{}
	user, err := s.uc.Get(ctx, req.GetId())
	if err != nil {
		reply.ErrCode = -1
		return reply, err
	}
	reply.User = &pb.UserData{
		Id:       user.ID,
		Username: user.UserName,
		Nickname: user.Nickname,
		HeadImg:  user.HeadImg,
	}
	return reply, nil
}
