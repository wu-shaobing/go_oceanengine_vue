# 中间件设计

## 概述

中间件是在请求处理流程中的拦截器，用于实现通用功能如认证、日志、限流等。本文档描述系统中各类中间件的设计与实现。

## 中间件架构

```
请求 → Recovery → Logger → RequestID → RateLimit → JWT → Casbin → Handler → 响应
```

## 中间件列表

### 1. 恢复中间件 (Recovery)

捕获 panic，防止服务崩溃。

```go
// internal/middleware/recovery.go
package middleware

import (
    "net/http"
    "runtime/debug"
    
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

// Recovery 恢复中间件
func Recovery(logger *zap.Logger) gin.HandlerFunc {
    return func(c *gin.Context) {
        defer func() {
            if err := recover(); err != nil {
                // 记录堆栈
                stack := string(debug.Stack())
                
                logger.Error("panic recovered",
                    zap.Any("error", err),
                    zap.String("stack", stack),
                    zap.String("path", c.Request.URL.Path),
                    zap.String("method", c.Request.Method),
                    zap.String("request_id", c.GetString("request_id")),
                )
                
                c.JSON(http.StatusInternalServerError, gin.H{
                    "code":    500,
                    "message": "服务器内部错误",
                })
                c.Abort()
            }
        }()
        c.Next()
    }
}
```

### 2. 日志中间件 (Logger)

记录请求和响应日志。

```go
// internal/middleware/logger.go
package middleware

import (
    "bytes"
    "io"
    "time"
    
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

// responseWriter 包装响应写入器
type responseWriter struct {
    gin.ResponseWriter
    body *bytes.Buffer
}

func (w *responseWriter) Write(b []byte) (int, error) {
    w.body.Write(b)
    return w.ResponseWriter.Write(b)
}

// Logger 日志中间件
func Logger(logger *zap.Logger) gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()
        path := c.Request.URL.Path
        query := c.Request.URL.RawQuery
        
        // 读取请求体
        var requestBody []byte
        if c.Request.Body != nil {
            requestBody, _ = io.ReadAll(c.Request.Body)
            c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
        }
        
        // 包装响应写入器
        blw := &responseWriter{body: bytes.NewBuffer(nil), ResponseWriter: c.Writer}
        c.Writer = blw
        
        c.Next()
        
        // 计算耗时
        latency := time.Since(start)
        
        // 构建日志字段
        fields := []zap.Field{
            zap.String("request_id", c.GetString("request_id")),
            zap.Int("status", c.Writer.Status()),
            zap.String("method", c.Request.Method),
            zap.String("path", path),
            zap.String("query", query),
            zap.String("ip", c.ClientIP()),
            zap.String("user_agent", c.Request.UserAgent()),
            zap.Duration("latency", latency),
        }
        
        // 仅在非生产环境记录请求体
        if gin.Mode() != gin.ReleaseMode && len(requestBody) > 0 && len(requestBody) < 4096 {
            fields = append(fields, zap.ByteString("request_body", requestBody))
        }
        
        // 错误响应记录响应体
        if c.Writer.Status() >= 400 {
            fields = append(fields, zap.String("response_body", blw.body.String()))
        }
        
        // 根据状态码选择日志级别
        switch {
        case c.Writer.Status() >= 500:
            logger.Error("server error", fields...)
        case c.Writer.Status() >= 400:
            logger.Warn("client error", fields...)
        default:
            logger.Info("request", fields...)
        }
    }
}
```

### 3. 请求ID中间件 (RequestID)

为每个请求生成唯一ID，便于追踪。

```go
// internal/middleware/request_id.go
package middleware

import (
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
)

const (
    RequestIDHeader = "X-Request-ID"
    RequestIDKey    = "request_id"
)

// RequestID 请求ID中间件
func RequestID() gin.HandlerFunc {
    return func(c *gin.Context) {
        // 优先使用客户端传递的 RequestID
        requestID := c.GetHeader(RequestIDHeader)
        if requestID == "" {
            requestID = uuid.New().String()
        }
        
        // 设置到上下文和响应头
        c.Set(RequestIDKey, requestID)
        c.Header(RequestIDHeader, requestID)
        
        c.Next()
    }
}

// GetRequestID 获取请求ID
func GetRequestID(c *gin.Context) string {
    return c.GetString(RequestIDKey)
}
```

