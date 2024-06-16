package strutils_test

import (
	"fmt"
	"strings"
	"testing"

	strutils "github.com/tekintian/go-str-utils"
)

func TestStrFn(t *testing.T) {

	cases := []struct {
		fn       func(string) string
		param1   string
		expected string
	}{ // 测试用例数据
		{fn: strutils.CamelStr, param1: "hello_world", expected: "HelloWorld"},
		{fn: strutils.CamelStr, param1: "hello          world    hi Xyz", expected: "HelloWorldHiXyz"},
		{fn: strutils.CamelStr, param1: "hello world你好", expected: "HelloWorld你好"},
		{fn: strutils.SmallCamelStr, param1: "hello-world", expected: "helloWorld"},
		{fn: strutils.SnakeStr, param1: "helloWorld", expected: "hello_world"},
		{fn: strutils.KebabStr, param1: "HelloWorld", expected: "hello-world"},
		{fn: strutils.Html2str, param1: "<b>hello world</b>\n\n\n\n<script>< / b>", expected: "hello world"},
		{fn: strutils.UcWords, param1: "hello world-How_are-you", expected: "Hello World How Are You"},
		{fn: strutils.UcWords, param1: "hello_world", expected: "Hello World"},
		{fn: strutils.UcWords, param1: "hello 你好 XYZ", expected: "Hello 你好 XYZ"},
		{fn: strutils.UcFirst, param1: "hello world", expected: "Hello world"},
		{fn: strutils.LcFirst, param1: "Hello World", expected: "hello World"},
		{fn: strutils.Title, param1: "hello golang", expected: "Hello Golang"},
		{fn: strutils.Title, param1: "你好golang1", expected: "你好Golang1"},
		{fn: strutils.Title, param1: " 你好 world", expected: "你好 World"},
		{fn: strutils.UnTitle, param1: "你好GoLang", expected: "你好goLang"},
		{fn: strings.Title, param1: "你好GoLang", expected: "你好GoLang"},
	}

	for _, c := range cases {
		fname := fmt.Sprintf("%p", c.fn)
		t.Run(fname, func(t *testing.T) {
			x := c.fn(c.param1)
			if x != c.expected {
				t.Fatalf("%v Expected %v, got %v", fname, c.expected, x)
			}
		})
	}
}

func TestTitle(t *testing.T) {
	list := []*ieData{
		{input: "_", expected: "_"},
		{input: "abc", expected: "Abc"},
		{input: "ABC", expected: "ABC"},
		{input: "", expected: ""},
		{input: " abc", expected: "Abc"},
		{input: "hello abc", expected: "Hello Abc"},
	}
	for _, v := range list {
		t.Run(v.input, func(t *testing.T) {
			ret := strutils.Title(v.input)
			if ret != v.expected {
				t.Fatalf("Expected %v, got %v", v.expected, ret)
			}
		})
	}
}

func TestUntitle(t *testing.T) {
	list := []*ieData{
		{input: "_", expected: "_"},
		{input: "Abc", expected: "abc"},
		{input: "ABC", expected: "aBC"},
		{input: "", expected: ""},
		{input: " Abc", expected: "abc"},
	}

	for _, v := range list {
		t.Run(v.input, func(t *testing.T) {
			ret := strutils.UnTitle(v.input)
			if ret != v.expected {
				t.Fatalf("Expected %v, got %v", v.expected, ret)
			}
		})
	}
}
