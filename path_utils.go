// path路径相关工具函数
package strutils

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

// 获取当前app的运行路径 如: /opt/app/gotms
func GetAppPath() string {
	var appPath string
	if exeFile, err := os.Executable(); err == nil {
		appPath = filepath.Dir(exeFile)
	} else {
		appPath, _ = os.Getwd()
	}
	return appPath
}

// 去除当前app的路径(如果包含)返回相对路径, 否则原样返回
func TrimAppPath(inPath string) string {
	if path, ok := strings.CutPrefix(inPath, GetAppPath()); ok {
		return strings.TrimLeft(path, string(os.PathSeparator)) // 注意这里去除左边的 /符号,返回的是一个相对路径
	}
	return inPath
}

// 获取inPath的绝对路径 会替换路径中的变量 ${var} 或者 $var 为环境变量的值
func AbsPathify(inPath string) string {
	if inPath == "$HOME" || strings.HasPrefix(inPath, "$HOME"+string(os.PathSeparator)) {
		homeDir, _ := os.UserHomeDir()
		inPath = homeDir + inPath[5:]
	}
	inPath = os.ExpandEnv(inPath) // 替换路径中的环境变量
	if filepath.IsAbs(inPath) {
		return filepath.Clean(inPath)
	}
	p, err := filepath.Abs(inPath)
	if err == nil {
		return filepath.Clean(p)
	}
	log.Printf("could not discover absolute path: %v", err.Error())
	return ""
}

