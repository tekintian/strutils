package strutils_test

import (
	"encoding/base64"
	"fmt"
	"strings"
	"testing"

	strutils "github.com/tekintian/strutils"
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

// base64字符判断单元测试用例
func TestJudgeBase64(t *testing.T) {
	cases := []struct {
		str      string
		expected bool
	}{
		{str: "00000000", expected: false},
		{str: "JTIyMTIzJTIy", expected: true},
		{str: "JTIyMSUyMg==", expected: true},
		{str: "123456 dfadsfdagf!", expected: false},
		{str: "1", expected: false},
		{str: "123", expected: false},
		{str: "1234", expected: false},
		{str: "0000", expected: false},
	}
	for _, v := range cases {
		t.Run(v.str, func(t *testing.T) {
			// 业务逻辑测试
			ret := strutils.JudgeBase64(v.str)
			if ret != v.expected {
				t.Fatalf(" %v test failed, expected %v, got %v", v.str, v.expected, ret)
			}
		})
	}
}

// 模糊测试命令: go test -fuzz=FuzzJudgeBase64 -fuzztime 30s
func FuzzJudgeBase64(f *testing.F) {
	// 测试种子语料
	testcases := []string{
		base64.StdEncoding.EncodeToString([]byte("hello world")),
		base64.StdEncoding.EncodeToString([]byte("你好中国")),
		base64.StdEncoding.EncodeToString([]byte("%$#@@*(*%$%")),
		base64.StdEncoding.EncodeToString([]byte("123")),
	}
	for _, tc := range testcases {
		f.Add(tc) // 提供种子语料库
	}
	f.Fuzz(func(t *testing.T, orig string) {
		// 首先需要对源语料 进行过滤, 因为这些情况 JudgeBase64 一定会返回false;
		judgeRes := strutils.JudgeBase64(orig)

		orig = strings.TrimSpace(orig)
		// 长度非4的倍数
		if (len(orig) < 4 || len(orig)%4 != 0) && !judgeRes {
			t.Skipf("测试数据 %v 长度不符合base64规则,跳过", orig)
		}
		re, _ := strutils.GetRegexp(`^\d+$`)
		//纯数字
		if re.MatchString(orig) && !judgeRes {
			t.Skipf("测试数据 %v 是纯数字,跳过", orig)
		}
		//包含非base64允许字符
		re, _ = strutils.GetRegexp(`^([0-9a-zA-Z+/=]+)$`)
		if !re.MatchString(orig) && !judgeRes {
			t.Skipf("测试数据 %v 包含非base64允许字符,跳过", orig)
		}

		// 其他情况
		if !judgeRes { // 判断结果为非base64
			// 再次进行验证, 将源进行base64解码,然后在对结果进行编码, 如果没有异常,且他们相等 那么说明JudgeBase64 判断失误! 否则判断成功
			b, err := base64.StdEncoding.DecodeString(orig)
			src := base64.StdEncoding.EncodeToString(b)
			if err == nil && fmt.Sprintf("%v", b) == fmt.Sprintf("%v", src) {
				t.Fatalf("%v JudgeBase64 Failed, expected true got false", orig)
			}
		}

	})

	/* FuzzJudgeBase64模糊测试结果
	> go test -fuzz=FuzzJudgeBase64 -fuzztime 30s
	fuzz: elapsed: 0s, gathering baseline coverage: 0/197 completed
	fuzz: elapsed: 0s, gathering baseline coverage: 197/197 completed, now fuzzing with 8 workers
	fuzz: elapsed: 3s, execs: 228550 (76179/sec), new interesting: 1 (total: 198)
	fuzz: elapsed: 6s, execs: 439994 (70460/sec), new interesting: 1 (total: 198)
	fuzz: elapsed: 9s, execs: 621605 (60546/sec), new interesting: 1 (total: 198)
	fuzz: elapsed: 12s, execs: 787822 (55397/sec), new interesting: 2 (total: 199)
	fuzz: elapsed: 15s, execs: 931280 (47820/sec), new interesting: 2 (total: 199)
	fuzz: elapsed: 18s, execs: 1044367 (37695/sec), new interesting: 2 (total: 199)
	fuzz: elapsed: 21s, execs: 1149541 (35062/sec), new interesting: 2 (total: 199)
	fuzz: elapsed: 24s, execs: 1230256 (26902/sec), new interesting: 2 (total: 199)
	fuzz: elapsed: 27s, execs: 1298371 (22711/sec), new interesting: 2 (total: 199)
	fuzz: elapsed: 30s, execs: 1348064 (16565/sec), new interesting: 2 (total: 199)
	fuzz: elapsed: 31s, execs: 1348064 (0/sec), new interesting: 2 (total: 199)
	PASS
	ok      github.com/tekintian/strutils       32.364s
	*/
}

// 准备gb2312编码的[]byte数据
var Gb2312Data = []byte{10, 10, 10, 10, 10, 123, 34, 105, 112, 34, 58, 34, 49, 49, 50, 46, 49, 49, 55, 46, 55, 53, 46, 57, 57, 34, 44, 34, 112, 114, 111, 34, 58, 34, 212, 198, 196, 207, 202, 161, 34, 44, 34, 112, 114, 111, 67, 111, 100, 101, 34, 58, 34, 53, 51, 48, 48, 48, 48, 34, 44, 34, 99, 105, 116, 121, 34, 58, 34, 192, 165, 195, 247, 202, 208, 34, 44, 34, 99, 105, 116, 121, 67, 111, 100, 101, 34, 58, 34, 53, 51, 48, 49, 48, 48, 34, 44, 34, 114, 101, 103, 105, 111, 110, 34, 58, 34, 34, 44, 34, 114, 101, 103, 105, 111, 110, 67, 111, 100, 101, 34, 58, 34, 48, 34, 44, 34, 97, 100, 100, 114, 34, 58, 34, 212, 198, 196, 207, 202, 161, 192, 165, 195, 247, 202, 208, 32, 181, 231, 208, 197, 34, 44, 34, 114, 101, 103, 105, 111, 110, 78, 97, 109, 101, 115, 34, 58, 34, 34, 44, 34, 101, 114, 114, 34, 58, 34, 34, 125, 10, 10, 10, 10}

func TestIsGbk(t *testing.T) {

	if !strutils.IsGbkData(Gb2312Data) {
		t.Errorf("condition failed, expected false, got %v , %v", strutils.IsGbkData(Gb2312Data), string(Gb2312Data))
	} else {
		t.Log("ok")
	}
}
