package strutils

import (
	crand "crypto/rand"
	"fmt"
	"strings"
)

// 字符串操作相关函数

// 截取指定长度的字符串 支持中文截取
//
//		start 开始截取的字符索引位置,从0开始
//		lengths 要截取的字符长度 可选参数, 默认从开始位置一直截取到字符串的末尾
//
//	 如: Substr("你好go语言!",2,4) 返回 "go语言"
//			Substr("Hello world",6,5) 返回 "world"
//
// 返回截取后的字符串
func Substr(str string, start int, lengths ...int) string {
	rstr := []rune(str) // 将字符串转换为rune切片,这样才能支持中文等双字节字符串的截取
	length := len(rstr) - start
	end := start + length
	if len(lengths) > 0 {
		// 防止end越界
		if lengths[0] < 0 {
			end = 0
		} else if lengths[0] <= end-start {
			end = lengths[0] + start
		}
	}
	// 防止start越界
	switch {
	case start < 0:
		start = 0
		if start+len(rstr) <= end {
			start = start + len(rstr)
		}
	case start > end:
		start = end
	}
	// 注意这里切片[start:end]里面的start:end规则是 开始索引和结束索引
	return string(rstr[start:end])
}

// Substr2 字符串截取 返回截取后的字符串 支持多字节字符
// start 为起始位置.若值是负数,返回的结果将从 str 结尾处向前数第 abs(start) 个字符开始.
// length 为截取的长度.若值时负数, str 末尾处的 abs(length) 个字符将会被省略.
// start/length的绝对值必须<=原字符串长度.
func Substr2(str string, start int, length ...int) string {
	if len(str) == 0 {
		return ""
	}
	// 将字符串转换为rune多字节字符数组
	runes := []rune(str)
	// 统计长度
	total := len(runes)

	var sublen, end int
	max := total // 最大的结束位置

	if len(length) == 0 {
		sublen = total
	} else {
		sublen = length[0]
	}
	if start < 0 {
		start = total + start
	}
	if sublen < 0 {
		sublen = total + sublen
		if sublen > 0 {
			max = sublen
		}
	}
	if start < 0 || sublen <= 0 || start >= max {
		return ""
	}
	end = start + sublen
	if end > max {
		end = max
	}
	return string(runes[start:end])
}

// 去除字符串中的空白字符包含 回车 换行 制表符等, 注意是字符串中的所有的空白符全部去除
func TrimWhiteSpace(s string) string {
	// 使用strings包中的Replacer对空白字符串进行批量 这里的规则都是成对的, 前面是查找字符串 后面是要替换的字符串
	replacer := strings.NewReplacer(" ", "", "\t", "", "\n", "", "\r", "", "\f", "")
	return replacer.Replace(s)
}

// 转换字符串为go语言中安全的命名样式, 即英文字母或者数字与 _ 组合,不能以数字开头
// SafeString会将所有的非字母或者数字全部转换为_  如果是数字开头则转换为_数字, 如123abc 转换为 _123abc
func SafeString(in string) string {
	if len(in) == 0 {
		return in
	}
	// 对输入的字符串的每个字符进行map映射操作处理
	data := strings.Map(func(r rune) rune {
		if isSafeRune(r) {
			return r
		}
		return '_' //将非安全命名字符全部替换为下划线 _
	}, in)
	// 判断第一个字符是否是数字
	firstStr := rune(data[0])
	if isNumber(firstStr) {
		return "_" + data
	}
	return data
}

// 判断是否是go语言中的安全命名字符  即字母,数字 或者下划线 _
func isSafeRune(r rune) bool {
	return isLetter(r) || isNumber(r) || r == '_'
}

func isLetter(r rune) bool {
	return 'A' <= r && r <= 'z'
}

func isNumber(r rune) bool {
	return '0' <= r && r <= '9'
}

// 字符串索引位置查找,找到返回对应的索引位置,  未找到返回 -1
func Index(slice []string, item string) int {
	for i := range slice {
		if slice[i] == item {
			return i
		}
	}
	return -1
}

