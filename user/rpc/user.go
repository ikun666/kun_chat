package main

import (
	"flag"
	"fmt"

	"github.com/ikun666/kun_chat/user/rpc/internal/config"
	"github.com/ikun666/kun_chat/user/rpc/internal/server"
	"github.com/ikun666/kun_chat/user/rpc/internal/svc"
	"github.com/ikun666/kun_chat/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	fmt.Println("rpc 1")
	conf.MustLoad(*configFile, &c)
	// time.Sleep(time.Duration(c.SleepTime) * time.Second)
	// fmt.Println("sleep time", c.SleepTime, "s")
	fmt.Println("rpc 2")
	ctx := svc.NewServiceContext(c)
	fmt.Println("rpc 3")
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
