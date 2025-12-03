# 日志设计

## 概述

本系统使用 Zap 作为日志框架，实现结构化日志记录，支持多输出、日志轮转和级别控制。

## 日志配置

```go
// pkg/logger/logger.go
package logger

import (
    "os"
    "time"
    
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "gopkg.in/natefinch/lumberjack.v2"
)

// Config 日志配置
type Config struct {
    Level      string `yaml:"level"`       // 日志级别: debug, info, warn, error
    Format     string `yaml:"format"`      // 格式: json, console
    Output     string `yaml:"output"`      // 输出: stdout, file, both
    Filename   string `yaml:"filename"`    // 日志文件路径
    MaxSize    int    `yaml:"max_size"`    // 单文件最大大小(MB)
    MaxBackups int    `yaml:"max_backups"` // 保留旧文件数量
    MaxAge     int    `yaml:"max_age"`     // 保留天数
    Compress   bool   `yaml:"compress"`    // 是否压缩
}

// DefaultConfig 默认配置
func DefaultConfig() *Config {
    return &Config{
        Level:      "info",
        Format:     "json",
        Output:     "both",
        Filename:   "logs/app.log",
        MaxSize:    100,
        MaxBackups: 30,
        MaxAge:     30,
        Compress:   true,
    }
}

// NewLogger 创建日志器
func NewLogger(cfg *Config) (*zap.Logger, error) {
    // 解析日志级别
    level, err := zapcore.ParseLevel(cfg.Level)
    if err != nil {
        return nil, err
    }
    
    // 编码配置
    encoderConfig := zapcore.EncoderConfig{
        TimeKey:        "time",
        LevelKey:       "level",
        NameKey:        "logger",
        CallerKey:      "caller",
        FunctionKey:    zapcore.OmitKey,
        MessageKey:     "msg",
        StacktraceKey:  "stacktrace",
        LineEnding:     zapcore.DefaultLineEnding,
        EncodeLevel:    zapcore.LowercaseLevelEncoder,
        EncodeTime:     zapcore.ISO8601TimeEncoder,
        EncodeDuration: zapcore.MillisDurationEncoder,
        EncodeCaller:   zapcore.ShortCallerEncoder,
    }
    
    // 选择编码器
    var encoder zapcore.Encoder
    if cfg.Format == "json" {
        encoder = zapcore.NewJSONEncoder(encoderConfig)
    } else {
        encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
        encoder = zapcore.NewConsoleEncoder(encoderConfig)
    }
    
    // 构建输出
    var cores []zapcore.Core
    
    // 控制台输出
    if cfg.Output == "stdout" || cfg.Output == "both" {
        consoleCore := zapcore.NewCore(
            encoder,
            zapcore.AddSync(os.Stdout),
            level,
        )
        cores = append(cores, consoleCore)
    }
    
    // 文件输出
    if cfg.Output == "file" || cfg.Output == "both" {
        fileWriter := &lumberjack.Logger{
            Filename:   cfg.Filename,
            MaxSize:    cfg.MaxSize,
            MaxBackups: cfg.MaxBackups,
            MaxAge:     cfg.MaxAge,
            Compress:   cfg.Compress,
        }
        
        fileCore := zapcore.NewCore(
            zapcore.NewJSONEncoder(encoderConfig), // 文件始终使用 JSON
            zapcore.AddSync(fileWriter),
            level,
        )
        cores = append(cores, fileCore)
    }
    
    // 合并核心
    core := zapcore.NewTee(cores...)
    
    // 创建日志器
    logger := zap.New(core,
        zap.AddCaller(),
        zap.AddCallerSkip(1),
        zap.AddStacktrace(zapcore.ErrorLevel),
    )
    
    return logger, nil
}

// 全局日志器
var globalLogger *zap.Logger

// Init 初始化全局日志器
func Init(cfg *Config) error {
    logger, err := NewLogger(cfg)
    if err != nil {
        return err
    }
    globalLogger = logger
    return nil
}

// L 获取全局日志器
func L() *zap.Logger {
    if globalLogger == nil {
        globalLogger, _ = zap.NewDevelopment()
    }
    return globalLogger
}

// S 获取 SugaredLogger
func S() *zap.SugaredLogger {
    return L().Sugar()
}

// Sync 同步日志
func Sync() error {
    return L().Sync()
}
```

---

## 日志字段规范

### 标准字段

```go
// pkg/logger/fields.go
package logger

import "go.uber.org/zap"

// 常用字段名
const (
    FieldRequestID   = "request_id"
    FieldUserID      = "user_id"
    FieldUsername    = "username"
    FieldPath        = "path"
    FieldMethod      = "method"
    FieldStatus      = "status"
    FieldLatency     = "latency"
    FieldIP          = "ip"
    FieldError       = "error"
    FieldModule      = "module"
    FieldAction      = "action"
    FieldAdvertiserID = "advertiser_id"
    FieldCampaignID   = "campaign_id"
)

// 常用字段构建器
func RequestID(id string) zap.Field {
    return zap.String(FieldRequestID, id)
}

func UserID(id int64) zap.Field {
    return zap.Int64(FieldUserID, id)
}

func Username(name string) zap.Field {
    return zap.String(FieldUsername, name)
}

func Module(name string) zap.Field {
    return zap.String(FieldModule, name)
}

func Action(name string) zap.Field {
    return zap.String(FieldAction, name)
}

func AdvertiserID(id int64) zap.Field {
    return zap.Int64(FieldAdvertiserID, id)
}

func CampaignID(id int64) zap.Field {
    return zap.Int64(FieldCampaignID, id)
}
```

