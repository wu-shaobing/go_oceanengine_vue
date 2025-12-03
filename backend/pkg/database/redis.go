package database

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"oceanengine-backend/config"
)

var rdb *redis.Client

// InitRedis 初始化 Redis 连接
func InitRedis(cfg *config.RedisConfig, logger *zap.Logger) (*redis.Client, error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:         cfg.Addr,
		Password:     cfg.Password,
		DB:           cfg.DB,
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
		DialTimeout:  cfg.DialTimeout,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	})

	// 测试连接
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect redis: %w", err)
	}

	logger.Info("redis connected successfully",
		zap.String("addr", cfg.Addr),
		zap.Int("db", cfg.DB),
	)

	return rdb, nil
}

// GetRedis 获取 Redis 客户端
func GetRedis() *redis.Client {
	return rdb
}

// CloseRedis 关闭 Redis 连接
func CloseRedis() error {
	if rdb != nil {
		return rdb.Close()
	}
	return nil
}

// RedisCache Redis 缓存操作封装
type RedisCache struct {
	client *redis.Client
}

// NewRedisCache 创建 Redis 缓存实例
func NewRedisCache(client *redis.Client) *RedisCache {
	return &RedisCache{client: client}
}

// Get 获取缓存
func (r *RedisCache) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

// GetObject 获取对象
func (r *RedisCache) GetObject(ctx context.Context, key string, dest interface{}) error {
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(val), dest)
}

// Set 设置缓存
func (r *RedisCache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	var val string
	switch v := value.(type) {
	case string:
		val = v
	case []byte:
		val = string(v)
	default:
		bytes, err := json.Marshal(value)
		if err != nil {
			return err
		}
		val = string(bytes)
	}
	return r.client.Set(ctx, key, val, expiration).Err()
}

// SetNX 设置缓存（不存在时）
func (r *RedisCache) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	var val string
	switch v := value.(type) {
	case string:
		val = v
	default:
		bytes, err := json.Marshal(value)
		if err != nil {
			return false, err
		}
		val = string(bytes)
	}
	return r.client.SetNX(ctx, key, val, expiration).Result()
}

// Delete 删除缓存
func (r *RedisCache) Delete(ctx context.Context, keys ...string) error {
	return r.client.Del(ctx, keys...).Err()
}

// Exists 检查是否存在
func (r *RedisCache) Exists(ctx context.Context, key string) (bool, error) {
	n, err := r.client.Exists(ctx, key).Result()
	return n > 0, err
}

// Expire 设置过期时间
func (r *RedisCache) Expire(ctx context.Context, key string, expiration time.Duration) error {
	return r.client.Expire(ctx, key, expiration).Err()
}

// Incr 自增
func (r *RedisCache) Incr(ctx context.Context, key string) (int64, error) {
	return r.client.Incr(ctx, key).Result()
}

// HGet 获取 Hash 字段
func (r *RedisCache) HGet(ctx context.Context, key, field string) (string, error) {
	return r.client.HGet(ctx, key, field).Result()
}

// HSet 设置 Hash 字段
func (r *RedisCache) HSet(ctx context.Context, key string, values ...interface{}) error {
	return r.client.HSet(ctx, key, values...).Err()
}

// HGetAll 获取所有 Hash 字段
func (r *RedisCache) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	return r.client.HGetAll(ctx, key).Result()
}

// HDel 删除 Hash 字段
func (r *RedisCache) HDel(ctx context.Context, key string, fields ...string) error {
	return r.client.HDel(ctx, key, fields...).Err()
}

// Client 获取原始客户端
func (r *RedisCache) Client() *redis.Client {
	return r.client
}
