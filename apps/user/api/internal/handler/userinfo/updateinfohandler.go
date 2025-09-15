package userinfo

import (
	"net/http"

	"Meeting/apps/user/api/internal/logic/userinfo"
	"Meeting/apps/user/api/internal/svc"
	"Meeting/apps/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 更新信息
func UpdateInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := userinfo.NewUpdateInfoLogic(r.Context(), svcCtx)
		resp, err := l.UpdateInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
