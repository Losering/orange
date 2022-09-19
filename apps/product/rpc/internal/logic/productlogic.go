package logic

import (
	"context"
	"fmt"

	"orange/apps/product/rpc/internal/svc"
	"orange/apps/product/rpc/model"
	"orange/apps/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductLogic {
	return &ProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductLogic) Product(in *product.ProductItemRequest) (*product.ProductItem, error) {
	v, err, _ := l.svcCtx.SingleGroup.Do(fmt.Sprintf("product:%d", in.ProductId), func() (interface{}, error) {
		return l.svcCtx.ProductModel.FindOne(l.ctx, uint64(in.ProductId))
	})
	if err != nil {
		return nil, err
	}
	p := v.(*model.Product)

	return &product.ProductItem{
		ProductId: int64(p.Id),
		Name:      p.Name,
		Stock:     p.Stock,
	}, nil
}
