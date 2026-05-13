package main

import (
	"context"
	"easy-swap/config"
	"easy-swap/dal"
	"easy-swap/internal/parser"
	"easy-swap/internal/scan"
	"easy-swap/logger"
	task "easy-swap/tasks"
	"easy-swap/utils"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func main() {
	// 初始化日志
	logger.InitZap()
	defer logger.Log.Sync()
	
	// 初始化配置
	config.InitViper()
	dal.InitDB()

	// 初始化redis
	utils.InitRedis("127.0.0.1:6379","",0)

	// 初始化 abi解析器
	if err := parser.Init();err != nil {
		logger.Log.Fatal("初始化 abi解析器失败",zap.Error(err))
	}

	// 初始化扫链服务
	scanner,err := scan.NewScanner()
	if err != nil {
		logger.Log.Fatal("创建扫链实例失败",zap.Error(err))
	}

	// 创建全局上下文，控制整个服务生命周期
	ctx,cancel := context.WithCancel(context.Background())
	defer cancel()

	// 启动定时任务 
	go func() {
		task.StartCalcJob()
	}()


	// 启动扫链服务,使用 ctx
	go func() {
		if err := scanner.Start(ctx);err != nil {
			logger.Log.Error("扫链服务退出",zap.Error(err))
			cancel()
		}
	}()
	// logger.Log.Info("✅ 扫链服务启动成功")

	quit := make(chan os.Signal,1)
	signal.Notify(quit,syscall.SIGINT,syscall.SIGTERM)

	// 循环监听信息
	for {
		select {
		case <-quit:
			logger.Log.Info("收到退出信息,准备退出...")
			// 通知所有协和退出
			cancel()
			// 等待协和处理完剩余数据
			time.Sleep(1 *time.Second)
			logger.Log.Info("✅ 扫链服务已安全退出")
			return
		case <-ctx.Done():
			// 上下文关闭（服务异常）
			logger.Log.Info("❌ 服务异常退出")
			return
		}
	}


}