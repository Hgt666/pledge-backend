package config

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
	"log"
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
}

type Contract struct {
	NFT    string `yaml:"nft"`
	Market string `yaml:"market"`
}

// 全局配置变量
var GlobalConfig *AppConfig
var (
	NftContractAddr    common.Address
	MarketContractAddr common.Address
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
	NftContractAddr = common.HexToAddress(cfg.Contract.NFT)
	MarketContractAddr = common.HexToAddress(cfg.Contract.Market)
}