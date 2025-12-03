package oauth

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	// StateKeyPrefix Redis key 前缀
	StateKeyPrefix = "oauth:state:"
	// StateExpiration State 过期时间 (10分钟)
	StateExpiration = 10 * time.Minute
)

// StateManager OAuth State 管理器
type StateManager struct {
	redis *redis.Client
}

// NewStateManager 创建 State 管理器
func NewStateManager(redisClient *redis.Client) *StateManager {
	return &StateManager{
		redis: redisClient,
	}
}

// GenerateState 生成随机 state
func (m *StateManager) GenerateState() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("failed to generate state: %w", err)
	}
	return hex.EncodeToString(bytes), nil
}

// SaveState 保存 state 到 Redis
func (m *StateManager) SaveState(ctx context.Context, state string, data map[string]string) error {
	key := StateKeyPrefix + state

	// 存储关联数据 (如用户ID、重定向URL等)
	if data == nil {
		data = make(map[string]string)
	}
	data["created_at"] = time.Now().Format(time.RFC3339)

	pipe := m.redis.Pipeline()
	for k, v := range data {
		pipe.HSet(ctx, key, k, v)
	}
	pipe.Expire(ctx, key, StateExpiration)

	_, err := pipe.Exec(ctx)
	return err
}

// ValidateState 验证并消费 state (一次性使用)
func (m *StateManager) ValidateState(ctx context.Context, state string) (map[string]string, error) {
	key := StateKeyPrefix + state

	// 获取 state 数据
	data, err := m.redis.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get state: %w", err)
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("invalid or expired state")
	}

	// 删除 state (一次性使用)
	if err := m.redis.Del(ctx, key).Err(); err != nil {
		// 只记录日志，不影响主流程
		fmt.Printf("failed to delete state: %v\n", err)
	}

	return data, nil
}

// ValidateStateOnly 仅验证 state 是否有效 (不删除)
func (m *StateManager) ValidateStateOnly(ctx context.Context, state string) (bool, error) {
	key := StateKeyPrefix + state
	exists, err := m.redis.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return exists > 0, nil
}

// DeleteState 手动删除 state
func (m *StateManager) DeleteState(ctx context.Context, state string) error {
	key := StateKeyPrefix + state
	return m.redis.Del(ctx, key).Err()
}

// GenerateAndSave 生成并保存 state
func (m *StateManager) GenerateAndSave(ctx context.Context, data map[string]string) (string, error) {
	state, err := m.GenerateState()
	if err != nil {
		return "", err
	}

	if err := m.SaveState(ctx, state, data); err != nil {
		return "", err
	}

	return state, nil
}
