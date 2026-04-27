package router

import (
	"easy-swap/api"
	"github.com/gin-gonic/gin"
)

// Register 注册全部路由
func Register(r *gin.Engine) {
	apiGroup := r.Group("/api")
	{
		// NFT相关接口
		apiGroup.GET("/nft/transfers", api.GetNftTransfers)
		// 订单相关接口
		apiGroup.GET("/order/list", api.GetOrderList)
		// 运维调试接口
		apiGroup.GET("/debug/rescan", api.RescanBlocks)
	}
}