// Code generated by goctl. DO NOT EDIT.
// Source: pay.proto

package server

import (
	"context"

	"mall/service/pay/rpc/internal/logic"
	"mall/service/pay/rpc/internal/svc"
	"mall/service/pay/rpc/pay"
)

type PayServer struct {
	svcCtx *svc.ServiceContext
	pay.UnimplementedPayServer
}

func NewPayServer(svcCtx *svc.ServiceContext) *PayServer {
	return &PayServer{
		svcCtx: svcCtx,
	}
}

func (s *PayServer) Create(ctx context.Context, in *pay.CreateRequest) (*pay.CreateResponse, error) {
	l := logic.NewCreateLogic(ctx, s.svcCtx)
	return l.Create(in)
}

func (s *PayServer) Detail(ctx context.Context, in *pay.DetailRequest) (*pay.DetailResponse, error) {
	l := logic.NewDetailLogic(ctx, s.svcCtx)
	return l.Detail(in)
}

func (s *PayServer) Callback(ctx context.Context, in *pay.CallbackRequest) (*pay.CallbackResponse, error) {
	l := logic.NewCallbackLogic(ctx, s.svcCtx)
	return l.Callback(in)
}
