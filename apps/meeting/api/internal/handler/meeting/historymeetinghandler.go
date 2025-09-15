package meeting

import (
	"net/http"

	"Meeting/apps/meeting/api/internal/logic/meeting"
	"Meeting/apps/meeting/api/internal/svc"
	"Meeting/apps/meeting/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 历史会议
func HistoryMeetingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HistoryMeetingReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := meeting.NewHistoryMeetingLogic(r.Context(), svcCtx)
		resp, err := l.HistoryMeeting(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
