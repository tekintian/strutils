package strutils_test

import (
	"testing"

	strutils "github.com/tekintian/go-str-utils"
)

func TestStrIsChinese(t *testing.T) {
	cases := []struct {
		str      string
		expected bool
	}{
		{str: "hello world!", expected: false},
		{str: "hello 中国!", expected: false},
		{str: "你好中国", expected: true},
		{str: "123456 dfadsfdagf!", expected: false},
	}
	for _, v := range cases {
		t.Run(v.str, func(t *testing.T) {
			// 业务逻辑测试
			ret := strutils.StrIsChinese(v.str)
			if ret != v.expected {
				t.Fatalf("StrIsChinese test failed, expected %v, got %v", v.expected, ret)
			}
		})
	}
}

func TestStrContainsChinese(t *testing.T) {
	cases := []struct {
		str      string
		expected bool
	}{
		{str: "hello world!", expected: false},
		{str: "hello 中国!", expected: true},
		{str: "你好aaasfadf中国", expected: true},
		{str: "你好", expected: true},
		{str: "123456 dfadsfdagf!", expected: false},
	}
	for _, v := range cases {
		t.Run(v.str, func(t *testing.T) {
			// 业务逻辑测试
			ret := strutils.StrContainsChinese(v.str)
			if ret != v.expected {
				t.Fatalf("%v test failed, expected %v, got %v", v.str, v.expected, ret)
			}
		})
	}
}
func TestStrContainsContinuousNum(t *testing.T) {
	cases := []struct {
		str      string
		expected bool
	}{
		{str: "hello world! 1", expected: false},
		{str: "hello 中国! 12", expected: true},
		{str: "123456 dfadsfdagf!", expected: true},
	}
	for _, v := range cases {
		t.Run(v.str, func(t *testing.T) {
			// 业务逻辑测试
			ret := strutils.StrContainsContinuousNum(v.str)
			if ret != v.expected {
				t.Fatalf(" %v test failed, expected %v, got %v", v.str, v.expected, ret)
			}
		})
	}
}
