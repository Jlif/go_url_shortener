package biz

import (
	"go_url_shortener/util"
	"log"
)

func GetSourceUrl(code string) string {
	// 查询语句
	query := "select url from shortener_url where short_url=?"

	// 执行查询
	rows, err := util.GetDB().Query(query, code)
	if err != nil {
		log.Println("查询记录时出现错误:", err)
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
	_, err := util.GetDB().Exec("insert into shortener_url (url,short_url) values (?,?)", sourceUrl, code)
	if err != nil {
		log.Println("保存记录时出现错误:", err)
	}
	return code
}
