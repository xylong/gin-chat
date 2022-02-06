package main

import (
	"fmt"
	"gin-chat/cache"
	"gin-chat/conf"
	"gin-chat/router"
	"gin-chat/service"
)

func main() {
	conf.Init()
	cache.Redis()

	go service.Manager.Start()

	r := router.NewRouter()
	r.Run(fmt.Sprintf(":%d", conf.ServiceConfig.HttpPort))
}
