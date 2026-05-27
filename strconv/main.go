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
}
