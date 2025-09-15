package meeting

import (
	"net/http"

	"Meeting/apps/meeting/api/internal/logic/meeting"
	"Meeting/apps/meeting/api/internal/svc"
	"Meeting/apps/meeting/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 加入会议
func JoinMeetingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.JoinMeetingReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := meeting.NewJoinMeetingLogic(r.Context(), svcCtx)
		resp, err := l.JoinMeeting(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
