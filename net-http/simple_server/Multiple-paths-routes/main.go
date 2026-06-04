package main

import (
	"fmt"
	"net/http"
)
/* 
要解决的问题

一个网站不可能只有一个页面。你要有首页、关于页、API 接口……你需要多个路径映射到不同的处理函数。
*/

func main() {
	// 路径1： 首页
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "这是首页！！！")
	})

	// 路径2： Hello页面
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "你好，世界！")
	})

	// 路径3： bye 页面
	http.HandleFunc("/bye", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "再见，世界！")
	})

	fmt.Println("服务器启动: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
