package middleware

import (
	"github.com/gin-gonic/gin"
	"go_url_shortener/config"
)

// 全局日志记录器

func LogHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 Gin 框架的请求日志
		requestLog := c.Request.Method + " " + c.Request.URL.Path
		config.SugarLogger.Debug(requestLog)

		// 继续处理请求
		c.Next()
	}
}
