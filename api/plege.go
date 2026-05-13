package api

import (
	"easy-swap/contract"
	"easy-swap/service"

	"github.com/gin-gonic/gin"
)


var pledgeService = service.NewPledgeService()

// GetPledgePoolList 获取借贷池列表
func GetPledgePoolList(c *gin.Context) {
	list,err := pledgeService.ListPledgePool()
	if err!= nil {
		Fail(c,"查询借贷池列表失败")
		return
	}
	Success(c,list)
}

// 获取借贷池的长度
func GetPledgePoolLength(c *gin.Context) {
	length,err := contract.PledgeInstance.PoolLength(nil)
	if err!= nil {
		Fail(c,"获取借贷池长度失败")
		return
	}
	Success(c,length)
}
