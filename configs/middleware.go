package configs

import (
	"github.com/gin-gonic/gin"
)

func BeforeAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "text/html; charset=UTF-8")
		c.Header("Server", "Ubuntu")
		c.Header("X-Powered-By", "Go")
		c.Next()
	}
}
