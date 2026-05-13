package utils

import (
	"crypto/sha256"
	"fmt"
	"time"
	"os"
	"path/filepath"
	"runtime"
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

// GetProjectRoot 获取当前项目源码根目录（开发/生产都不会错）
func GetProjectRoot() string {
	_, filename, _, _ := runtime.Caller(0)
	// 这里回退到项目根目录，根据你的目录结构调整
	// 假设 utils 包在项目根目录下
	root := filepath.Join(filepath.Dir(filename), "../")
	return root
}

// ReadABIFile 读取 abi 文件，绝对不会错
func ReadABIFile(filename string) ([]byte, error) {
	root := GetProjectRoot()
	path := filepath.Join(root, "abi","pledge", filename)
	return os.ReadFile(path)
}