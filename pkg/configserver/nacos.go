package configserver

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

type NacosConfig struct {
	Host      string `mapstructure:"host"`
	Port      uint64 `mapstructure:"port"`
	Namespace string `mapstructure:"namespace"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	DataId    string `mapstructure:"dataid"`
	Group     string `mapstructure:"group"`
}

type NacosServer struct {
	config *NacosConfig
}

func NewNacosServer(c *NacosConfig) *NacosServer {
	return &NacosServer{
		config: c,
	}
}

func (s *NacosServer) FromJson() ([]byte, error) {
	sc := []constant.ServerConfig{
		{
			IpAddr: s.config.Host,
			Port:   s.config.Port,
		},
	}

	cc := constant.ClientConfig{
		NamespaceId:         s.config.Namespace, // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "nacos/log",
		CacheDir:            "nacos/cache",
		LogLevel:            "debug",
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	if err != nil {
		return nil, err
	}
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: s.config.DataId,
		Group:  s.config.Group})

	if err != nil {
		return nil, err
	}
	return []byte(content), nil
}
