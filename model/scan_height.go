package model

import "gorm.io/gorm"

// ScanHeight 扫链区块高度记录表
type ScanHeight struct {
	gorm.Model
	LastBlock uint64 `gorm:"comment:最后扫描区块高度"`
	LastBlockHash string `gorm:"comment:最后扫描区块hash"`
}