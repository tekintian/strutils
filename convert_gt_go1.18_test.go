//go:build go1.18
// +build go1.18

// go 条件编译文件
// 本文件仅在go运行环境版本大于等于 1.18时才会生效
package strutils_test

import (
	"testing"

	strutils "github.com/tekintian/strutils"
)

func TestStrToNumber(t *testing.T) {
	var intDefval int64 = 0 // 默认值,StrToNumber返回就就是这个默认值的类型
	ival := strutils.StrToNumber("123.888abc", intDefval)
	if ival != 123 {
		t.Fatalf("Expected output to be 123, got %v", ival)
	}

	var intDefval2 int64 = 0 // 默认值,StrToNumber返回就就是这个默认值的类型
	ival2 := strutils.StrToNumber("123,456.888abc", intDefval2)
	if ival2 != 123456 {
		t.Fatalf("Expected output to be 123456, got %v", ival)
	}

	var fDefval float64 = 0 // 默认值,StrToNumber返回就就是这个默认值的类型
	fval := strutils.StrToNumber("hi123.888abc", fDefval)
	if fval != 123.888 {
		t.Fatalf("Expected output to be 123.888, got %v", ival)
	}
}
