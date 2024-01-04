package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		// 核心处理方式
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Header("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE")
		c.Set("content-type", "application/json")
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
