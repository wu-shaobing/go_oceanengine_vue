package integration

import (
	"bytes"
	"encoding/json"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"oceanengine-backend/config"
	adModel "oceanengine-backend/internal/app/ad/model"
	adminModel "oceanengine-backend/internal/app/admin/model"
	advModel "oceanengine-backend/internal/app/advertiser/model"
	audienceModel "oceanengine-backend/internal/app/audience/model"
	campaignModel "oceanengine-backend/internal/app/campaign/model"
	creativeModel "oceanengine-backend/internal/app/creative/model"
	mediaModel "oceanengine-backend/internal/app/media/model"
	reportModel "oceanengine-backend/internal/app/report/model"
	"oceanengine-backend/internal/router"
	"oceanengine-backend/pkg/auth"
)

// TestServer 测试服务器
type TestServer struct {
	Router     *gin.Engine
	DB         *gorm.DB
	JWTManager *auth.JWTManager
	Logger     *zap.Logger
}

// NewTestServer 创建测试服务器
func NewTestServer(t *testing.T) *TestServer {
	// 设置测试模式
	gin.SetMode(gin.TestMode)

	// 创建内存SQLite数据库用于测试
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}

	// 自动迁移测试表 - 系统管理表
	err = db.AutoMigrate(
		&adminModel.User{},
		&adminModel.Role{},
		&adminModel.Menu{},
		&adminModel.RoleMenu{},
		&adminModel.OperationLog{},
	)
	if err != nil {
		t.Fatalf("Failed to migrate admin tables: %v", err)
	}

	// 自动迁移测试表 - 广告业务表
	err = db.AutoMigrate(
		&advModel.Advertiser{},
		&advModel.AdvertiserFund{},
		&campaignModel.Campaign{},
		&adModel.Ad{},
		&creativeModel.Creative{},
		&mediaModel.MaterialImage{},
		&mediaModel.MaterialVideo{},
		&audienceModel.AudiencePackage{},
		&audienceModel.CustomAudience{},
		&reportModel.AdvertiserReport{},
		&reportModel.CampaignReport{},
		&reportModel.AdReport{},
		&reportModel.ExportTask{},
	)
	if err != nil {
		t.Fatalf("Failed to migrate business tables: %v", err)
	}

	// 创建测试日志
	logger, _ := zap.NewDevelopment()

	// 创建JWT管理器
	jwtManager := auth.NewJWTManager(&config.JWTConfig{
		SecretKey:     "test-secret-key-for-integration-tests",
		Issuer:        "oceanengine-test",
		AccessExpire:  2 * time.Hour,
		RefreshExpire: 168 * time.Hour,
	})

	// 创建Ocean配置
	oceanCfg := &config.OceanConfig{
		AppID:       os.Getenv("OCEAN_APP_ID"),
		Secret:      os.Getenv("OCEAN_SECRET"),
		RedirectURI: "http://localhost:8080/api/v1/advertisers/oauth/callback",
		BaseURL:     "https://ad.oceanengine.com/open_api",
		Timeout:     30 * time.Second,
	}

	// 创建路由
	r := router.NewRouter(db, logger, jwtManager, oceanCfg)
	engine := r.Setup(gin.TestMode)

	return &TestServer{
		Router:     engine,
		DB:         db,
		JWTManager: jwtManager,
		Logger:     logger,
	}
}

// SeedTestData 填充测试数据
func (ts *TestServer) SeedTestData(t *testing.T) {
	// 创建测试角色
	role := &adminModel.Role{
		Name:      "超级管理员",
		Key:       "admin",
		Sort:      1,
		Status:    1,
		DataScope: 1,
		Remark:    "测试角色",
	}
	if err := ts.DB.Create(role).Error; err != nil {
		t.Fatalf("Failed to create test role: %v", err)
	}

	// 创建测试用户（密码: admin123）
	hashedPassword, _ := auth.HashPassword("admin123")
	user := &adminModel.User{
		Username: "admin",
		Password: hashedPassword,
		Nickname: "管理员",
		Email:    "admin@test.com",
		Phone:    "13800138000",
		RoleID:   role.ID,
		Status:   1,
	}
	if err := ts.DB.Create(user).Error; err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}

	// 创建测试菜单
	menus := []adminModel.Menu{
		{Name: "系统管理", Path: "/system", Component: "Layout", Sort: 1, Status: 1, Type: 1},
		{Name: "广告管理", Path: "/ads", Component: "Layout", Sort: 2, Status: 1, Type: 1},
	}
	for _, menu := range menus {
		ts.DB.Create(&menu)
	}
}

// GenerateTestToken 生成测试Token
func (ts *TestServer) GenerateTestToken(userID uint, username string) (string, error) {
	claims := &auth.Claims{
		UserID:    int64(userID),
		Username:  username,
		RoleKey:   "admin",
		RoleID:    1,
		DataScope: "1",
	}
	return ts.JWTManager.GenerateToken(claims)
}

// MakeRequest 发送测试请求
func (ts *TestServer) MakeRequest(method, path string, body interface{}, token string) *httptest.ResponseRecorder {
	var reqBody *bytes.Buffer
	if body != nil {
		jsonBytes, _ := json.Marshal(body)
		reqBody = bytes.NewBuffer(jsonBytes)
	} else {
		reqBody = bytes.NewBuffer(nil)
	}

	req, _ := http.NewRequest(method, path, reqBody)
	req.Header.Set("Content-Type", "application/json")
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	w := httptest.NewRecorder()
	ts.Router.ServeHTTP(w, req)
	return w
}

// ParseResponse 解析响应
func ParseResponse(w *httptest.ResponseRecorder, v interface{}) error {
	return json.Unmarshal(w.Body.Bytes(), v)
}

// Response 通用响应结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// LoginResponse 登录响应
type LoginResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		ExpiresIn    int64  `json:"expires_in"`
	} `json:"data"`
}

// Cleanup 清理测试资源
func (ts *TestServer) Cleanup() {
	sqlDB, _ := ts.DB.DB()
	if sqlDB != nil {
		sqlDB.Close()
	}
}

// isRedisAvailable 检查Redis是否可用
func isRedisAvailable() bool {
	redisHost := os.Getenv("REDIS_HOST")
	if redisHost == "" {
		redisHost = "localhost"
	}
	redisPort := os.Getenv("REDIS_PORT")
	if redisPort == "" {
		redisPort = "6379"
	}

	// 尝试连接Redis
	conn, err := net.DialTimeout("tcp", redisHost+":"+redisPort, 1*time.Second)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

// isExternalAPIAvailable 检查外部API是否可用
func isExternalAPIAvailable() bool {
	// 检查是否配置了真实的API凭证
	appID := os.Getenv("OCEAN_APP_ID")
	return appID != ""
}
