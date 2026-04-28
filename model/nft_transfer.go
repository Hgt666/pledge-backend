package model

import "gorm.io/gorm"

// NftTransfer NFT转账记录
type NftTransfer struct {
	gorm.Model
	TokenID     string `gorm:"index;comment:NFT代币ID"`
	FromAddress string `gorm:"comment:转出地址"`
	ToAddress   string `gorm:"comment:转入地址"`
	TxHash      string `gorm:"uniqueIndex;comment:交易哈希，防重复;type:varchar(128)"`
	BlockNumber uint64 `gorm:"comment:区块高度"`
}