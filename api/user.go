package api

import (
	"gin-chat/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

// Register 注册
func Register(ctx *gin.Context) {
	var s service.RegisterService

	err := ctx.ShouldBind(&s)
	if err != nil {
		ctx.JSONP(http.StatusBadRequest, ErrorResponse(err))
		logrus.Info("register", err)
		return
	}

	result := s.Register()
	ctx.JSON(http.StatusOK, result)
}