### 4. 限流中间件 (RateLimit)

基于令牌桶算法的限流。

```go
// internal/middleware/ratelimit.go
package middleware

import (
    "net/http"
    "sync"
    "time"
    
    "github.com/gin-gonic/gin"
    "golang.org/x/time/rate"
)

// RateLimitConfig 限流配置
type RateLimitConfig struct {
    Rate  rate.Limit // 每秒请求数
    Burst int        // 突发请求数
}

// IPRateLimiter IP限流器
type IPRateLimiter struct {
    limiters map[string]*rate.Limiter
    mu       sync.RWMutex
    config   *RateLimitConfig
}

func NewIPRateLimiter(config *RateLimitConfig) *IPRateLimiter {
    return &IPRateLimiter{
        limiters: make(map[string]*rate.Limiter),
        config:   config,
    }
}

func (l *IPRateLimiter) getLimiter(ip string) *rate.Limiter {
    l.mu.RLock()
    limiter, exists := l.limiters[ip]
    l.mu.RUnlock()
    
    if exists {
        return limiter
    }
    
    l.mu.Lock()
    defer l.mu.Unlock()
    
    // 双重检查
    if limiter, exists = l.limiters[ip]; exists {
        return limiter
    }
    
    limiter = rate.NewLimiter(l.config.Rate, l.config.Burst)
    l.limiters[ip] = limiter
    return limiter
}

// RateLimit 限流中间件
func RateLimit(limiter *IPRateLimiter) gin.HandlerFunc {
    return func(c *gin.Context) {
        ip := c.ClientIP()
        l := limiter.getLimiter(ip)
        
        if !l.Allow() {
            c.JSON(http.StatusTooManyRequests, gin.H{
                "code":    429,
                "message": "请求过于频繁，请稍后再试",
            })
            c.Abort()
            return
        }
        
        c.Next()
    }
}

// Redis 分布式限流
type RedisRateLimiter struct {
    redis  *redis.Client
    rate   int           // 请求数
    window time.Duration // 时间窗口
}

func NewRedisRateLimiter(redis *redis.Client, rate int, window time.Duration) *RedisRateLimiter {
    return &RedisRateLimiter{
        redis:  redis,
        rate:   rate,
        window: window,
    }
}

// Allow 检查是否允许请求
func (l *RedisRateLimiter) Allow(ctx context.Context, key string) (bool, error) {
    // 使用滑动窗口计数
    now := time.Now().UnixNano()
    windowStart := now - l.window.Nanoseconds()
    
    pipe := l.redis.Pipeline()
    
    // 移除窗口外的记录
    pipe.ZRemRangeByScore(ctx, key, "0", fmt.Sprintf("%d", windowStart))
    // 添加当前请求
    pipe.ZAdd(ctx, key, &redis.Z{Score: float64(now), Member: now})
    // 计数
    countCmd := pipe.ZCard(ctx, key)
    // 设置过期时间
    pipe.Expire(ctx, key, l.window)
    
    _, err := pipe.Exec(ctx)
    if err != nil {
        return false, err
    }
    
    return countCmd.Val() <= int64(l.rate), nil
}

// RedisRateLimit Redis 限流中间件
func RedisRateLimit(limiter *RedisRateLimiter) gin.HandlerFunc {
    return func(c *gin.Context) {
        key := fmt.Sprintf("ratelimit:%s", c.ClientIP())
        
        allowed, err := limiter.Allow(c.Request.Context(), key)
        if err != nil {
            c.Next()
            return
        }
        
        if !allowed {
            c.JSON(http.StatusTooManyRequests, gin.H{
                "code":    429,
                "message": "请求过于频繁，请稍后再试",
            })
            c.Abort()
            return
        }
        
        c.Next()
    }
}
```

### 5. CORS 中间件

