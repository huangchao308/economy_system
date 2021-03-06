// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.0.5

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type GiftHTTPServer interface {
	GetReceivingHistory(context.Context, *GetReceivingHistoryRequest) (*GetReceivingHistoryReply, error)
	GetSendingHistory(context.Context, *GetSendingHistoryRequest) (*GetSendingHistoryReply, error)
	GiveGift(context.Context, *GiveGiftRequest) (*GiveGiftReply, error)
}

func RegisterGiftHTTPServer(s *http.Server, srv GiftHTTPServer) {
	r := s.Route("/")
	r.POST("/gift/give_gift", _Gift_GiveGift0_HTTP_Handler(srv))
	r.GET("/gift/sending_history/{user_id}", _Gift_GetSendingHistory0_HTTP_Handler(srv))
	r.GET("/gift/receiving_history/{user_id}", _Gift_GetReceivingHistory0_HTTP_Handler(srv))
}

func _Gift_GiveGift0_HTTP_Handler(srv GiftHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GiveGiftRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.gift.v1.Gift/GiveGift")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GiveGift(ctx, req.(*GiveGiftRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GiveGiftReply)
		return ctx.Result(200, reply)
	}
}

func _Gift_GetSendingHistory0_HTTP_Handler(srv GiftHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetSendingHistoryRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.gift.v1.Gift/GetSendingHistory")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetSendingHistory(ctx, req.(*GetSendingHistoryRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetSendingHistoryReply)
		return ctx.Result(200, reply)
	}
}

func _Gift_GetReceivingHistory0_HTTP_Handler(srv GiftHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetReceivingHistoryRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.gift.v1.Gift/GetReceivingHistory")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetReceivingHistory(ctx, req.(*GetReceivingHistoryRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetReceivingHistoryReply)
		return ctx.Result(200, reply)
	}
}

type GiftHTTPClient interface {
	GetReceivingHistory(ctx context.Context, req *GetReceivingHistoryRequest, opts ...http.CallOption) (rsp *GetReceivingHistoryReply, err error)
	GetSendingHistory(ctx context.Context, req *GetSendingHistoryRequest, opts ...http.CallOption) (rsp *GetSendingHistoryReply, err error)
	GiveGift(ctx context.Context, req *GiveGiftRequest, opts ...http.CallOption) (rsp *GiveGiftReply, err error)
}

type GiftHTTPClientImpl struct {
	cc *http.Client
}

func NewGiftHTTPClient(client *http.Client) GiftHTTPClient {
	return &GiftHTTPClientImpl{client}
}

func (c *GiftHTTPClientImpl) GetReceivingHistory(ctx context.Context, in *GetReceivingHistoryRequest, opts ...http.CallOption) (*GetReceivingHistoryReply, error) {
	var out GetReceivingHistoryReply
	pattern := "/gift/receiving_history/{user_id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.gift.v1.Gift/GetReceivingHistory"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *GiftHTTPClientImpl) GetSendingHistory(ctx context.Context, in *GetSendingHistoryRequest, opts ...http.CallOption) (*GetSendingHistoryReply, error) {
	var out GetSendingHistoryReply
	pattern := "/gift/sending_history/{user_id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.gift.v1.Gift/GetSendingHistory"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *GiftHTTPClientImpl) GiveGift(ctx context.Context, in *GiveGiftRequest, opts ...http.CallOption) (*GiveGiftReply, error) {
	var out GiveGiftReply
	pattern := "/gift/give_gift"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.gift.v1.Gift/GiveGift"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
