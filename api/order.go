package api

import (
	"easy-swap/service"
	"github.com/gin-gonic/gin"
)

var orderService = service.NewOrderService()

// GetOrderList 获取在售挂单列表
func GetOrderList(c *gin.Context) {
	list, err := orderService.GetOnSaleList()
	if err != nil {
		Fail(c, "查询订单列表失败")
		return
	}
	Success(c, list)
}