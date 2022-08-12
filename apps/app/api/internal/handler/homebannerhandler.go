package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"orange/apps/app/api/internal/logic"
	"orange/apps/app/api/internal/svc"
)

func HomeBannerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewHomeBannerLogic(r.Context(), svcCtx)
		resp, err := l.HomeBanner()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
