package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux() // mux 是一个 HTTP 请求 multiplexer（多路复用器），它将 URL 路径映射到处理函数。

	// 健康检查路由
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	})

	// 用户路由。
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "用户列表")
	})

	fmt.Println("服务器启动: http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("服务器启动失败:", err)
	}
}
