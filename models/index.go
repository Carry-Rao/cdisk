package models

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// GenerateMainScript 生成JavaScript加载脚本
func models(baseDir string) error {
	// 确保基础目录存在
	if _, err := os.Stat(baseDir); os.IsNotExist(err) {
		return fmt.Errorf("base directory %s does not exist", baseDir)
	}

	// 读取JS目录列表
	jsDir := filepath.Join(baseDir, "js")
	files, err := ioutil.ReadDir(jsDir)
	if err != nil {
		return fmt.Errorf("failed to read directory %s: %v", jsDir, err)
	}

	// 创建或覆盖main.js文件
	mainJSPath := filepath.Join(jsDir, "main.js")
	mainFile, err := os.Create(mainJSPath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %v", mainJSPath, err)
	}
	defer mainFile.Close()

	// 写入初始化注释
	if _, err := mainFile.WriteString("// Auto-generated by Go script\n"); err != nil {
		return fmt.Errorf("failed to write to file %s: %v", mainJSPath, err)
	}

	// 遍历目录列表
	for _, file := range files {
		if !file.IsDir() {
			continue
		}

		// 安全处理目录名
		safeDirName := sanitizeFileName(file.Name())
		if safeDirName == "" {
			continue // 跳过无效目录名
		}

		// 构建script标签
		scriptPath := filepath.Join("js", safeDirName, "main.js")
		scriptTag := fmt.Sprintf(`document.getElementById("html").appendChild(document.createElement("script")).src = "%s";`, scriptPath)

		// 写入文件
		if _, err := mainFile.WriteString(scriptTag + "\n"); err != nil {
			return fmt.Errorf("failed to write to file %s: %v", mainJSPath, err)
		}
	}

	return nil
}

// sanitizeFileName 清理文件名，移除不安全字符
func sanitizeFileName(name string) string {
	// 移除路径分隔符
	name = strings.ReplaceAll(name, "/", "")
	name = strings.ReplaceAll(name, "\\", "")

	// 移除其他不安全字符
	unsafeChars := []string{"..", ":", "*", "?", "\"", "<", ">", "|"}
	for _, char := range unsafeChars {
		name = strings.ReplaceAll(name, char, "")
	}

	return name
}
