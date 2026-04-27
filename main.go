package main

import (
	"easy-swap/config"
	"easy-swap/dal"
	"easy-swap/logger"   // 新增
	"easy-swap/internal/parser"
	"easy-swap/internal/scan"
	"easy-swap/router"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化日志
	logger.InitZap()
	defer logger.Log.Sync()

	logger.Log.Info("🚀 服务启动开始")

	// 配置
	config.InitViper()

	// DB
	dal.InitDB()

	// ABI
	if err := parser.Init(); err != nil {
		logger.Log.Fatal("parser 初始化失败", zap.Error(err))
	}

	// 扫链
	scanner, err := scan.NewScanner()
	if err != nil {
		logger.Log.Fatal("扫链服务启动失败", zap.Error(err))
	}
	go scanner.Start()

	// 路由
	r := gin.Default()
	router.Register(r)

	port := ":" + config.GlobalConfig.App.Port
	logger.Log.Info("服务启动成功", zap.String("port", port))

	if err := r.Run(port); err != nil {
		logger.Log.Fatal("服务启动失败", zap.Error(err))
	}
}