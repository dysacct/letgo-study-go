package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /search", func(w http.ResponseWriter, r *http.Request) {
		// 从 URL 查询参数中读取 q,
		keyword := r.URL.Query().Get("q")
		if keyword == "" {
			http.Error(w, "miss query parameter: q", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "search keyword: %s\n", keyword)
		w.Write([]byte("search keyword " + keyword))
	})

	fmt.Println("服务器启动: http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
