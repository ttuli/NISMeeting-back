package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"Meeting/apps/meeting/model"
	datastruct "Meeting/apps/meeting/rpc/internal/dataStruct"
	"Meeting/apps/meeting/rpc/internal/svc"
	"Meeting/apps/meeting/rpc/meeting"
	"Meeting/pkg/ctxdata"

	"github.com/livekit/protocol/livekit"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type JoinMeetingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewJoinMeetingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JoinMeetingLogic {
	return &JoinMeetingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *JoinMeetingLogic) JoinMeeting(in *meeting.JoinMeetingReq) (*meeting.JoinMeetingResp, error) {
	res, err := l.svcCtx.RoomServiceClient.ListRooms(l.ctx, &livekit.ListRoomsRequest{
		Names: []string{
			in.MeetingId,
		},
	})
	if err != nil {
		l.Logger.Error(err)
		return nil, errors.New("查找会议房间失败")
	}
	if len(res.Rooms) == 0 {
		return nil, errors.New("会议不存在或已结束")
	}
	t := ""
	var inf datastruct.MeetingInfo
	err = json.Unmarshal([]byte(res.Rooms[0].Metadata), &inf)
	if err != nil {
		fmt.Println("json error", err)
		return nil, err
	}
	if inf.MeetingPassword != in.MeetingPassword {
		return nil, fmt.Errorf("密码错误")
	}

	err = l.svcCtx.BaseMeetingModel.Transaction(func(tx *gorm.DB) error {
		err = tx.Create(&model.MeetingHistory{
			MeetingId: in.MeetingId,
			UserId:    in.UserId,
		}).Error
		if err != nil {
			l.Logger.Error(err)
			return err
		}
		t, err = ctxdata.GetLiveKitToken(l.svcCtx.Config.RoomServiceClient.ApiKey,
			l.svcCtx.Config.RoomServiceClient.SecretKey,
			in.UserId, in.MeetingId, map[string]string{
				"uid":  in.UserId,
				"name": in.UserName,
			})
		if err != nil {
			l.Logger.Error(err)
			return err
		}
		return nil
	})
	if err != nil {
		l.Logger.Error(err)
		return nil, errors.New("加入会议失败")
	}
	return &meeting.JoinMeetingResp{
		Token: t,
	}, nil
}
