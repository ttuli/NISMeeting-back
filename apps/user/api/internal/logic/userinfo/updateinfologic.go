package userinfo

import (
	"context"

	"Meeting/apps/user/api/internal/svc"
	"Meeting/apps/user/api/internal/types"
	"Meeting/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新信息
func NewUpdateInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateInfoLogic {
	return &UpdateInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateInfoLogic) UpdateInfo(req *types.UpdateInfoReq) (resp *types.UpdateInfoResp, err error) {
	u := req.User
	_, err = l.svcCtx.UpdateInfo(l.ctx, &user.UpdateInfoReq{
		User: &user.UserEntity{
			UserId:            u.UserId,
			NickName:          u.NickName,
			AreaName:          u.AreaName,
			AreaCode:          u.AreaCode,
			Phone:             u.Phone,
			PersonalSignature: u.PersonalSignature,
			Sex:               int32(u.Sex),
		},
	})
	if err != nil {
		return nil, err
	}
	return &types.UpdateInfoResp{
		User: types.UserEntity{
			UserId:            u.UserId,
			NickName:          u.NickName,
			AreaName:          u.AreaName,
			AreaCode:          u.AreaCode,
			PersonalSignature: u.PersonalSignature,
			Phone:             u.Phone,
			Sex:               int(u.Sex),
		},
	}, nil
}
