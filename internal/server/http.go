package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"

	coinAPI "economy_system/api/coin/v1"
	giftAPI "economy_system/api/gift/v1"
	userAPI "economy_system/api/user/v1"
	"economy_system/internal/conf"
	"economy_system/internal/service"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, user *service.UserService, coin *service.CoinService, gift *service.GiftService,
	logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Logger(logger),
		http.Middleware(
			recovery.Recovery(),
			validate.Validator(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	userAPI.RegisterUserHTTPServer(srv, user)
	coinAPI.RegisterCoinHTTPServer(srv, coin)
	giftAPI.RegisterGiftHTTPServer(srv, gift)

	return srv
}
