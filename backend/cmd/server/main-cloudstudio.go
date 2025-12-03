//go:build cloudstudio
// +build cloudstudio

package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"oceanengine-backend/config"
	"oceanengine-backend/internal/router"
	"oceanengine-backend/pkg/auth"
	"oceanengine-backend/pkg/database"
	"oceanengine-backend/pkg/logger"
)

func main() {
	// 加载CloudStudio配置
	configPath := "config/settings.yml"
	if _, err := os.Stat("config/settings-cloudstudio.yml"); err == nil {
		configPath = "config/settings-cloudstudio.yml"
	}

	cfg, err := config.Load(configPath)
	if err != nil {
		fmt.Printf("加载配置失败: %v\n", err)
		os.Exit(1)
	}

	// 初始化日志
	log, err := logger.Init(&cfg.Logger)
	if err != nil {
		fmt.Printf("初始化日志失败: %v\n", err)
		os.Exit(1)
	}
	defer log.Sync()

	log.Info("启动 OceanEngine 后台管理系统 (CloudStudio版)")

	// 初始化数据库
	db, err := database.Init(&cfg.Database, log)
	if err != nil {
		log.Fatal(fmt.Sprintf("初始化数据库失败: %v", err))
	}
	log.Info("数据库连接成功")

	// 初始化 Redis (可选)
	if cfg.Redis.Addr != "" {
		_, err := database.InitRedis(&cfg.Redis, log)
		if err != nil {
			log.Warn(fmt.Sprintf("初始化 Redis 失败: %v", err))
		} else {
			log.Info("Redis 连接成功")
		}
	}

	// 初始化 JWT 管理器
	jwtManager := auth.NewJWTManager(&cfg.JWT)

	// 设置路由
	r := router.NewRouter(db, log, jwtManager, &cfg.Ocean)
	engine := r.Setup(cfg.Server.Mode)

	// 创建 HTTP 服务器
	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	srv := &http.Server{
		Addr:         addr,
		Handler:      engine,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
	}

	// 启动服务器（非阻塞）
	go func() {
		log.Info(fmt.Sprintf("服务器启动于 %s", addr))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(fmt.Sprintf("服务器启动失败: %v", err))
		}
	}()

	// 优雅关闭
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("正在关闭服务器...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(fmt.Sprintf("服务器关闭失败: %v", err))
	}

	log.Info("服务器已关闭")
}
