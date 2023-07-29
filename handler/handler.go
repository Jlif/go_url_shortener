package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go_url_shortener/biz"
	"go_url_shortener/config"
	"net/http"
)

func GetUrlHandler(c *gin.Context) {
	code := c.Param("code")
	config.Logger.Debug(" 传入的 code 为：", zap.String("code", code))

	if len(code) > 6 {
		c.JSON(http.StatusOK, "invalid short url, please check...")
		return
	}
	url := biz.GetSourceUrl(code)
	if url == "" {
		c.JSON(http.StatusOK, "can't find mapping url, please check your short url.")
		return
	}
	// 重定向
	c.Redirect(http.StatusMovedPermanently, url)
}

type SaveUrlParam struct {
	Url string `json:"url"`
}

func SaveUrlHandler(c *gin.Context) {
	// 读取请求 Body 数据
	var param SaveUrlParam
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": " 无法解析 JSON 对象 "})
		return
	}

	// 打印解析后的 JSON 对象
	config.SugarLogger.Debug(" 解析后的 JSON 对象:", zap.Any("param", param))
	biz.SaveSourceUrl(param.Url)
	c.JSON(200, gin.H{"message": " 成功 "})
}
