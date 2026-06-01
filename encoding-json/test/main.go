package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type ServerInfo struct {
	ServerPort any    `json:"server_port"`
	DbHost     string `json:"db_host"`
	LogLevel   string `json:"log_level"`
}

func main() {
	configFile := "config.json"
	data, err := os.ReadFile(configFile)
	if err != nil {
		fmt.Printf("读取文件失败: %v\n", err)
		os.Exit(1)
	}

	// 解析json数据
	var config ServerInfo
	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Printf("解析 JSON 数据失败: %v\n", err)
	}

	// 处理 server_port 的动态类型转换
	var finalPort int

	switch v := config.ServerPort.(type) {
	case string:
		p, err := strconv.Atoi(v)
		if err != nil {
			fmt.Printf("错误: server_port 字符串无法转换为数字: %v\n", err)
			os.Exit(1)
		}
		finalPort = p
	case float64:
		finalPort = int(v)
	case int:
		finalPort = v
	default:
		fmt.Printf("错误：不支持的 server_port 类型: %T\n", v)
		os.Exit(1)
	}

	// 4. 打印所有配置项
	fmt.Println("====== 配置项解析成功 ======")
	fmt.Printf("数据库地址 (db_host):   %s\n", config.DbHost)
	fmt.Printf("日志级别 (log_level): %s\n", config.LogLevel)
	fmt.Printf("服务端口 (server_port): %d (已成功转换为 int)\n", finalPort)
}
