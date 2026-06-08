package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		// 读取路径参数，构造一个演示用户
		id := r.PathValue("id")
		user := User{
			ID:   1,
			Name: "user-" + id,
		}

		// JSON API 应明确设置响应类型
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(user)
	})

	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}

	
}
