package main

import (
	"fmt"
	"strings"
)

func main() {
	// --- 场景1 ： URL 路由匹配 （web 框架底层）
	fmt.Println("==== URL匹配 ====")
	path := "/api/v1/users/12345/profile/"

	// 判断是否匹配某路由前缀
	if strings.HasPrefix(path, "/api/v1") {
		fmt.Println("匹配 API V1 路由")
	}

	// 拆解 URL 路径
	parts := strings.Split(path, "/")
	fmt.Printf("路径拆解: %v\n", parts)

	// 提取用户ID(倒数第二段)
	// fmt.Println(len(parts))
	if len(parts) >= 2 {
		userID := parts[len(parts)-2]
		fmt.Printf("提取用户ID: %s\n", userID)
	}

	// 场景2 ：处理用户输入（清理空格）
	fmt.Println("\n==== 清理用户输入 ====")

	rawEmail := "   Alice@Example.COM   "
	// 真是项目： 用户输入总是要做 trim + 规范化
	cleanEmail := strings.ToLower(strings.TrimSpace(rawEmail))
	fmt.Printf("原始输入: %q\n", rawEmail)
	fmt.Printf("清理后: %q\n", cleanEmail)

	// --- 场景3 ： CSV 数据解析 ----

	fmt.Println("\n==== CSV 数据解析 ====")

	csvLine := "Alice,25,Enginerr,Beijing"
	fields := strings.Split(csvLine, ",")
	fmt.Printf("CSV字段: %v\n", fields)
	fmt.Printf("姓名: %s, 年龄: %s, 职业: %s, 城市: %s\n", fields[0], fields[1], fields[2], fields[3])

	// ---- 场景4 ： 高性能字符串拼接（日志/报告生成）
	fmt.Println("\n==== string.Builder 字符串拼接 ====")

	var sb strings.Builder // 创建一个 Builder 实例

	// 模拟拼接一个 JSON 日志
	sb.WriteString(`{"level":"info"`)
	sb.WriteString(`,"msg":"request completed"`)
	sb.WriteString(`,"duration_ms":42`)
	sb.WriteString("}\n")

	result := sb.String() // 一次获取最终字符串
	fmt.Printf("生成的日志: %s", result)
}
