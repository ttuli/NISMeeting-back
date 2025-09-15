package logic

import (
	"context"
	"errors"

	"Meeting/apps/user/model"
	"Meeting/apps/user/rpc/internal/svc"
	"Meeting/apps/user/rpc/user"
	"Meeting/pkg/encrypt"
	"Meeting/pkg/wuid"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.RegisterResp, error) {
	u, err := l.svcCtx.UserModel.GetByPhone(l.ctx, in.Phone)
	if err != nil {
		return nil, err
	}
	if u != nil {
		return nil, errors.New("手机号已注册")
	}
	genPassword, err := encrypt.GenPasswordHash([]byte(in.Password))
	if err != nil {
		return nil, err
	}
	uid := wuid.GenUid(l.svcCtx.Config.Mysql.WriteDataSource)
	err = l.svcCtx.Create(l.ctx, &model.UserEntity{
		UserId:   uid,
		Phone:    in.Phone,
		Password: string(genPassword),
		NickName: "用户" + uid,
		Sex:      1,
	})
	if err != nil {
		return nil, err
	}

	return &user.RegisterResp{}, nil
}
