package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// ServeMux pattern 可以包含HTTP 方法。
	mux.HandleFunc("GET /users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "list users")
	})

	// 同一个路径，不同方法，映射到不同处理逻辑
	mux.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "create user")
	})

	fmt.Println("服务器启动: http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
