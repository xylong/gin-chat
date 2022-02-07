package main

import (
	"fmt"
	"gin-chat/cache"
	"gin-chat/conf"
	"gin-chat/router"
)

func main() {
	conf.Init()
	cache.Redis()

	r := router.NewRouter()
	r.Run(fmt.Sprintf(":%d", conf.ServiceConfig.HttpPort))
}
