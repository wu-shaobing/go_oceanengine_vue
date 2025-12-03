package middleware

import (
	"net/http"
	"sync"

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

// NewIPRateLimiter 创建 IP 限流器
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

// DefaultRateLimitConfig 默认限流配置
func DefaultRateLimitConfig() *RateLimitConfig {
	return &RateLimitConfig{
		Rate:  100, // 每秒100个请求
		Burst: 200, // 最大突发200个
	}
}
