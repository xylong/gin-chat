package router

import (
	"gin-chat/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

// NewRouter 创建路由
func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery(), gin.Logger())

	v1 := r.Group("v1")
	{
		v1.GET("ping", func(context *gin.Context) {
			context.String(http.StatusOK, "pong")
		})

		v1.POST("users", api.Register)
	}

	return r
}
