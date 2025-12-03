# 缓存设计

## 概述

本系统使用 Redis 作为缓存层，用于提升系统性能、减少数据库压力，以及实现分布式锁、限流等功能。

## 缓存策略

### 缓存层次

```
┌─────────────────┐
│   Local Cache   │  <- 本地内存缓存（高频热点）
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│   Redis Cache   │  <- 分布式缓存（共享数据）
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│    Database     │  <- 持久化存储
└─────────────────┘
```

### 缓存键命名规范

```
格式: {模块}:{类型}:{标识}

示例:
- user:info:123          用户信息
- advertiser:info:456    广告主信息
- campaign:list:789      广告系列列表
- token:access:xxx       访问令牌
- lock:sync:123          同步锁
- rate:limit:ip          限流计数
```

---

## Redis 客户端封装

```go
// pkg/cache/redis.go
package cache

import (
    "context"
    "encoding/json"
    "time"
    
    "github.com/redis/go-redis/v9"
)

// RedisClient Redis 客户端封装
type RedisClient struct {
    client *redis.Client
}

// Config Redis 配置
type Config struct {
    Addr         string        `yaml:"addr"`
    Password     string        `yaml:"password"`
    DB           int           `yaml:"db"`
    PoolSize     int           `yaml:"pool_size"`
    MinIdleConns int           `yaml:"min_idle_conns"`
    DialTimeout  time.Duration `yaml:"dial_timeout"`
    ReadTimeout  time.Duration `yaml:"read_timeout"`
    WriteTimeout time.Duration `yaml:"write_timeout"`
}

func NewRedisClient(cfg *Config) (*RedisClient, error) {
    client := redis.NewClient(&redis.Options{
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
    
    if err := client.Ping(ctx).Err(); err != nil {
        return nil, err
    }
    
    return &RedisClient{client: client}, nil
}

// Get 获取缓存
func (r *RedisClient) Get(ctx context.Context, key string) (string, error) {
    return r.client.Get(ctx, key).Result()
}

// GetObject 获取对象
func (r *RedisClient) GetObject(ctx context.Context, key string, dest interface{}) error {
    val, err := r.client.Get(ctx, key).Result()
    if err != nil {
        return err
    }
    return json.Unmarshal([]byte(val), dest)
}

// Set 设置缓存
func (r *RedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
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
func (r *RedisClient) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
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
func (r *RedisClient) Delete(ctx context.Context, keys ...string) error {
    return r.client.Del(ctx, keys...).Err()
}

// Exists 检查是否存在
func (r *RedisClient) Exists(ctx context.Context, key string) (bool, error) {
    n, err := r.client.Exists(ctx, key).Result()
    return n > 0, err
}

// Expire 设置过期时间
func (r *RedisClient) Expire(ctx context.Context, key string, expiration time.Duration) error {
    return r.client.Expire(ctx, key, expiration).Err()
}

// Incr 自增
func (r *RedisClient) Incr(ctx context.Context, key string) (int64, error) {
    return r.client.Incr(ctx, key).Result()
}

// HGet 获取 Hash 字段
func (r *RedisClient) HGet(ctx context.Context, key, field string) (string, error) {
    return r.client.HGet(ctx, key, field).Result()
}

// HSet 设置 Hash 字段
func (r *RedisClient) HSet(ctx context.Context, key string, values ...interface{}) error {
    return r.client.HSet(ctx, key, values...).Err()
}

// HGetAll 获取所有 Hash 字段
func (r *RedisClient) HGetAll(ctx context.Context, key string) (map[string]string, error) {
    return r.client.HGetAll(ctx, key).Result()
}

// Client 获取原始客户端
func (r *RedisClient) Client() *redis.Client {
    return r.client
}
```

---

## 缓存管理器

