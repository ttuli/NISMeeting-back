package logic

import (
	"context"

	"Meeting/apps/meeting/rpc/internal/svc"
	"Meeting/apps/meeting/rpc/meeting"

	"github.com/zeromicro/go-zero/core/logx"
)

type HistoryMeetingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHistoryMeetingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HistoryMeetingLogic {
	return &HistoryMeetingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HistoryMeetingLogic) HistoryMeeting(in *meeting.HistoryMeetingReq) (*meeting.HistoryMeetingResp, error) {
	list, err := l.svcCtx.BaseMeetingModel.GetMeetingById(in.UserId)
	if err != nil {
		return nil, err
	}

	var res []*meeting.MeetingEntity
	for _, v := range list {
		res = append(res, &meeting.MeetingEntity{
			MeetingId:       v.MeetingEntity.MeetingId,
			MeetingName:     v.MeetingEntity.MeetingName,
			MeetingPassword: v.MeetingEntity.MeetingPassword,
			StartTime:       v.MeetingEntity.StartTime,
			EndTime:         v.MeetingEntity.EndTime,
			Description:     v.MeetingEntity.Description,
			HostId:          v.MeetingEntity.HostId,
			HostName:        v.MeetingEntity.HostName,
		})
	}
	return &meeting.HistoryMeetingResp{
		List: res,
	}, nil
}
