package logic

import (
	"context"

	"orange/apps/user/rpc/internal/svc"
	"orange/apps/user/rpc/model"
	"orange/apps/user/rpc/user"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户信息
func (l *UserInfoLogic) UserInfo(in *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	umsMember, err := l.svcCtx.UserModel.FindOne(l.ctx, uint64(in.Id))
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "用户不存在")
		}
		return nil, status.Error(500, err.Error())
	}
	var resp user.UserInfo
	_ = copier.Copy(&resp, umsMember)
	resp.CreateTime = umsMember.CreateTime.Unix()
	resp.UpdateTime = umsMember.UpdateTime.Unix()
	return &user.UserInfoResponse{
		User: &resp,
	}, nil
}
