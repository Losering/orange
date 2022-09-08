package user

import (
	"context"
	"fmt"

	"orange/apps/app/api/internal/svc"
	"orange/apps/app/api/internal/types"
	"orange/apps/user/rpc/user"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type RegistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegistLogic {
	return &RegistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegistLogic) Regist(req *types.RegistReq) (resp *types.RegistResp, err error) {
	var registReq user.RegistRequest
	registReq.Username = req.Username
	registReq.Password = req.Password
	res, err := l.svcCtx.UserRPC.Regist(l.ctx, &registReq)
	fmt.Println(res)
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	return &types.RegistResp{}, nil
}
