package strutils_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

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
		{str: "hi123你好88", expected: []int64{123}},
		{str: "hi123, 你好88", expected: []int64{123, 88}},
		{str: "hi123,你好88", expected: []int64{123, 88}},
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

func TestGbkToUtf8(t *testing.T) {
	utf8Data, err := strutils.GbkToUtf8(Gb2312Data)
	if err != nil {
		t.Error(err)
	}
	t.Logf("utf8 data: %v", string(utf8Data))
}

func TestUtf8ToGbk(t *testing.T) {
	utf8Data, err := strutils.GbkToUtf8(Gb2312Data)
	if err != nil {
		t.Error(err)
	}
	gbkData, err := strutils.Utf8ToGbk(utf8Data)
	if err != nil {
		t.Error(err)
	}
	if fmt.Sprintf("%v", gbkData) != fmt.Sprintf("%v", Gb2312Data) {
		t.Fatal("Utf8ToGbk编码转换失败, got false, want true")
	} else {
		t.Log("ok")
	}
}

func TestAnyToStr(t *testing.T) {
	now := time.Now()
	testCases := []struct {
		input  interface{}
		output string
	}{
		{input: now, output: now.Format(time.RFC3339)},
		{input: [2]int{123, 456}, output: "123 456"},
		{input: []int{123, 456}, output: "123 456"},
		{input: map[string]interface{}{"aaa": 123, "bbb": 456}, output: "aaa:123 bbb:456"},
		{input: 123, output: "123"},
		{input: 12.3, output: "12.3"},
	}
	for _, v := range testCases {
		str := strutils.AnyToStr(v.input)
		if !strings.HasPrefix(str, v.output) {
			t.Errorf("Testing failure , expected %v, got %v", v.output, str)
		}
	}

}

func TestStrToInt64(t *testing.T)  {
	testCases := []struct {
		input  string
		output float64
	}{
		// {input: "123", output: 123},
		// {input:"1,2,3", output: 123},
		// {input:"abc123", output:123},
		{input: "12.888", output: 12.888},
	}
	for _, v := range testCases {
		ival:= strutils.StrToFloat64(v.input)
		if ival != v.output {
			t.Fatalf("Expected output to be %v, got %v", v.output, ival)
		}
	}
}
