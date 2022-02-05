package main

import (
	"fmt"
	"gin-chat/conf"
)

func main() {
	conf.Init()
	fmt.Println(conf.ServiceConfig, conf.MysqlConfig, conf.MongoConfig)
}
