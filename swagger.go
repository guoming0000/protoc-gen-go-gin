package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	fullPath := os.Args[1]
	newFilePath := strings.ReplaceAll(fullPath, ".proto", ".swagger.proto")
	buf := &bytes.Buffer{}

	// Read the file
	// 打开文件
	file, err := os.Open(fullPath)
	if err != nil {
		fmt.Println("无法打开文件:", err)
		return
	}
	defer file.Close()

	// 创建一个Scanner对象
	scanner := bufio.NewScanner(file)

	// 逐行读取文件内容
	rawMethods := []string{}
	re := regexp.MustCompile(`returns\s*\((\w+)\)`)
	for scanner.Scan() {
		rawLine := scanner.Text()
		line := strings.Trim(rawLine, " ")
		// 忽略注释行
		if strings.HasPrefix(line, "//") {
			buf.WriteString(rawLine + "\n")
			continue
		}
		match := re.FindStringSubmatch(line)
		if len(match) > 1 {
			replyStr := match[1]
			rawMethods = append(rawMethods, replyStr)
			newLine := strings.ReplaceAll(rawLine, "("+replyStr+")", "(T"+replyStr+")")
			buf.WriteString(newLine + "\n")
		} else {
			buf.WriteString(rawLine + "\n")
		}
	}

	for _, m := range rawMethods {
		buf.WriteString("message T" + m + " {\n")
		buf.WriteString("   int32 code = 1; // binding:\"required\"\n")
		buf.WriteString("   string msg = 2; // binding:\"required\"\n")
		buf.WriteString("   " + m + " data = 3; // binding:\"required\"\n")
		buf.WriteString("}\n")
	}
	if os.WriteFile(newFilePath, buf.Bytes(), 0666) == nil {
		fmt.Println(newFilePath + " 文件生成成功")
	}
}
