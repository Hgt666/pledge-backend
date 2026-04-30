package dao

import (
	"easy-swap/dal"
	"easy-swap/model"
)

// ScanDao 扫链高度数据操作
type ScanDao struct{}

func NewScanDao() *ScanDao {
	return &ScanDao{}
}

// InitRecord 初始化扫描高度记录
func (d *ScanDao) InitRecord() error {
	var count int64
	dal.DB.Model(&model.ScanHeight{}).Count(&count)
	if count == 0 {
		return dal.DB.Create(&model.ScanHeight{LastBlock: 8923483}).Error
	}
	return nil
}

// GetLastBlock 获取最后扫描区块
func (d *ScanDao) GetLastBlock() (uint64, error) {
	var sh model.ScanHeight
	err := dal.DB.First(&sh).Error
	return sh.LastBlock, err
}

// 获取最后扫描区块hash
func (d *ScanDao) GetLastBlockHash() (string, error) {
	var sh model.ScanHeight
	err := dal.DB.First(&sh).Error
	return sh.LastBlockHash,err
}

// SaveLastBlock 保存最新扫描区块
func (d *ScanDao) SaveLastBlock(num uint64) error {
	return dal.DB.Model(&model.ScanHeight{}).
		Where("id = 1").
		Update("last_block", num).Error
}