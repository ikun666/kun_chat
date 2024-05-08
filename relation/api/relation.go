package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/ikun666/kun_chat/relation/api/internal/config"
	"github.com/ikun666/kun_chat/relation/api/internal/handler"
	"github.com/ikun666/kun_chat/relation/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/relation.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	time.Sleep(time.Second * 10)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
