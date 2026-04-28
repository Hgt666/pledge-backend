package model

import "gorm.io/gorm"

// ScanHeight 扫链区块高度记录表
type ScanHeight struct {
	gorm.Model
	LastBlock uint64 `gorm:"default:9280861;comment:最后扫描区块高度"`
}