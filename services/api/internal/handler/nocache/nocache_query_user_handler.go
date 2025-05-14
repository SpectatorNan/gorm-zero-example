package nocache

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gorm-zero-example/services/api/internal/logic/nocache"
	"gorm-zero-example/services/api/internal/svc"
	"gorm-zero-example/services/api/internal/types"
)

func NocacheQueryUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := nocache.NewNocacheQueryUserLogic(r.Context(), svcCtx)
		resp, err := l.NocacheQueryUser(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
