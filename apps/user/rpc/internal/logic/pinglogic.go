package logic

import (
	"context"

	"Meeting/apps/user/rpc/internal/svc"
	"Meeting/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *user.PingReq) (*user.PingResp, error) {
	err := l.svcCtx.Redis.SetexCtx(l.ctx, in.Id, "1", l.svcCtx.Config.ExpireTime)
	if err != nil {
		return nil, err
	}
	return &user.PingResp{}, nil
}
