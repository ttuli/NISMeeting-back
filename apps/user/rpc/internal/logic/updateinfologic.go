package logic

import (
	"context"

	"Meeting/apps/user/model"
	"Meeting/apps/user/rpc/internal/svc"
	"Meeting/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateInfoLogic {
	return &UpdateInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateInfoLogic) UpdateInfo(in *user.UpdateInfoReq) (*user.UpdateInfoResp, error) {
	u := in.User
	err := l.svcCtx.UserModel.Update(l.ctx, &model.UserEntity{
		UserId:            u.UserId,
		Phone:             u.Phone,
		PersonalSignature: u.PersonalSignature,
		NickName:          u.NickName,
		AreaName:          u.AreaName,
		AreaCode:          u.AreaCode,
		Sex:               int(u.Sex),
	})
	if err != nil {
		return nil, err
	}
	return &user.UpdateInfoResp{}, nil
}
