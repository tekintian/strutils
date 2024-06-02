package strutils_test

import (
	"testing"

	strutils "github.com/tekintian/go-str-utils"
)

func TestHtml2str(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{ // 定义测试用例
		// {input: "<ssss>hello world<./>", expected: "hello world"},
		// {input: "hello<h1> world</h1>", expected: "hello world"},
		// {input: "<zx>hello world     hi<hh>", expected: "hello world     hi"},
		{input: `
		 <script>
      (function() {
        const theme = document.cookie.match(/prefers-color-scheme=(light|dark|auto)/)?.[1]
        if (theme) {
          document.querySelector('html').setAttribute('data-theme', theme);
        }
      }())
    </script>
    <meta charset="utf-8">
	  <title>strings package \n\r\t- strings - Go Packages</title>
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="Description" content="Package strings implements simple functions to manipulate UTF-8 encoded strings.">
		`, expected: "strings package - strings - Go Packages"},
	}
	// test
	for _, v := range cases {
		t.Run(v.input, func(t *testing.T) {
			ret := strutils.Html2str(v.input)
			if ret != v.expected {
				t.Fatalf("test failed, expected %v got %v", v.expected, ret)
			}
		})
	}
}
