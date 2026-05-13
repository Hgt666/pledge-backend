package model

import "gorm.io/gorm"

type MultiSign struct {
	gorm.Model
	ChainID   uint64 `json:"chain_id"  gorm:"column:chain_id;comment:链ID"`
	Signers   string `json:"signers"  gorm:"column:signers;comment:多签地址列表，逗号分隔"`
	Threshold uint   `json:"threshold"  gorm:"column:threshold;comment:多签阈值"`
	Status    uint   `json:"status"  gorm:"column:status;comment:状态 0-待审批 1-已通过"`
}
