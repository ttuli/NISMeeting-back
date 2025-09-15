package configserver

import (
	"errors"

	"github.com/zeromicro/go-zero/core/conf"
)

type configServer interface {
	FromJson() ([]byte, error)
}

type ConfigServer struct {
	configServer
	configName string
}

func NewConfigServer(configName string, s configServer) *ConfigServer {
	return &ConfigServer{
		configServer: s,
		configName:   configName,
	}
}

func (s ConfigServer) MustLoad(v any) error {
	if s.configName == "" {
		return errors.New("未设置配置文件")
	}

	if s.configServer == nil {
		conf.MustLoad(s.configName, v)
		return nil
	}

	data, err := s.configServer.FromJson()
	if err != nil {
		return err
	}

	return conf.LoadConfigFromJsonBytes(data, v)
}
