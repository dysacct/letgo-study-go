package main

import (
	"fmt"
	"net/http"
)

// Go 1.22+ 的 ServeMux 支持 {name} 路径变量，并通过 r.PathValue("name") 读取

func main() {
	mux := http.NewServeMux()

	// {id} 是路径的参数，匹配/users/1, /users/abc 等路径。
	mux.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id") // 从路径参数中获取 id 的值
		fmt.Fprintf(w, "user id: %s\n", id)
	})

	fmt.Println("server listening on http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}

}
