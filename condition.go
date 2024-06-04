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

func StrContainsChinese(str string)  bool {
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
