package handler

import (
	"gorm-zero-example/app/respx"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gorm-zero-example/services/api/internal/logic"
	"gorm-zero-example/services/api/internal/svc"
	"gorm-zero-example/services/api/internal/types"
)

func queryUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewQueryUserLogic(r.Context(), svcCtx)
		resp, err := l.QueryUser(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, respx.NewSuccessResponse(resp))
		}
	}
}