```go
// pkg/cache/manager.go
package cache

import (
    "context"
    "fmt"
    "time"
)

// CacheManager 缓存管理器
type CacheManager struct {
    redis  *RedisClient
    local  *LocalCache
    prefix string
}

func NewCacheManager(redis *RedisClient, prefix string) *CacheManager {
    return &CacheManager{
        redis:  redis,
        local:  NewLocalCache(10000, 5*time.Minute),
        prefix: prefix,
    }
}

// key 生成完整键名
func (m *CacheManager) key(parts ...string) string {
    key := m.prefix
    for _, part := range parts {
        key += ":" + part
    }
    return key
}

// Get 获取缓存（先本地后 Redis）
func (m *CacheManager) Get(ctx context.Context, key string, dest interface{}) error {
    fullKey := m.key(key)
    
    // 1. 尝试本地缓存
    if val, ok := m.local.Get(fullKey); ok {
        if err := json.Unmarshal(val.([]byte), dest); err == nil {
            return nil
        }
    }
    
    // 2. 从 Redis 获取
    err := m.redis.GetObject(ctx, fullKey, dest)
    if err != nil {
        return err
    }
    
    // 3. 写入本地缓存
    bytes, _ := json.Marshal(dest)
    m.local.Set(fullKey, bytes)
    
    return nil
}

// Set 设置缓存
func (m *CacheManager) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
    fullKey := m.key(key)
    
    // 写入 Redis
    if err := m.redis.Set(ctx, fullKey, value, ttl); err != nil {
        return err
    }
    
    // 写入本地缓存
    bytes, _ := json.Marshal(value)
    m.local.SetWithTTL(fullKey, bytes, ttl)
    
    return nil
}

// Delete 删除缓存
func (m *CacheManager) Delete(ctx context.Context, keys ...string) error {
    fullKeys := make([]string, len(keys))
    for i, key := range keys {
        fullKeys[i] = m.key(key)
        m.local.Delete(fullKeys[i])
    }
    return m.redis.Delete(ctx, fullKeys...)
}

// GetOrSet 获取或设置缓存
func (m *CacheManager) GetOrSet(ctx context.Context, key string, dest interface{}, ttl time.Duration, loader func() (interface{}, error)) error {
    // 尝试获取
    if err := m.Get(ctx, key, dest); err == nil {
        return nil
    }
    
    // 加载数据
    data, err := loader()
    if err != nil {
        return err
    }
    
    // 设置缓存
    if err := m.Set(ctx, key, data, ttl); err != nil {
        return err
    }
    
    // 复制到目标
    bytes, _ := json.Marshal(data)
    return json.Unmarshal(bytes, dest)
}

// InvalidatePattern 按模式删除缓存
func (m *CacheManager) InvalidatePattern(ctx context.Context, pattern string) error {
    fullPattern := m.key(pattern)
    
    // 使用 SCAN 避免阻塞
    var cursor uint64
    for {
        keys, nextCursor, err := m.redis.client.Scan(ctx, cursor, fullPattern, 100).Result()
        if err != nil {
            return err
        }
        
        if len(keys) > 0 {
            m.redis.Delete(ctx, keys...)
            for _, key := range keys {
                m.local.Delete(key)
            }
        }
        
        cursor = nextCursor
        if cursor == 0 {
            break
        }
    }
    
    return nil
}
```

---

## 本地缓存

```go
// pkg/cache/local.go
package cache

import (
    "sync"
    "time"
)

// LocalCache 本地内存缓存
type LocalCache struct {
    data       map[string]*cacheItem
    mu         sync.RWMutex
    maxSize    int
    defaultTTL time.Duration
}

type cacheItem struct {
    value    interface{}
    expireAt time.Time
}

func NewLocalCache(maxSize int, defaultTTL time.Duration) *LocalCache {
    cache := &LocalCache{
        data:       make(map[string]*cacheItem),
        maxSize:    maxSize,
        defaultTTL: defaultTTL,
    }
    
    // 启动清理协程
    go cache.cleanup()
    
    return cache
}

// Get 获取缓存
func (c *LocalCache) Get(key string) (interface{}, bool) {
    c.mu.RLock()
    defer c.mu.RUnlock()
    
    item, ok := c.data[key]
    if !ok {
        return nil, false
    }
    
    if time.Now().After(item.expireAt) {
        return nil, false
    }
    
    return item.value, true
}

// Set 设置缓存
func (c *LocalCache) Set(key string, value interface{}) {
    c.SetWithTTL(key, value, c.defaultTTL)
}

// SetWithTTL 设置缓存（带 TTL）
func (c *LocalCache) SetWithTTL(key string, value interface{}, ttl time.Duration) {
    c.mu.Lock()
    defer c.mu.Unlock()
    
    // 检查容量
    if len(c.data) >= c.maxSize {
        c.evict()
    }
    
    c.data[key] = &cacheItem{
        value:    value,
        expireAt: time.Now().Add(ttl),
    }
}

// Delete 删除缓存
func (c *LocalCache) Delete(key string) {
    c.mu.Lock()
    defer c.mu.Unlock()
    delete(c.data, key)
}

// evict 淘汰过期条目
func (c *LocalCache) evict() {
    now := time.Now()
    for key, item := range c.data {
        if now.After(item.expireAt) {
            delete(c.data, key)
        }
    }
}

// cleanup 定期清理
func (c *LocalCache) cleanup() {
    ticker := time.NewTicker(time.Minute)
    for range ticker.C {
        c.mu.Lock()
        c.evict()
        c.mu.Unlock()
    }
}
```

---

## 分布式锁

