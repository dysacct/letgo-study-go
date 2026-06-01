package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ---- 数据模型定义 -------

// 用户注册请求（前端 POST 过来的 JSON）
type SignupRequest struct {
	Username string `json:"username"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// 用户信息响应（后端返回给前端的JSON）
type UserResponse struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	// omitempty ： 如果Role 为空字符串 ， JSON 中不显示这个字段
	Role string `json:"role,omitempty"`
	// json: "-" 密码绝对不返回给前端
	Password string `json:"-"`
}

func main() {
	// ---  场景 1 ：解析请求（Unmarshal） ---
	fmt.Println("====== 解析前端请求 ======")

	// 假设前端 POST 锅里的JSON body
	requestBody := `{
		"username": "alice",
		"password": "secret123",
		"email": "alice@example.com",
		"age": 25
	}`

	var req SignupRequest
	// json.Unmarshal: JSON -> GO struct
	err := json.Unmarshal([]byte(requestBody), &req)
	if err != nil {
		fmt.Println("解释失败:", err)
		return
	}
	fmt.Printf("收到注册请求: %+v\n", req) // +v是 Go 语言中格式化输出结构体的字段和值
	// 数据验证
	if req.Username == "" {
		fmt.Println("错误: 用户名不能为空")
		return
	}

	// --- 场景2 ： 构造响应（marshal） ---
	fmt.Println("\n===== 构造API相应 =====")

	// 模拟创建用户后返回
	user := UserResponse{
		ID:       10001,
		Username: req.Username,
		Email:    req.Email,
		Role:     "",
		Password: "secret123", // 这个字段会被 json:"-" 忽略掉
	}

	// json.Marshal: Go srtruct -> JSON
	respBytes, err := json.Marshal(user)
	if err != nil {
		fmt.Println("JSON 序列化失败:", err)
		return
	}
	fmt.Println("响应 JSON:", string(respBytes))

	// --- 场景3： 美化输出（调试用） ---
	fmt.Println("\n===== 美化输出 =====")

	prettyJSON, _ := json.MarshalIndent(user, "", "    ") // 第二个参数是前缀，第三个参数是缩进
	fmt.Println("美化后的 JSON:\n", string(prettyJSON))

	// ---- 场景4 ： 使用 ENcoder （适合直接写 HTTP Respond） ----
	fmt.Println("\n==== json.Encoder =====")

	var sb strings.Builder // 用于构建字符串
	encoder := json.NewEncoder(&sb)
	// SetEscapeHTML(false): 不转义HTML字符（&,<,>）
	encoder.SetEscapeHTML(false)
	encoder.Encode(user) // 直接写到strings.Builder
	fmt.Println("Encoder 输出的 JSON:\n", sb.String())
}
