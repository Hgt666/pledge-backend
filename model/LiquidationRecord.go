package model

import "gorm.io/gorm"

type LiquidationRecord struct {
	gorm.Model
	ChainID          uint64  `json:"chain_id"  gorm:"column:chain_id;comment:链ID"`
	PoolID           string  `json:"pool_id"  gorm:"column:pool_id;comment:池ID"`
	UserAddress      string  `json:"user_address"  gorm:"column:user_address;comment:用户钱包地址"`
	PledgeLoanID     uint    `json:"pledge_loan_id"  gorm:"column:pledge_loan_id;comment:关联的借贷记录ID"`
	TxHash           string  `json:"tx_hash"  gorm:"column:tx_hash;comment:清算交易hash;uniqueIndex;size:66"`
	CollateralAmount string  `json:"collateral_amount"  gorm:"column:collateral_amount;comment:清算金额"`
	Price            float64 `json:"price"  gorm:"column:price;comment:清算时的价格"`
	Status           uint8   `json:"status"  gorm:"column:status;comment:状态 0-进行中 1-已清算"`
}
