# 错误处理方案

## 概述

统一的错误处理机制是保证系统稳定性和用户体验的关键。本文档描述系统中错误的定义、传递和处理方式。

## 错误码设计

### 错误码规范

```
错误码格式: XXYYYY
- XX: 模块代码 (00-99)
- YYYY: 错误序号 (0001-9999)

模块代码:
- 00: 通用错误
- 10: 认证授权
- 20: 用户管理
- 30: 广告主管理
- 31: 广告系列管理
- 32: 广告组管理
- 33: 创意管理
- 34: 素材管理
- 35: 人群管理
- 40: 数据报表
- 50: 系统管理
- 90: Ocean Engine API
```

### 错误码定义

```go
// pkg/errcode/errcode.go
package errcode

// 通用错误码 (00xxxx)
const (
    Success           = 0      // 成功
    ErrUnknown        = 1      // 未知错误
    ErrInvalidParam   = 100001 // 参数错误
    ErrNotFound       = 100002 // 资源不存在
    ErrAlreadyExists  = 100003 // 资源已存在
    ErrPermissionDeny = 100004 // 权限不足
    ErrInternalServer = 100005 // 服务器内部错误
    ErrServiceUnavail = 100006 // 服务不可用
    ErrTimeout        = 100007 // 请求超时
    ErrTooManyRequest = 100008 // 请求过于频繁
)

// 认证授权错误码 (10xxxx)
const (
    ErrUnauthorized     = 100100 // 未登录
    ErrTokenInvalid     = 100101 // Token 无效
    ErrTokenExpired     = 100102 // Token 已过期
    ErrLoginFailed      = 100103 // 登录失败
    ErrCaptchaInvalid   = 100104 // 验证码错误
    ErrPasswordWrong    = 100105 // 密码错误
    ErrAccountLocked    = 100106 // 账号已锁定
    ErrAccountDisabled  = 100107 // 账号已禁用
    ErrRefreshTokenInvalid = 100108 // 刷新 Token 无效
)

// 用户管理错误码 (20xxxx)
const (
    ErrUserNotFound     = 200001 // 用户不存在
    ErrUserExists       = 200002 // 用户已存在
    ErrUsernameInvalid  = 200003 // 用户名格式错误
    ErrPasswordInvalid  = 200004 // 密码格式错误
    ErrEmailInvalid     = 200005 // 邮箱格式错误
    ErrPhoneInvalid     = 200006 // 手机号格式错误
)

// 广告主管理错误码 (30xxxx)
const (
    ErrAdvertiserNotFound   = 300001 // 广告主不存在
    ErrAdvertiserExists     = 300002 // 广告主已存在
    ErrAdvertiserDisabled   = 300003 // 广告主已禁用
    ErrAdvertiserAuthFailed = 300004 // 广告主授权失败
    ErrAdvertiserSyncFailed = 300005 // 广告主同步失败
)

// 广告系列错误码 (31xxxx)
const (
    ErrCampaignNotFound   = 310001 // 广告系列不存在
    ErrCampaignExists     = 310002 // 广告系列已存在
    ErrCampaignCreateFail = 310003 // 创建广告系列失败
    ErrCampaignUpdateFail = 310004 // 更新广告系列失败
    ErrCampaignDeleteFail = 310005 // 删除广告系列失败
    ErrCampaignStatusFail = 310006 // 状态变更失败
)

// Ocean Engine API 错误码 (90xxxx)
const (
    ErrOEAPIFailed       = 900001 // API 调用失败
    ErrOETokenInvalid    = 900002 // Ocean Engine Token 无效
    ErrOETokenExpired    = 900003 // Ocean Engine Token 过期
    ErrOERateLimit       = 900004 // 请求频率限制
    ErrOEParamInvalid    = 900005 // 参数错误
    ErrOEResourceNotFound = 900006 // 资源不存在
)
```

### 错误消息映射

