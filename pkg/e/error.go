package e

// ErrCode 错误码
type ErrCode int64

// CustomError 自定义错误
//type CustomError struct {
//	Code    ErrCode `json:"code"`
//	Message string  `json:"message"`
//}

// NewCustomError 实例话自定义错误
//func NewCustomError(code ErrCode) error {
//	return errors.Wrap(&CustomError{
//		Code:    code,
//Message: code.String(),
//}, "")
//}

// Error 获取错误码对应信息
//func (e *CustomError) Error() string {
//	return e.Code.String()
//}

//go:generate stringer -type ErrCode -linecomment
const (
	OK        ErrCode = 0     // 成功
	WsUpgrade ErrCode = 10001 // 协议升级失败
)
