package model

import (

	"gorm.io/gorm"
)

// MarketOrder NFT挂单订单
type MarketOrder struct {
	gorm.Model
	OrderKey          string `gorm:"uniqueIndex;type:varchar(64);comment:合约订单Key"`
	Side              uint8    `gorm:"comment:0 List 1Bid"`
	SaleKind          uint8    `gorm:"comment:0 FixedPriceForCollection 1 FixedPriceForItem  "`
	Maker             string   `gorm:"type:varchar(64);comment:订单创建者"`
	Taker             string   `gorm:"type:varchar(64);comment:购买者"`
	TokenID           string   `gorm:"index;type:varchar(64);comment:NFT代币ID"`
	CollectionAddress string   `gorm:"type:varchar(64);comment:NFT合约地址"`
	Amount            string `gorm:"comment:NFT数量"`
	Price             string `gorm:"comment:挂单价格"`
	OrderStatus       uint8    `gorm:"comment:0挂单 1成交 2取消"`
	TxHash            string   `gorm:"uniqueIndex;type:varchar(128);comment:交易哈希"`
	BlockNumber       uint64   `gorm:"comment:区块高度"`
}
