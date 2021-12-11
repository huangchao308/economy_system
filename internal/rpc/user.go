package rpc

import (
	"context"
	"sync"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	userPb "economy_system/api/user/v1"
	"economy_system/internal/conf"
)

type UserRpc interface {
	GetUser(ctx context.Context, userIDs []uint64) (map[uint64]*userPb.UserData, error)
}

type UserRpcImpl struct {
	client userPb.UserClient
	log    *log.Helper
}

func NewUserRpc(c *conf.Client, logger log.Logger) UserRpc {
	endpoint := c.Grpc.Addr
	conn, err := grpc.DialInsecure(context.Background(), grpc.WithEndpoint(endpoint))
	if err != nil {
		panic(err)
	}
	client := userPb.NewUserClient(conn)
	return &UserRpcImpl{client: client, log: log.NewHelper(logger)}
}

func (ur *UserRpcImpl) GetUser(ctx context.Context, userIDs []uint64) (map[uint64]*userPb.UserData, error) {
	ur.log.Infof("GetUser input: %v", userIDs)
	users := make([]*userPb.UserData, 0, len(userIDs))
	wg := sync.WaitGroup{}

	for _, id := range userIDs {
		wg.Add(1)
		go func(iCtx context.Context, uID uint64) {
			defer func() {
				wg.Done()
			}()
			userReply, e := ur.client.GetUser(iCtx, &userPb.GetUserRequest{Id: uID})
			if e != nil {
				ur.log.Errorf("GetUser err: %+v", e)
			}
			users = append(users, userReply.GetUser())
		}(ctx, id)
	}
	wg.Wait()

	ur.log.Infof("users: %v", users)
	result := make(map[uint64]*userPb.UserData, len(userIDs))
	for _, user := range users {
		result[user.Id] = user
	}

	return result, nil
}
