package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"Meeting/apps/meeting/model"
	datastruct "Meeting/apps/meeting/rpc/internal/dataStruct"
	"Meeting/apps/meeting/rpc/internal/svc"
	"Meeting/apps/meeting/rpc/meeting"
	"Meeting/pkg/ctxdata"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/livekit/protocol/livekit"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type CreateMeetingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateMeetingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMeetingLogic {
	return &CreateMeetingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateMeetingLogic) CreateMeeting(in *meeting.CreateMeetingReq) (*meeting.CreateMeetingResp, error) {
	uid := uuid.New()
	now := time.Now().Unix()
	var r *livekit.Room
	t := ""
	err := l.svcCtx.BaseMeetingModel.Transaction(func(tx *gorm.DB) error {
		m := make([]*model.SingleMember, 0)
		m = append(m, &model.SingleMember{
			Uid:  in.HostId,
			Name: in.HostName,
		})
		meta := datastruct.MeetingInfo{
			MeetingId:       uid.String(),
			MeetingPassword: in.MeetingPassword,
			MeetingName:     in.MeetingName,
			Description:     in.Description,
			HostId:          in.HostId,
			HostName:        in.HostName,
			StartTime:       now,
		}

		d, err := json.Marshal(meta)
		if err != nil {
			return err
		}
		room, err := l.svcCtx.RoomServiceClient.CreateRoom(l.ctx, &livekit.CreateRoomRequest{
			Name:             uid.String(),
			DepartureTimeout: 10,
			EmptyTimeout:     10,
			Metadata:         string(d),
		})
		r = room
		if err != nil {
			return err
		}
		err = tx.Create(&model.MeetingEntity{
			MeetingId:       uid.String(),
			MeetingName:     in.MeetingName,
			Description:     in.Description,
			HostId:          in.HostId,
			HostName:        in.HostName,
			MeetingPassword: in.MeetingPassword,
			StartTime:       now,
		}).Error
		if err != nil {
			return err
		}
		err = tx.Create(&model.MeetingHistory{
			MeetingId: uid.String(),
			UserId:    in.HostId,
		}).Error
		if err != nil {
			return err
		}
		token, err := ctxdata.GetLiveKitToken(l.svcCtx.Config.RoomServiceClient.ApiKey,
			l.svcCtx.Config.RoomServiceClient.SecretKey,
			in.HostId,
			uid.String(),
			map[string]string{
				"uid":  in.HostId,
				"name": in.HostName,
			})
		if err != nil {
			return err
		}
		t = token
		return nil
	})
	if err != nil || r == nil || t == "" {
		fmt.Println("")
		fmt.Println(err)
		fmt.Println("")
		return nil, errors.New("创建会议失败")
	}

	info := &meeting.MeetingEntity{
		MeetingId: uid.String(),
		StartTime: now,
	}
	copier.Copy(info, in)
	return &meeting.CreateMeetingResp{
		Info:  info,
		Token: t,
	}, nil
}
