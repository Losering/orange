// Code generated by goctl. DO NOT EDIT!
// Source: order.proto

package server

import (
	"context"

	"orange/apps/order/rpc/internal/logic"
	"orange/apps/order/rpc/internal/svc"
	"orange/apps/order/rpc/order"
)

type OrderServer struct {
	svcCtx *svc.ServiceContext
	order.UnimplementedOrderServer
}

func NewOrderServer(svcCtx *svc.ServiceContext) *OrderServer {
	return &OrderServer{
		svcCtx: svcCtx,
	}
}

func (s *OrderServer) Orders(ctx context.Context, in *order.OrdersRequest) (*order.OrdersResponse, error) {
	l := logic.NewOrdersLogic(ctx, s.svcCtx)
	return l.Orders(in)
}

func (s *OrderServer) CreateOrder(ctx context.Context, in *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	l := logic.NewCreateOrderLogic(ctx, s.svcCtx)
	return l.CreateOrder(in)
}

func (s *OrderServer) CreateOrderCheck(ctx context.Context, in *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	l := logic.NewCreateOrderCheckLogic(ctx, s.svcCtx)
	return l.CreateOrderCheck(in)
}

func (s *OrderServer) RollbackOrder(ctx context.Context, in *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	l := logic.NewRollbackOrderLogic(ctx, s.svcCtx)
	return l.RollbackOrder(in)
}

func (s *OrderServer) CreateOrderDTM(ctx context.Context, in *order.AddOrderReq) (*order.AddOrderResp, error) {
	l := logic.NewCreateOrderDTMLogic(ctx, s.svcCtx)
	return l.CreateOrderDTM(in)
}

func (s *OrderServer) CreateOrderDTMRevert(ctx context.Context, in *order.AddOrderReq) (*order.AddOrderResp, error) {
	l := logic.NewCreateOrderDTMRevertLogic(ctx, s.svcCtx)
	return l.CreateOrderDTMRevert(in)
}

func (s *OrderServer) GetOrderById(ctx context.Context, in *order.GetOrderByIdReq) (*order.GetOrderByIdResp, error) {
	l := logic.NewGetOrderByIdLogic(ctx, s.svcCtx)
	return l.GetOrderById(in)
}