```go
// pkg/errcode/message.go
package errcode

var messages = map[int]string{
    Success:           "成功",
    ErrUnknown:        "未知错误",
    ErrInvalidParam:   "参数错误",
    ErrNotFound:       "资源不存在",
    ErrAlreadyExists:  "资源已存在",
    ErrPermissionDeny: "权限不足",
    ErrInternalServer: "服务器内部错误",
    ErrServiceUnavail: "服务不可用",
    ErrTimeout:        "请求超时",
    ErrTooManyRequest: "请求过于频繁",
    
    ErrUnauthorized:     "请先登录",
    ErrTokenInvalid:     "Token 无效",
    ErrTokenExpired:     "Token 已过期",
    ErrLoginFailed:      "登录失败",
    ErrCaptchaInvalid:   "验证码错误",
    ErrPasswordWrong:    "密码错误",
    ErrAccountLocked:    "账号已锁定，请稍后再试",
    ErrAccountDisabled:  "账号已禁用",
    
    ErrUserNotFound:     "用户不存在",
    ErrUserExists:       "用户已存在",
    ErrUsernameInvalid:  "用户名格式错误",
    ErrPasswordInvalid:  "密码格式不符合要求",
    
    ErrAdvertiserNotFound:   "广告主不存在",
    ErrAdvertiserExists:     "广告主已存在",
    ErrAdvertiserDisabled:   "广告主已禁用",
    ErrAdvertiserAuthFailed: "广告主授权失败",
    
    ErrCampaignNotFound:   "广告系列不存在",
    ErrCampaignCreateFail: "创建广告系列失败",
    ErrCampaignUpdateFail: "更新广告系列失败",
    
    ErrOEAPIFailed:    "Ocean Engine API 调用失败",
    ErrOETokenInvalid: "广告主授权已失效，请重新授权",
    ErrOERateLimit:    "请求过于频繁，请稍后再试",
}

// Message 获取错误消息
func Message(code int) string {
    if msg, ok := messages[code]; ok {
        return msg
    }
    return "未知错误"
}
```

---

## 错误结构定义

```go
// pkg/errcode/error.go
package errcode

import (
    "fmt"
    "net/http"
)

// AppError 应用错误
type AppError struct {
    Code      int         `json:"code"`      // 错误码
    Message   string      `json:"message"`   // 错误消息
    Details   interface{} `json:"details,omitempty"` // 详细信息
    RequestID string      `json:"request_id,omitempty"` // 请求ID
    cause     error       // 原始错误
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
    e.cause = err
    return e
}

// WithDetails 设置详细信息
func (e *AppError) WithDetails(details interface{}) *AppError {
    e.Details = details
    return e
}

// WithRequestID 设置请求ID
func (e *AppError) WithRequestID(requestID string) *AppError {
    e.RequestID = requestID
    return e
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

// Is 判断错误码
func Is(err error, code int) bool {
    if appErr, ok := err.(*AppError); ok {
        return appErr.Code == code
    }
    return false
}
```

---

## 错误响应格式

```go
// pkg/response/response.go
package response

import (
    "github.com/gin-gonic/gin"
)

// Response 统一响应结构
type Response struct {
    Code      int         `json:"code"`
    Message   string      `json:"message"`
    Data      interface{} `json:"data,omitempty"`
    RequestID string      `json:"request_id,omitempty"`
}

// Success 成功响应
func Success(c *gin.Context, data interface{}) {
    c.JSON(200, Response{
        Code:      0,
        Message:   "成功",
        Data:      data,
        RequestID: c.GetString("request_id"),
    })
}

// Error 错误响应
func Error(c *gin.Context, err error) {
    requestID := c.GetString("request_id")
    
    if appErr, ok := err.(*errcode.AppError); ok {
        appErr.RequestID = requestID
        c.JSON(appErr.HTTPStatus(), Response{
            Code:      appErr.Code,
            Message:   appErr.Message,
            RequestID: requestID,
        })
        return
    }
    
    // 未知错误
    c.JSON(500, Response{
        Code:      errcode.ErrInternalServer,
        Message:   "服务器内部错误",
        RequestID: requestID,
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
    })
}

// PageData 分页数据
type PageData struct {
    List     interface{} `json:"list"`
    Total    int64       `json:"total"`
    Page     int         `json:"page"`
    PageSize int         `json:"page_size"`
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
```

---

## 错误处理示例

### 1. Service 层错误处理

