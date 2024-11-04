package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Vulnerability struct {
	Severity string
	Age      int
}

func main() {
	// 打开文件
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var vulnerabilities []Vulnerability

	// 读取文件内容
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line) // 分割每行
		if len(parts) == 2 {
			age := 0
			fmt.Sscanf(parts[1], "%d", &age)
			vulnerabilities = append(vulnerabilities, Vulnerability{
				Severity: parts[0],
				Age:      age,
			})
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// 统计按期修复情况
	totalCount := map[string]int{"HIGH": 0, "MEDIUM": 0, "LOW": 0}
	onTimeCount := map[string]int{"HIGH": 0, "MEDIUM": 0, "LOW": 0}

	// 计算按期修复情况并打印结果
	fmt.Println("修复情况如下:")
	for _, v := range vulnerabilities {
		totalCount[v.Severity]++
		isOnTime := checkRepairDeadline(v.Severity, v.Age)
		status := ""
		if isOnTime {
			onTimeCount[v.Severity]++
			status = "[按期修复]"
		} else {
			status = "[未按期修复]"
		}
		fmt.Printf("%s\t%d %s\n", v.Severity, v.Age, status)
	}

	// 输出按期修复率
	fmt.Println("\n按期修复率如下:")
	for severity, total := range totalCount {
		if total > 0 {
			rate := (float64(onTimeCount[severity]) / float64(total)) * 100
			fmt.Printf("%s 按期修复率: %.2f%%\n", severity, rate)
		} else {
			fmt.Printf("%s 没有数据\n", severity)
		}
	}
}

// 检查是否按期修复
func checkRepairDeadline(severity string, age int) bool {
	switch severity {
	case "CRITICAL":
		return age <= 3
	case "HIGH":
		return age <= 7
	case "MEDIUM":
		return age <= 10
	case "LOW":
		return age <= 20
	default:
		return false
	}
}
