package userinfo

import (
	"context"

	"Meeting/apps/user/api/internal/svc"
	"Meeting/apps/user/rpc/user"
	"Meeting/pkg/ctxdata"

	"github.com/zeromicro/go-zero/core/logx"
)

type OfflineLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 下线
func NewOfflineLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OfflineLogic {
	return &OfflineLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OfflineLogic) Offline() error {
	_, err := l.svcCtx.User.Offline(l.ctx, &user.PingReq{
		Id: ctxdata.GetUId(l.ctx),
	})
	return err
}
