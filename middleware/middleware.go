package middleware

import (
	"easy-swap/api"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"runtime/debug"
	"time"
)

// 1. CORS 跨域中间件（前端必用）
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Expose-Headers", "*")

		// 放行OPTIONS预检
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

// 2. 全局日志中间件（记录请求）
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		clientIP := c.ClientIP()
		method := c.Request.Method

		// 处理请求
		c.Next()

		// 打印日志
		cost := time.Since(start)
		status := c.Writer.Status()
		log.Printf("[%s] %s | %d | %s | %s",
			method, path, status, clientIP, cost)
	}
}

// 3. 全局异常捕获（防止panic崩服务）
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 打印堆栈
				log.Println("panic:", err)
				debug.PrintStack()
				// 返回统一错误
				api.Fail(c, "服务器内部错误")
				c.Abort()
			}
		}()
		c.Next()
	}
}

// 4. 请求ID追踪（企业级必备）
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 生成唯一ID
		requestID := uuid.NewString()
		c.Set("request_id", requestID)
		c.Header("X-Request-ID", requestID)
		c.Next()
	}
}