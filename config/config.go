package config

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
)

// 全局配置结构体
type AppConfig struct {
	App      App      `yaml:"app"`
	MySQL    MySQL    `yaml:"mysql"`
	Chain    Chain    `yaml:"chain"`
	Contract Contract `yaml:"contract"`
}

type App struct {
	Name string `yaml:"name"`
	Port string `yaml:"port"`
}

type MySQL struct {
	DSN string `yaml:"dsn"`
}

type Chain struct {
	RPC          string `yaml:"rpc"`
	SafeBlock    uint64 `yaml:"safe_block"`
	ScanInterval int    `yaml:"scan_interval"`
	ChainID      uint64  `yaml:"chain_id"`
}

type Contract struct {
	Pledge string `yaml:"pledge"`
}

// 全局配置变量
var GlobalConfig *AppConfig
var (
	PledgeContractAddr common.Address
)

// InitViper 初始化viper加载yaml配置
func InitViper() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("读取配置文件失败: ", err)
	}

	// 解析配置到结构体
	var cfg AppConfig
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatal("解析配置文件失败: ", err)
	}
	GlobalConfig = &cfg

	// 初始化合约地址
	PledgeContractAddr = common.HexToAddress(cfg.Contract.Pledge)
}
