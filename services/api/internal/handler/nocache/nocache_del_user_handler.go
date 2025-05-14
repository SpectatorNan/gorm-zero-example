package nocache

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gorm-zero-example/services/api/internal/logic/nocache"
	"gorm-zero-example/services/api/internal/svc"
	"gorm-zero-example/services/api/internal/types"
)

func NocacheDelUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := nocache.NewNocacheDelUserLogic(r.Context(), svcCtx)
		err := l.NocacheDelUser(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
