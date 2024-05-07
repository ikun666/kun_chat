package main

import (
	"flag"
	"net/http"

	"github.com/ikun666/kun_chat/chat/testws/handler"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
)

var (
	port    = flag.Int("port", 8200, "the port to listen")
	timeout = flag.Int64("timeout", 0, "timeout of milliseconds")
	cpu     = flag.Int64("cpu", 500, "cpu threshold")
)

func main() {
	flag.Parse()

	logx.Disable()
	engine := rest.MustNewServer(rest.RestConf{
		ServiceConf: service.ServiceConf{
			Log: logx.LogConf{
				Mode: "console",
			},
		},
		Host:         "localhost",
		Port:         *port,
		Timeout:      *timeout,
		CpuThreshold: *cpu,
	})
	defer engine.Stop()

	engine.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/api/chat/ws",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			handler.Chat(w, r)
		},
	})
	engine.AddRoute(rest.Route{
		Method: http.MethodPost,
		Path:   "/api/chat/getmsg",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			handler.GetMsg(w, r)
		},
	})

	engine.Start()
}
