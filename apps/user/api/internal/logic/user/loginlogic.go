package user

import (
	"context"
	"fmt"

	"Meeting/apps/user/api/internal/svc"
	"Meeting/apps/user/api/internal/types"
	"Meeting/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户登入
func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	res, err := l.svcCtx.User.Login(l.ctx, &user.LoginReq{
		Phone:    req.Phone,
		Password: req.Password,
	})
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	u := res.User
	return &types.LoginResp{
		Token: res.Token,
		User: types.UserEntity{
			UserId:            u.UserId,
			NickName:          u.NickName,
			Phone:             u.Phone,
			Sex:               int(u.Sex),
			PersonalSignature: u.PersonalSignature,
			CreateTime:        u.CreateTime,
			AreaName:          u.AreaName,
			AreaCode:          u.AreaCode,
		},
	}, nil
}
