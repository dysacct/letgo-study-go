package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 当用户访问 / 时，执行这个函数
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// w 是响应出口，往里面写什么，客户端收到什么
		fmt.Fprintln(w, "Hello , GO HTTP!")
	})
	fmt.Println("open http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
