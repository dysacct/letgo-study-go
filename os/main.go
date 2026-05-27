package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("\n==== 命令行参数 ====")
	fmt.Println("程序名称:", os.Args[0])
	if len(os.Args) > 1 {
		fmt.Println("传入参数:", os.Args[1:])
	} else {
		fmt.Println("未传入参数")
	}

	fmt.Println("\n==== 环境变量 ====")
	api := GetEnvExample()
	fmt.Println("API_KEY:", api)

	fmt.Println("\n==== 文件目录操作 ====")
	uploadDir := "./uploads/images"
	err := os.MkdirAll(uploadDir, 0755) // 0755 = rwxr-xr-x
	if err != nil {
		fmt.Println("创建目录失败:", err)
		os.Exit(1)
	}
	fmt.Println("已创建目录:", uploadDir)

	// 写配置文件（真实项目启动时可能会动态生成配置）
	// fmt.Println("\n==== 写配置文件 ====")
	configContent := "server_port=8080\nlog_level=info\n"
	err = os.WriteFile("./config.conf", []byte(configContent), 0644)
	if err != nil {
		fmt.Println("写配置文件失败:", err)
	} else {
		fmt.Println("已写入配置文件: config.conf")
	}

	// 检查文件/目录是否存在
	absPath, _ := filepath.Abs("./config.conf")
	fmt.Println("上传文件绝对路径:", absPath)
}

func GetEnvExample() string {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	apiKey := os.Getenv("API_DEEPSEEK")
	if apiKey == "" {
		fmt.Println("未找到环境变量 API_DEEPSEEK")
		return ""
	} else {
		// fmt.Println("API_DEEPSEEK:", apiKey)
		return apiKey
	}
}
