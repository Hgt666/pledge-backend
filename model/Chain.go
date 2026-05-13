package model

import "gorm.io/gorm"

type Chain struct {
	gorm.Model
	ChainID   uint64 `json:"chain_id"  gorm:"column:chain_id;comment:链ID;uniqueIndex"`
	ChainName string `json:"chain_name"  gorm:"column:chain_name;comment:链名称"`
	Symbol    string `json:"symbol"  gorm:"column:symbol;comment:链符号"`
	RpcUrl    string `json:"rpc_url"  gorm:"column:rpc_url;comment:RPC地址"`
	Exporer   string `json:"exporer"  gorm:"column:exporer;comment:区块浏览器"`
	Status    uint   `json:"status"  gorm:"column:status;comment:状态 0-禁用 1-启用"`
}
