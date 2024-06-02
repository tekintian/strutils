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
	txt := "你好go语言!"
	if ret := strutils.Substr(txt, 4); ret == "你好go" {
		t.Log("ok")
	} else {
		t.Fatalf("test failed, expected: 你好go got: %v", ret)
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
