package main

import (
	"log"
	"os"

	"github.com/bububa/oceanengine/server/internal/config"
	"github.com/bububa/oceanengine/server/internal/handler"
	"github.com/bububa/oceanengine/server/internal/middleware"
	"github.com/bububa/oceanengine/server/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// 加载 .env 文件 (如果存在)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// 加载配置
	cfg := config.Load()

	// 验证必要配置
	if cfg.OceanEngine.AppID == 0 || cfg.OceanEngine.AppSecret == "" {
		log.Println("Warning: OCEANENGINE_APP_ID or OCEANENGINE_APP_SECRET not set")
	}

	// 设置Gin模式
	gin.SetMode(cfg.Server.Mode)

	// 创建服务
	oceanEngineSvc := service.NewOceanEngineService(
		cfg.OceanEngine.AppID,
		cfg.OceanEngine.AppSecret,
	)

	// 创建Handler
	h := handler.NewHandler(oceanEngineSvc)

	// 创建路由
	r := gin.New()

	// 中间件
	r.Use(middleware.Logger())
	r.Use(middleware.Recovery())
	r.Use(middleware.CORS())

	// 路由注册
	setupRoutes(r, h)

	// 启动服务
	addr := ":" + cfg.Server.Port
	log.Printf("Server starting on %s", addr)

	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
		os.Exit(1)
	}
}

func setupRoutes(r *gin.Engine, h *handler.Handler) {
	// 健康检查
	r.GET("/health", h.Health)

	// API v1
	v1 := r.Group("/api/v1")
	{
		// OAuth
		oauth := v1.Group("/oauth")
		{
			oauth.POST("/auth_url", h.GetAuthURL)
			oauth.POST("/access_token", h.GetAccessToken)
			oauth.POST("/refresh_token", h.RefreshToken)
		}

		// 广告主
		advertiser := v1.Group("/advertiser")
		{
			advertiser.GET("/list", h.GetAdvertisers)
			advertiser.POST("/info", h.GetAdvertiserInfo)
		}
	}
}
