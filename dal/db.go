package dal

import (
	"easy-swap/config"
	"easy-swap/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB 全局数据库实例
var DB *gorm.DB

// InitDB 初始化mysql连接
func InitDB() {
	dsn := config.GlobalConfig.MySQL.DSN
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("MySQL连接失败: ", err)
	}

	// 自动创建数据表
	err = db.AutoMigrate(
		&model.NftTransfer{},
		&model.MarketOrder{},
		&model.ScanHeight{},
	)
	if err != nil {
		log.Fatal("数据表迁移失败: ", err)
	}

	DB = db
	log.Println("✅ 数据库初始化完成")
}