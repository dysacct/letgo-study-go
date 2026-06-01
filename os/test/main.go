package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("错误：请提供文件路径。")
		fmt.Println("用法：wordcount <文件路径>")
		os.Exit(1)
	}
	filePath := os.Args[1]

	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("读取文件失败: %v\n", err)
		os.Exit(1)
	}
	content := string(data)

	fmt.Println(len(content))
	if len(content) == 0 {
		printResult(0, 0, 0)
		return
	}

	// 3. 统计行数
	// 使用 strings.Split 按换行符切分
	lines := strings.Split(content, "\n")
	lineCount := len(lines)
	// 如果文件以换行符结尾，Split 会多切出一个空字符串，需要减 1
	if strings.HasSuffix(content, "\n") {
		lineCount--
	}

	// 4. 统计单词数
	// 题目要求用 strings.Split。因为单词可能被空格或换行符分隔，
	// 我们先将换行符替换为空格，统一处理。
	cleanedContent := strings.ReplaceAll(content, "\n", " ")
	cleanedContent = strings.ReplaceAll(cleanedContent, "\r", " ")

	rawWords := strings.Split(cleanedContent, " ")
	wordCount := 0
	// 过滤掉由于连续空格切分出来的空字符串
	for _, word := range rawWords {
		if word != "" {
			wordCount++
		}
	}

	// 5. 统计字符数
	// Go 语言中 string 的 len() 返回的是字节数。
	// 统计字符数（如中文字符占3字节）需要转换成 rune 切片。
	charCount := len([]rune(content))

	// 6. 输出结果
	printResult(lineCount, wordCount, charCount)
}

func printResult(lines, words, chars int) {
	output := "行数： " + strconv.Itoa(lines) +
		"  单词数： " + strconv.Itoa(words) +
		"  字符数： " + strconv.Itoa(chars)

	fmt.Println(output)

}
