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
		apiGroup.GET("/ping",api.Ping)
		apiGroup.GET("/pledge/pools",api.GetPledgePoolList)
		apiGroup.GET("/getPoolLength",api.GetPledgePoolLength)
	}
}