package main

// ---- 数据模型定义 -------

// 用户注册请求（前端 POST 过来的 JSON）
type User struct {
	Name     string `json:"name"`
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
	// ---  场景 1 ：
}
