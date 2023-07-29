package middleware

import (
	"github.com/gin-gonic/gin"
	"go_url_shortener/config"
)

func CommonMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		config.SugarLogger.Info(" 执行了中间件 ")
		c.Next()
	}
}
