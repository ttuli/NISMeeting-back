package logic

import (
	"context"

	"Meeting/apps/user/rpc/internal/svc"
	"Meeting/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type OfflineLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOfflineLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OfflineLogic {
	return &OfflineLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OfflineLogic) Offline(in *user.PingReq) (*user.PingResp, error) {
	_, err := l.svcCtx.Redis.DelCtx(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &user.PingResp{}, nil
}