跨域资源共享配置。

```go
// internal/middleware/cors.go
package middleware

import (
    "github.com/gin-gonic/gin"
)

// CORSConfig CORS 配置
type CORSConfig struct {
    AllowOrigins     []string
    AllowMethods     []string
    AllowHeaders     []string
    ExposeHeaders    []string
    AllowCredentials bool
    MaxAge           int
}

// DefaultCORSConfig 默认配置
func DefaultCORSConfig() *CORSConfig {
    return &CORSConfig{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Request-ID"},
        ExposeHeaders:    []string{"Content-Length", "X-Request-ID"},
        AllowCredentials: true,
        MaxAge:           86400,
    }
}

// CORS 跨域中间件
func CORS(config *CORSConfig) gin.HandlerFunc {
    return func(c *gin.Context) {
        origin := c.Request.Header.Get("Origin")
        
        // 检查是否允许的源
        allowed := false
        for _, o := range config.AllowOrigins {
            if o == "*" || o == origin {
                allowed = true
                break
            }
        }
        
        if allowed {
            c.Header("Access-Control-Allow-Origin", origin)
            c.Header("Access-Control-Allow-Methods", strings.Join(config.AllowMethods, ", "))
            c.Header("Access-Control-Allow-Headers", strings.Join(config.AllowHeaders, ", "))
            c.Header("Access-Control-Expose-Headers", strings.Join(config.ExposeHeaders, ", "))
            c.Header("Access-Control-Max-Age", strconv.Itoa(config.MaxAge))
            
            if config.AllowCredentials {
                c.Header("Access-Control-Allow-Credentials", "true")
            }
        }
        
        // 处理预检请求
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(http.StatusNoContent)
            return
        }
        
        c.Next()
    }
}
```

### 6. 超时中间件 (Timeout)

请求超时控制。

```go
// internal/middleware/timeout.go
package middleware

import (
    "context"
    "net/http"
    "time"
    
    "github.com/gin-gonic/gin"
)

// Timeout 超时中间件
func Timeout(timeout time.Duration) gin.HandlerFunc {
    return func(c *gin.Context) {
        // 创建带超时的上下文
        ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
        defer cancel()
        
        // 替换请求上下文
        c.Request = c.Request.WithContext(ctx)
        
        // 创建结果通道
        done := make(chan struct{})
        
        go func() {
            c.Next()
            close(done)
        }()
        
        select {
        case <-done:
            // 正常完成
        case <-ctx.Done():
            c.JSON(http.StatusGatewayTimeout, gin.H{
                "code":    504,
                "message": "请求超时",
            })
            c.Abort()
        }
    }
}
```

### 7. 操作日志中间件

记录用户操作日志。

```go
// internal/middleware/operation_log.go
package middleware

import (
    "bytes"
    "encoding/json"
    "io"
    "strings"
    "time"
    
    "github.com/gin-gonic/gin"
)

// OperationLogConfig 操作日志配置
type OperationLogConfig struct {
    SkipPaths     []string // 跳过的路径
    MaxBodySize   int      // 最大记录请求体大小
    EnableRequest bool     // 是否记录请求体
}

// OperationLog 操作日志中间件
func OperationLog(logService *service.OperationLogService, config *OperationLogConfig) gin.HandlerFunc {
    return func(c *gin.Context) {
        // 检查是否跳过
        path := c.Request.URL.Path
        for _, skip := range config.SkipPaths {
            if strings.HasPrefix(path, skip) {
                c.Next()
                return
            }
        }
        
        // 仅记录写操作
        if c.Request.Method == "GET" {
            c.Next()
            return
        }
        
        start := time.Now()
        
        // 读取请求体
        var requestBody string
        if config.EnableRequest && c.Request.Body != nil {
            body, _ := io.ReadAll(c.Request.Body)
            c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
            if len(body) <= config.MaxBodySize {
                requestBody = string(body)
            }
        }
        
        // 包装响应
        blw := &responseWriter{body: bytes.NewBuffer(nil), ResponseWriter: c.Writer}
        c.Writer = blw
        
        c.Next()
        
        // 异步记录日志
        go func() {
            userID, _ := c.Get("user_id")
            username, _ := c.Get("username")
            
            log := &model.OperationLog{
                UserID:       userID.(int64),
                Username:     username.(string),
                Method:       c.Request.Method,
                Path:         path,
                Query:        c.Request.URL.RawQuery,
                RequestBody:  requestBody,
                ResponseBody: blw.body.String(),
                StatusCode:   c.Writer.Status(),
                ClientIP:     c.ClientIP(),
                UserAgent:    c.Request.UserAgent(),
                Latency:      time.Since(start).Milliseconds(),
                CreatedAt:    time.Now(),
            }
            
            logService.Create(context.Background(), log)
        }()
    }
}
```

