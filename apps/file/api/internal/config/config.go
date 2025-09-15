package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}

	OssConf struct {
		Region     string
		BucketName string
		Dir        string
		Product    string
	}
}
