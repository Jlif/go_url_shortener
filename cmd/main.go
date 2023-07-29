package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go_url_shortener/handler"
	"go_url_shortener/middleware"
	"log"
)

func main() {
	router := gin.Default()

	// 增添跨域支持
	router.Use(cors.Default())

	// 中间件
	router.Use(middleware.CommonMiddleware())

	// 业务路由
	router.GET("/:code", handler.GetUrlHandler)
	router.POST("/save", handler.SaveUrlHandler)

	// 监听 8080 端口并启动服务器
	err := router.Run("127.0.0.1:8080")
	if err != nil {
		log.Fatal(" 服务启动失败 ")
	}
}
