package database

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"oceanengine-backend/config"
)

var db *gorm.DB

// Init 初始化数据库连接
func Init(cfg *config.DatabaseConfig, logger *zap.Logger) (*gorm.DB, error) {
	// 配置 GORM 日志
	logLevel := gormlogger.Silent
	switch cfg.LogLevel {
	case "error":
		logLevel = gormlogger.Error
	case "warn":
		logLevel = gormlogger.Warn
	case "info":
		logLevel = gormlogger.Info
	}

	gormConfig := &gorm.Config{
		Logger: gormlogger.Default.LogMode(logLevel),
	}

	var err error
	var dbDSN string

	// 根据驱动类型选择连接方式
	if cfg.Driver == "sqlite" {
		// SQLite 连接
		dbDSN = cfg.Database
		if dbDSN == "" {
			dbDSN = "./oceanengine.db"
		}
		db, err = gorm.Open(sqlite.Open(dbDSN), gormConfig)
		if err != nil {
			return nil, fmt.Errorf("failed to connect sqlite database: %w", err)
		}
	} else {
		// 默认使用 MySQL
		dbDSN = cfg.GetDSN()
		db, err = gorm.Open(mysql.Open(dbDSN), gormConfig)
		if err != nil {
			return nil, fmt.Errorf("failed to connect mysql database: %w", err)
		}

		// 获取底层 sql.DB 并配置连接池
		sqlDB, err := db.DB()
		if err != nil {
			return nil, fmt.Errorf("failed to get sql.DB: %w", err)
		}

		sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
		sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
		sqlDB.SetConnMaxLifetime(cfg.ConnMaxLifetime)
	}

	// 测试连接
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql.DB: %w", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	logger.Info("database connected successfully",
		zap.String("driver", cfg.Driver),
		zap.String("database", cfg.Database),
	)

	return db, nil
}

// GetDB 获取数据库实例
func GetDB() *gorm.DB {
	return db
}

// Close 关闭数据库连接
func Close() error {
	if db != nil {
		sqlDB, err := db.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}
	return nil
}

// AutoMigrate 自动迁移
func AutoMigrate(models ...interface{}) error {
	return db.AutoMigrate(models...)
}

// Transaction 事务执行
func Transaction(fn func(tx *gorm.DB) error) error {
	return db.Transaction(fn)
}

// Paginate 分页查询
func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}
		if pageSize <= 0 {
			pageSize = 10
		}
		if pageSize > 100 {
			pageSize = 100
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

// BaseModel 基础模型
type BaseModel struct {
	ID        uint64         `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// BaseModelWithUser 带用户信息的基础模型
type BaseModelWithUser struct {
	BaseModel
	CreatedBy uint64 `gorm:"default:0" json:"created_by"`
	UpdatedBy uint64 `gorm:"default:0" json:"updated_by"`
}
