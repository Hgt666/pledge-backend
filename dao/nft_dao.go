package dao

import (
	"easy-swap/dal"
	"easy-swap/model"
)

// NFTDao NFT数据操作
type NFTDao struct{}

func NewNFTDao() *NFTDao {
	return &NFTDao{}
}

// CreateTransfer 新增NFT转账记录
func (d *NFTDao) CreateTransfer(transfer *model.NftTransfer) error {
	return dal.DB.Create(transfer).Error
}

// ListTransfers 查询NFT转账列表
func (d *NFTDao) ListTransfers(limit int) ([]model.NftTransfer, error) {
	var list []model.NftTransfer
	err := dal.DB.
		Order("id DESC").
		Limit(limit).
		Find(&list).Error
	return list, err
}

// DeleteByBlockGreater 删除指定区块之后的数据（分叉回滚）
func (d *NFTDao) DeleteByBlockGreater(block uint64) error {
	return dal.DB.Where("block_number > ?", block).Delete(&model.NftTransfer{}).Error
}