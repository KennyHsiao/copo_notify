package main

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"log"

	"com.copo/copo_notify/notify/internal/config"
	"com.copo/copo_notify/notify/internal/handler"
	"com.copo/copo_notify/notify/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var (
	configFile = flag.String("f", "etc/notify.yaml", "the config file")
	envFile    = flag.String("env", "etc/.env", "the env file")
)

func main() {
	flag.Parse()
	err := godotenv.Load(*envFile)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
