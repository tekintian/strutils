package strutils_test

import (
	"fmt"
	"testing"

	strutils "github.com/tekintian/go-str-utils"
)

func TestStr2Int(t *testing.T) {
	cases := []struct {
		str      string
		expected int
	}{
		{str: "12.3", expected: 12},
		{str: "hi123你好", expected: 123},
		{str: "2,56.9", expected: 256},
		{str: "hello2.56aa", expected: 2},
		{str: "hello aa 89", expected: 89},
	}
	for _, v := range cases {
		t.Run(v.str, func(t *testing.T) {
			ret := strutils.Str2Int(v.str)
			if ret != v.expected {
				t.Fatalf("Str2Int %v expected %v got %v", v.str, v.expected, ret)
			}
		})
	}
}
func TestStr2Int64Slice(t *testing.T) {
	cases := []struct {
		str      string
		expected []int64
	}{
		{str: "12.3你好, 456", expected: []int64{12, 456}},
		{str: "hi123你好88", expected: []int64{123, 88}},
	}
	for _, v := range cases {
		t.Run(v.str, func(t *testing.T) {
			ret := strutils.Str2Int64Slice(v.str)
			// 这里判断2个切片是否相等,我们直接将他转换为字符串后比较; go版本大于1.21的话可以使用 slices包里面的Equal函数比较
			if fmt.Sprintf("%v", ret) != fmt.Sprintf("%v", v.expected) {
				t.Fatalf("Str2Int64Slice %v expected %v got %v", v.str, v.expected, ret)
			}
		})
	}
}
func TestStr2Float64(t *testing.T) {
	cases := []struct {
		str      string
		expected float64
	}{
		{str: "12.3", expected: 12.3},
		{str: "hi123你好", expected: 123},
		{str: "2,56.9", expected: 256.9},
		{str: "hello2.56aa", expected: 2.56},
		{str: "hello aa 89", expected: 89},
	}
	for _, v := range cases {
		t.Run(v.str, func(t *testing.T) {
			ret := strutils.Str2Float64(v.str)
			if ret != v.expected {
				t.Fatalf("Str2Int %v expected %v got %v", v.str, v.expected, ret)
			}
		})
	}
}
