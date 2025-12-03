package middleware

import (
	"net/http"
	"os"
	"strconv"
	"strings"

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

// DefaultCORSConfig 默认配置（可通过环境变量覆盖）
// CORS_ALLOWED_ORIGINS= https://example.com,https://admin.example.com
func DefaultCORSConfig() *CORSConfig {
	cfg := &CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Request-ID"},
		ExposeHeaders:    []string{"Content-Length", "X-Request-ID"},
		AllowCredentials: true,
		MaxAge:           86400,
	}

	if origins := strings.TrimSpace(os.Getenv("CORS_ALLOWED_ORIGINS")); origins != "" {
		parts := strings.Split(origins, ",")
		var cleaned []string
		for _, p := range parts {
			if o := strings.TrimSpace(p); o != "" {
				cleaned = append(cleaned, o)
			}
		}
		if len(cleaned) > 0 {
			cfg.AllowOrigins = cleaned
		}
	}
	return cfg
}

// CORS 跨域中间件
func CORS(config *CORSConfig) gin.HandlerFunc {
	if config == nil {
		config = DefaultCORSConfig()
	}

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

		if allowed && origin != "" {
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
