package config

import (
	"os"
	"strconv"
)

// Config 应用配置
type Config struct {
	// Server 服务配置
	Server ServerConfig

	// OceanEngine SDK配置
	OceanEngine OceanEngineConfig
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Port string
	Mode string // debug, release, test
}

// OceanEngineConfig 巨量引擎配置
type OceanEngineConfig struct {
	AppID     uint64
	AppSecret string
}

// Load 从环境变量加载配置
func Load() *Config {
	return &Config{
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
			Mode: getEnv("SERVER_MODE", "release"),
		},
		OceanEngine: OceanEngineConfig{
			AppID:     getEnvUint64("OCEANENGINE_APP_ID", 0),
			AppSecret: getEnv("OCEANENGINE_APP_SECRET", ""),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvUint64(key string, defaultValue uint64) uint64 {
	if value := os.Getenv(key); value != "" {
		if v, err := strconv.ParseUint(value, 10, 64); err == nil {
			return v
		}
	}
	return defaultValue
}
