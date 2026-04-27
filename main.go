package main

import (
	"easy-swap/config"
	"easy-swap/dal"
	"easy-swap/internal/parser"
	"easy-swap/internal/scan"
	"easy-swap/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化Viper读取yaml配置
	config.InitViper()

	// 初始化数据库连接
	dal.InitDB()

	// 初始化合约ABI事件解析器
	if err := parser.Init(); err != nil {
		log.Fatal("ABI解析器初始化失败: ", err)
	}

	// 初始化并启动后台扫链服务
	scanner, err := scan.NewScanner()
	if err != nil {
		log.Fatal("扫链服务初始化失败: ", err)
	}
	go scanner.Start()

	// 初始化gin路由
	r := gin.Default()
	router.Register(r)

	// 启动HTTP服务
	addr := ":" + config.GlobalConfig.App.Port
	log.Printf("🚀 服务启动成功，监听端口 %s", addr)
	_ = r.Run(addr)
}