package model

import "gorm.io/gorm"

// MarketOrder NFT挂单订单
type MarketOrder struct {
	gorm.Model
	OrderId     string `gorm:"uniqueIndex;type:varchar(64);comment:合约订单ID"`
	TokenID     string `gorm:"index;type:varchar(64);comment:NFT代币ID"`
	Seller      string `gorm:"type:varchar(64);comment:卖家地址"`
	Price       string `gorm:"type:varchar(64);comment:挂单价格"`
	Status      uint8  `gorm:"comment:0挂单 1成交 2取消"`
	TxHash      string `gorm:"uniqueIndex;type:varchar(128);comment:交易哈希"`
	BlockNumber uint64 `gorm:"comment:区块高度"`
}