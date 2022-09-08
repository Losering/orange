package user

import (
	"net/http"

	"orange/apps/app/api/internal/logic/user"
	"orange/apps/app/api/internal/svc"
	"orange/apps/app/api/internal/types"
	"orange/pkg/result"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}
		// if err := validator.New().StructCtx(r.Context(), req); err != nil {
		// 	result.ParamErrorResult(r, w, err)
		// 	return
		// }

		l := user.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		result.HttpResult(r, w, resp, err)
	}
}