### 8. 压缩中间件 (Gzip)

响应压缩。

```go
// internal/middleware/gzip.go
package middleware

import (
    "compress/gzip"
    "strings"
    
    "github.com/gin-gonic/gin"
)

type gzipWriter struct {
    gin.ResponseWriter
    writer *gzip.Writer
}

func (g *gzipWriter) Write(data []byte) (int, error) {
    return g.writer.Write(data)
}

// Gzip 压缩中间件
func Gzip(level int) gin.HandlerFunc {
    return func(c *gin.Context) {
        // 检查客户端是否支持 gzip
        if !strings.Contains(c.Request.Header.Get("Accept-Encoding"), "gzip") {
            c.Next()
            return
        }
        
        // 创建 gzip writer
        gz, err := gzip.NewWriterLevel(c.Writer, level)
        if err != nil {
            c.Next()
            return
        }
        defer gz.Close()
        
        c.Header("Content-Encoding", "gzip")
        c.Header("Vary", "Accept-Encoding")
        
        c.Writer = &gzipWriter{ResponseWriter: c.Writer, writer: gz}
        c.Next()
    }
}
```

### 9. 幂等性中间件

防止重复提交。

```go
// internal/middleware/idempotency.go
package middleware

import (
    "net/http"
    "time"
    
    "github.com/gin-gonic/gin"
    "github.com/redis/go-redis/v9"
)

const (
    IdempotencyKeyHeader = "X-Idempotency-Key"
    IdempotencyTTL       = 24 * time.Hour
)

// Idempotency 幂等性中间件
func Idempotency(redis *redis.Client) gin.HandlerFunc {
    return func(c *gin.Context) {
        // 仅对写操作检查
        if c.Request.Method == "GET" || c.Request.Method == "OPTIONS" {
            c.Next()
            return
        }
        
        // 获取幂等键
        key := c.GetHeader(IdempotencyKeyHeader)
        if key == "" {
            c.Next()
            return
        }
        
        ctx := c.Request.Context()
        redisKey := fmt.Sprintf("idempotency:%s", key)
        
        // 尝试设置键
        set, err := redis.SetNX(ctx, redisKey, "processing", IdempotencyTTL).Result()
        if err != nil {
            c.Next()
            return
        }
        
        if !set {
            // 键已存在，检查状态
            val, err := redis.Get(ctx, redisKey).Result()
            if err == nil && val != "processing" {
                // 返回缓存的响应
                c.JSON(http.StatusOK, json.RawMessage(val))
                c.Abort()
                return
            }
            
            // 正在处理中
            c.JSON(http.StatusConflict, gin.H{
                "code":    409,
                "message": "请求正在处理中，请勿重复提交",
            })
            c.Abort()
            return
        }
        
        // 包装响应写入器以捕获响应
        blw := &responseWriter{body: bytes.NewBuffer(nil), ResponseWriter: c.Writer}
        c.Writer = blw
        
        c.Next()
        
        // 缓存成功响应
        if c.Writer.Status() == http.StatusOK {
            redis.Set(ctx, redisKey, blw.body.String(), IdempotencyTTL)
        } else {
            // 失败则删除键，允许重试
            redis.Del(ctx, redisKey)
        }
    }
}
```

---

## 中间件注册

