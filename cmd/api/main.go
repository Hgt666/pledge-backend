package main

import (
	"easy-swap/config"
	"easy-swap/dal"
	"easy-swap/logger"
	"easy-swap/router"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	logger.InitZap()
	defer logger.Log.Sync()

	config.InitViper()
	dal.InitDB()

	r := gin.Default()
	router.Register(r)

	port := ":" + config.GlobalConfig.App.Port
	logger.Log.Info("API服务启动", zap.String("port", port))
	_ = r.Run(port)
}