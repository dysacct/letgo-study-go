package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println("Current time:", t)
	// 转换成时间戳
	timestamp := t.Unix()
	fmt.Println("Timestamp:", timestamp)

	// 分别获取名字段
	fmt.Printf("年:%d 月:%d 日:%d 时:%d 分:%d 秒:%d\n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

	// ---- 2. 时间格式化（Go 特有方式） ----
	fmt.Println("\n===== 时间格式化 =====")
	// 记住这个魔法时间: 2006-01-02 15:04:05
	fmt.Println("标准格式:", t.Format("2006-01-02 15:04:05"))
	fmt.Println("日期格式:", t.Format("2006/01/02"))
	fmt.Println("时间格式:", t.Format("15:04:05"))
	fmt.Println("RFC3339:", t.Format(time.RFC3339)) // 国际标准
	fmt.Println("中文格式:", t.Format("2006年01月02日 15:04:05"))

	// 判断Token是否过期
	fmt.Println("\n===== 判断Token是否过期 =====")
	expireTime := t.Add(-1 * time.Hour)
	fmt.Println(expireTime)
	tokenExpired := 2 * time.Hour
	fmt.Println(tokenExpired)
	if time.Since(expireTime) > tokenExpired {
		fmt.Println("Token 已过期")
	} else {
		fmt.Println("Token 仍然有效")
	}

	// ---- 3. 时间解析 ----
	fmt.Println("\n===== 时间解析 =====")

	// 前端传来 204-01-15 14:30:25
	inputTime := "2024-01-15 14:30:25"
	tim, err := time.Parse("2006-01-02 15:04:05", inputTime)
	if err != nil {
		fmt.Println("时间解析错误:", err)
	} else {
		fmt.Println("解析后的时间:", tim)
		fmt.Printf("是星期几: %s\n", tim.Weekday())
	}

	// 用 time.ParseInLocation 解析带时区的时间字符串
	loc, _ := time.LoadLocation("Asia/Shanghai")
	t2, _ := time.ParseInLocation("2006-01-02 15:04:05", inputTime, loc)
	fmt.Println("解析后的时间（上海时区）:", t2)
}
