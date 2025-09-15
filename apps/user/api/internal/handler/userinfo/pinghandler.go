package userinfo

import (
	"net/http"

	"Meeting/apps/user/api/internal/logic/userinfo"
	"Meeting/apps/user/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// ping
func PingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := userinfo.NewPingLogic(r.Context(), svcCtx)
		err := l.Ping()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
