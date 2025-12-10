package middleware

import (
	"bytes"
	"io"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"oceanengine-backend/internal/app/admin/model"
)

// OperationLogConfig 操作日志配置
type OperationLogConfig struct {
	MaxBodySize     int      // 最大请求/响应体记录大小（字节）
	SkipPaths       []string // 跳过的路径
	SkipMethods     []string // 跳过的方法
	OnlyError       bool     // 仅记录错误请求
	SensitiveFields []string // 敏感字段列表
}

// DefaultOperationLogConfig 默认配置
func DefaultOperationLogConfig() *OperationLogConfig {
	return &OperationLogConfig{
		MaxBodySize: 4096, // 4KB
		SkipPaths: []string{
			"/health",
			"/api/v1/auth/captcha",
			"/api/v1/system/logs", // 避免日志接口本身被记录
		},
		SkipMethods: []string{"GET", "OPTIONS", "HEAD"},
		OnlyError:   false,
		SensitiveFields: []string{
			"password", "old_password", "new_password",
			"token", "access_token", "refresh_token",
			"secret", "secret_key", "api_key",
			"captcha", "captcha_code",
		},
	}
}

// operationLogWriter 包装响应写入器以捕获响应
type operationLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *operationLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// OperationLog 操作日志中间件
func OperationLog(db *gorm.DB, config *OperationLogConfig) gin.HandlerFunc {
	if config == nil {
		config = DefaultOperationLogConfig()
	}

	return func(c *gin.Context) {
		// 检查是否跳过
		if shouldSkip(c, config) {
			c.Next()
			return
		}

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
		blw := &operationLogWriter{body: bytes.NewBuffer(nil), ResponseWriter: c.Writer}
		c.Writer = blw

		c.Next()

		// 计算耗时
		latency := time.Since(start)
		status := c.Writer.Status()

		// 仅记录错误时跳过成功请求
		if config.OnlyError && status < 400 {
			return
		}

		// 提取用户信息
		userID := uint64(0)
		username := ""
		if uid, exists := c.Get("user_id"); exists {
			userID = uint64(uid.(int64))
		}
		if uname, exists := c.Get("username"); exists {
			username = uname.(string)
		}

		// 解析模块和操作
		module, action := parseModuleAndAction(c.Request.Method, path)

		// 脱敏处理
		sanitizedBody := sanitizeBody(string(requestBody), config.SensitiveFields, config.MaxBodySize)
		sanitizedResponse := truncateString(blw.body.String(), config.MaxBodySize)

		// 提取错误信息
		errorMsg := ""
		if status >= 400 {
			errorMsg = extractErrorMessage(blw.body.String())
		}

		// 构建日志记录
		log := &model.OperationLog{
			UserID:    userID,
			Username:  username,
			Module:    module,
			Action:    action,
			Method:    c.Request.Method,
			Path:      path,
			Query:     query,
			Body:      sanitizedBody,
			Response:  sanitizedResponse,
			IP:        c.ClientIP(),
			UserAgent: truncateString(c.Request.UserAgent(), 500),
			Status:    status,
			Latency:   latency.Milliseconds(),
			ErrorMsg:  errorMsg,
			CreatedAt: time.Now(),
		}

		// 异步写入数据库
		go func(log *model.OperationLog) {
			if err := db.Create(log).Error; err != nil {
				// 日志写入失败不影响主流程，仅输出到标准错误
				// 生产环境可接入监控告警
			}
		}(log)
	}
}

// shouldSkip 检查是否跳过日志记录
func shouldSkip(c *gin.Context, config *OperationLogConfig) bool {
	// 检查路径
	path := c.Request.URL.Path
	for _, skip := range config.SkipPaths {
		if strings.HasPrefix(path, skip) {
			return true
		}
	}

	// 检查方法
	method := c.Request.Method
	for _, skip := range config.SkipMethods {
		if method == skip {
			return true
		}
	}

	return false
}

// parseModuleAndAction 从路径解析模块和操作
func parseModuleAndAction(method, path string) (module, action string) {
	// 移除 /api/v1 前缀
	path = strings.TrimPrefix(path, "/api/v1/")
	path = strings.TrimPrefix(path, "/api/")

	// 解析模块
	parts := strings.Split(path, "/")
	if len(parts) > 0 {
		module = parts[0]
	}

	// 模块名映射
	moduleMap := map[string]string{
		"auth":        "认证管理",
		"system":      "系统管理",
		"advertisers": "广告主管理",
		"campaigns":   "广告系列",
		"ads":         "广告组",
		"creatives":   "创意管理",
		"reports":     "数据报表",
		"materials":   "素材管理",
		"audiences":   "人群定向",
		"qianchuan":   "千川电商",
		"enterprise":  "企业号",
		"local":       "本地推",
		"star":        "星图达人",
	}
	if mapped, ok := moduleMap[module]; ok {
		module = mapped
	}

	// 解析操作
	switch method {
	case "POST":
		action = "创建"
		if strings.Contains(path, "login") {
			action = "登录"
		} else if strings.Contains(path, "logout") {
			action = "登出"
		} else if strings.Contains(path, "sync") {
			action = "同步"
		} else if strings.Contains(path, "reset-password") {
			action = "重置密码"
		} else if strings.Contains(path, "change-password") {
			action = "修改密码"
		}
	case "PUT":
		action = "更新"
		if strings.Contains(path, "status") {
			action = "更新状态"
		} else if strings.Contains(path, "menus") {
			action = "分配菜单"
		}
	case "DELETE":
		action = "删除"
	case "GET":
		action = "查询"
	default:
		action = method
	}

	// 细化系统管理模块的操作
	if len(parts) > 1 && parts[0] == "system" {
		subModule := parts[1]
		subModuleMap := map[string]string{
			"users": "用户",
			"roles": "角色",
			"menus": "菜单",
			"logs":  "日志",
		}
		if mapped, ok := subModuleMap[subModule]; ok {
			action = action + mapped
		}
	}

	return module, action
}

// sanitizeBody 脱敏处理请求体
func sanitizeBody(body string, sensitiveFields []string, maxSize int) string {
	if body == "" {
		return ""
	}

	// 构建正则表达式匹配敏感字段
	for _, field := range sensitiveFields {
		// 匹配 JSON 格式: "field": "value" 或 "field":"value"
		patterns := []string{
			`"` + field + `"\s*:\s*"[^"]*"`,
			`"` + field + `"\s*:\s*'[^']*'`,
		}
		for _, pattern := range patterns {
			re := regexp.MustCompile(`(?i)` + pattern)
			body = re.ReplaceAllString(body, `"`+field+`":"******"`)
		}
	}

	return truncateString(body, maxSize)
}

// truncateString 截断字符串
func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "...[truncated]"
}

// extractErrorMessage 从响应中提取错误信息
func extractErrorMessage(response string) string {
	// 尝试提取 message 字段
	re := regexp.MustCompile(`"message"\s*:\s*"([^"]*)"`)
	matches := re.FindStringSubmatch(response)
	if len(matches) > 1 {
		return matches[1]
	}
	return truncateString(response, 500)
}
