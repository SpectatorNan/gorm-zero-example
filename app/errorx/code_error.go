package errorx

import (
	"gorm-zero-example/app/respx"
)

const defaultCode = 10001

// 00 公共模块
const (
	DefaultCode              = 10001
	NotFoundResourceCode     = 20001 // 资源不存在
	ExistsResourceCode       = 20002 // 资源已存在
	DBErrorCode              = 20003 // db 操作异常
	UnLoginCode              = 401
	MarshalCode              = 20004 // 序列化 反序列化异常
	UserPasswordNotMatchCode = 30001 // 密码不正确
)

type CodeError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Reason  string      `json:"reason"`
}

func (e *CodeError) Error() string {
	return e.Message
}

func NewCodeError(code int, message string, reason string, data interface{}) error {
	return &CodeError{Code: code, Message: message, Data: data, Reason: reason}
}

func NewMsgReasonCodeError(code int, message, reason string) error {
	if len(message) < 1 {
		message = "操作异常"
	}
	return NewCodeError(code, message, reason, nil)
}

func NewMsgCodeError(code int, message string) error {
	return NewCodeError(code, message, message, nil)
}

func NewReasonError(message, reason string) error {
	return NewCodeError(defaultCode, message, reason, nil)
}
func NewDefaultError(message string) error {
	return NewCodeError(defaultCode, message, message, nil)
}
func NewCodeReason(code int, reason string) error {
	return NewCodeError(code, "操作异常", reason, nil)
}
func NewUnknownError(reason string) error {
	return NewCodeError(defaultCode, "操作异常", reason, nil)
}

func (e *CodeError) DataInfo() *respx.Response {
	return &respx.Response{
		Code:    e.Code,
		Message: e.Message,
		Reason:  e.Reason,
	}
}
