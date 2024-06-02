package strutils_test

import (
	"fmt"
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