---

## 上下文日志

```go
// pkg/logger/context.go
package logger

import (
    "context"
    
    "go.uber.org/zap"
)

type ctxKey struct{}

// WithLogger 设置日志器到上下文
func WithLogger(ctx context.Context, logger *zap.Logger) context.Context {
    return context.WithValue(ctx, ctxKey{}, logger)
}

// FromContext 从上下文获取日志器
func FromContext(ctx context.Context) *zap.Logger {
    if logger, ok := ctx.Value(ctxKey{}).(*zap.Logger); ok {
        return logger
    }
    return L()
}

// WithFields 创建带字段的日志器
func WithFields(ctx context.Context, fields ...zap.Field) context.Context {
    logger := FromContext(ctx).With(fields...)
    return WithLogger(ctx, logger)
}

// Debug 记录 Debug 日志
func Debug(ctx context.Context, msg string, fields ...zap.Field) {
    FromContext(ctx).Debug(msg, fields...)
}

// Info 记录 Info 日志
func Info(ctx context.Context, msg string, fields ...zap.Field) {
    FromContext(ctx).Info(msg, fields...)
}

// Warn 记录 Warn 日志
func Warn(ctx context.Context, msg string, fields ...zap.Field) {
    FromContext(ctx).Warn(msg, fields...)
}

// Error 记录 Error 日志
func Error(ctx context.Context, msg string, fields ...zap.Field) {
    FromContext(ctx).Error(msg, fields...)
}
```

---

## 日志分类

### 1. 访问日志

```go
// pkg/logger/access.go
package logger

import (
    "time"
    
    "go.uber.org/zap"
)

// AccessLog 访问日志
type AccessLog struct {
    logger *zap.Logger
}

func NewAccessLog() *AccessLog {
    return &AccessLog{
        logger: L().Named("access"),
    }
}

// Log 记录访问日志
func (l *AccessLog) Log(
    requestID string,
    method string,
    path string,
    query string,
    status int,
    latency time.Duration,
    clientIP string,
    userAgent string,
    userID int64,
) {
    l.logger.Info("http request",
        zap.String("request_id", requestID),
        zap.String("method", method),
        zap.String("path", path),
        zap.String("query", query),
        zap.Int("status", status),
        zap.Duration("latency", latency),
        zap.String("client_ip", clientIP),
        zap.String("user_agent", userAgent),
        zap.Int64("user_id", userID),
    )
}
```

### 2. 业务日志

```go
// pkg/logger/business.go
package logger

import "go.uber.org/zap"

// BusinessLog 业务日志
type BusinessLog struct {
    logger *zap.Logger
}

func NewBusinessLog(module string) *BusinessLog {
    return &BusinessLog{
        logger: L().Named("business").With(zap.String("module", module)),
    }
}

// Info 记录业务信息
func (l *BusinessLog) Info(action string, fields ...zap.Field) {
    l.logger.Info(action, fields...)
}

// Warn 记录业务警告
func (l *BusinessLog) Warn(action string, fields ...zap.Field) {
    l.logger.Warn(action, fields...)
}

// Error 记录业务错误
func (l *BusinessLog) Error(action string, err error, fields ...zap.Field) {
    fields = append(fields, zap.Error(err))
    l.logger.Error(action, fields...)
}

// 使用示例
var advertiserLog = NewBusinessLog("advertiser")

func (s *AdvertiserService) Create(ctx context.Context, req *dto.CreateReq) error {
    advertiserLog.Info("create_advertiser",
        zap.Int64("advertiser_id", req.AdvertiserID),
        zap.String("name", req.Name),
    )
    // ...
}
```

### 3. 审计日志

```go
// pkg/logger/audit.go
package logger

import (
    "time"
    
    "go.uber.org/zap"
)

// AuditLog 审计日志
type AuditLog struct {
    logger *zap.Logger
}

func NewAuditLog() *AuditLog {
    return &AuditLog{
        logger: L().Named("audit"),
    }
}

// Log 记录审计日志
func (l *AuditLog) Log(
    userID int64,
    username string,
    action string,
    resource string,
    resourceID string,
    detail string,
    clientIP string,
) {
    l.logger.Info("audit",
        zap.Int64("user_id", userID),
        zap.String("username", username),
        zap.String("action", action),
        zap.String("resource", resource),
        zap.String("resource_id", resourceID),
        zap.String("detail", detail),
        zap.String("client_ip", clientIP),
        zap.Time("timestamp", time.Now()),
    )
}
```

### 4. SDK 调用日志

