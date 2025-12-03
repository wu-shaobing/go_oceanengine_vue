package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"oceanengine-backend/config"
)

var (
	ErrTokenExpired     = errors.New("token has expired")
	ErrTokenMalformed   = errors.New("token is malformed")
	ErrTokenInvalid     = errors.New("token is invalid")
	ErrTokenNotValidYet = errors.New("token is not active yet")
)

// Claims 自定义声明
type Claims struct {
	UserID    int64  `json:"user_id"`
	Username  string `json:"username"`
	RoleKey   string `json:"role_key"`
	RoleID    int64  `json:"role_id"`
	DataScope string `json:"data_scope"`
	jwt.RegisteredClaims
}

// JWTManager JWT 管理器
type JWTManager struct {
	config *config.JWTConfig
}

// NewJWTManager 创建 JWT 管理器
func NewJWTManager(cfg *config.JWTConfig) *JWTManager {
	return &JWTManager{config: cfg}
}

// GenerateToken 生成访问 Token
func (j *JWTManager) GenerateToken(claims *Claims) (string, error) {
	claims.Issuer = j.config.Issuer
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(j.config.AccessExpire))
	claims.IssuedAt = jwt.NewNumericDate(time.Now())
	claims.NotBefore = jwt.NewNumericDate(time.Now())

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.config.SecretKey))
}

// GenerateRefreshToken 生成刷新 Token
func (j *JWTManager) GenerateRefreshToken(userID int64) (string, error) {
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.config.Issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.config.RefreshExpire)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.config.SecretKey))
}

// ParseToken 解析 Token
func (j *JWTManager) ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.config.SecretKey), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrTokenExpired
		}
		if errors.Is(err, jwt.ErrTokenMalformed) {
			return nil, ErrTokenMalformed
		}
		if errors.Is(err, jwt.ErrTokenNotValidYet) {
			return nil, ErrTokenNotValidYet
		}
		return nil, ErrTokenInvalid
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrTokenInvalid
}

// GetConfig 获取配置
func (j *JWTManager) GetConfig() *config.JWTConfig {
	return j.config
}

// GetAccessExpire 获取访问令牌过期时间
func (j *JWTManager) GetAccessExpire() time.Duration {
	return j.config.AccessExpire
}

// GetRefreshExpire 获取刷新令牌过期时间
func (j *JWTManager) GetRefreshExpire() time.Duration {
	return j.config.RefreshExpire
}
