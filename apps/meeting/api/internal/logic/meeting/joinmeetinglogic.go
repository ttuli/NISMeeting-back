package meeting

import (
	"context"

	"Meeting/apps/meeting/api/internal/svc"
	"Meeting/apps/meeting/api/internal/types"
	"Meeting/apps/meeting/rpc/meeting"

	"github.com/zeromicro/go-zero/core/logx"
)

type JoinMeetingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 加入会议
func NewJoinMeetingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JoinMeetingLogic {
	return &JoinMeetingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JoinMeetingLogic) JoinMeeting(req *types.JoinMeetingReq) (resp *types.JoinMeetingResp, err error) {
	res, err := l.svcCtx.JoinMeeting(l.ctx, &meeting.JoinMeetingReq{
		MeetingId:       req.MeetingId,
		UserId:          req.UserId,
		UserName:        req.UserName,
		MeetingPassword: req.MeetingPassword,
	})
	if err != nil {
		return nil, err
	}
	return &types.JoinMeetingResp{
		Token: res.Token,
	}, err
}
