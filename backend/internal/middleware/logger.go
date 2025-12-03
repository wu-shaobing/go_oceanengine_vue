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

		// 添加用户ID（如果存在）
		if userID, exists := c.Get("user_id"); exists {
			fields = append(fields, zap.Any("user_id", userID))
		}

		// 仅在非生产环境记录请求体
		if gin.Mode() != gin.ReleaseMode && len(requestBody) > 0 && len(requestBody) < 4096 {
			fields = append(fields, zap.ByteString("request_body", requestBody))
		}

		// 错误响应记录响应体
		if c.Writer.Status() >= 400 && blw.body.Len() < 4096 {
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
