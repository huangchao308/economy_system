package rpc

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewUserRpc)

//type Rpc struct {
//	c      *conf.Client
//	logger *log.Logger
//}

//func NewRpc(c *conf.Client, logger log.Logger) {
//
//}
