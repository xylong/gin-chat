package router

import (
	"gin-chat/api"
	"gin-chat/service/ws"
	"github.com/gin-gonic/gin"
)

// NewRouter 创建路由
func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery(), gin.Logger())

	v1 := r.Group("v1")
	{
		v1.POST("users", api.Register)

		v1.GET("im", ws.Ws)
		v1.POST("all", ws.All)
	}

	return r
}