```go
// internal/app/advertiser/service/advertiser.go
package service

func (s *AdvertiserService) GetByID(ctx context.Context, id int64) (*dto.AdvertiserResp, error) {
    advertiser, err := s.repo.GetByID(ctx, id)
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, errcode.New(errcode.ErrAdvertiserNotFound)
        }
        return nil, errcode.Wrap(errcode.ErrInternalServer, err)
    }
    
    return dto.ToAdvertiserResp(advertiser), nil
}

func (s *AdvertiserService) Create(ctx context.Context, req *dto.AdvertiserCreateReq) (int64, error) {
    // 检查是否已存在
    exists, err := s.repo.ExistsByAdvertiserID(ctx, req.AdvertiserID)
    if err != nil {
        return 0, errcode.Wrap(errcode.ErrInternalServer, err)
    }
    if exists {
        return 0, errcode.New(errcode.ErrAdvertiserExists)
    }
    
    // 调用 Ocean Engine API 验证
    info, err := s.oceanSDK.Advertiser.GetInfo(ctx, req.AdvertiserID)
    if err != nil {
        if oceanengine.IsTokenError(err) {
            return 0, errcode.New(errcode.ErrOETokenInvalid)
        }
        return 0, errcode.Wrap(errcode.ErrOEAPIFailed, err)
    }
    
    // 创建广告主
    advertiser := &model.Advertiser{
        AdvertiserID: req.AdvertiserID,
        Name:         info.Name,
        // ...
    }
    
    if err := s.repo.Create(ctx, advertiser); err != nil {
        return 0, errcode.Wrap(errcode.ErrInternalServer, err)
    }
    
    return advertiser.ID, nil
}
```

### 2. API 层错误处理

```go
// internal/app/advertiser/api/advertiser.go
package api

func (a *AdvertiserAPI) GetByID(c *gin.Context) {
    // 参数解析
    id, err := strconv.ParseInt(c.Param("id"), 10, 64)
    if err != nil {
        response.Error(c, errcode.New(errcode.ErrInvalidParam))
        return
    }
    
    // 调用服务
    data, err := a.service.GetByID(c.Request.Context(), id)
    if err != nil {
        response.Error(c, err)
        return
    }
    
    response.Success(c, data)
}

func (a *AdvertiserAPI) Create(c *gin.Context) {
    var req dto.AdvertiserCreateReq
    if err := c.ShouldBindJSON(&req); err != nil {
        response.Error(c, errcode.NewWithMessage(errcode.ErrInvalidParam, err.Error()))
        return
    }
    
    // 参数验证
    if err := a.validator.Struct(&req); err != nil {
        response.ErrorWithDetails(c, errcode.New(errcode.ErrInvalidParam), parseValidationErrors(err))
        return
    }
    
    id, err := a.service.Create(c.Request.Context(), &req)
    if err != nil {
        response.Error(c, err)
        return
    }
    
    response.Success(c, gin.H{"id": id})
}
```

### 3. 参数验证错误处理

```go
// pkg/validator/validator.go
package validator

import (
    "github.com/go-playground/validator/v10"
)

// ValidationError 验证错误
type ValidationError struct {
    Field   string `json:"field"`
    Message string `json:"message"`
}

// ParseValidationErrors 解析验证错误
func ParseValidationErrors(err error) []ValidationError {
    var errors []ValidationError
    
    if validationErrors, ok := err.(validator.ValidationErrors); ok {
        for _, e := range validationErrors {
            errors = append(errors, ValidationError{
                Field:   e.Field(),
                Message: getValidationMessage(e),
            })
        }
    }
    
    return errors
}

func getValidationMessage(e validator.FieldError) string {
    switch e.Tag() {
    case "required":
        return "该字段不能为空"
    case "min":
        return fmt.Sprintf("最小长度为 %s", e.Param())
    case "max":
        return fmt.Sprintf("最大长度为 %s", e.Param())
    case "email":
        return "邮箱格式不正确"
    case "url":
        return "URL 格式不正确"
    default:
        return "格式不正确"
    }
}
```

---

## 错误日志记录

