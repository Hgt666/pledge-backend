package utils

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

// 全局 Redis 客户端（整个项目用这一个）
var Rdb *redis.Client

// InitRedis 初始化 Redis
func InitRedis(addr string, password string, db int) {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     addr,     // 127.0.0.1:6379
		Password: password, // 没有密码就空字符串
		DB:       db,       // 0
		PoolSize: 10,       // 连接池

		DialTimeout:  5 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	})

	// 测试连接
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := Rdb.Ping(ctx).Result()
	if err != nil {
		panic("redis 连接失败: " + err.Error())
	}

	println("✅ redis 连接成功")
}