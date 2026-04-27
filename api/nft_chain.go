package api

import (
	"easy-swap/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

var nftChainSvc = service.NewNFTChainService()

// 获取NFT名称和符号
func GetNFTInfo(c *gin.Context) {
	name, symbol, err := nftChainSvc.GetNFTInfo()
	if err != nil {
		Fail(c, "获取失败")
		return
	}
	Success(c, gin.H{"name": name, "symbol": symbol})
}

// 获取某个tokenId的拥有者
func GetNFTOwner(c *gin.Context) {
	tokenIdStr := c.Query("tokenId")
	tokenId, err := strconv.ParseUint(tokenIdStr, 10, 64)
	if err != nil {
		Fail(c, "参数错误")
		return
	}

	owner, err := nftChainSvc.GetOwner(tokenId)
	if err != nil {
		Fail(c, "查询失败")
		return
	}
	Success(c, gin.H{"owner": owner})
}

// 获取地址的NFT余额
func GetNFTBalance(c *gin.Context) {
	addr := c.Query("address")
	balance, err := nftChainSvc.GetBalance(addr)
	if err != nil {
		Fail(c, "查询失败")
		return
	}
	Success(c, gin.H{"balance": balance})
}