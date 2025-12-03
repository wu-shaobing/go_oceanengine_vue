package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// Cache 缓存接口
type Cache interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Delete(ctx context.Context, keys ...string) error
	Exists(ctx context.Context, key string) (bool, error)
	Expire(ctx context.Context, key string, expiration time.Duration) error

	// JSON 操作
	GetJSON(ctx context.Context, key string, dest interface{}) error
	SetJSON(ctx context.Context, key string, value interface{}, expiration time.Duration) error

	// Hash 操作
	HGet(ctx context.Context, key, field string) (string, error)
	HSet(ctx context.Context, key string, values ...interface{}) error
	HGetAll(ctx context.Context, key string) (map[string]string, error)
	HDel(ctx context.Context, key string, fields ...string) error

	// List 操作
	LPush(ctx context.Context, key string, values ...interface{}) error
	RPush(ctx context.Context, key string, values ...interface{}) error
	LRange(ctx context.Context, key string, start, stop int64) ([]string, error)
	LLen(ctx context.Context, key string) (int64, error)

	// Set 操作
	SAdd(ctx context.Context, key string, members ...interface{}) error
	SMembers(ctx context.Context, key string) ([]string, error)
	SIsMember(ctx context.Context, key string, member interface{}) (bool, error)
	SRem(ctx context.Context, key string, members ...interface{}) error

	// 分布式锁
	Lock(ctx context.Context, key string, expiration time.Duration) (bool, error)
	Unlock(ctx context.Context, key string) error

	// 计数器
	Incr(ctx context.Context, key string) (int64, error)
	IncrBy(ctx context.Context, key string, value int64) (int64, error)
	Decr(ctx context.Context, key string) (int64, error)
}

// RedisCache Redis 缓存实现
type RedisCache struct {
	client *redis.Client
	prefix string
}

// NewRedisCache 创建 Redis 缓存
func NewRedisCache(client *redis.Client, prefix string) *RedisCache {
	return &RedisCache{
		client: client,
		prefix: prefix,
	}
}

// key 生成带前缀的 key
func (c *RedisCache) key(key string) string {
	if c.prefix == "" {
		return key
	}
	return fmt.Sprintf("%s:%s", c.prefix, key)
}

// Get 获取字符串值
func (c *RedisCache) Get(ctx context.Context, key string) (string, error) {
	return c.client.Get(ctx, c.key(key)).Result()
}

// Set 设置字符串值
func (c *RedisCache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return c.client.Set(ctx, c.key(key), value, expiration).Err()
}

// Delete 删除 key
func (c *RedisCache) Delete(ctx context.Context, keys ...string) error {
	prefixedKeys := make([]string, len(keys))
	for i, key := range keys {
		prefixedKeys[i] = c.key(key)
	}
	return c.client.Del(ctx, prefixedKeys...).Err()
}

// Exists 检查 key 是否存在
func (c *RedisCache) Exists(ctx context.Context, key string) (bool, error) {
	n, err := c.client.Exists(ctx, c.key(key)).Result()
	return n > 0, err
}

// Expire 设置过期时间
func (c *RedisCache) Expire(ctx context.Context, key string, expiration time.Duration) error {
	return c.client.Expire(ctx, c.key(key), expiration).Err()
}

// GetJSON 获取 JSON 值
func (c *RedisCache) GetJSON(ctx context.Context, key string, dest interface{}) error {
	data, err := c.client.Get(ctx, c.key(key)).Bytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(data, dest)
}

// SetJSON 设置 JSON 值
func (c *RedisCache) SetJSON(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.client.Set(ctx, c.key(key), data, expiration).Err()
}

// HGet 获取 Hash 字段
func (c *RedisCache) HGet(ctx context.Context, key, field string) (string, error) {
	return c.client.HGet(ctx, c.key(key), field).Result()
}

// HSet 设置 Hash 字段
func (c *RedisCache) HSet(ctx context.Context, key string, values ...interface{}) error {
	return c.client.HSet(ctx, c.key(key), values...).Err()
}

// HGetAll 获取 Hash 所有字段
func (c *RedisCache) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	return c.client.HGetAll(ctx, c.key(key)).Result()
}

// HDel 删除 Hash 字段
func (c *RedisCache) HDel(ctx context.Context, key string, fields ...string) error {
	return c.client.HDel(ctx, c.key(key), fields...).Err()
}

// LPush 左侧插入列表
func (c *RedisCache) LPush(ctx context.Context, key string, values ...interface{}) error {
	return c.client.LPush(ctx, c.key(key), values...).Err()
}

// RPush 右侧插入列表
func (c *RedisCache) RPush(ctx context.Context, key string, values ...interface{}) error {
	return c.client.RPush(ctx, c.key(key), values...).Err()
}

// LRange 获取列表范围
func (c *RedisCache) LRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	return c.client.LRange(ctx, c.key(key), start, stop).Result()
}

// LLen 获取列表长度
func (c *RedisCache) LLen(ctx context.Context, key string) (int64, error) {
	return c.client.LLen(ctx, c.key(key)).Result()
}

// SAdd 添加集合成员
func (c *RedisCache) SAdd(ctx context.Context, key string, members ...interface{}) error {
	return c.client.SAdd(ctx, c.key(key), members...).Err()
}

// SMembers 获取集合所有成员
func (c *RedisCache) SMembers(ctx context.Context, key string) ([]string, error) {
	return c.client.SMembers(ctx, c.key(key)).Result()
}

// SIsMember 检查是否是集合成员
func (c *RedisCache) SIsMember(ctx context.Context, key string, member interface{}) (bool, error) {
	return c.client.SIsMember(ctx, c.key(key), member).Result()
}

// SRem 移除集合成员
func (c *RedisCache) SRem(ctx context.Context, key string, members ...interface{}) error {
	return c.client.SRem(ctx, c.key(key), members...).Err()
}

// Lock 获取分布式锁
func (c *RedisCache) Lock(ctx context.Context, key string, expiration time.Duration) (bool, error) {
	return c.client.SetNX(ctx, c.key("lock:"+key), "1", expiration).Result()
}

// Unlock 释放分布式锁
func (c *RedisCache) Unlock(ctx context.Context, key string) error {
	return c.client.Del(ctx, c.key("lock:"+key)).Err()
}

// Incr 自增
func (c *RedisCache) Incr(ctx context.Context, key string) (int64, error) {
	return c.client.Incr(ctx, c.key(key)).Result()
}

// IncrBy 自增指定值
func (c *RedisCache) IncrBy(ctx context.Context, key string, value int64) (int64, error) {
	return c.client.IncrBy(ctx, c.key(key), value).Result()
}

// Decr 自减
func (c *RedisCache) Decr(ctx context.Context, key string) (int64, error) {
	return c.client.Decr(ctx, c.key(key)).Result()
}
