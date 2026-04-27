package api

import (
	"easy-swap/service"
	"github.com/gin-gonic/gin"
)

var nftService = service.NewNFTService()

// GetNftTransfers 获取NFT转账记录列表
func GetNftTransfers(c *gin.Context) {
	list, err := nftService.GetTransferList(30)
	if err != nil {
		Fail(c, "查询NFT流转记录失败")
		return
	}
	Success(c, list)
}