package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go_url_shortener/config"
	"go_url_shortener/handler"
	"go_url_shortener/middleware"
)

func main() {
	config.InitLogger()
	defer config.SugarLogger.Sync()
	config.InitDB()

	// 不使用默认的，自己自定义日志支持
	router := gin.New()
	router.Use(gin.Recovery())
	// 增添跨域支持
	router.Use(cors.Default())
	// 中间件
	router.Use(middleware.CommonMiddleware())
	router.Use(middleware.LogHandler())

	// 业务路由
	router.GET("/:code", handler.GetUrlHandler)
	router.POST("/save", handler.SaveUrlHandler)

	// 监听 8080 端口并启动服务器
	err := router.Run("127.0.0.1:8080")
	if err != nil {
		config.SugarLogger.Error(" 服务启动失败，请检查相关配置 ")
	}
	config.SugarLogger.Info(" 服务成功启动...")
}
