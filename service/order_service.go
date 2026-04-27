package service

import (
	"easy-swap/dao"
	"easy-swap/model"
)

// OrderService 订单业务逻辑
type OrderService struct {
	orderDao *dao.OrderDao
}

func NewOrderService() *OrderService {
	return &OrderService{
		orderDao: dao.NewOrderDao(),
	}
}

// CreateOrder 创建挂单
func (s *OrderService) CreateOrder(order *model.MarketOrder) error {
	return s.orderDao.CreateOrder(order)
}

// CancelOrder 取消订单
func (s *OrderService) CancelOrder(orderId string) error {
	return s.orderDao.UpdateStatus(orderId, 2)
}

// FillOrder 订单成交
func (s *OrderService) FillOrder(orderId string) error {
	return s.orderDao.UpdateStatus(orderId, 1)
}

// GetOnSaleList 获取在售订单
func (s *OrderService) GetOnSaleList() ([]model.MarketOrder, error) {
	return s.orderDao.ListOnSale()
}

// RollbackByBlock 订单数据回滚
func (s *OrderService) RollbackByBlock(block uint64) error {
	return s.orderDao.DeleteByBlockGreater(block)
}