```go
// pkg/errcode/logger.go
package errcode

import (
    "go.uber.org/zap"
)

// LogError 记录错误日志
func LogError(logger *zap.Logger, err error, fields ...zap.Field) {
    if appErr, ok := err.(*AppError); ok {
        allFields := append(fields,
            zap.Int("code", appErr.Code),
            zap.String("message", appErr.Message),
        )
        
        if appErr.cause != nil {
            allFields = append(allFields, zap.Error(appErr.cause))
        }
        
        // 根据错误码选择日志级别
        switch {
        case appErr.Code >= 100000 && appErr.Code < 200000:
            // 认证错误，警告级别
            logger.Warn("auth error", allFields...)
        case appErr.Code >= 900000:
            // 外部 API 错误
            logger.Error("external api error", allFields...)
        case appErr.Code == ErrInternalServer:
            // 内部错误，错误级别
            logger.Error("internal error", allFields...)
        default:
            // 业务错误，信息级别
            logger.Info("business error", allFields...)
        }
    } else {
        logger.Error("unknown error", append(fields, zap.Error(err))...)
    }
}
```

---

## 错误恢复与重试

```go
// pkg/retry/retry.go
package retry

import (
    "context"
    "time"
)

// Config 重试配置
type Config struct {
    MaxRetries int
    Delay      time.Duration
    MaxDelay   time.Duration
    Multiplier float64
}

// DefaultConfig 默认配置
func DefaultConfig() *Config {
    return &Config{
        MaxRetries: 3,
        Delay:      100 * time.Millisecond,
        MaxDelay:   5 * time.Second,
        Multiplier: 2.0,
    }
}

// Retryable 可重试函数
type Retryable func(ctx context.Context) error

// ShouldRetry 是否应该重试
type ShouldRetry func(err error) bool

// Do 执行重试
func Do(ctx context.Context, fn Retryable, shouldRetry ShouldRetry, config *Config) error {
    if config == nil {
        config = DefaultConfig()
    }
    
    var lastErr error
    delay := config.Delay
    
    for i := 0; i <= config.MaxRetries; i++ {
        err := fn(ctx)
        if err == nil {
            return nil
        }
        
        lastErr = err
        
        // 检查是否应该重试
        if !shouldRetry(err) {
            return err
        }
        
        // 最后一次不等待
        if i == config.MaxRetries {
            break
        }
        
        // 等待重试
        select {
        case <-ctx.Done():
            return ctx.Err()
        case <-time.After(delay):
        }
        
        // 指数退避
        delay = time.Duration(float64(delay) * config.Multiplier)
        if delay > config.MaxDelay {
            delay = config.MaxDelay
        }
    }
    
    return lastErr
}

// DefaultShouldRetry 默认重试判断
func DefaultShouldRetry(err error) bool {
    if appErr, ok := err.(*errcode.AppError); ok {
        // 以下错误可重试
        switch appErr.Code {
        case errcode.ErrTimeout,
            errcode.ErrServiceUnavail,
            errcode.ErrOERateLimit:
            return true
        }
    }
    return false
}
```

### 使用示例

```go
func (s *AdvertiserService) SyncWithRetry(ctx context.Context, advertiserID int64) error {
    return retry.Do(ctx, func(ctx context.Context) error {
        return s.SyncAdvertiserInfo(ctx, advertiserID)
    }, retry.DefaultShouldRetry, &retry.Config{
        MaxRetries: 3,
        Delay:      time.Second,
        MaxDelay:   10 * time.Second,
        Multiplier: 2.0,
    })
}
```

---

## 错误响应示例

### 成功响应
```json
{
    "code": 0,
    "message": "成功",
    "data": {
        "id": 1,
        "name": "测试广告主"
    },
    "request_id": "550e8400-e29b-41d4-a716-446655440000"
}
```

### 参数错误
```json
{
    "code": 100001,
    "message": "参数错误",
    "data": [
        {"field": "name", "message": "该字段不能为空"},
        {"field": "email", "message": "邮箱格式不正确"}
    ],
    "request_id": "550e8400-e29b-41d4-a716-446655440000"
}
```

### 认证错误
```json
{
    "code": 100102,
    "message": "Token 已过期",
    "request_id": "550e8400-e29b-41d4-a716-446655440000"
}
```

### 业务错误
```json
{
    "code": 300001,
    "message": "广告主不存在",
    "request_id": "550e8400-e29b-41d4-a716-446655440000"
}
```

### 服务器错误
```json
{
    "code": 100005,
    "message": "服务器内部错误",
    "request_id": "550e8400-e29b-41d4-a716-446655440000"
}
```
