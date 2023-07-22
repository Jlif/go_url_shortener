package main

import (
	"fmt"
	"go_url_shortener/handler"
	"net/http"
)

func main() {
	// 注册路由处理函数
	http.HandleFunc("/", handler.GetUrl)
	http.HandleFunc("/save", handler.SaveUrl)

	// 监听端口并启动服务器
	port := "8080"
	fmt.Printf("服务器已启动，监听端口：%s\n", port)
	err := http.ListenAndServe("127.0.0.1:"+port, nil)
	if err != nil {
		fmt.Println("启动服务器时出现错误:", err)
	}
}
