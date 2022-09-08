package logic

import (
	"context"

	"orange/apps/user/rpc/internal/svc"
	"orange/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserReceiveAddressInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserReceiveAddressInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserReceiveAddressInfoLogic {
	return &GetUserReceiveAddressInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据主键id,查询收获地址
func (l *GetUserReceiveAddressInfoLogic) GetUserReceiveAddressInfo(in *user.UserReceiveAddressInfoReq) (*user.UserReceiveAddress, error) {
	// todo: add your logic here and delete this line

	return &user.UserReceiveAddress{}, nil
}