// 反转字符串 支持所有字符串的反转 中英文符号等都支持
//
//	str 要进行反转的字符串
//
// 返回反转后的字符串
// @author: TekinTian <tekintian@gmail.com>
func ReverseStr(str string) string {
	ss := func(s string) *[]rune {
		var b []rune
		for _, k := range s {
			// 这里利用栈的特性 先进后出 来反转字符串 注意这个defer里面的东西只有当函数体执行完毕后才会被执行
			defer func(v rune) {
				b = append(b, v)
			}(k)
		}
		return &b
	}(str)
	return string(*ss)
}

// 切割并清理字符串, 返回非空的字符串切片
//
//	str 要处理的字符串
//	sep 房子付出的分隔符
//	 cutsets 可选参数 要清理的字符串 默认 空格
//
// 返回切割后非空的字符串切片
func StrSplitAndTrim(str, sep string, cutsets ...string) []string {
	cutset := " "
	if len(cutsets) > 0 {
		cutset = cutsets[0]
	}
	tmps := strings.Split(str, sep)
	ss := make([]string, 0)
	for _, v := range tmps {
		if vt := strings.Trim(v, cutset); vt != "" {
			ss = append(ss, vt)
		}
	}
	return ss
}

// 字符串 子串索引 获取
func StrIndex(str, substr string) int {
	// 子串在字符串的字节位置
	result := strings.Index(str, substr)
	if result >= 0 {
		// 获得子串之前的字符串并转换成[]byte
		prefix := []byte(str)[0:result]
		// 将子串之前的字符串转换成[]rune
		rs := []rune(string(prefix))
		// 获得子串之前的字符串的长度，便是子串在字符串的字符位置
		result = len(rs)
	}
	return result
}

// 反转字符串模板
func ReverseStrTpl(temp, str string, res map[string]interface{}) {
	index := StrIndex(temp, "{")
	ei := StrIndex(temp, "}") + 1
	next := strings.Trim(temp[ei:], " ")
	nextContain := StrIndex(next, "{")
	nextIndexValue := next
	if nextContain != -1 {
		nextIndexValue = Substr(next, 0, nextContain)
	}
	key := temp[index+1 : ei-1]
	// 如果后面没有内容了，则取字符串的长度即可
	var valueLastIndex int
	if nextIndexValue == "" {
		valueLastIndex = len([]rune(str))
	} else {
		valueLastIndex = StrIndex(str, nextIndexValue)
	}
	value := strings.Trim(Substr(str, index, valueLastIndex), " ")
	res[key] = value
	// 如果后面的还有需要解析的，则递归调用解析
	if nextContain != -1 {
		sstr := Substr(str, StrIndex(str, value)+len([]rune(value)), len([]rune(str)))
		ReverseStrTpl(next, sstr, res)
	}
}

// 判断字符串haystack是否包含子串needle 可以指定偏移位置
// 返回索引位置 -1 不包含,其他包含
func Contains(haystack, needle string, startOffset ...int) int {
	length := len(haystack)
	offset := 0
	if len(startOffset) > 0 {
		offset = startOffset[0]
	}
	if length == 0 || offset > length || -offset > length {
		return -1
	}

	if offset < 0 {
		offset += length
	}
	pos := strings.Index(strings.ToLower(haystack[offset:]), strings.ToLower(needle))
	if pos == -1 {
		return -1
	}
	return pos + offset
}

// 正则查找替换字符串
func ReplaceString(pattern, replace, src string) (string, error) {
	if r, err := GetRegexp(pattern); err == nil {
		return r.ReplaceAllString(src, replace), nil
	} else {
		return "", err
	}
}

// GetUuid 获取36位UUID RFC4122
func GetUuid() string {
	u := make([]byte, 16)
	n, err := crand.Read(u)
	if err != nil {
		return fmt.Sprintf("%v", n) // 异常情况下,返回字符串 0
	}
	//sets the version bits
	u[6] = (u[6] & 0x0f) | (3 << 4)
	//sets the variant bits
	u[8] = u[8]&(0xff>>2) | (0x02 << 6)
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%12x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:16])
}
