package model

import "gorm.io/gorm"

type LendingPool struct {
	gorm.Model
	ChainID         uint64 `json:"chain_id"  gorm:"column:chain_id;comment:链ID"`
	MultiSignID     uint   `json:"multi_sign_id"  gorm:"column:multi_sign_id;comment:多签ID"`
	PoolID          string `json:"pool_id"  gorm:"column:pool_id;comment:池ID;uniqueIndex;size:255"`
	PoolAddress     string `json:"pool_address"  gorm:"column:pool_address;comment:池地址"`
	Name            string `json:"name"  gorm:"column:name;comment:池名称"`
	CollateralToken string `json:"collateral_token"  gorm:"column:collateral_token;comment:抵押物地址"`
	LoanToken       string `json:"loan_token"  gorm:"column:loan_token;comment:借贷币地址"`
	Duration        uint64 `json:"duration"  gorm:"column:duration;comment:借贷期限"`
	LTV             uint64 `json:"ltv"  gorm:"column:ltv;comment:贷款价值比,抵押率"`
	LiquidationRate uint64 `json:"liquidation_rate"  gorm:"column:liquidation_rate;comment:清算率"`
	ARP             uint64 `json:"arp"  gorm:"column:arp;comment:年化利率"`
	TotalDeposit    string `json:"total_deposit"  gorm:"column:total_deposit;comment:总存款"`
	TotalLoan       string `json:"total_loans"  gorm:"column:total_loans;comment:总借款"`
	Status          uint   `json:"status"  gorm:"column:status;comment:状态 0-Match 1-Execution 2-Finish 3-Liquidation 4-Undone"`
}
