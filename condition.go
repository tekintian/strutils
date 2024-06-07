package strutils

import (
	"fmt"
	"log"
)

// 字符串条件判断相关函数

func StrIsChinese(str string) bool {
	reg, err := GetRegexp("^([\u4e00-\u9fa5]+)$")
	if err != nil {
		log.Fatal(err)
		return false
	}
	return reg.MatchString(str)
}

func StrContainsChinese(str string) bool {
	reg, err := GetRegexp("([\u4e00-\u9fa5]+)")
	if err != nil {
		log.Fatal(err)
		return false
	}
	return reg.MatchString(str)
}

// 判断字符串是否包含连续的数字, minCs 最小连续数字的长度,默认 2
func StrContainsContinuousNum(str string, minCs ...uint16) bool {
	var minlen uint16 = 2 // 最低的连续数组长度 2
	if len(minCs) > 0 && minCs[0] > 0 {
		minlen = minCs[0]
	}
	reg, _ := GetRegexp(fmt.Sprintf(`(\d{%v,})`, minlen))

	return reg.MatchString(str)
}

// 判断字符串是否为空或者是空白字符
func IsEmptyStringOrWhiteSpace(s string) bool {
	v := TrimWhiteSpace(s)
	return len(v) == 0
}

// 字符串包含检查 runes 这个为你要检查的字符串rune 可以是多个,只要有一个包含即返回true
func ContainsAny(s string, runes ...rune) bool {
	if len(runes) == 0 {
		return true
	}
	tmp := make(map[rune]byte, len(runes))
	for _, r := range runes {
		tmp[r] = 1
	}

	for _, r := range s {
		if _, ok := tmp[r]; ok {
			return true
		}
	}
	return false
}

// 检测字符串是否包含空白字符
func ContainsWhiteSpace(s string) bool {
	wrs := []rune{'\n', '\t', '\f', '\v', ' '}
	return ContainsAny(s, wrs...)
}
