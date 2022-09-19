package svc

import (
	"orange/apps/product/rpc/internal/config"
	"orange/apps/product/rpc/model"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"golang.org/x/sync/singleflight"
)

type ServiceContext struct {
	Config        config.Config
	ProductModel  model.ProductModel
	SingleGroup   singleflight.Group
	CategoryModel model.CategoryModel
	BizRedis      *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:        c,
		ProductModel:  model.NewProductModel(sqlConn, c.CacheRedis),
		CategoryModel: model.NewCategoryModel(sqlConn, c.CacheRedis),
		BizRedis:      redis.New(c.BizRedis.Host, redis.WithPass(c.BizRedis.Pass)),
	}
}
