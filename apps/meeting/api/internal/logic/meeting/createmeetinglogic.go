package meeting

import (
	"context"

	"Meeting/apps/meeting/api/internal/svc"
	"Meeting/apps/meeting/api/internal/types"
	"Meeting/apps/meeting/rpc/meeting"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMeetingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户注册
func NewCreateMeetingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMeetingLogic {
	return &CreateMeetingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateMeetingLogic) CreateMeeting(req *types.CreateMeetingReq) (resp *types.CreateMeetingResp, err error) {
	res, err := l.svcCtx.Meeting.CreateMeeting(l.ctx, &meeting.CreateMeetingReq{
		MeetingName:     req.MeetingName,
		MeetingPassword: req.MeetingPassword,
		Description:     req.Description,
		HostName:        req.HostName,
		HostId:          req.HostId,
	})
	if err != nil {
		return nil, err
	}
	m := types.MeetingEntity{}
	copier.Copy(&m, res.Info)
	return &types.CreateMeetingResp{
		Info:  m,
		Token: res.Token,
	}, nil
}
