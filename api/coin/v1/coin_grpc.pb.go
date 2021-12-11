// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CoinClient is the client API for Coin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CoinClient interface {
	Recharge(ctx context.Context, in *RechargeCoinRequest, opts ...grpc.CallOption) (*RechargeCoinReply, error)
	ConsumeCoin(ctx context.Context, in *ConsumeCoinRequest, opts ...grpc.CallOption) (*ConsumeCoinReply, error)
	GetCoinBalance(ctx context.Context, in *GetCoinBalanceRequest, opts ...grpc.CallOption) (*GetCoinBalanceReply, error)
}

type coinClient struct {
	cc grpc.ClientConnInterface
}

func NewCoinClient(cc grpc.ClientConnInterface) CoinClient {
	return &coinClient{cc}
}

func (c *coinClient) Recharge(ctx context.Context, in *RechargeCoinRequest, opts ...grpc.CallOption) (*RechargeCoinReply, error) {
	out := new(RechargeCoinReply)
	err := c.cc.Invoke(ctx, "/api.coin.v1.Coin/Recharge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coinClient) ConsumeCoin(ctx context.Context, in *ConsumeCoinRequest, opts ...grpc.CallOption) (*ConsumeCoinReply, error) {
	out := new(ConsumeCoinReply)
	err := c.cc.Invoke(ctx, "/api.coin.v1.Coin/ConsumeCoin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coinClient) GetCoinBalance(ctx context.Context, in *GetCoinBalanceRequest, opts ...grpc.CallOption) (*GetCoinBalanceReply, error) {
	out := new(GetCoinBalanceReply)
	err := c.cc.Invoke(ctx, "/api.coin.v1.Coin/GetCoinBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoinServer is the server API for Coin service.
// All implementations must embed UnimplementedCoinServer
// for forward compatibility
type CoinServer interface {
	Recharge(context.Context, *RechargeCoinRequest) (*RechargeCoinReply, error)
	ConsumeCoin(context.Context, *ConsumeCoinRequest) (*ConsumeCoinReply, error)
	GetCoinBalance(context.Context, *GetCoinBalanceRequest) (*GetCoinBalanceReply, error)
	mustEmbedUnimplementedCoinServer()
}

// UnimplementedCoinServer must be embedded to have forward compatible implementations.
type UnimplementedCoinServer struct {
}

func (UnimplementedCoinServer) Recharge(context.Context, *RechargeCoinRequest) (*RechargeCoinReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Recharge not implemented")
}
func (UnimplementedCoinServer) ConsumeCoin(context.Context, *ConsumeCoinRequest) (*ConsumeCoinReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConsumeCoin not implemented")
}
func (UnimplementedCoinServer) GetCoinBalance(context.Context, *GetCoinBalanceRequest) (*GetCoinBalanceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCoinBalance not implemented")
}
func (UnimplementedCoinServer) mustEmbedUnimplementedCoinServer() {}

// UnsafeCoinServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CoinServer will
// result in compilation errors.
type UnsafeCoinServer interface {
	mustEmbedUnimplementedCoinServer()
}

func RegisterCoinServer(s grpc.ServiceRegistrar, srv CoinServer) {
	s.RegisterService(&Coin_ServiceDesc, srv)
}

func _Coin_Recharge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RechargeCoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoinServer).Recharge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.coin.v1.Coin/Recharge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoinServer).Recharge(ctx, req.(*RechargeCoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coin_ConsumeCoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConsumeCoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoinServer).ConsumeCoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.coin.v1.Coin/ConsumeCoin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoinServer).ConsumeCoin(ctx, req.(*ConsumeCoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coin_GetCoinBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCoinBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoinServer).GetCoinBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.coin.v1.Coin/GetCoinBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoinServer).GetCoinBalance(ctx, req.(*GetCoinBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Coin_ServiceDesc is the grpc.ServiceDesc for Coin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Coin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.coin.v1.Coin",
	HandlerType: (*CoinServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Recharge",
			Handler:    _Coin_Recharge_Handler,
		},
		{
			MethodName: "ConsumeCoin",
			Handler:    _Coin_ConsumeCoin_Handler,
		},
		{
			MethodName: "GetCoinBalance",
			Handler:    _Coin_GetCoinBalance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/coin/v1/coin.proto",
}
