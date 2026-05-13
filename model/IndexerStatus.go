package model

import "gorm.io/gorm"

type IndexerStatus struct {
	gorm.Model
	ChainID     uint64 `json:"chain_id"  gorm:"column:chain_id;comment:链ID;uniqueIndex"`
	LastBlock   uint64 `json:"last_block"  gorm:"column:last_block;comment:最后处理的区块高度"`
	BlockHash   string `json:"block_hash"  gorm:"column:block_hash;comment:最后处理的区块hash"`
	ScanStatus  int    `json:"scan_status"  gorm:"column:scan_status;comment:扫描状态 0-未扫描 1-扫描中 2-扫描完成"`
}