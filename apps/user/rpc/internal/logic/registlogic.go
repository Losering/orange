package logic

import (
	"context"

	"orange/apps/user/rpc/internal/svc"
	"orange/apps/user/rpc/model"
	"orange/apps/user/rpc/user"
	"orange/pkg/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegistLogic {
	return &RegistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 注册
func (l *RegistLogic) Regist(in *user.RegistRequest) (*user.RegistResponse, error) {
	md5ByString, err := tools.Md5ByString(in.Password)

	userModel := model.User{
		Username: in.Username,
		Password: md5ByString,
		Phone:    "15600356209",
		Question: "test question",
		Answer:   "answer",
	}
	userDB, err := l.svcCtx.UserModel.Insert(l.ctx, &userModel)
	if err != nil {
		return nil, err
	}
	_, err = userDB.LastInsertId()
	return &user.RegistResponse{
		Id:       int64(12345),
		Username: in.Username,
		Phone:    "15600356209",
	}, nil
}
