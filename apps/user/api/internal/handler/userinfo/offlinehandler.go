package userinfo

import (
	"net/http"

	"Meeting/apps/user/api/internal/logic/userinfo"
	"Meeting/apps/user/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 下线
func OfflineHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := userinfo.NewOfflineLogic(r.Context(), svcCtx)
		err := l.Offline()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
