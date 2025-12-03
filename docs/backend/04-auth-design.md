# 认证授权设计

## 概述

本系统采用 JWT（JSON Web Token）进行用户认证，结合 Casbin 实现 RBAC（基于角色的访问控制）权限管理。

## 认证方案

### JWT Token 设计

```go
// pkg/auth/jwt.go
package auth

import (
    "errors"
    "time"
    
    "github.com/golang-jwt/jwt/v5"
)

var (
    ErrTokenExpired     = errors.New("token has expired")
    ErrTokenMalformed   = errors.New("token is malformed")
    ErrTokenInvalid     = errors.New("token is invalid")
    ErrTokenNotValidYet = errors.New("token is not active yet")
)

// Claims 自定义声明
type Claims struct {
    UserID   int64  `json:"user_id"`
    Username string `json:"username"`
    RoleKey  string `json:"role_key"`
    RoleID   int64  `json:"role_id"`
    DataScope string `json:"data_scope"` // 数据权限范围
    jwt.RegisteredClaims
}

// JWTConfig JWT 配置
type JWTConfig struct {
    SecretKey     string        `yaml:"secret_key"`
    Issuer        string        `yaml:"issuer"`
    AccessExpire  time.Duration `yaml:"access_expire"`  // 访问令牌过期时间
    RefreshExpire time.Duration `yaml:"refresh_expire"` // 刷新令牌过期时间
}

// JWTManager JWT 管理器
type JWTManager struct {
    config *JWTConfig
}

func NewJWTManager(config *JWTConfig) *JWTManager {
    return &JWTManager{config: config}
}

// GenerateToken 生成 Token
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
```

### Token 双令牌机制

```
Access Token:
- 有效期: 2 小时
- 用于访问受保护资源
- 不存储在服务端

Refresh Token:
- 有效期: 7 天
- 用于刷新 Access Token
- 存储在 Redis 中，支持单点登出
```

### Token 存储（Redis）

```go
// pkg/auth/token_store.go
package auth

import (
    "context"
    "fmt"
    "time"
    
    "github.com/redis/go-redis/v9"
)

const (
    tokenKeyPrefix   = "token:"
    refreshKeyPrefix = "refresh:"
)

// TokenStore Token 存储
type TokenStore struct {
    redis  *redis.Client
    config *JWTConfig
}

func NewTokenStore(redis *redis.Client, config *JWTConfig) *TokenStore {
    return &TokenStore{redis: redis, config: config}
}

// StoreRefreshToken 存储刷新 Token
func (s *TokenStore) StoreRefreshToken(ctx context.Context, userID int64, refreshToken string) error {
    key := fmt.Sprintf("%s%d", refreshKeyPrefix, userID)
    return s.redis.Set(ctx, key, refreshToken, s.config.RefreshExpire).Err()
}

// GetRefreshToken 获取刷新 Token
func (s *TokenStore) GetRefreshToken(ctx context.Context, userID int64) (string, error) {
    key := fmt.Sprintf("%s%d", refreshKeyPrefix, userID)
    return s.redis.Get(ctx, key).Result()
}

// DeleteRefreshToken 删除刷新 Token（登出）
func (s *TokenStore) DeleteRefreshToken(ctx context.Context, userID int64) error {
    key := fmt.Sprintf("%s%d", refreshKeyPrefix, userID)
    return s.redis.Del(ctx, key).Err()
}

// IsTokenBlacklisted 检查 Token 是否在黑名单
func (s *TokenStore) IsTokenBlacklisted(ctx context.Context, tokenID string) (bool, error) {
    key := fmt.Sprintf("blacklist:%s", tokenID)
    exists, err := s.redis.Exists(ctx, key).Result()
    return exists > 0, err
}

// BlacklistToken 将 Token 加入黑名单
func (s *TokenStore) BlacklistToken(ctx context.Context, tokenID string, expireAt time.Time) error {
    key := fmt.Sprintf("blacklist:%s", tokenID)
    ttl := time.Until(expireAt)
    if ttl <= 0 {
        return nil
    }
    return s.redis.Set(ctx, key, "1", ttl).Err()
}
```

