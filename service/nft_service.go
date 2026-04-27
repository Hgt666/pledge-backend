package service

import (
	"easy-swap/dao"
	"easy-swap/model"
)

// NFTService NFT业务逻辑
type NFTService struct {
	nftDao *dao.NFTDao
}

func NewNFTService() *NFTService {
	return &NFTService{
		nftDao: dao.NewNFTDao(),
	}
}

// SaveTransfer 保存NFT转账数据
func (s *NFTService) SaveTransfer(transfer *model.NftTransfer) error {
	return s.nftDao.CreateTransfer(transfer)
}

// GetTransferList 获取转账列表
func (s *NFTService) GetTransferList(limit int) ([]model.NftTransfer, error) {
	return s.nftDao.ListTransfers(limit)
}

// RollbackByBlock 区块分叉数据回滚
func (s *NFTService) RollbackByBlock(block uint64) error {
	return s.nftDao.DeleteByBlockGreater(block)
}