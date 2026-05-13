package model

import "gorm.io/gorm"

type PriceCache struct {
	gorm.Model
	Symbol  string  `json:"symbol"  gorm:"column:symbol;comment:代币符号;uniqueIndex:idx_symbol_chainID;size:255"`
	ChainID uint64  `json:"chain_id"  gorm:"column:chain_id;comment:链ID;uniqueIndex:idx_symbol_chainID;size:255"`
	Price   float64 `json:"price"  gorm:"column:price;comment:美元价格"`
	Source  string  `json:"source"  gorm:"column:source;comment:价格来源"`
}
