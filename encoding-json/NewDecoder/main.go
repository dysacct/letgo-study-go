package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type UserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	// 安全检查： 确保是 POST 请求
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 2. 生产环境防御： 限制读取的最大字节数（防止恶意大文件大到撑破连接）
	r.Body = http.MaxBytesReader(w, r.Body, 1048576) // 1MB
	defer r.Body.Close()                             // 确保请求体被关闭
	var req UserRequest

	// 3 . 创建Decoder 并进行流式解析
	decoder := json.NewDecoder(r.Body)

	// 可选进阶：如果请求中包含结构体未定义的字段，直接报错（防止可无端传错参数）
	// decoder.DisallowUnkonwFields()
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Invalid JSON: %v", err), http.StatusBadRequest)
		return
	}

	// 4. 业务逻辑处理
	fmt.Printf("成功解析数据: %+v\n", req)

	// 5. 返回响应
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

func main() {
	http.HandleFunc("/create-user", createUserHandler)

	fmt.Println("Server is running on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server failed: ", err)
	}
}
