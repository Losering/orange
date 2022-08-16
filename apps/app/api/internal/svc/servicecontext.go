package svc

import (
	"orange/apps/app/api/internal/config"
	"orange/apps/order/rpc/orderclient"
	"orange/apps/product/rpc/productclient"
	"orange/apps/user/rpc/userclient"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	OrderRPC   orderclient.Order
	ProductRPC productclient.Product
	UserRPC    userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		ProductRPC: productclient.NewProduct(zrpc.MustNewClient(c.ProductRPC)),
		OrderRPC:   orderclient.NewOrder(zrpc.MustNewClient(c.OrderRPC)),
		UserRPC:    userclient.NewUser(zrpc.MustNewClient(c.UserRPC)),
	}
}
