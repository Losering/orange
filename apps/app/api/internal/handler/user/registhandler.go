package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"orange/apps/app/api/internal/logic/user"
	"orange/apps/app/api/internal/svc"
	"orange/apps/app/api/internal/types"
)

func RegistHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegistReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewRegistLogic(r.Context(), svcCtx)
		resp, err := l.Regist(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
