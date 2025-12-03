package response

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"oceanengine-backend/pkg/errcode"
)

// Response 统一响应结构
type Response struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
	RequestID string      `json:"request_id,omitempty"`
	Timestamp int64       `json:"timestamp"`
}

// PageData 分页数据
type PageData struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
}

// Success 成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:      errcode.Success,
		Message:   "成功",
		Data:      data,
		RequestID: c.GetString("request_id"),
		Timestamp: time.Now().UnixMilli(),
	})
}

// SuccessWithMessage 成功响应（带自定义消息）
func SuccessWithMessage(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:      errcode.Success,
		Message:   message,
		Data:      data,
		RequestID: c.GetString("request_id"),
		Timestamp: time.Now().UnixMilli(),
	})
}

// SuccessWithPage 分页成功响应
func SuccessWithPage(c *gin.Context, list interface{}, total int64, page, pageSize int) {
	Success(c, PageData{
		List:     list,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	})
}

// Error 错误响应
func Error(c *gin.Context, err error) {
	requestID := c.GetString("request_id")

	if appErr, ok := err.(*errcode.AppError); ok {
		c.JSON(appErr.HTTPStatus(), Response{
			Code:      appErr.Code,
			Message:   appErr.Message,
			RequestID: requestID,
			Timestamp: time.Now().UnixMilli(),
		})
		return
	}

	// 未知错误
	c.JSON(http.StatusInternalServerError, Response{
		Code:      errcode.ErrInternalServer,
		Message:   "服务器内部错误",
		RequestID: requestID,
		Timestamp: time.Now().UnixMilli(),
	})
}

// ErrorWithDetails 带详情的错误响应
func ErrorWithDetails(c *gin.Context, err *errcode.AppError, details interface{}) {
	requestID := c.GetString("request_id")

	c.JSON(err.HTTPStatus(), Response{
		Code:      err.Code,
		Message:   err.Message,
		Data:      details,
		RequestID: requestID,
		Timestamp: time.Now().UnixMilli(),
	})
}

// ErrorWithCode 根据错误码响应
func ErrorWithCode(c *gin.Context, code int) {
	err := errcode.New(code)
	Error(c, err)
}

// ErrorWithMessage 带自定义消息的错误响应
func ErrorWithMessage(c *gin.Context, code int, message string) {
	requestID := c.GetString("request_id")
	err := errcode.NewWithMessage(code, message)

	c.JSON(err.HTTPStatus(), Response{
		Code:      err.Code,
		Message:   err.Message,
		RequestID: requestID,
		Timestamp: time.Now().UnixMilli(),
	})
}

// BadRequest 参数错误响应
func BadRequest(c *gin.Context, message string) {
	ErrorWithMessage(c, errcode.ErrInvalidParam, message)
}

// Unauthorized 未授权响应
func Unauthorized(c *gin.Context, message string) {
	if message == "" {
		message = "请先登录"
	}
	ErrorWithMessage(c, errcode.ErrUnauthorized, message)
}

// Forbidden 禁止访问响应
func Forbidden(c *gin.Context, message string) {
	if message == "" {
		message = "权限不足"
	}
	ErrorWithMessage(c, errcode.ErrPermissionDeny, message)
}

// NotFound 资源不存在响应
func NotFound(c *gin.Context, message string) {
	if message == "" {
		message = "资源不存在"
	}
	ErrorWithMessage(c, errcode.ErrNotFound, message)
}

// InternalError 内部错误响应
func InternalError(c *gin.Context, message string) {
	if message == "" {
		message = "服务器内部错误"
	}
	ErrorWithMessage(c, errcode.ErrInternalServer, message)
}

// OK 成功响应（无数据）
func OK(c *gin.Context) {
	Success(c, nil)
}

// OKWithData 成功响应（带数据）
func OKWithData(c *gin.Context, data interface{}) {
	Success(c, data)
}

// OKWithList 成功响应（分页数据）
func OKWithList(c *gin.Context, list interface{}, total int64, page, pageSize int) {
	SuccessWithPage(c, list, total, page, pageSize)
}

// Fail 错误响应
func Fail(c *gin.Context, err error) {
	Error(c, err)
}