```go
// pkg/cache/lock.go
package cache

import (
    "context"
    "errors"
    "time"
    
    "github.com/google/uuid"
)

var (
    ErrLockFailed  = errors.New("failed to acquire lock")
    ErrLockExpired = errors.New("lock expired")
)

// DistributedLock 分布式锁
type DistributedLock struct {
    redis   *RedisClient
    key     string
    token   string
    ttl     time.Duration
    stopCh  chan struct{}
}

// NewLock 创建分布式锁
func (r *RedisClient) NewLock(key string, ttl time.Duration) *DistributedLock {
    return &DistributedLock{
        redis:  r,
        key:    "lock:" + key,
        token:  uuid.New().String(),
        ttl:    ttl,
        stopCh: make(chan struct{}),
    }
}

// Lock 加锁
func (l *DistributedLock) Lock(ctx context.Context) error {
    ok, err := l.redis.SetNX(ctx, l.key, l.token, l.ttl)
    if err != nil {
        return err
    }
    if !ok {
        return ErrLockFailed
    }
    
    // 启动续期协程
    go l.keepAlive(ctx)
    
    return nil
}

// TryLock 尝试加锁
func (l *DistributedLock) TryLock(ctx context.Context, timeout time.Duration) error {
    deadline := time.Now().Add(timeout)
    
    for time.Now().Before(deadline) {
        ok, err := l.redis.SetNX(ctx, l.key, l.token, l.ttl)
        if err != nil {
            return err
        }
        if ok {
            go l.keepAlive(ctx)
            return nil
        }
        
        select {
        case <-ctx.Done():
            return ctx.Err()
        case <-time.After(100 * time.Millisecond):
        }
    }
    
    return ErrLockFailed
}

// Unlock 解锁
func (l *DistributedLock) Unlock(ctx context.Context) error {
    close(l.stopCh)
    
    // 使用 Lua 脚本保证原子性
    script := `
        if redis.call("get", KEYS[1]) == ARGV[1] then
            return redis.call("del", KEYS[1])
        else
            return 0
        end
    `
    
    _, err := l.redis.client.Eval(ctx, script, []string{l.key}, l.token).Result()
    return err
}

// keepAlive 续期
func (l *DistributedLock) keepAlive(ctx context.Context) {
    ticker := time.NewTicker(l.ttl / 3)
    defer ticker.Stop()
    
    for {
        select {
        case <-l.stopCh:
            return
        case <-ctx.Done():
            return
        case <-ticker.C:
            // 续期
            script := `
                if redis.call("get", KEYS[1]) == ARGV[1] then
                    return redis.call("expire", KEYS[1], ARGV[2])
                else
                    return 0
                end
            `
            l.redis.client.Eval(ctx, script, []string{l.key}, l.token, int(l.ttl.Seconds()))
        }
    }
}

// WithLock 使用锁执行函数
func (r *RedisClient) WithLock(ctx context.Context, key string, ttl time.Duration, fn func() error) error {
    lock := r.NewLock(key, ttl)
    
    if err := lock.Lock(ctx); err != nil {
        return err
    }
    defer lock.Unlock(ctx)
    
    return fn()
}
```

---

## 缓存使用示例

### 广告主缓存

```go
// internal/app/advertiser/cache/advertiser.go
package cache

const (
    advertiserInfoKey   = "advertiser:info:%d"
    advertiserListKey   = "advertiser:list"
    advertiserInfoTTL   = 10 * time.Minute
    advertiserListTTL   = 5 * time.Minute
)

type AdvertiserCache struct {
    manager *cache.CacheManager
}

func NewAdvertiserCache(manager *cache.CacheManager) *AdvertiserCache {
    return &AdvertiserCache{manager: manager}
}

// GetInfo 获取广告主信息
func (c *AdvertiserCache) GetInfo(ctx context.Context, id int64) (*model.Advertiser, error) {
    key := fmt.Sprintf(advertiserInfoKey, id)
    
    var advertiser model.Advertiser
    err := c.manager.Get(ctx, key, &advertiser)
    if err != nil {
        return nil, err
    }
    
    return &advertiser, nil
}

// SetInfo 设置广告主信息
func (c *AdvertiserCache) SetInfo(ctx context.Context, advertiser *model.Advertiser) error {
    key := fmt.Sprintf(advertiserInfoKey, advertiser.ID)
    return c.manager.Set(ctx, key, advertiser, advertiserInfoTTL)
}

// DeleteInfo 删除广告主缓存
func (c *AdvertiserCache) DeleteInfo(ctx context.Context, id int64) error {
    key := fmt.Sprintf(advertiserInfoKey, id)
    return c.manager.Delete(ctx, key)
}

// GetOrLoadInfo 获取或加载广告主信息
func (c *AdvertiserCache) GetOrLoadInfo(ctx context.Context, id int64, loader func() (*model.Advertiser, error)) (*model.Advertiser, error) {
    key := fmt.Sprintf(advertiserInfoKey, id)
    
    var advertiser model.Advertiser
    err := c.manager.GetOrSet(ctx, key, &advertiser, advertiserInfoTTL, func() (interface{}, error) {
        return loader()
    })
    if err != nil {
        return nil, err
    }
    
    return &advertiser, nil
}
```

