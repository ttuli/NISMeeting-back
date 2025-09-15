package meeting

import (
	"context"
	"errors"
	"fmt"

	"Meeting/apps/meeting/api/internal/svc"
	"Meeting/apps/meeting/api/internal/types"
	"Meeting/apps/meeting/rpc/meeting"

	"github.com/zeromicro/go-zero/core/logx"
)

type HistoryMeetingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 历史会议
func NewHistoryMeetingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HistoryMeetingLogic {
	return &HistoryMeetingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HistoryMeetingLogic) HistoryMeeting(req *types.HistoryMeetingReq) (resp *types.HistoryMeetingResp, err error) {
	fmt.Println("")
	fmt.Println(*req)
	fmt.Println("")
	res, err := l.svcCtx.HistoryMeeting(l.ctx, &meeting.HistoryMeetingReq{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, errors.New("获取历史会议失败")
	}

	var list []*types.MeetingEntity
	for _, v := range res.List {
		list = append(list, &types.MeetingEntity{
			MeetingId:   v.MeetingId,
			MeetingName: v.MeetingName,
			Description: v.Description,
			HostId:      v.HostId,
			HostName:    v.HostName,
			StartTime:   v.StartTime,
			EndTime:     v.EndTime,
		})
	}
	return &types.HistoryMeetingResp{
		List: list,
	}, nil
}
