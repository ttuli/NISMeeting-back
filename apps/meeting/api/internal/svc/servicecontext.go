package svc

import (
	"Meeting/apps/meeting/api/internal/config"
	"Meeting/apps/meeting/rpc/meetingclient"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	meetingclient.Meeting
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		Meeting: meetingclient.NewMeeting(zrpc.MustNewClient(c.MeetingRpc)),
	}
}
