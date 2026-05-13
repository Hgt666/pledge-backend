package main

import (
	"easy-swap/config"
	"easy-swap/contract"
	"easy-swap/dal"
	_ "easy-swap/docs" // 👈 你的项目名/docs
	"easy-swap/logger"
	"easy-swap/router"

	"github.com/swaggo/files" // 👈 正确写法
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @title           NFT Market API
// @version         1.0
// @description     NFT 交易平台接口文档
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1
func main() {
	logger.InitZap()
	defer logger.Log.Sync()

	config.InitViper()
	dal.InitDB()

	r := gin.Default()
	// Swagger 接口文档路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Register(r)
	contract.InitContract()

	port := ":" + config.GlobalConfig.App.Port
	logger.Log.Info("API服务启动", zap.String("port", port))
	_ = r.Run(port)
}