---

## 登录流程

### 登录服务

```go
// internal/app/sys/service/auth.go
package service

type AuthService struct {
    userRepo   repository.UserRepository
    jwtManager *auth.JWTManager
    tokenStore *auth.TokenStore
    captcha    *captcha.Captcha
}

// LoginReq 登录请求
type LoginReq struct {
    Username    string `json:"username" binding:"required"`
    Password    string `json:"password" binding:"required"`
    CaptchaID   string `json:"captcha_id" binding:"required"`
    CaptchaCode string `json:"captcha_code" binding:"required"`
}

// LoginResp 登录响应
type LoginResp struct {
    AccessToken  string `json:"access_token"`
    RefreshToken string `json:"refresh_token"`
    ExpiresIn    int64  `json:"expires_in"`
}

// Login 用户登录
func (s *AuthService) Login(ctx context.Context, req *LoginReq, clientIP string) (*LoginResp, error) {
    // 1. 验证码校验
    if !s.captcha.Verify(req.CaptchaID, req.CaptchaCode) {
        return nil, errors.New("验证码错误")
    }
    
    // 2. 获取用户
    user, err := s.userRepo.GetByUsername(ctx, req.Username)
    if err != nil {
        return nil, errors.New("用户名或密码错误")
    }
    
    // 3. 检查用户状态
    if user.Status != StatusEnabled {
        return nil, errors.New("账号已被禁用")
    }
    
    // 4. 验证密码
    if !auth.VerifyPassword(req.Password, user.Password) {
        // 记录登录失败
        s.recordLoginFail(ctx, user.ID, clientIP)
        return nil, errors.New("用户名或密码错误")
    }
    
    // 5. 生成 Token
    claims := &auth.Claims{
        UserID:    user.ID,
        Username:  user.Username,
        RoleKey:   user.Role.RoleKey,
        RoleID:    user.RoleID,
        DataScope: user.Role.DataScope,
    }
    
    accessToken, err := s.jwtManager.GenerateToken(claims)
    if err != nil {
        return nil, fmt.Errorf("生成 Token 失败: %w", err)
    }
    
    refreshToken, err := s.jwtManager.GenerateRefreshToken(user.ID)
    if err != nil {
        return nil, fmt.Errorf("生成刷新 Token 失败: %w", err)
    }
    
    // 6. 存储刷新 Token
    if err := s.tokenStore.StoreRefreshToken(ctx, user.ID, refreshToken); err != nil {
        return nil, fmt.Errorf("存储 Token 失败: %w", err)
    }
    
    // 7. 记录登录日志
    s.recordLoginSuccess(ctx, user.ID, clientIP)
    
    return &LoginResp{
        AccessToken:  accessToken,
        RefreshToken: refreshToken,
        ExpiresIn:    int64(s.jwtManager.config.AccessExpire.Seconds()),
    }, nil
}

// RefreshToken 刷新 Token
func (s *AuthService) RefreshToken(ctx context.Context, refreshToken string) (*LoginResp, error) {
    // 1. 解析刷新 Token
    claims, err := s.jwtManager.ParseToken(refreshToken)
    if err != nil {
        return nil, errors.New("刷新 Token 无效")
    }
    
    // 2. 验证刷新 Token 是否有效
    storedToken, err := s.tokenStore.GetRefreshToken(ctx, claims.UserID)
    if err != nil || storedToken != refreshToken {
        return nil, errors.New("刷新 Token 已失效")
    }
    
    // 3. 获取最新用户信息
    user, err := s.userRepo.GetByID(ctx, claims.UserID)
    if err != nil {
        return nil, errors.New("用户不存在")
    }
    
    // 4. 生成新 Token
    newClaims := &auth.Claims{
        UserID:    user.ID,
        Username:  user.Username,
        RoleKey:   user.Role.RoleKey,
        RoleID:    user.RoleID,
        DataScope: user.Role.DataScope,
    }
    
    accessToken, err := s.jwtManager.GenerateToken(newClaims)
    if err != nil {
        return nil, err
    }
    
    newRefreshToken, err := s.jwtManager.GenerateRefreshToken(user.ID)
    if err != nil {
        return nil, err
    }
    
    // 5. 更新刷新 Token
    if err := s.tokenStore.StoreRefreshToken(ctx, user.ID, newRefreshToken); err != nil {
        return nil, err
    }
    
    return &LoginResp{
        AccessToken:  accessToken,
        RefreshToken: newRefreshToken,
        ExpiresIn:    int64(s.jwtManager.config.AccessExpire.Seconds()),
    }, nil
}

// Logout 登出
func (s *AuthService) Logout(ctx context.Context, userID int64, tokenID string, expireAt time.Time) error {
    // 1. 删除刷新 Token
    if err := s.tokenStore.DeleteRefreshToken(ctx, userID); err != nil {
        return err
    }
    
    // 2. 将当前 Token 加入黑名单
    return s.tokenStore.BlacklistToken(ctx, tokenID, expireAt)
}
```

