package strutils

import (
	"io/fs"
	"os"
	"path/filepath"
)

// 检查指定的路径是否为空 即指定的文件夹下没有文件
func IsEmptyDir(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		return true
	}
	if stat.IsDir() {
		file, err := os.Open(path)
		if err != nil {
			return true
		}
		defer file.Close()
		names, err := file.Readdirnames(-1)
		if err != nil {
			return true
		}
		return len(names) == 0
	} else {
		return stat.Size() == 0
	}
}

// 确保给定的路径的目录存在，如果不存在则自动创建对应目录 默认目录权限标识 0777
//
//	path  要判断是否存在的路径
//	perms 目录权限标识  默认 0777可任意读写, 可自行指定 如 0755,  0666等
//
// 如果目录创建失败在返回相关的异常信息,否则返回nil
func DirMustExist(path string, perms ...fs.FileMode) error {
	// 判断文件/路径是否存在
	if flag := IsExist(path); flag == 0 {
		fdir := filepath.Dir(path)             //获取目录
		if flagd := IsExist(fdir); flagd < 1 { //如果目录不存在，则自动创建目录
			if flagd == -1 { // 获取到的 fdir是空字符串 表示就是当前路径, 可以直接返回了
				return nil
			}
			var perm fs.FileMode = 0777
			if len(perms) > 0 {
				perm = perms[0]
			}
			if err := os.MkdirAll(fdir, perm); err != nil {
				return err
			}
		}
	}
	return nil
}

// DirSize 获取目录大小(bytes字节).
func GetDirSize(fname string) int64 {
	var size int64
	//filepath.Walk压测很慢
	_ = filepath.Walk(fname, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size
}

// IsDir 判断给定的文件是否是目录 是 true 否则 false (包括文件不存在等都是false)
func IsDir(fname string) bool {
	if f, err := os.Lstat(fname); err == nil {
		return f.IsDir()
	}
	return false
}