```go
// internal/router/router.go
package router

import (
    "github.com/gin-gonic/gin"
)

func SetupRouter(
    logger *zap.Logger,
    jwtManager *auth.JWTManager,
    tokenStore *auth.TokenStore,
    enforcer *auth.CasbinEnforcer,
    redis *redis.Client,
) *gin.Engine {
    r := gin.New()
    
    // 全局中间件
    r.Use(middleware.Recovery(logger))
    r.Use(middleware.RequestID())
    r.Use(middleware.Logger(logger))
    r.Use(middleware.CORS(middleware.DefaultCORSConfig()))
    
    // 限流中间件
    rateLimiter := middleware.NewIPRateLimiter(&middleware.RateLimitConfig{
        Rate:  100,
        Burst: 200,
    })
    r.Use(middleware.RateLimit(rateLimiter))
    
    // 公开路由
    public := r.Group("/api/v1")
    {
        public.POST("/auth/login", authAPI.Login)
        public.POST("/auth/captcha", authAPI.Captcha)
    }
    
    // 需要认证的路由
    protected := r.Group("/api/v1")
    protected.Use(middleware.JWTAuth(jwtManager, tokenStore))
    protected.Use(middleware.CasbinAuth(enforcer))
    protected.Use(middleware.OperationLog(logService, &middleware.OperationLogConfig{
        SkipPaths:     []string{"/api/v1/report"},
        MaxBodySize:   4096,
        EnableRequest: true,
    }))
    {
        // 用户管理
        protected.GET("/user", userAPI.List)
        protected.POST("/user", userAPI.Create)
        // ... 其他路由
    }
    
    return r
}
```

---

## 中间件执行顺序图

```
┌─────────────────────────────────────────────────────────────────────────┐
│                              请求进入                                    │
└─────────────────────────────────────────────────────────────────────────┘
                                    │
                                    ▼
┌─────────────────────────────────────────────────────────────────────────┐
│                         Recovery (异常恢复)                              │
│                     捕获 panic，记录日志，返回500                          │
└─────────────────────────────────────────────────────────────────────────┘
                                    │
                                    ▼
┌─────────────────────────────────────────────────────────────────────────┐
│                         RequestID (请求追踪)                             │
│                      生成/获取请求唯一标识                                │
└─────────────────────────────────────────────────────────────────────────┘
                                    │
                                    ▼
┌─────────────────────────────────────────────────────────────────────────┐
│                           Logger (日志记录)                              │
│                       记录请求/响应详细信息                               │
└─────────────────────────────────────────────────────────────────────────┘
                                    │
                                    ▼
┌─────────────────────────────────────────────────────────────────────────┐
│                             CORS (跨域)                                  │
│                        处理跨域请求和预检                                 │
└─────────────────────────────────────────────────────────────────────────┘
                                    │
                                    ▼
┌─────────────────────────────────────────────────────────────────────────┐
│                          RateLimit (限流)                                │
│                         检查请求频率限制                                  │
└─────────────────────────────────────────────────────────────────────────┘
                                    │
                                    ▼
┌─────────────────────────────────────────────────────────────────────────┐
│                           JWTAuth (认证)                                 │
│                      解析验证 Token，提取用户信息                          │
└─────────────────────────────────────────────────────────────────────────┘
                                    │
                                    ▼
┌─────────────────────────────────────────────────────────────────────────┐
│                          CasbinAuth (授权)                               │
│                       检查用户是否有访问权限                              │
└─────────────────────────────────────────────────────────────────────────┘
                                    │
                                    ▼
┌─────────────────────────────────────────────────────────────────────────┐
│                       OperationLog (操作日志)                            │
│                        记录用户操作行为                                   │
└─────────────────────────────────────────────────────────────────────────┘
                                    │
                                    ▼
┌─────────────────────────────────────────────────────────────────────────┐
│                            Handler (处理器)                              │
│                          业务逻辑处理                                    │
└─────────────────────────────────────────────────────────────────────────┘
                                    │
                                    ▼
┌─────────────────────────────────────────────────────────────────────────┐
│                              响应返回                                    │
└─────────────────────────────────────────────────────────────────────────┘
```
