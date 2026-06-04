package main

import (
	"encoding/json"
	"net/http"
)

/*
真实的 API 不返回纯文本，而是返回 JSON。浏览器/前端需要结构化数据。
*/
func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		// 第 1 步 ： 告诉浏览器 我返回的是JSON数据
		w.Header().Set("Content-Type", "application/json")

		// 第 2 步 ： 构造要返回的数据
		data := map[string]string{
			"message": "pong",
			"status":  "success",
		}

		// 第 3 步 ： 将数据编码成 JSON 并写入响应
		json.NewEncoder(w).Encode(data)
	})

	http.ListenAndServe(":8081", nil)
}
