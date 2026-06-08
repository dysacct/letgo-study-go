package main

import "net/http"

type CreateUserRequest struct {
	Name string `json:"name"`
}

type CreteUserResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request) {
		// 限制请求体最大 1MB， 防止客户端提交较大的body。
		r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
		defer r.Body.Close()
	})
}
