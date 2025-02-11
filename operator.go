package main

import (
	"flag"
	"fmt"
	"operator-manager/internal/middleware"

	"operator-manager/internal/config"
	"operator-manager/internal/handler"
	"operator-manager/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/operator.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	server.Use(middleware.NewEarlyCheckMiddleware(c).Handle)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
