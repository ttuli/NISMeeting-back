package logic

import (
	"context"
	"errors"
	"fmt"
	"time"

	"Meeting/apps/user/rpc/internal/svc"
	"Meeting/apps/user/rpc/user"
	"Meeting/pkg/ctxdata"
	"Meeting/pkg/encrypt"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginResp, error) {
	u, err := l.svcCtx.UserModel.GetByPhone(l.ctx, in.Phone)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, errors.New("手机号未注册")
	}
	if !encrypt.ValidatePasswordHash(in.Password, u.Password) {
		return nil, errors.New("密码错误")
	}

	now := time.Now().Unix()
	token, err := ctxdata.GetJwtToken(l.svcCtx.Config.JwtAuth.AccessSecret, now, l.svcCtx.Config.JwtAuth.AccessExpire,
		u.UserId)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("登录失败")
	}
	r, err := l.svcCtx.Redis.GetCtx(l.ctx, u.UserId)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("登录失败")
	}
	if r == "1" {
		return nil, errors.New("账号已登录")
	}
	err = l.svcCtx.Redis.SetexCtx(l.ctx, u.UserId, "1", l.svcCtx.Config.ExpireTime)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("登录失败")
	}
	return &user.LoginResp{
		Token: token,
		User: &user.UserEntity{
			UserId:            u.UserId,
			Phone:             u.Phone,
			PersonalSignature: u.PersonalSignature,
			AreaName:          u.AreaName,
			AreaCode:          u.AreaCode,
			NickName:          u.NickName,
			Sex:               int32(u.Sex),
			CreateTime:        u.CreateTime.Unix(),
		},
	}, nil
}
