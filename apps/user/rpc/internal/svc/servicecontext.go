package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"Meeting/apps/user/model"
	"Meeting/apps/user/rpc/internal/config"
)

type ServiceContext struct {
	Config       config.Config
	ReadSqlConn  sqlx.SqlConn
	WriteSqlConn sqlx.SqlConn

	*model.UserModel
	*redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	ReadSqlConn := sqlx.NewMysql(c.Mysql.ReadDataSource)
	WriteSqlConn := sqlx.NewMysql(c.Mysql.WriteDataSource)
	return &ServiceContext{
		Config:       c,
		ReadSqlConn:  ReadSqlConn,
		WriteSqlConn: WriteSqlConn,

		UserModel: model.NewUserModel(ReadSqlConn, WriteSqlConn),
		Redis:     redis.MustNewRedis(c.Redisx),
	}
}
