package router

import (
	"easy-swap/api"
	"easy-swap/middleware"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	r.Use(
		middleware.Recovery(),
		middleware.CORS(),
		middleware.ZapLogger(), // 替换
		middleware.RequestID(),
	)

	apiGroup := r.Group("/api")
	{
		apiGroup.GET("/nft/transfers", api.GetNftTransfers)
		apiGroup.GET("/nft/info", api.GetNFTInfo)
		apiGroup.GET("/nft/owner", api.GetNFTOwner)
		apiGroup.GET("/nft/balance", api.GetNFTBalance)
		apiGroup.GET("/order/list", api.GetOrderList)
		apiGroup.GET("/debug/rescan", api.RescanBlocks)
	}
}