package service

import (
	"fmt"
	"gin-chat/model"
	"gin-chat/serializer"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

const (
	// PasswordCost 密码强度
	PasswordCost = 12
)

// RegisterService 用户注册
type RegisterService struct {
	Username string `json:"username" form:"username" binding:"required,min=2,max=10"`
	Password string `json:"password" form:"password" binding:"required,min=6,max=16"`
}

func (service *RegisterService) Register() serializer.Response {
	user := &model.User{}
	res := model.DB.Where("username=?", "").First(user)

	if res.RowsAffected != 0 {
		return serializer.Response{
			Status: http.StatusBadRequest,
			Msg:    "用户已存在",
		}
	}
	user.Username = service.Username

	bytes, err := bcrypt.GenerateFromPassword([]byte(service.Password), PasswordCost)
	if err != nil {
		return serializer.Response{
			Status: http.StatusBadRequest,
			Msg:    fmt.Sprintf(res.Error.Error()),
		}
	}
	user.Password = string(bytes)

	res = model.DB.Create(user)
	if res.Error != nil {
		return serializer.Response{
			Status: http.StatusInternalServerError,
			Msg:    fmt.Sprintf(res.Error.Error()),
		}
	}

	return serializer.Response{
		Status: http.StatusOK,
		Msg:    "ok",
	}
}
