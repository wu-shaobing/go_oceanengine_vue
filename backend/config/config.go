package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/viper"
)

// Config 应用配置
type Config struct {
	Server    ServerConfig    `mapstructure:"server"`
	Database  DatabaseConfig  `mapstructure:"database"`
	Redis     RedisConfig     `mapstructure:"redis"`
	JWT       JWTConfig       `mapstructure:"jwt"`
	Logger    LoggerConfig    `mapstructure:"logger"`
	Ocean     OceanConfig     `mapstructure:"ocean"`
	Qianchuan QianchuanConfig `mapstructure:"qianchuan"`
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Mode         string        `mapstructure:"mode"`
	Host         string        `mapstructure:"host"`
	Port         int           `mapstructure:"port"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"`
}

// GetAddr 获取服务地址
func (c *ServerConfig) GetAddr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Driver          string        `mapstructure:"driver"`
	Host            string        `mapstructure:"host"`
	Port            int           `mapstructure:"port"`
	Database        string        `mapstructure:"database"`
	Username        string        `mapstructure:"username"`
	Password        string        `mapstructure:"password"`
	Charset         string        `mapstructure:"charset"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns"`
	MaxOpenConns    int           `mapstructure:"max_open_conns"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`
	LogLevel        string        `mapstructure:"log_level"`
}

// GetDSN 获取数据库连接字符串
func (c *DatabaseConfig) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		c.Username, c.Password, c.Host, c.Port, c.Database, c.Charset)
}

// RedisConfig Redis配置
type RedisConfig struct {
	Addr         string        `mapstructure:"addr"`
	Password     string        `mapstructure:"password"`
	DB           int           `mapstructure:"db"`
	PoolSize     int           `mapstructure:"pool_size"`
	MinIdleConns int           `mapstructure:"min_idle_conns"`
	DialTimeout  time.Duration `mapstructure:"dial_timeout"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"`
}

// JWTConfig JWT配置
type JWTConfig struct {
	SecretKey     string        `mapstructure:"secret_key"`
	Issuer        string        `mapstructure:"issuer"`
	AccessExpire  time.Duration `mapstructure:"access_expire"`
	RefreshExpire time.Duration `mapstructure:"refresh_expire"`
}

