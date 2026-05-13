package model

import "gorm.io/gorm"

type Token struct {
	gorm.Model
	ChainID      uint64 `json:"chain_id"  gorm:"column:chain_id;comment:链ID"`
	TokenAddress string `json:"token_address"  gorm:"column:token_address;comment:代币地址;uniqueIndex;size:66"`
	TokenName    string `json:"token_name"  gorm:"column:token_name;comment:代币名称"`
	TokenSymbol  string `json:"token_symbol"  gorm:"column:token_symbol;comment:代币符号"`
	Decimals     uint8  `json:"decimals"  gorm:"column:decimals;comment:小数位"`
	Status       uint   `json:"status"  gorm:"column:status;comment:状态 0-禁用 1-启用"`
	IsStable     bool   `json:"is_stable"  gorm:"column:is_stable;comment:是否稳定币"`
}
