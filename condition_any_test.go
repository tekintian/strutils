package strutils_test

import (
	"testing"
	"time"

	"github.com/tekintian/strutils"
)

func TestAnyIsBlank(t *testing.T) {
	// 定义测试用例数据
	var ztime time.Time // 定义一个时间变量 默认 零值
	cases := []struct {
		in  interface{} // 输入参数
		out bool        // 期望结果
	}{
		{in: "", out: true},
		{in: ztime, out: true},
		{in: time.Now(), out: false},
	}
	for _, v := range cases {
		isBlank := strutils.AnyIsBlank(v.in)
		if isBlank != v.out { //如果返回的结果和预期的不一样,测试失败
			t.Fatalf("failed test %v: expected %v, got %v", v.in, v.out, isBlank)
		}
	}
}
