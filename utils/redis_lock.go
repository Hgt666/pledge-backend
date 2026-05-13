package utils

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"time"

	"github.com/go-redis/redis/v8"
)

// RedisLock 分布式锁
type RedisLock struct {
	client    *redis.Client
	key       string
	value     string // 唯一值，防止误删
	expire    time.Duration
}

// NewRedisLock 创建锁
func NewRedisLock(client *redis.Client, key string, expire time.Duration) *RedisLock {
	return &RedisLock{
		client: client,
		key:    key,
		value:  uniqueID(), // 生成唯一ID
		expire: expire,
	}
}

// Lock 加锁
func (l *RedisLock) Lock() (bool, error) {
	return l.client.SetNX(context.Background(), l.key, l.value, l.expire).Result()
}

// Unlock 解锁（LUA脚本保证原子性）
func (l *RedisLock) Unlock() error {
	script := `
	if redis.call("get", KEYS[1]) == ARGV[1] then
		return redis.call("del", KEYS[1])
	else
		return 0
	end
`
	_, err := l.client.Eval(context.Background(), script, []string{l.key}, l.value).Result()
	return err
}

// 生成唯一ID
func uniqueID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}