### 密码处理

```go
// pkg/auth/password.go
package auth

import (
    "golang.org/x/crypto/bcrypt"
)

const bcryptCost = 10

// HashPassword 加密密码
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
    return string(bytes), err
}

// VerifyPassword 验证密码
func VerifyPassword(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
```

---

## 授权方案 (RBAC)

### Casbin 模型定义

```ini
# config/rbac_model.conf
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act || r.sub == "admin"
```

### Casbin 适配器

```go
// pkg/auth/casbin.go
package auth

import (
    "github.com/casbin/casbin/v2"
    gormadapter "github.com/casbin/gorm-adapter/v3"
    "gorm.io/gorm"
)

// CasbinEnforcer Casbin 执行器
type CasbinEnforcer struct {
    enforcer *casbin.Enforcer
}

// NewCasbinEnforcer 创建执行器
func NewCasbinEnforcer(db *gorm.DB, modelPath string) (*CasbinEnforcer, error) {
    adapter, err := gormadapter.NewAdapterByDB(db)
    if err != nil {
        return nil, err
    }
    
    enforcer, err := casbin.NewEnforcer(modelPath, adapter)
    if err != nil {
        return nil, err
    }
    
    // 加载策略
    if err := enforcer.LoadPolicy(); err != nil {
        return nil, err
    }
    
    return &CasbinEnforcer{enforcer: enforcer}, nil
}

// CheckPermission 检查权限
func (c *CasbinEnforcer) CheckPermission(roleKey, path, method string) (bool, error) {
    return c.enforcer.Enforce(roleKey, path, method)
}

// AddPolicy 添加策略
func (c *CasbinEnforcer) AddPolicy(roleKey, path, method string) (bool, error) {
    return c.enforcer.AddPolicy(roleKey, path, method)
}

// RemovePolicy 移除策略
func (c *CasbinEnforcer) RemovePolicy(roleKey, path, method string) (bool, error) {
    return c.enforcer.RemovePolicy(roleKey, path, method)
}

// AddRoleForUser 为用户添加角色
func (c *CasbinEnforcer) AddRoleForUser(userKey, roleKey string) (bool, error) {
    return c.enforcer.AddGroupingPolicy(userKey, roleKey)
}

// GetRolesForUser 获取用户角色
func (c *CasbinEnforcer) GetRolesForUser(userKey string) ([]string, error) {
    return c.enforcer.GetRolesForUser(userKey)
}

// ReloadPolicy 重新加载策略
func (c *CasbinEnforcer) ReloadPolicy() error {
    return c.enforcer.LoadPolicy()
}
```

### 权限同步服务

