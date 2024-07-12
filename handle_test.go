package strutils_test

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"

	strutils "github.com/tekintian/go-str-utils"
)

func TestSubstr(t *testing.T) {
	cases := []struct {
		str     string
		start   int
		lengths int
		out     string
	}{
		{str: "530102001", start: 2, lengths: 1, out: "0"},
		{str: "530102001", start: 2, lengths: 3, out: "010"},
		{str: "你好go语言!", start: 2, lengths: 2, out: "go"},
		{str: "你好go语言!", start: 4, lengths: 12, out: "语言!"},
		{str: "hello world!", start: 6, lengths: 5, out: "world"},
	}
	for _, v := range cases {
		if ret := strutils.Substr(v.str, v.start, v.lengths); ret != v.out {
			t.Fatalf("Substr %v failed, want %v , got %v ", v.str, v.out, ret)
		}
	}
}

// 模糊测试命令: go test -fuzz=FuzzSubStr -fuzztime 30s
func FuzzSubStr(f *testing.F) {
	testcases := []string{"Hello world", "你好,中国"}
	for _, tc := range testcases {
		f.Add(tc) // 提供种子语料库
	}
	// 随机种子
	rand.NewSource(time.Now().UnixNano())
	// 执行测试
	f.Fuzz(func(t *testing.T, orig string) {
		if orig == "" { // 空字符串不参与测试,因为 rand.Intn(0) 会抛panic异常
			return
		}
		ro := []rune(orig)
		l := rand.Intn(len(ro)) // 获取随机获取的字符串长度,这个随机数必须是 0 -- 字符串长度
		// 执行函数获取结果 ret
		ret := strutils.Substr(orig, l)

		// 获取随机长度的原始字符串
		var buf strings.Builder
		for i := 0; i < l; i++ {
			if i >= len(ro) {
				break
			}
			fmt.Fprintf(&buf, "%s", string(ro[i]))
		}
		// 和返回结果对比
		if ret != buf.String() {
			t.Fatalf("test failed, expected %v, got %v", orig[:l], ret)
		}
	})
}

type ieData struct {
	input    string
	expected string
}

func TestSafeString(t *testing.T) {
	list := []*ieData{
		{input: " hello \n\r\tworld", expected: "_hello____world"},
		{input: "_", expected: "_"},
		{input: "a-b-c", expected: "a_b_c"},
		{input: "123abc", expected: "_123abc"},
		{input: "汉abc", expected: "_abc"},
		{input: "汉a字", expected: "_a_"},
		{input: "a_B C", expected: "a_B_C"},
		{input: "A#B#C", expected: "A_B_C"},
		{input: "_123", expected: "_123"},
		{input: "", expected: ""},
		{input: "\t", expected: "_"},
		{input: "\n", expected: "_"},
	}
	for _, v := range list {
		t.Run(v.input, func(t *testing.T) {
			ret := strutils.SafeString(v.input)
			if ret != v.expected {
				t.Fatalf("Expected %v, got %v", v.expected, ret)
			}
		})
	}
}

func TestTrimWhiteSpace(t *testing.T) {
	list := []*ieData{
		{input: " hello \n\r\tworld", expected: "helloworld"},
		{input: " a-b-c\r\t", expected: "a-b-c"},
		{input: "\t\r\n\n12 3a \nbc", expected: "123abc"},
	}
	for _, v := range list {
		t.Run(v.input, func(t *testing.T) {
			ret := strutils.TrimWhiteSpace(v.input)
			if ret != v.expected {
				t.Fatalf("Expected %v, got %v", v.expected, ret)
			}
		})
	}
}

func TestIndex(t *testing.T) {
	list := []string{"a", "b", "c"}

	if ret := strutils.Index(list, "b"); ret != 1 {
		t.Fatalf("Expected 1, got %v", ret)
	}
	if ret := strutils.Index(list, "d"); ret != -1 {
		t.Fatalf("Expected -1, got %v", ret)
	}
}

func TestReverseStr(t *testing.T) {
	testCases := []struct {
		In  string
		Out string
	}{
		{In: "你好Hello", Out: "olleH好你"},
		{In: "Hello world", Out: "dlrow olleH"},
	}
	for _, v := range testCases {
		rstr := strutils.ReverseStr(v.In)
		if rstr != v.Out {
			t.Fatalf("ReverseStr %v Failed!  expected: %v, got: %v", v.In, v.Out, rstr)
		}
	}

}
