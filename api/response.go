package api

import "github.com/gin-gonic/gin"

// Success 统一成功响应
func Success(c *gin.Context, data any) {
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": data,
	})
}

// Fail 统一失败响应
func Fail(c *gin.Context, msg string) {
	c.JSON(200, gin.H{
		"code": 500,
		"msg":  msg,
	})
}