```go
// internal/app/sys/service/casbin.go
package service

type CasbinService struct {
    enforcer *auth.CasbinEnforcer
    menuRepo repository.MenuRepository
    roleRepo repository.RoleRepository
}

// SyncRolePermissions 同步角色权限
func (s *CasbinService) SyncRolePermissions(ctx context.Context, roleID int64) error {
    // 1. 获取角色信息
    role, err := s.roleRepo.GetByID(ctx, roleID)
    if err != nil {
        return err
    }
    
    // 2. 获取角色关联的菜单
    menus, err := s.menuRepo.GetByRoleID(ctx, roleID)
    if err != nil {
        return err
    }
    
    // 3. 清除旧权限
    s.enforcer.enforcer.DeletePermissionsForUser(role.RoleKey)
    
    // 4. 添加新权限
    for _, menu := range menus {
        if menu.Permission != "" && menu.Path != "" {
            // 解析权限: sys:user:list -> GET /api/v1/sys/user
            path, method := s.parsePermission(menu.Permission, menu.Path)
            s.enforcer.AddPolicy(role.RoleKey, path, method)
        }
    }
    
    return nil
}

// parsePermission 解析权限字符串
func (s *CasbinService) parsePermission(permission, path string) (string, string) {
    // 简单映射: :list -> GET, :add -> POST, :edit -> PUT, :delete -> DELETE
    method := "GET"
    if strings.HasSuffix(permission, ":add") {
        method = "POST"
    } else if strings.HasSuffix(permission, ":edit") {
        method = "PUT"
    } else if strings.HasSuffix(permission, ":delete") {
        method = "DELETE"
    }
    
    return path, method
}
```

---

## 中间件实现

### JWT 认证中间件

```go
// internal/middleware/jwt.go
package middleware

import (
    "strings"
    
    "github.com/gin-gonic/gin"
)

// JWTAuth JWT 认证中间件
func JWTAuth(jwtManager *auth.JWTManager, tokenStore *auth.TokenStore) gin.HandlerFunc {
    return func(c *gin.Context) {
        // 1. 获取 Token
        token := c.GetHeader("Authorization")
        if token == "" {
            c.JSON(401, gin.H{"code": 401, "message": "请先登录"})
            c.Abort()
            return
        }
        
        // 2. 去除 Bearer 前缀
        token = strings.TrimPrefix(token, "Bearer ")
        
        // 3. 解析 Token
        claims, err := jwtManager.ParseToken(token)
        if err != nil {
            msg := "Token 无效"
            if err == auth.ErrTokenExpired {
                msg = "Token 已过期"
            }
            c.JSON(401, gin.H{"code": 401, "message": msg})
            c.Abort()
            return
        }
        
        // 4. 检查黑名单
        if claims.ID != "" {
            blacklisted, _ := tokenStore.IsTokenBlacklisted(c.Request.Context(), claims.ID)
            if blacklisted {
                c.JSON(401, gin.H{"code": 401, "message": "Token 已失效"})
                c.Abort()
                return
            }
        }
        
        // 5. 设置用户信息到上下文
        c.Set("user_id", claims.UserID)
        c.Set("username", claims.Username)
        c.Set("role_key", claims.RoleKey)
        c.Set("role_id", claims.RoleID)
        c.Set("data_scope", claims.DataScope)
        c.Set("claims", claims)
        
        c.Next()
    }
}
```

### Casbin 权限中间件

```go
// internal/middleware/casbin.go
package middleware

import (
    "github.com/gin-gonic/gin"
)

// CasbinAuth Casbin 权限中间件
func CasbinAuth(enforcer *auth.CasbinEnforcer) gin.HandlerFunc {
    return func(c *gin.Context) {
        // 1. 获取角色
        roleKey, exists := c.Get("role_key")
        if !exists {
            c.JSON(401, gin.H{"code": 401, "message": "请先登录"})
            c.Abort()
            return
        }
        
        // 2. 获取请求路径和方法
        path := c.Request.URL.Path
        method := c.Request.Method
        
        // 3. 检查权限
        ok, err := enforcer.CheckPermission(roleKey.(string), path, method)
        if err != nil {
            c.JSON(500, gin.H{"code": 500, "message": "权限检查失败"})
            c.Abort()
            return
        }
        
        if !ok {
            c.JSON(403, gin.H{"code": 403, "message": "没有访问权限"})
            c.Abort()
            return
        }
        
        c.Next()
    }
}
```

