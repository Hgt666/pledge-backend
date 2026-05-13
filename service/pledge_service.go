package service

import (
	"easy-swap/dao"
	"easy-swap/model"
)

// PledgeService 借贷业务逻辑
type PledgeService struct {
	pledgeDao *dao.PledgeDao
}

func NewPledgeService() *PledgeService {
	return &PledgeService{
		pledgeDao: dao.NewPledgeDao(),
	}
}

// 查询借贷池列表
func (s *PledgeService) ListPledgePool() ([]model.LendingPool, error) {
	return s.pledgeDao.GetPledgePoolList()
}

// 保存 DepositLend 事件数据
func (s *PledgeService) SaveDepositLend(event *model.UserDeposit) error {
	return s.pledgeDao.SaveDespositLend(event)
}

// 保存 DepositBorrow 事件数据
func (s *PledgeService) SaveDepositBorrow(event *model.UserPledgeLoan) error {
	return s.pledgeDao.SaveDespositBorrow(event)
}

// 保存 LendingPool 事件数据
func (s *PledgeService) SaveLendingPool(event *model.LendingPool) error {
	return s.pledgeDao.SaveLendingPool(event)
}