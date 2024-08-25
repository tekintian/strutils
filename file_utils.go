package strutils

import (
	"fmt"
	"os"
)

// 获取文件内容数据并将文件内容中的变量替换为环境变量
// filename 要读取的文件路径,相对于当前项目根目录的相对路径 或者绝对路径
// 返回使用os.ExpandEnv将文件内容中的变量替换为环境变量的值后返回
func GetFileEnvStr(filename string) (string, error) {
	if data, err := os.ReadFile(filename); err == nil {
		return os.ExpandEnv(string(data)), nil
	} else {
		return "", fmt.Errorf("读取文件 %v 异常! %v", filename, err)
	}
}