// LoggerConfig 日志配置
type LoggerConfig struct {
	Level      string `mapstructure:"level"`
	Format     string `mapstructure:"format"`
	Output     string `mapstructure:"output"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
	Compress   bool   `mapstructure:"compress"`
}

// OceanConfig Ocean Engine配置 (巨量广告 - 代理商)
type OceanConfig struct {
	AppID        string        `mapstructure:"app_id"`
	Secret       string        `mapstructure:"secret"`
	RedirectURI  string        `mapstructure:"redirect_uri"`
	BaseURL      string        `mapstructure:"base_url"`
	AuthURL      string        `mapstructure:"auth_url"`
	Timeout      time.Duration `mapstructure:"timeout"`
	RetryCount   int           `mapstructure:"retry_count"`
	MaterialAuth bool          `mapstructure:"material_auth"` // 是否启用素材授权
}

// QianchuanConfig 巨量千川配置
type QianchuanConfig struct {
	AppID        string        `mapstructure:"app_id"`
	Secret       string        `mapstructure:"secret"`
	RedirectURI  string        `mapstructure:"redirect_uri"`
	BaseURL      string        `mapstructure:"base_url"`
	AuthURL      string        `mapstructure:"auth_url"`
	Timeout      time.Duration `mapstructure:"timeout"`
	RetryCount   int           `mapstructure:"retry_count"`
	MaterialAuth bool          `mapstructure:"material_auth"` // 是否启用素材授权
}

var cfg *Config

// Load 加载配置
func Load(configPath string) (*Config, error) {
	v := viper.New()
	v.SetConfigFile(configPath)
	v.SetConfigType("yaml")

	// 支持环境变量覆盖（嵌套字段 => 大写+下划线）
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	cfg = &Config{}
	if err := v.Unmarshal(cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	// 兼容常见环境变量别名（优先级高于配置文件）
	// 数据库
	if host := os.Getenv("DB_HOST"); host != "" {
		cfg.Database.Host = host
	}
	if port := os.Getenv("DB_PORT"); port != "" {
		if p, err := strconv.Atoi(port); err == nil {
			cfg.Database.Port = p
		}
	}
	if user := os.Getenv("DB_USER"); user != "" {
		cfg.Database.Username = user
	}
	if pass := os.Getenv("DB_PASSWORD"); pass != "" {
		cfg.Database.Password = pass
	}
	if name := os.Getenv("DB_NAME"); name != "" {
		cfg.Database.Database = name
	}

	// Redis
	if rpass := os.Getenv("REDIS_PASSWORD"); rpass != "" {
		cfg.Redis.Password = rpass
	}

	// JWT
	if sk := os.Getenv("JWT_SECRET_KEY"); sk != "" {
		cfg.JWT.SecretKey = sk
	}

	// Ocean Engine
	if appID := os.Getenv("OCEAN_APP_ID"); appID != "" {
		cfg.Ocean.AppID = appID
	}
	if secret := os.Getenv("OCEAN_SECRET"); secret != "" {
		cfg.Ocean.Secret = secret
	}

	// 设置默认值
	setDefaults(cfg)

	return cfg, nil
}

// Get 获取配置实例
func Get() *Config {
	return cfg
}

// setDefaults 设置默认值
func setDefaults(c *Config) {
	if c.Server.Mode == "" {
		c.Server.Mode = "debug"
	}
	if c.Server.Host == "" {
		c.Server.Host = "0.0.0.0"
	}
	if c.Server.Port == 0 {
		c.Server.Port = 8080
	}
	if c.Server.ReadTimeout == 0 {
		c.Server.ReadTimeout = 30 * time.Second
	}
	if c.Server.WriteTimeout == 0 {
		c.Server.WriteTimeout = 30 * time.Second
	}
	if c.Database.Charset == "" {
		c.Database.Charset = "utf8mb4"
	}
	if c.Database.MaxIdleConns == 0 {
		c.Database.MaxIdleConns = 10
	}
	if c.Database.MaxOpenConns == 0 {
		c.Database.MaxOpenConns = 100
	}
	if c.Database.ConnMaxLifetime == 0 {
		c.Database.ConnMaxLifetime = time.Hour
	}
	if c.Redis.PoolSize == 0 {
		c.Redis.PoolSize = 100
	}
	if c.JWT.Issuer == "" {
		c.JWT.Issuer = "oceanengine"
	}
	if c.JWT.AccessExpire == 0 {
		c.JWT.AccessExpire = 2 * time.Hour
	}
	if c.JWT.RefreshExpire == 0 {
		c.JWT.RefreshExpire = 7 * 24 * time.Hour
	}
	if c.Logger.Level == "" {
		c.Logger.Level = "info"
	}
	if c.Logger.Format == "" {
		c.Logger.Format = "json"
	}
	if c.Logger.Output == "" {
		c.Logger.Output = "stdout"
	}
	if c.Ocean.BaseURL == "" {
		c.Ocean.BaseURL = "https://ad.oceanengine.com/open_api"
	}
	if c.Ocean.AuthURL == "" {
		c.Ocean.AuthURL = "https://open.oceanengine.com/audit/oauth.html"
	}
	if c.Ocean.Timeout == 0 {
		c.Ocean.Timeout = 30 * time.Second
	}
	if c.Ocean.RetryCount == 0 {
		c.Ocean.RetryCount = 3
	}
	// 千川默认值
	if c.Qianchuan.BaseURL == "" {
		c.Qianchuan.BaseURL = "https://ad.oceanengine.com/open_api"
	}
	if c.Qianchuan.AuthURL == "" {
		c.Qianchuan.AuthURL = "https://qianchuan.jinritemai.com/openapi/qc/audit/oauth.html"
	}
	if c.Qianchuan.Timeout == 0 {
		c.Qianchuan.Timeout = 30 * time.Second
	}
	if c.Qianchuan.RetryCount == 0 {
		c.Qianchuan.RetryCount = 3
	}
}