### 数据权限中间件

```go
// internal/middleware/datascope.go
package middleware

import (
    "github.com/gin-gonic/gin"
)

// 数据权限范围
const (
    DataScopeAll       = "1" // 全部数据权限
    DataScopeCustom    = "2" // 自定义数据权限
    DataScopeDept      = "3" // 本部门数据权限
    DataScopeDeptBelow = "4" // 本部门及以下数据权限
    DataScopeSelf      = "5" // 仅本人数据权限
)

// DataScope 数据权限中间件
func DataScope() gin.HandlerFunc {
    return func(c *gin.Context) {
        dataScope, _ := c.Get("data_scope")
        userID, _ := c.Get("user_id")
        
        // 构建数据权限条件
        var scopeCondition string
        
        switch dataScope.(string) {
        case DataScopeAll:
            // 全部数据，不需要条件
            scopeCondition = ""
        case DataScopeSelf:
            // 仅本人数据
            scopeCondition = fmt.Sprintf("created_by = %d", userID)
        case DataScopeDept:
            // 本部门数据 (需要查询用户所属部门)
            // scopeCondition = fmt.Sprintf("dept_id = %d", deptID)
        case DataScopeDeptBelow:
            // 本部门及以下
            // 需要递归查询下级部门
        case DataScopeCustom:
            // 自定义权限
            // 需要查询角色关联的部门
        }
        
        c.Set("scope_condition", scopeCondition)
        c.Next()
    }
}
```

---

## 权限数据结构

### 角色表

```sql
CREATE TABLE sys_role (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    role_name VARCHAR(128) NOT NULL COMMENT '角色名称',
    role_key VARCHAR(128) NOT NULL UNIQUE COMMENT '角色标识',
    role_sort INT DEFAULT 0 COMMENT '排序',
    data_scope CHAR(1) DEFAULT '1' COMMENT '数据范围:1全部,2自定义,3本部门,4本部门及以下,5仅本人',
    status CHAR(1) DEFAULT '0' COMMENT '状态:0正常,1停用',
    remark VARCHAR(500) COMMENT '备注',
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL
);
```

### 菜单表

```sql
CREATE TABLE sys_menu (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    parent_id BIGINT DEFAULT 0 COMMENT '父菜单ID',
    menu_name VARCHAR(50) NOT NULL COMMENT '菜单名称',
    menu_type CHAR(1) NOT NULL COMMENT 'M目录,C菜单,F按钮',
    sort INT DEFAULT 0 COMMENT '排序',
    path VARCHAR(200) COMMENT '路由地址',
    component VARCHAR(255) COMMENT '组件路径',
    permission VARCHAR(100) COMMENT '权限标识',
    icon VARCHAR(100) COMMENT '图标',
    visible CHAR(1) DEFAULT '0' COMMENT '0显示,1隐藏',
    status CHAR(1) DEFAULT '0' COMMENT '0正常,1停用',
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL
);
```

### 角色菜单关联表

```sql
CREATE TABLE sys_role_menu (
    role_id BIGINT NOT NULL,
    menu_id BIGINT NOT NULL,
    PRIMARY KEY (role_id, menu_id)
);
```

### Casbin 规则表（自动创建）

```sql
CREATE TABLE casbin_rule (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    ptype VARCHAR(100),
    v0 VARCHAR(100),
    v1 VARCHAR(100),
    v2 VARCHAR(100),
    v3 VARCHAR(100),
    v4 VARCHAR(100),
    v5 VARCHAR(100)
);
```

---

## API 权限配置示例

