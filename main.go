package main

import (
	"fmt"
	"gin-chat/cache"
	"gin-chat/conf"
)

func main() {
	conf.Init()
	fmt.Println(conf.ServiceConfig, conf.MysqlConfig, conf.MongoConfig)
	cache.Redis()
}
