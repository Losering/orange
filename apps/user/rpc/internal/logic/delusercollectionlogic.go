package logic

import (
	"context"

	"orange/apps/user/rpc/internal/svc"
	"orange/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelUserCollectionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelUserCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelUserCollectionLogic {
	return &DelUserCollectionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  删除收藏
func (l *DelUserCollectionLogic) DelUserCollection(in *user.UserCollectionDelReq) (*user.UserCollectionDelRes, error) {
	// todo: add your logic here and delete this line

	return &user.UserCollectionDelRes{}, nil
}
