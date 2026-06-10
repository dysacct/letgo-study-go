package main

import (
	"encoding/json"
	"net/http"
)

type CreateUserRequest struct {
	Name string `json:"name"`
}

type CreateUserResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request) {
		// 限制请求体最大 1MB， 防止客户端提交较大的body。
		r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
		defer r.Body.Close()

		var req CreateUserRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid json body", http.StatusBadRequest)
			return
		}
		if req.Name == "" {
			http.Error(w, "name sis required", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusAccepted)
		_ = json.NewEncoder(w).Encode(CreateUserResponse{ID: 1001, Name: req.Name})
	})

	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
