package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"gorm-zero-example/app/errorx"

	"gorm-zero-example/services/api/internal/config"
	"gorm-zero-example/services/api/internal/handler"
	"gorm-zero-example/services/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "services/api/etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	httpx.SetErrorHandler(errorx.ErrHandle)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	logx.DisableStat()
	for _, route := range server.Routes() {
		fmt.Println(fmt.Sprintf("Method: %s Path: %s, handler: %v", route.Method, route.Path, route.Handler))
	}
	fmt.Println(fmt.Sprintf("routes count: %d", len(server.Routes())))
	server.Start()
}
