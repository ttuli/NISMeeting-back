package userinfo

import (
	"context"

	"Meeting/apps/user/api/internal/svc"
	"Meeting/apps/user/rpc/user"
	"Meeting/pkg/ctxdata"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// ping
func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PingLogic) Ping() error {
	_, err := l.svcCtx.User.Ping(l.ctx, &user.PingReq{
		Id: ctxdata.GetUId(l.ctx),
	})
	return err
}
