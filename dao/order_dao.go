package dao

import (
	"easy-swap/dal"
	"easy-swap/model"
)

// OrderDao 订单数据操作
type OrderDao struct{}

func NewOrderDao() *OrderDao {
	return &OrderDao{}
}

// CreateOrder 创建挂单记录
func (d *OrderDao) CreateOrder(order *model.MarketOrder) error {
	return dal.DB.Create(order).Error
}

// UpdateStatus 修改订单状态
func (d *OrderDao) UpdateStatus(orderId string, status uint8) error {
	return dal.DB.Model(&model.MarketOrder{}).
		Where("order_id = ?", orderId).
		Update("status", status).Error
}

// ListOnSale 查询在售订单
func (d *OrderDao) ListOnSale() ([]model.MarketOrder, error) {
	var list []model.MarketOrder
	err := dal.DB.
		Where("status = 0").
		Order("id DESC").
		Find(&list).Error
	return list, err
}

// DeleteByBlockGreater 分叉回滚删除订单
func (d *OrderDao) DeleteByBlockGreater(block uint64) error {
	return dal.DB.Where("block_number > ?", block).Delete(&model.MarketOrder{}).Error
}