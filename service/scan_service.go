package service

import (
	"easy-swap/dao"
)

// ScanService 扫链业务逻辑
type ScanService struct {
	scanDao *dao.ScanDao
}

func NewScanService() *ScanService {
	return &ScanService{
		scanDao: dao.NewScanDao(),
	}
}

// Init 初始化扫描记录
func (s *ScanService) Init() {
	_ = s.scanDao.InitRecord()
}

// GetLastBlock 获取最后扫描高度
func (s *ScanService) GetLastBlock() uint64 {
	block, _ := s.scanDao.GetLastBlock()
	return block
}

// 获取最后扫描高度hash
func ( s *ScanService) GetLastBlockHash() (string,error) {
	lastBlockHash,err := s.scanDao.GetLastBlockHash()
	return lastBlockHash,err

}

// SaveLastBlock 保存扫描高度
func (s *ScanService) SaveLastBlock(num uint64) {
	_ = s.scanDao.SaveLastBlock(num)
}