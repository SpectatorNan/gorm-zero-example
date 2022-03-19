package errorx

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
	"gorm-zero-example/app/respx"
	"net/http"
)

func JwtUnAuthorizedHandle() rest.RunOption {
	return rest.WithUnauthorizedCallback(func(w http.ResponseWriter, r *http.Request, err error) {
		logx.Info("===========jwt=WithUnauthorizedCallback=====================")
		httpx.Error(w, NewMsgCodeError(UnLoginCode, "请登录"))
	})
}

func ErrHandle(err error) (int, interface{}) {
	switch e := err.(type) {
	case *CodeError:
		return http.StatusOK, e.DataInfo()
	default:
		st, ok := status.FromError(err)
		if ok {
			return http.StatusOK, &respx.Response{
				Code:    int(st.Code()),
				Message: st.Message(),
				Data:    nil,
				Reason:  err.Error(),
			}
		}
		return http.StatusOK, &respx.Response{
			Code:    10001,
			Message: "server error",
			Data:    nil,
			Reason:  err.Error(),
		}
	}
}
