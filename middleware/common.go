package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func CommonMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(" 执行了中间件 ")
		c.Next()
	}
}
