package e

import "github.com/pkg/errors"

// ErrCode 错误码
type ErrCode int64

// CustomError 自定义错误
type CustomError struct {
	Code    ErrCode `json:"code"`
	Message string  `json:"message"`
}

// NewCustomError 实例话自定义错误
func NewCustomError(code ErrCode) error {
	return errors.Wrap(&CustomError{
		Code:    code,
		Message: code.String(),
	}, "")
}

// Error 获取错误码对应信息
func (e *CustomError) Error() string {
	return e.Code.String()
}

//go:generate stringer -type ErrCode -linecomment
const (
	OK             ErrCode = 0     // 成功
	WsParseMsg     ErrCode = 10001 // 解析内容信息
	WsRecord       ErrCode = 10002 // 发送信息，请求历史记录
	WsRecordEnd    ErrCode = 10003 // 没有更多记录了
	WsOnlineReply  ErrCode = 10004 // 在线应答成功
	WsOfflineReply ErrCode = 10004 // 离线应答成功
	WsMsgLimit     ErrCode = 10005 // 请求受到限制
)
