package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"oceanengine-backend/config"
)

var globalLogger *zap.Logger

// Init 初始化日志器
func Init(cfg *config.LoggerConfig) (*zap.Logger, error) {
	// 解析日志级别
	level, err := zapcore.ParseLevel(cfg.Level)
	if err != nil {
		level = zapcore.InfoLevel
	}

	// 编码配置
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.MillisDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// 选择编码器
	var encoder zapcore.Encoder
	if cfg.Format == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	// 构建输出
	var cores []zapcore.Core

	// 控制台输出
	if cfg.Output == "stdout" || cfg.Output == "both" {
		consoleCore := zapcore.NewCore(
			encoder,
			zapcore.AddSync(os.Stdout),
			level,
		)
		cores = append(cores, consoleCore)
	}

	// 文件输出
	if cfg.Output == "file" || cfg.Output == "both" {
		if cfg.Filename != "" {
			fileWriter := &lumberjack.Logger{
				Filename:   cfg.Filename,
				MaxSize:    cfg.MaxSize,
				MaxBackups: cfg.MaxBackups,
				MaxAge:     cfg.MaxAge,
				Compress:   cfg.Compress,
			}

			// 文件始终使用 JSON
			jsonEncoder := zapcore.NewJSONEncoder(encoderConfig)
			fileCore := zapcore.NewCore(
				jsonEncoder,
				zapcore.AddSync(fileWriter),
				level,
			)
			cores = append(cores, fileCore)
		}
	}

	// 合并核心
	core := zapcore.NewTee(cores...)

	// 创建日志器
	globalLogger = zap.New(core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)

	return globalLogger, nil
}

// L 获取全局日志器
func L() *zap.Logger {
	if globalLogger == nil {
		globalLogger, _ = zap.NewDevelopment()
	}
	return globalLogger
}

// S 获取 SugaredLogger
func S() *zap.SugaredLogger {
	return L().Sugar()
}

// Sync 同步日志
func Sync() error {
	if globalLogger != nil {
		return globalLogger.Sync()
	}
	return nil
}

// Named 创建命名日志器
func Named(name string) *zap.Logger {
	return L().Named(name)
}

// With 添加字段
func With(fields ...zap.Field) *zap.Logger {
	return L().With(fields...)
}

// Debug 记录调试日志
func Debug(msg string, fields ...zap.Field) {
	L().Debug(msg, fields...)
}

// Info 记录信息日志
func Info(msg string, fields ...zap.Field) {
	L().Info(msg, fields...)
}

// Warn 记录警告日志
func Warn(msg string, fields ...zap.Field) {
	L().Warn(msg, fields...)
}

// Error 记录错误日志
func Error(msg string, fields ...zap.Field) {
	L().Error(msg, fields...)
}

// Fatal 记录致命错误日志
func Fatal(msg string, fields ...zap.Field) {
	L().Fatal(msg, fields...)
}

// 常用字段名
const (
	FieldRequestID    = "request_id"
	FieldUserID       = "user_id"
	FieldUsername     = "username"
	FieldPath         = "path"
	FieldMethod       = "method"
	FieldStatus       = "status"
	FieldLatency      = "latency"
	FieldIP           = "ip"
	FieldError        = "error"
	FieldModule       = "module"
	FieldAction       = "action"
	FieldAdvertiserID = "advertiser_id"
	FieldCampaignID   = "campaign_id"
)

// 常用字段构建器
func RequestID(id string) zap.Field {
	return zap.String(FieldRequestID, id)
}

func UserID(id int64) zap.Field {
	return zap.Int64(FieldUserID, id)
}

func Username(name string) zap.Field {
	return zap.String(FieldUsername, name)
}

func Module(name string) zap.Field {
	return zap.String(FieldModule, name)
}

func Action(name string) zap.Field {
	return zap.String(FieldAction, name)
}

func AdvertiserID(id int64) zap.Field {
	return zap.Int64(FieldAdvertiserID, id)
}

func CampaignID(id int64) zap.Field {
	return zap.Int64(FieldCampaignID, id)
}
