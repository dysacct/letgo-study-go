package main

import "fmt"

// 定义一个结构体，模拟项目中的"用户"数据模型
type User struct {
	Name  string
	Age   int
	Email string
}

func main() {
	// ---- 基础打印 ----
	fmt.Println("===== 基础打印 =====")

	// Println：自动加空格、自动换行
	fmt.Println("Hello", "World", 2026)
	// 输出: Hello World 2026

	// Printf：用占位符精确控制输出格式
	name := "小明"
	score := 95.5
	fmt.Printf("学生 %s 的成绩是 %.1f 分\n", name, score)
	// 输出: 学生 小明 的成绩是 95.5 分

	// ---- 调试神器 %v / %+v / %#v ----
	fmt.Println("===== 调试神器 =====")

	u := User{Name: "Alice", Age: 25, Email: "alice@example.com"}

	fmt.Printf("  %%v:  %v\n", u)
	// 输出:   %v:  {Alice 25 alice@example.com}

	fmt.Printf(" %%+v: %+v\n", u)
	// 输出:  %+v: {Name:Alice Age:25 Email:alice@example.com}

	fmt.Printf(" %%#v: %#v\n", u)
	// 输出:  %#v: main.User{Name:"Alice", Age:25, Email:"alice@example.com"}

	fmt.Printf("  %%T: %T\n", u)
	// 输出:   %T: main.User

	// ---- Sprintf：不打印，返回格式化后的字符串 ----
	fmt.Println("===== Sprintf 实战 =====")

	// 真实场景：构造 HTTP 响应消息
	statusCode := 200
	message := fmt.Sprintf("状态码: %d, 操作成功", statusCode)
	fmt.Println(message)
	// 输出: 状态码: 200, 操作成功

	// 真实场景：拼接 API URL
	userID := 12345
	apiURL := fmt.Sprintf("/api/v1/users/%d/profile", userID)
	fmt.Println("请求URL:", apiURL)
	// 输出: 请求URL: /api/v1/users/12345/profile
}
