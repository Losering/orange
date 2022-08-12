package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"orange/apps/app/api/internal/logic"
	"orange/apps/app/api/internal/svc"
	"orange/apps/app/api/internal/types"
)

func CategoryListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CategoryListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCategoryListLogic(r.Context(), svcCtx)
		resp, err := l.CategoryList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
