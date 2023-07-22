package handler

import (
	"encoding/json"
	"fmt"
	"go_url_shortener/biz"
	"net/http"
)

// w表示response对象，返回给客户端的内容都在对象里处理
// r表示客户端请求对象，包含了请求头，请求参数等等
func GetUrl(w http.ResponseWriter, r *http.Request) {
	uri := r.RequestURI
	fmt.Printf("传入的uri路径为：%s\n", uri)
	//去掉'/'
	uri = uri[1:]
	if len(uri) > 6 {
		fmt.Fprintf(w, "invalid short url, please check...")
		return
	}
	url := biz.GetSourceUrl(uri)
	if url == "" {
		fmt.Fprintf(w, "can't find mapping url, please check your short url.")
		return
	}
	http.Redirect(w, r, url, http.StatusFound)
}

type SaveUrlParam struct {
	Url string `json:"url"`
}

func SaveUrl(w http.ResponseWriter, r *http.Request) {
	uri := r.RequestURI
	fmt.Printf("传入的uri路径为：%s\n", uri)
	// 读取请求Body数据
	var param SaveUrlParam
	err := json.NewDecoder(r.Body).Decode(&param)
	if err != nil {
		http.Error(w, "无法解析JSON对象", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// 打印解析后的JSON对象
	fmt.Printf("解析后的JSON对象:%v\n", param)
	biz.SaveSourceUrl(param.Url)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "{\"success\":true}")
}
