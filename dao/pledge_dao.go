package dao

import (
	"easy-swap/dal"
	"easy-swap/model"

	"gorm.io/gorm/clause"
)

// PledgeDao 借贷数据操作
type PledgeDao struct{}

func NewPledgeDao() *PledgeDao {
	return &PledgeDao{}
}



// 获取借贷池列表
func (p *PledgeDao) GetPledgePoolList() (pools []model.LendingPool, err error) {
	var poolList []model.LendingPool
	err = dal.DB.Find(&poolList).Error
	if err != nil {
		return nil, err
	}
	return poolList, nil
}

// SaveDespositLend 保存 DepositLend 事件数据
func (p *PledgeDao) SaveDespositLend(event *model.UserDeposit) error {
	// 冪等创建
	err := dal.DB.Clauses(
		clause.OnConflict{
			Columns:   []clause.Column{{Name: "tx_hash"}}, // 根据 tx_hash 判断重复
			DoNothing: true,
		},
	).Create(event).Error
	if err != nil {
		return err
	}
	return nil
}


// SaveDespositBorrow 保存 DepositBorrow 事件数据
func (p *PledgeDao) SaveDespositBorrow(event *model.UserPledgeLoan) error {
	// 冪等创建
	err := dal.DB.Clauses(
		clause.OnConflict{
			Columns:   []clause.Column{{Name: "tx_hash"}}, // 根据 tx_hash 判断重复
			DoNothing: true,
		},
	).Create(event).Error
	if err != nil {
		return err
	}
	return nil
}


// SaveLendingPool 保存 LendingPool 事件数据
func (p *PledgeDao) SaveLendingPool(event *model.LendingPool) error {
	// 冪等创建
	err := dal.DB.Clauses(
		clause.OnConflict{
			Columns:   []clause.Column{{Name: "pool_id"}}, // 根据 pool_id 判断重复
			DoNothing: true,
		},
	).Create(event).Error
	if err != nil {
		return err
	}
	return nil
}