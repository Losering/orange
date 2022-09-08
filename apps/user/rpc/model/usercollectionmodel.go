package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserCollectionModel = (*customUserCollectionModel)(nil)

type (
	// UserCollectionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserCollectionModel.
	UserCollectionModel interface {
		userCollectionModel
	}

	customUserCollectionModel struct {
		*defaultUserCollectionModel
	}
)

// NewUserCollectionModel returns a model for the database table.
func NewUserCollectionModel(conn sqlx.SqlConn, c cache.CacheConf) UserCollectionModel {
	return &customUserCollectionModel{
		defaultUserCollectionModel: newUserCollectionModel(conn, c),
	}
}