```go
// pkg/logger/sdk.go
package logger

import (
    "time"
    
    "go.uber.org/zap"
)

// SDKLog SDK 调用日志
type SDKLog struct {
    logger *zap.Logger
}

func NewSDKLog() *SDKLog {
    return &SDKLog{
        logger: L().Named("sdk"),
    }
}

// LogRequest 记录请求
func (l *SDKLog) LogRequest(
    requestID string,
    api string,
    advertiserID int64,
    request interface{},
) {
    l.logger.Debug("sdk request",
        zap.String("request_id", requestID),
        zap.String("api", api),
        zap.Int64("advertiser_id", advertiserID),
        zap.Any("request", request),
    )
}

// LogResponse 记录响应
func (l *SDKLog) LogResponse(
    requestID string,
    api string,
    latency time.Duration,
    code int,
    response interface{},
    err error,
) {
    fields := []zap.Field{
        zap.String("request_id", requestID),
        zap.String("api", api),
        zap.Duration("latency", latency),
        zap.Int("code", code),
    }
    
    if err != nil {
        fields = append(fields, zap.Error(err))
        l.logger.Error("sdk response", fields...)
    } else {
        fields = append(fields, zap.Any("response", response))
        l.logger.Debug("sdk response", fields...)
    }
}
```

---

## 日志采样

```go
// pkg/logger/sampler.go
package logger

import (
    "time"
    
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

// NewSampledLogger 创建采样日志器
func NewSampledLogger(base *zap.Logger, tick time.Duration, first, thereafter int) *zap.Logger {
    sampler := zap.WrapCore(func(core zapcore.Core) zapcore.Core {
        return zapcore.NewSamplerWithOptions(
            core,
            tick,       // 采样间隔
            first,      // 首次记录数
            thereafter, // 之后每 tick 记录数
        )
    })
    return base.WithOptions(sampler)
}

// 使用示例
func main() {
    logger := NewLogger(DefaultConfig())
    
    // 每秒最多记录 100 条日志，之后每秒记录 10 条
    sampledLogger := NewSampledLogger(logger, time.Second, 100, 10)
}
```

---

## 敏感信息脱敏

```go
// pkg/logger/mask.go
package logger

import (
    "regexp"
    "strings"
)

var (
    phoneRegex    = regexp.MustCompile(`1[3-9]\d{9}`)
    emailRegex    = regexp.MustCompile(`[\w.-]+@[\w.-]+\.\w+`)
    idCardRegex   = regexp.MustCompile(`\d{17}[\dXx]`)
    tokenRegex    = regexp.MustCompile(`(access_token|refresh_token|password)["']?\s*[=:]\s*["']?[\w.-]+`)
)

// MaskPhone 脱敏手机号
func MaskPhone(s string) string {
    return phoneRegex.ReplaceAllStringFunc(s, func(phone string) string {
        return phone[:3] + "****" + phone[7:]
    })
}

// MaskEmail 脱敏邮箱
func MaskEmail(s string) string {
    return emailRegex.ReplaceAllStringFunc(s, func(email string) string {
        parts := strings.Split(email, "@")
        if len(parts) != 2 {
            return email
        }
        name := parts[0]
        if len(name) > 2 {
            name = name[:2] + "***"
        }
        return name + "@" + parts[1]
    })
}

// MaskIDCard 脱敏身份证
func MaskIDCard(s string) string {
    return idCardRegex.ReplaceAllStringFunc(s, func(id string) string {
        return id[:6] + "********" + id[14:]
    })
}

// MaskToken 脱敏 Token
func MaskToken(s string) string {
    return tokenRegex.ReplaceAllString(s, "$1=***")
}

// MaskSensitive 综合脱敏
func MaskSensitive(s string) string {
    s = MaskPhone(s)
    s = MaskEmail(s)
    s = MaskIDCard(s)
    s = MaskToken(s)
    return s
}
```

---

## 日志输出示例

### JSON 格式

```json
{
    "time": "2024-01-15T10:30:45.123Z",
    "level": "info",
    "logger": "access",
    "caller": "middleware/logger.go:42",
    "msg": "http request",
    "request_id": "550e8400-e29b-41d4-a716-446655440000",
    "method": "POST",
    "path": "/api/v1/campaign",
    "status": 200,
    "latency": 125,
    "client_ip": "192.168.1.100",
    "user_id": 1001
}
```

### 业务日志示例

```json
{
    "time": "2024-01-15T10:30:45.123Z",
    "level": "info",
    "logger": "business",
    "module": "campaign",
    "msg": "create_campaign",
    "advertiser_id": 12345,
    "campaign_id": 67890,
    "campaign_name": "测试活动"
}
```

### 错误日志示例

```json
{
    "time": "2024-01-15T10:30:45.123Z",
    "level": "error",
    "logger": "sdk",
    "msg": "sdk response",
    "request_id": "550e8400-e29b-41d4-a716-446655440000",
    "api": "/2/campaign/create/",
    "latency": 1500,
    "code": 40001,
    "error": "token invalid",
    "stacktrace": "..."
}
```
