package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	Mysql struct {
		ReadDataSource  string
		WriteDataSource string
	}

	Redisx redis.RedisConf

	ExpireTime int

	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
}
