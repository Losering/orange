package logic

import (
	"context"
	"fmt"
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
		fmt.Println("user rpc login", err)
		if err == model.ErrNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DataNoExistError), "根据username查询用户信息失败，username:%s,err:%v", in.Username, err)
		}
		return nil, err
	}
	// verify user password
	md5ByString, err := tools.Md5ByString(in.Password)
	if !(md5ByString == userDB.Password) {
		fmt.Println("userdb pwd", userDB.Password)
		fmt.Println("md5ByString pwd", md5ByString)
		return nil, errors.Wrap(xerr.NewErrMsg("账号或密码错误"), "密码错误")
	}
	//return sql
	var resp user.LoginResponse
	_ = copier.Copy(&resp, userDB)
	return &resp, nil
}
