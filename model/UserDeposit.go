package model

import "gorm.io/gorm"

type UserDeposit struct {
	gorm.Model
	ChainID     uint64 `json:"chain_id"  gorm:"column:chain_id;comment:链ID"`
	PoolID      string `json:"pool_id"  gorm:"column:pool_id;comment:池ID"`
	// 联合唯一索引
	UserAddress string `json:"user_address"  gorm:"column:user_address;comment:用户地址;not null;size:66"`
	TxHash      string `json:"tx_hash"  gorm:"column:tx_hash;comment:交易hash;uniqueIndex;not null;size:66"`
	BlockNumber uint64 `json:"block_number"  gorm:"column:block_number;comment:区块高度"`
	Amount      string `json:"amount"  gorm:"column:amount;comment:存款金额"`
	Interest    string `json:"interest"  gorm:"column:interest;comment:利息"`
	Duration    uint64 `json:"duration"  gorm:"column:duration;comment:存款期限"`
	IsClaimed   bool   `json:"is_claimed"  gorm:"column:is_claimed;comment:是否已领取"`
	Status      uint8  `json:"status"  gorm:"column:status;comment:状态 0-进行中 1-已完成 2-已撤销"`
}
