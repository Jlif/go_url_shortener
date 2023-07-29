package biz

import (
	"go.uber.org/zap"
	"go_url_shortener/config"
	"go_url_shortener/util"
)

func GetSourceUrl(code string) string {
	// 查询语句
	query := "select url from shortener_url where short_url=?"

	// 执行查询
	rows, err := config.GetDB().Query(query, code)
	if err != nil {
		config.SugarLogger.Error(" 查询记录时出现错误:", zap.Error(err))
	}

	var url string

	if rows != nil {
		for rows.Next() {
			err = rows.Scan(&url)
		}
	}
	return url
}

func SaveSourceUrl(sourceUrl string) string {
	code := util.ShortUrl32(sourceUrl)
	_, err := config.GetDB().Exec("insert into shortener_url (url,short_url) values (?,?)", sourceUrl, code)
	if err != nil {
		config.SugarLogger.Error(" 保存记录时出现错误:", zap.Error(err))
	}
	return code
}
