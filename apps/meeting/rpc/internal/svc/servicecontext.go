package svc

import (
	"Meeting/apps/meeting/model"
	"Meeting/apps/meeting/rpc/internal/config"

	lksdk "github.com/livekit/server-sdk-go/v2"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config config.Config
	model.BaseMeetingModel

	*redis.Redis
	*lksdk.RoomServiceClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:            c,
		Redis:             redis.MustNewRedis(c.Redisx),
		BaseMeetingModel:  model.NewMeetingModel(c.Mysql.ReadDataSource, c.Mysql.WriteDataSource),
		RoomServiceClient: lksdk.NewRoomServiceClient(c.RoomServiceClient.Url, c.RoomServiceClient.ApiKey, c.RoomServiceClient.SecretKey),
	}
}
