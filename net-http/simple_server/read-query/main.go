package main

import (
	"encoding/json"
	"net/http"
)


/*
r.URL.Path      → "/echo"              （路径部分）
r.URL.RawQuery  → "msg=hello&count=3"  （原始 query 字符串）
r.URL.Query()   → map[string][]string{   （解析好的 map）
                     "msg":   ["hello"],
                     "count": ["3"],
                  }
*/
func main() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		// 第一步 ： 从URL中提取出query参数
		msg := r.URL.Query().Get("msg")
		// r.URL -> URL 对象
		// .Query() -> 解析？后面的部分，返回map[string]string
		// .Get("msg") -> 获取msg参数的值*（如果有多个，只返回一个）

		// 第二步： 如果没有传msg 给个默认值
		if msg == "" {
			msg = "没有传递msg参数"
		}

		// 第三步： 返回JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"echo": msg,
		})
	})

	http.ListenAndServe(":8080", nil)

}
