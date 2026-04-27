package router

import (
	"easy-swap/api"
	"github.com/gin-gonic/gin"
	"easy-swap/middleware"
)

// Register 注册全部路由
func Register(r *gin.Engine) {
	// 全局注册中间件（顺序不能乱）
	r.Use(
		middleware.Recovery(),   // 异常捕获
		middleware.CORS(),       // 跨域
		middleware.Logger(),     // 日志
		middleware.RequestID(),  // 请求ID
	)
	apiGroup := r.Group("/api")
	{
		// NFT相关接口
		apiGroup.GET("/nft/transfers", api.GetNftTransfers)
		// 订单相关接口
		apiGroup.GET("/order/list", api.GetOrderList)
		// 运维调试接口
		apiGroup.GET("/debug/rescan", api.RescanBlocks)

		apiGroup.GET("/nft/info", api.GetNFTInfo)
		apiGroup.GET("/nft/owner", api.GetNFTOwner)
		apiGroup.GET("/nft/balance", api.GetNFTBalance)
	}
}
