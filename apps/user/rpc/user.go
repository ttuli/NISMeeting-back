package main

import (
	"flag"
	"fmt"

	"Meeting/apps/user/rpc/internal/config"
	"Meeting/apps/user/rpc/internal/server"
	"Meeting/apps/user/rpc/internal/svc"
	"Meeting/apps/user/rpc/user"
	"Meeting/pkg/configserver"

	"github.com/spf13/viper"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

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
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserServer(grpcServer, server.NewUserServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
