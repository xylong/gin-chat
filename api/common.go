package api

import (
	"encoding/json"
	"fmt"
	"gin-chat/serializer"
	"github.com/go-playground/validator/v10"
	"net/http"
)

// ErrorResponse 错误响应
func ErrorResponse(err error) serializer.Response {
	if _, ok := err.(validator.ValidationErrors); ok {
		return serializer.Response{
			Status: http.StatusBadRequest,
			Data:   nil,
			Msg:    "",
			Error:  fmt.Sprintf(err.Error()),
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.Response{
			Status: http.StatusBadRequest,
			Msg:    "JSON类型不匹配",
			Error:  fmt.Sprint(err),
		}
	}

	return serializer.Response{
		Status: http.StatusBadRequest,
		Msg:    "参数错误",
		Error:  fmt.Sprint(err),
	}
}
