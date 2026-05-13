package model

import "gorm.io/gorm"

type UserPledgeLoan struct {
	gorm.Model
	ChainID          uint64 `json:"chain_id"  gorm:"column:chain_id;comment:链ID"`
	PoolID           string `json:"pool_id"  gorm:"column:pool_id;comment:池ID"`
	UserAddress      string `json:"user_address"  gorm:"column:user_address;comment:用户钱包地址"`
	TxHash           string `json:"tx_hash"  gorm:"column:tx_hash;comment:交易hash;uniqueIndex;size:66"`
	BlockNumber      uint64 `json:"block_number"  gorm:"column:block_number;comment:区块高度"`
	Duration         uint64 `json:"duration"  gorm:"column:duration;comment:借贷期限"`
	CollateralAmount string `json:"collateral_amount"  gorm:"column:collateral_amount;comment:质押数量"`
	LoanAmount       string `json:"loan_amount"  gorm:"column:loan_amount;comment:借款数量"`
	InterestAmount   string `json:"interest_amount"  gorm:"column:interest_amount;comment:总利息"`
	HealthRate       uint8  `json:"health_rate"  gorm:"column:health_rate;comment:健康率"`
	Status           uint8  `json:"status"  gorm:"column:status;comment:状态 0-进行中 1-已结清 2-已清算"`
	Version          uint8  `json:"version"  gorm:"column:version;comment:版本号;version升级时,旧数据不删除,通过版本号区分新旧数据"`
}
