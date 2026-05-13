package dao

import (
	// "easy-swap/config"
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
	dal.DB.Model(&model.IndexerStatus{}).Count(&count)
	if count == 0 {
		return dal.DB.Create(&model.IndexerStatus{LastBlock: 10817916,}).Error
	}
	return nil
}

// GetLastBlock 获取最后扫描区块
func (d *ScanDao) GetLastBlock() (uint64, error) {
	var sh model.IndexerStatus
	err := dal.DB.First(&sh).Error
	return sh.LastBlock, err
}

// 获取最后扫描区块hash
func (d *ScanDao) GetLastBlockHash() (string, error) {
	var sh model.IndexerStatus
	err := dal.DB.First(&sh).Error
	return sh.BlockHash, err
}

// SaveLastBlock 保存最新扫描区块
func (d *ScanDao) SaveLastBlock(num uint64,block_hash string) error {
	return dal.DB.Model(&model.IndexerStatus{}).
		Where("id = 1").
		Updates(map[string]interface{}{
			"last_block": num,
			"block_hash": block_hash,
		}).Error
}