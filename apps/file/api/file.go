package main

import (
	"flag"
	"fmt"
	"net/http"

	"Meeting/apps/file/api/internal/config"
	"Meeting/apps/file/api/internal/handler"
	"Meeting/apps/file/api/internal/svc"
	"Meeting/pkg/configserver"

	"github.com/spf13/viper"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/file.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	nacosConfig := &configserver.NacosConfig{}
	v := viper.New()
	v.SetConfigFile(*configFile)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := v.Unmarshal(nacosConfig); err != nil {
		panic(err)
	}
	configServer := configserver.NewConfigServer(*configFile, configserver.NewNacosServer(nacosConfig))
	err := configServer.MustLoad(&c)
	if err != nil {
		panic(err)
	}

	server := rest.MustNewServer(c.RestConf, rest.WithCustomCors(func(header http.Header) {
		header.Set("Access-Control-Allow-Origin", "http://localhost:5173")
		header.Add("Access-Control-Allow-Headers", "Content-Type, Authorization,X-Requested-With")
		header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		header.Set("Access-Control-Allow-Credentials", "true")
		header.Set("Access-Control-Expose-Headers", "Content-Length, Content-Type")
	}, nil, "*"))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
