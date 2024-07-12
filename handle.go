package strutils

import "strings"

// 字符串操作相关函数

// 截取指定长度的字符串 支持中文截取
//
//		start 开始截取的字符位置,从1开始 包含1
//		lengths 要截取的字符长度 可选参数, 默认从开始位置一直截取到字符串的末尾
//
//	 如: Substr("你好go语言!",3,4) 返回 "go语言"
//			Substr("Hello world",7,5) 返回 "world"
//
// 返回截取后的字符串
func Substr(str string, start int, lengths ...int) string {
	rstr := []rune(str) // 将字符串转换为rune切片,这样才能支持中文等双字节字符串的截取
	start--             // 切片里面的上下标都是索引 所以这里需要 -1
	length := len(rstr) - start
	endIdx := start + length
	if len(lengths) > 0 {
		// 防止end越界
		if lengths[0] < 0 {
			endIdx = 0
		} else if lengths[0] <= endIdx - start {
			endIdx = lengths[0] + start
		}
	}
	// 防止start越界
	switch {
	case start < 0:
		start = 0
	case start > endIdx:
		start = endIdx
	}
	// 注意这里切片[start:end]里面的start:end规则是 开始索引和结束索引
	return string(rstr[start:endIdx])
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
