package api

import (
	"easy-swap/service"
	"github.com/gin-gonic/gin"
)

var orderService = service.NewOrderService()

// GetOrderList 获取在售挂单列表
// GetNftList
// @Summary 获取 NFT 列表
// @Description 返回所有 NFT 数据
// @Tags NFT
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":200,"data":[],"msg":"success"}"
// @Router /order/list [get]
func GetOrderList(c *gin.Context) {
	list, err := orderService.GetOnSaleList()
	if err != nil {
		Fail(c, "查询订单列表失败")
		return
	}
	Success(c, list)
}