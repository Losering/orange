package logic

import (
	"context"
	"orange/apps/user/rpc/internal/svc"
	"orange/apps/user/rpc/model"
	"orange/apps/user/rpc/user"
	"orange/pkg/tools"
	"orange/pkg/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 登录
func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	//verfi user exists
	userDB, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.Wrapf(xerr.NewErrCodeMsg(xerr.UserNoErr, "账号不存在"), "账号不存在")
		}
		return nil, err
	}
	// verify user password
	md5ByString, err := tools.Md5ByString(in.Password)
	if !(md5ByString == userDB.Password) {
		return nil, errors.Wrap(xerr.NewErrCodeMsg(xerr.PasswordErr, "账号或密码错误"), "密码错误")
	}
	//return sql
	var resp user.LoginResponse
	_ = copier.Copy(&resp, userDB)
	return &resp, nil
}
