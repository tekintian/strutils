package strutils

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/url"
	"strings"
)

// 字符串条件判断相关函数

// 判断字符串是否是中文字符串
func StrIsChinese(str string) bool {
	reg, err := GetRegexp("^([\u4e00-\u9fa5]+)$")
	if err != nil {
		log.Fatal(err)
		return false
	}
	return reg.MatchString(str)
}

// 判断字符串是否包含中文字符
func StrContainsChinese(str string) bool {
	reg, err := GetRegexp("([\u4e00-\u9fa5]+)")
	if err != nil {
		log.Fatal(err)
		return false
	}
	return reg.MatchString(str)
}

// 判断字符串str是否为数字 整数或者浮点数都算是数字
func StrIsNumber(str string) bool {
	re, _ := GetRegexp(`^(\d+|.\d+|\d+.\d+)$`)
	return re.MatchString(str)
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

// 判断字符串是否是base64字符 可准确判断是或者否   纯数字被视为非base64字符
//
// Base64规定: 4 个 Base64 字符表示 3 个二进制字节，因为：64 * 64 * 64 * 64 = 256 * 256 * 256。
// 因为，Base64 将三个字节转化成四个字节，因此 Base64 编码后的文本，会比原文本大出三分之一左右。
// Base64共有65个字符,他们分别是: 大写字母 A-Z、小写字母 a-z、数字 0-9、符号 "+"、"/"（再加上作为垫字的 "="，
// 垫字是当生成的 Base64 字符串的个数不是 4 的倍数时，添加在尾部的字符），作为一个基本字符集。然后，其他所有符号都转换成这个字符集中的字符。
// @author: TekinTian <tekintian@gmail.com>
// @see https://dev.tekin.cn/
func JudgeBase64(str string) bool {
	// 先断言字符串的长度是否符合base64规范, 即必须是4的倍数; 把这个放在开始,因为这个判断的时间复杂度最低
	if len(str) < 4 || len(str)%4 != 0 {
		return false
	}
	// 在用正则来判断 字符串的规则是否符合
	// 纯数字 这个是一个特例,他不是base64但是 某些情况下可以通过base64编码和解码 这里将他们全部视为非base64字符
	re, _ := GetRegexp(`^\d+$`)
	if re.MatchString(str) {
		return false
	}
	// 正则后向非断言  (?!\d+$) 断言后面非连续的纯数字
	// 判断字符串是否符合base64的编码格式 即是由64个字符和垫字符组成
	re, _ = GetRegexp(`^([A-Za-z0-9+/]{4})*([A-Za-z0-9+/]{4}|[A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{2}==)$`)
	matched := re.MatchString(str)
	if !matched {
		return false
	}
	// 在使用自身带的解码和转码 在查看他么的值是否相等,如果相等说明是base64 否则不是
	deStr, err := base64.StdEncoding.DecodeString(str)
	if err != nil { // 无法解码,肯定不是base64
		return false
	}
	bs64Str := base64.StdEncoding.EncodeToString(deStr)
	//如果解码后再转码和原来的字符一样说明是base64 否则不是
	return str == bs64Str
}

// JudgeBase64 别名
func IsBase64Str(str string) bool {
	return JudgeBase64(str)
}

// 判断数据是否是gbk编码
func IsGbkData(data []byte) bool {
	length := len(data)
	var i int = 0
	for i < length {
		// // ASCII 编码的范围:  十进制 => 0 - 127 。  十六进制： 0x00  -  0x7F 。
		if data[i] <= 127 {
			i++
			continue
		} else {
			// GB2312编码的范围: 十进制 => 高位字节：161 - 247, 十六进制：0xA1 - 0xF7
			// 低位字节：161 - 254 , 十六进制：0xA1 - 0xFE
			if data[i] >= 129 &&
				data[i] <= 254 &&
				data[i+1] >= 64 &&
				data[i+1] <= 254 &&
				data[i+1] <= 247 {
				i += 2
				continue
			} else {
				return false
			}
		}
	}
	return true
}

// 判断字符串是否是gbk编码
func IsGbkStr(str string) bool {
	if str == "" {
		return false
	}
	return IsGbkData([]byte(str))
}

// 验证是否有效的URL
func IsValidUrl(urlStr string) bool {
	ssurl, _ := url.PathUnescape(urlStr)
	ssurl = strings.TrimSpace(ssurl)
	regex, _ := GetRegexp(`^(https?:\/\/)?([\w\./&^#_!-=+$@~*?]+)`)
	return regex.MatchString(ssurl)
}
