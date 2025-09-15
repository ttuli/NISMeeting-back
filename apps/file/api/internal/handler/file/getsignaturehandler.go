package file

import (
	"net/http"

	"Meeting/apps/file/api/internal/logic/file"
	"Meeting/apps/file/api/internal/svc"
	"Meeting/apps/file/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// getSignature
func GetSignatureHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetSignatureReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := file.NewGetSignatureLogic(r.Context(), svcCtx)
		resp, err := l.GetSignature(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
