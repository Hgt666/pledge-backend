package main

import (
	"easy-swap/config"
	"easy-swap/dal"
	"easy-swap/internal/parser"
	"easy-swap/internal/scan"
	"easy-swap/logger"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

func main() {
	logger.InitZap()
	defer logger.Log.Sync()

	config.InitViper()
	dal.InitDB()

	if err := parser.Init(); err != nil {
		logger.Log.Fatal("parser init failed", zap.Error(err))
	}

	scanner, err := scan.NewScanner()
	if err != nil {
		logger.Log.Fatal("scan init failed", zap.Error(err))
	}

	go scanner.Start()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Log.Info("扫链服务退出")
}