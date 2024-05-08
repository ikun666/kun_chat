package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/ikun666/kun_chat/user/api/internal/config"
	"github.com/ikun666/kun_chat/user/api/internal/handler"
	"github.com/ikun666/kun_chat/user/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	fmt.Println("api 1")
	conf.MustLoad(*configFile, &c)
	fmt.Println("api 2")
	time.Sleep(time.Second * 10)
	server := rest.MustNewServer(c.RestConf)
	fmt.Println("api 3")
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	fmt.Println("api 4")
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
