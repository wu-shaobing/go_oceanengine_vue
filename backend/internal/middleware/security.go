package middleware

import "github.com/gin-gonic/gin"

// SecurityHeaders 添加通用安全响应头（生产环境建议启用 HSTS 并自定义 CSP）
func SecurityHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-Frame-Options", "DENY")
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-XSS-Protection", "1; mode=block")
		// 如启用 HTTPS，可在生产开启 HSTS：
		// c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		c.Header("Referrer-Policy", "strict-origin-when-cross-origin")
		c.Header("Permissions-Policy", "geolocation=(), microphone=(), camera=()")
		// 最小 CSP；可按需放宽
		c.Header("Content-Security-Policy", "default-src 'self'")
		c.Next()
	}
}