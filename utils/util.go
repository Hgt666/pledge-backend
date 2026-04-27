package utils

import (
	"crypto/sha256"
	"fmt"
	"time"
)

// 全局超时配置
const (
	RPCTimeout    = 8 * time.Second
	DBOperateTime = 5 * time.Second
)

// HashString 字符串sha256哈希
func HashString(s string) string {
	hash := sha256.Sum256([]byte(s))
	return fmt.Sprintf("%x", hash)
}

// CheckError 统一错误判断
func CheckError(err error, tip string) bool {
	if err != nil {
		fmt.Printf("[ERROR] %s: %v \n", tip, err)
		return true
	}
	return false
}