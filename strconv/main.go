package main

import (
	"fmt"
	"strconv"
)

func main() {
	// ---  场景1 ： HTTP Query 参数解析（Web 后端最频繁的操作） ---
	fmt.Println("=== Query 参数解析 ===")

	// 假设 HTTP 请求 URL :/api/user?page=3&size=20&active=true&score=95.5
	// 后端收到的全部是字符串
	pageStr := "3"
	sizeStr := "20"
	activeStr := "true"
	scoreStr := "95.5"

	sextStr := "ff" // 16 进制字符串
	// 把字符串转为真正的数据类型
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		fmt.Printf("解析 page 失败: %v\n", err)
		page = 1 // 失败使用默认值
	}
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		fmt.Printf("解析 size 失败: %v\n", err)
		size = 10 // 失败使用默认值
	}
	active, err := strconv.ParseBool(activeStr)
	if err != nil {
		fmt.Printf("解析 active 失败: %v\n", err)
		active = false // 失败使用默认值
	}
	score, err := strconv.ParseFloat(scoreStr, 64)
	if err != nil {
		fmt.Printf("解析 score 失败: %v\n", err)
		score = 0.0 // 失败使用默认值
	}

	sext, err := strconv.ParseInt(sextStr, 16, 64)
	if err != nil {
		fmt.Printf("解析 sext 失败: %v\n", err)
		sext = 0 // 失败使用默认值
	}

	fmt.Printf("解析结果 - page: %d, size: %d, active: %t, score: %.2f, sext: %d\n", page, size, active, score, sext)

	// 计算偏移量（SQL分页用）
	offset := (page - 1) * size
	fmt.Printf("SQL: SELECT * FROM users LIMIT %d OFFSET %d\n", size, offset)

	// ---- 场景2： 构造响应数据 ----
	fmt.Println("\n==== 构造响应 ==== ")

	// 把计算结果转回字符串 (用于 HTTP JSON 响应)
	totalCount := 142
	respBody := fmt.Sprintf(
		`{"page":%d, "size":%d,"total":%s}`,
		page, size, strconv.Itoa(totalCount), // Itoa: int -> string
	)
	fmt.Println("响应数据:", respBody)

	// ---- 场景3 ： 进制转换（处理颜色、权限位等） ----
	fmt.Println("\n==== 进制转换 ====")

	// 前端传来十六进制颜色值 "#FF5733"
	hexColor := "FF5733"
	// 转成十进制数字
	rgb, _ := strconv.ParseInt(hexColor, 16, 64)
	fmt.Printf("颜色 #%s → 十进制: %d\n", hexColor, rgb)

	// 权限码: 二进制1101 = 十进制13
	perm, _ := strconv.ParseInt("1101", 2, 64)
	fmt.Printf("二进制 1101 → 十进制: %d\n", perm)
	fmt.Printf("二进制 1101 → 十六进制: %s\n", strconv.FormatInt(perm, 16))

	price := 3.10000

	s := strconv.FormatFloat(price, 'f', -1, 64)

	fmt.Println(s)
}