### 在 Service 中使用缓存

```go
// internal/app/advertiser/service/advertiser.go
package service

type AdvertiserService struct {
    repo  repository.AdvertiserRepository
    cache *cache.AdvertiserCache
}

func (s *AdvertiserService) GetByID(ctx context.Context, id int64) (*dto.AdvertiserResp, error) {
    // 使用缓存
    advertiser, err := s.cache.GetOrLoadInfo(ctx, id, func() (*model.Advertiser, error) {
        return s.repo.GetByID(ctx, id)
    })
    if err != nil {
        return nil, err
    }
    
    return dto.ToAdvertiserResp(advertiser), nil
}

func (s *AdvertiserService) Update(ctx context.Context, id int64, req *dto.AdvertiserUpdateReq) error {
    // 更新数据库
    if err := s.repo.Update(ctx, id, req); err != nil {
        return err
    }
    
    // 删除缓存
    return s.cache.DeleteInfo(ctx, id)
}
```

---

## 缓存预热

```go
// pkg/cache/warmup.go
package cache

import (
    "context"
    "sync"
)

// Warmer 缓存预热器
type Warmer struct {
    tasks []WarmupTask
}

// WarmupTask 预热任务
type WarmupTask struct {
    Name   string
    Loader func(ctx context.Context) error
}

func NewWarmer() *Warmer {
    return &Warmer{}
}

// Register 注册预热任务
func (w *Warmer) Register(name string, loader func(ctx context.Context) error) {
    w.tasks = append(w.tasks, WarmupTask{Name: name, Loader: loader})
}

// Warmup 执行预热
func (w *Warmer) Warmup(ctx context.Context, concurrency int) error {
    sem := make(chan struct{}, concurrency)
    var wg sync.WaitGroup
    var mu sync.Mutex
    var errors []error
    
    for _, task := range w.tasks {
        wg.Add(1)
        go func(t WarmupTask) {
            defer wg.Done()
            sem <- struct{}{}
            defer func() { <-sem }()
            
            if err := t.Loader(ctx); err != nil {
                mu.Lock()
                errors = append(errors, fmt.Errorf("%s: %w", t.Name, err))
                mu.Unlock()
            }
        }(task)
    }
    
    wg.Wait()
    
    if len(errors) > 0 {
        return fmt.Errorf("warmup errors: %v", errors)
    }
    return nil
}
```

### 使用示例

```go
func main() {
    // 创建预热器
    warmer := cache.NewWarmer()
    
    // 注册广告主预热
    warmer.Register("advertisers", func(ctx context.Context) error {
        advertisers, err := advertiserRepo.ListAll(ctx)
        if err != nil {
            return err
        }
        for _, adv := range advertisers {
            advertiserCache.SetInfo(ctx, adv)
        }
        return nil
    })
    
    // 注册配置预热
    warmer.Register("configs", func(ctx context.Context) error {
        configs, err := configRepo.ListAll(ctx)
        if err != nil {
            return err
        }
        for _, cfg := range configs {
            configCache.Set(ctx, cfg.Key, cfg.Value)
        }
        return nil
    })
    
    // 执行预热
    if err := warmer.Warmup(context.Background(), 5); err != nil {
        log.Printf("cache warmup failed: %v", err)
    }
}
```

---

## 缓存监控

```go
// pkg/cache/metrics.go
package cache

import (
    "github.com/prometheus/client_golang/prometheus"
)

var (
    cacheHits = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "cache_hits_total",
            Help: "Total number of cache hits",
        },
        []string{"cache_type", "key_pattern"},
    )
    
    cacheMisses = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "cache_misses_total",
            Help: "Total number of cache misses",
        },
        []string{"cache_type", "key_pattern"},
    )
    
    cacheLatency = prometheus.NewHistogramVec(
        prometheus.HistogramOpts{
            Name:    "cache_operation_duration_seconds",
            Help:    "Cache operation latency",
            Buckets: []float64{.001, .005, .01, .025, .05, .1, .25, .5, 1},
        },
        []string{"operation"},
    )
)

func init() {
    prometheus.MustRegister(cacheHits, cacheMisses, cacheLatency)
}
```
