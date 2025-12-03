package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"oceanengine-backend/pkg/auth"
)

// JWTAuth JWT 认证中间件
func JWTAuth(jwtManager *auth.JWTManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 获取 Token
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    100100,
				"message": "请先登录",
			})
			c.Abort()
			return
		}

		// 2. 去除 Bearer 前缀
		token = strings.TrimPrefix(token, "Bearer ")
		token = strings.TrimSpace(token)

		// 3. 解析 Token
		claims, err := jwtManager.ParseToken(token)
		if err != nil {
			msg := "Token 无效"
			code := 100101
			if err == auth.ErrTokenExpired {
				msg = "Token 已过期"
				code = 100102
			}
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    code,
				"message": msg,
			})
			c.Abort()
			return
		}

		// 4. 设置用户信息到上下文
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("role_key", claims.RoleKey)
		c.Set("role_id", claims.RoleID)
		c.Set("data_scope", claims.DataScope)
		c.Set("claims", claims)

		c.Next()
	}
}

// GetUserID 获取当前用户ID
func GetUserID(c *gin.Context) int64 {
	userID, exists := c.Get("user_id")
	if !exists {
		return 0
	}
	return userID.(int64)
}

// GetUsername 获取当前用户名
func GetUsername(c *gin.Context) string {
	username, exists := c.Get("username")
	if !exists {
		return ""
	}
	return username.(string)
}

// GetRoleKey 获取当前角色标识
func GetRoleKey(c *gin.Context) string {
	roleKey, exists := c.Get("role_key")
	if !exists {
		return ""
	}
	return roleKey.(string)
}

// GetRoleID 获取当前角色ID
func GetRoleID(c *gin.Context) int64 {
	roleID, exists := c.Get("role_id")
	if !exists {
		return 0
	}
	return roleID.(int64)
}

// GetClaims 获取完整的Claims
func GetClaims(c *gin.Context) *auth.Claims {
	claims, exists := c.Get("claims")
	if !exists {
		return nil
	}
	return claims.(*auth.Claims)
}
