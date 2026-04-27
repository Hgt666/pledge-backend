package api

import (
	"easy-swap/internal/scan"
	"github.com/gin-gonic/gin"
	"strconv"
)

// RescanBlocks 手动区块回溯重扫
func RescanBlocks(c *gin.Context) {
	stepStr := c.DefaultQuery("step", "100")
	step, err := strconv.ParseUint(stepStr, 10, 64)
	if err != nil {
		Fail(c, "参数格式错误")
		return
	}
	scan.Global.ReScan(step)
	Success(c, "区块回溯任务已执行，正在重新扫描")
}