```yaml
# 路由权限配置
permissions:
  # 广告主管理
  - path: /api/v1/advertiser
    method: GET
    permission: ad:advertiser:list
  - path: /api/v1/advertiser
    method: POST
    permission: ad:advertiser:add
  - path: /api/v1/advertiser/:id
    method: PUT
    permission: ad:advertiser:edit
  - path: /api/v1/advertiser/:id
    method: DELETE
    permission: ad:advertiser:delete
  
  # 广告系列管理
  - path: /api/v1/campaign
    method: GET
    permission: ad:campaign:list
  - path: /api/v1/campaign
    method: POST
    permission: ad:campaign:add
  
  # 数据报表
  - path: /api/v1/report/*
    method: GET
    permission: data:report:view
  - path: /api/v1/report/export
    method: POST
    permission: data:report:export
```

---

## 安全措施

### 1. 密码安全

```go
// 密码复杂度验证
func ValidatePassword(password string) error {
    if len(password) < 8 {
        return errors.New("密码长度不能少于8位")
    }
    
    var hasUpper, hasLower, hasNumber, hasSpecial bool
    for _, c := range password {
        switch {
        case unicode.IsUpper(c):
            hasUpper = true
        case unicode.IsLower(c):
            hasLower = true
        case unicode.IsNumber(c):
            hasNumber = true
        case unicode.IsPunct(c) || unicode.IsSymbol(c):
            hasSpecial = true
        }
    }
    
    if !hasUpper || !hasLower || !hasNumber || !hasSpecial {
        return errors.New("密码必须包含大小写字母、数字和特殊字符")
    }
    
    return nil
}
```

### 2. 登录安全

```go
// 登录失败处理
const (
    MaxLoginAttempts = 5           // 最大尝试次数
    LockDuration     = time.Hour   // 锁定时长
)

func (s *AuthService) recordLoginFail(ctx context.Context, userID int64, clientIP string) {
    key := fmt.Sprintf("login_fail:%d", userID)
    
    count, _ := s.redis.Incr(ctx, key).Result()
    s.redis.Expire(ctx, key, LockDuration)
    
    if count >= MaxLoginAttempts {
        // 锁定账户
        s.userRepo.UpdateStatus(ctx, userID, StatusLocked)
    }
}

func (s *AuthService) checkLoginLock(ctx context.Context, userID int64) error {
    key := fmt.Sprintf("login_fail:%d", userID)
    count, _ := s.redis.Get(ctx, key).Int()
    
    if count >= MaxLoginAttempts {
        return errors.New("账号已被锁定，请稍后再试")
    }
    return nil
}
```

### 3. Token 安全

```go
// Token 绑定客户端信息
type Claims struct {
    // ... 其他字段
    ClientIP    string `json:"client_ip"`
    UserAgent   string `json:"user_agent"`
    DeviceID    string `json:"device_id"`
}

// 验证时检查客户端一致性
func (j *JWTManager) ValidateClient(claims *Claims, clientIP, userAgent string) bool {
    // 可选：严格模式下检查 IP
    // if claims.ClientIP != clientIP {
    //     return false
    // }
    return true
}
```

---

## 验证码

```go
// pkg/captcha/captcha.go
package captcha

import (
    "github.com/mojocn/base64Captcha"
)

type Captcha struct {
    store  base64Captcha.Store
    driver base64Captcha.Driver
}

func NewCaptcha(redis *redis.Client) *Captcha {
    store := NewRedisStore(redis)
    driver := base64Captcha.NewDriverDigit(80, 240, 6, 0.7, 80)
    
    return &Captcha{
        store:  store,
        driver: driver,
    }
}

// Generate 生成验证码
func (c *Captcha) Generate() (id, b64s string, err error) {
    captcha := base64Captcha.NewCaptcha(c.driver, c.store)
    return captcha.Generate()
}

// Verify 验证验证码
func (c *Captcha) Verify(id, code string) bool {
    return c.store.Verify(id, code, true)
}
```
