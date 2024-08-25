package strutils

import (
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
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

// 获取远程要下载的文件的大小
func GetRemoteFileSize(furl string) (int64, error) {
	// 先发送一个Head请求 来获取要下载的文件头信息(主要是文件大小的信息)
	resp, err := http.Head(furl)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	// 从返回的头信息中获取要下载的文件大小 resp.Header.Get("Content-Length")
	ct, err := strconv.Atoi(resp.Header.Get("Content-Length"))
	return int64(ct), err
}

// 判断给定的文件/路径是否存在 返回数字 0 表示不存在; 1 目录存在; 2 文件存在; -1 表示path为空
//
//	path 要检测是否存在的路径 文件夹或者文件路径, 如果路径为空返回 -1
func IsExist(path string) int8 {
	var flag int8 // 设定返回标志  0 文件或者目录不存在,大于0表示存在; 1 是目录; 2是文件
	if path == "" {
		return -1
	}
	// 查看文件是否存在  err不为nil即路径错误 表示文件或者目录不存在
	if fi, err := os.Stat(path); err == nil {
		// 判断是否是路径
		if fi.IsDir() {
			flag = 1 // 目录存在
		} else {
			flag = 2 // 文件存在
		}
	}
	return flag
}

// WriteFile 将内容写入文件 自动创建相关目录
// fpath为文件路径;data为内容;perm为权限,默认为0655.
func WriteFile(fname string, data []byte, perms ...os.FileMode) error {
	// 确保要写入的文件路径存在
	if err := DirMustExist(fname, os.ModePerm); err != nil {
		return err
	}
	var perm os.FileMode = 0655 // 文件权限标识
	if len(perms) > 0 {
		perm = perms[0]
	}
	return os.WriteFile(fname, data, perm)
}

// GetFileMode 获取路径的权限模式.
func GetFileMode(fname string) (os.FileMode, error) {
	if fi, err := os.Lstat(fname); err != nil {
		return 0, err
	} else {
		return fi.Mode(), nil
	}
}

// AppendFile 插入文件内容.若文件不存在,则自动创建.
func AppendFile(fname string, data []byte) error {
	var (
		err  error
		file *os.File
	)
	dir := path.Dir(fname)
	if err = os.MkdirAll(dir, os.ModePerm); err == nil {
		var filePerm os.FileMode
		filePerm, err = GetFileMode(fname)
		if err != nil {
			file, err = os.Create(fname)
		} else {
			file, err = os.OpenFile(fname, os.O_APPEND|os.O_CREATE|os.O_WRONLY, filePerm)
		}
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = file.Write(data)
	}
	return err
}

// GetMime 获取文件mime类型;fast为true时根据后缀快速获取;为false时读取文件头获取.
func GetFileMime(fname string, fast bool) string {
	var res string
	if fast {
		suffix := filepath.Ext(fname)
		//若unix系统中没有相关的mime.types文件时,将返回空
		res = mime.TypeByExtension(suffix)
	} else {
		srcFile, err := os.Open(fname)
		if err == nil {
			buffer := make([]byte, 512)
			_, err = srcFile.Read(buffer)
			if err == nil {
				res = http.DetectContentType(buffer)
			}
		}
	}

	return res
}

// GetFileSize 获取文件大小(bytes字节);注意:文件不存在或无法访问时返回-1 .
func GetFileSize(fname string) int64 {
	f, err := os.Stat(fname)
	if nil != err {
		return -1
	}
	return f.Size()
}

// IsFile 判断给定的文件是否是文件 是 true, 否则false (包括文件不存在等都是false)
func IsFile(fname string) bool {
	if f, err := os.Lstat(fname); err == nil {
		return !f.IsDir()
	}
	return false
}

// IsLink 判断给定的文件是否是软链接
func IsLink(fpath string) bool {
	if f, err := os.Lstat(fpath); err == nil {
		return f.Mode() == os.ModeSymlink
	}
	return false
}

// IsBinaryFile 是否二进制文件(且存在)
// fname 要判断的文件名 包含路径
// 是二进制文件 返回 true , 否则返回false
// @author tekintian <tekintian@gmail.com>
func IsBinaryFile(fname string) bool {
	f, err := os.Open(fname)
	if err != nil {
		return false
	}
	var maxR int64
	// 通过文件统计信息 获取文件大小
	if fi, err := f.Stat(); err == nil {
		maxR = fi.Size()
	}
	if maxR == 0 { // 如果文件内容为空,字节返回false
		return false
	}
	if maxR > 32 { // 最多读取32个字节 判断是否binary应该是足够了
		maxR = 32
	}
	// 只读取指定大小的数据 LimitReader
	reader := io.LimitReader(f, maxR)
	fbs := make([]byte, maxR)
	if _, err := reader.Read(fbs); err != nil {
		return false
	}
	for _, v := range fbs {
		// 读取的字节切片中如果有 v==0 说明这个文件是二进制文件. 注意字符串0的byte是48 十六进制 0x30 !!!
		if v == 0 {
			return true
		}
	}
	return false
}

// GetExt 获取文件的小写扩展名,不包括点.  如: fname为 xxxx.jpg 则返回 jpg
func GetFileExt(fname string) string {
	ext := filepath.Ext(fname)
	if ext != "" {
		return strings.ToLower(ext[1:])
	}
	return ext
}

// IsImg 是否图片文件(仅检查后缀).
func IsImg(fname string) bool {
	ext := GetFileExt(fname)
	switch ext {
	case "jpg", "jpeg", "gif", "png", "svg", "ico", "webp", "tff", "bmp":
		return true
	default:
		return false
	}
}
