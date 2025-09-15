package meeting

import (
	"net/http"

	"Meeting/apps/meeting/api/internal/logic/meeting"
	"Meeting/apps/meeting/api/internal/svc"
	"Meeting/apps/meeting/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 用户注册
func CreateMeetingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateMeetingReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := meeting.NewCreateMeetingLogic(r.Context(), svcCtx)
		resp, err := l.CreateMeeting(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
