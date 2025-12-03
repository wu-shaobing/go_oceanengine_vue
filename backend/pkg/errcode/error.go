package errcode

import (
	"fmt"
	"net/http"
)

// AppError 应用错误
type AppError struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Details   interface{} `json:"details,omitempty"`
	RequestID string      `json:"request_id,omitempty"`
	cause     error
}

// Error 实现 error 接口
func (e *AppError) Error() string {
	if e.cause != nil {
		return fmt.Sprintf("code=%d, message=%s, cause=%v", e.Code, e.Message, e.cause)
	}
	return fmt.Sprintf("code=%d, message=%s", e.Code, e.Message)
}

// Unwrap 实现错误链
func (e *AppError) Unwrap() error {
	return e.cause
}

// WithCause 设置原始错误
func (e *AppError) WithCause(err error) *AppError {
	newErr := *e
	newErr.cause = err
	return &newErr
}

// WithDetails 设置详细信息
func (e *AppError) WithDetails(details interface{}) *AppError {
	newErr := *e
	newErr.Details = details
	return &newErr
}

// WithRequestID 设置请求ID
func (e *AppError) WithRequestID(requestID string) *AppError {
	newErr := *e
	newErr.RequestID = requestID
	return &newErr
}

// HTTPStatus 获取 HTTP 状态码
func (e *AppError) HTTPStatus() int {
	switch {
	case e.Code == Success:
		return http.StatusOK
	case e.Code >= 100100 && e.Code < 100200:
		return http.StatusUnauthorized
	case e.Code == ErrPermissionDeny:
		return http.StatusForbidden
	case e.Code == ErrNotFound:
		return http.StatusNotFound
	case e.Code == ErrTooManyRequest:
		return http.StatusTooManyRequests
	case e.Code == ErrInvalidParam || e.Code == ErrAlreadyExists:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

// Cause 获取原始错误
func (e *AppError) Cause() error {
	return e.cause
}

// New 创建新错误
func New(code int) *AppError {
	return &AppError{
		Code:    code,
		Message: Message(code),
	}
}

// NewWithMessage 创建带自定义消息的错误
func NewWithMessage(code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

// Wrap 包装错误
func Wrap(code int, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: Message(code),
		cause:   err,
	}
}

// WrapWithMessage 包装错误（带自定义消息）
func WrapWithMessage(code int, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		cause:   err,
	}
}

// Is 判断错误码
func Is(err error, code int) bool {
	if appErr, ok := err.(*AppError); ok {
		return appErr.Code == code
	}
	return false
}

// GetCode 获取错误码
func GetCode(err error) int {
	if appErr, ok := err.(*AppError); ok {
		return appErr.Code
	}
	return ErrUnknown
}

// IsAppError 判断是否为 AppError
func IsAppError(err error) bool {
	_, ok := err.(*AppError)
	return ok
}
