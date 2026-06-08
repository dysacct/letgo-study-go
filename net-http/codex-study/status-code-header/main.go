package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request) {
		// Header 必须在WriteHeader 或写body钱设置
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Requst-Handled-By", "net-http-demo")

		// 创建资源通常返回201 Created
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "user created")
	})

	fmt.Println("服务器启动: http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
