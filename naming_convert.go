package strutils

import (
	"bytes"
	"fmt"
	"strings"
)

// naming conventions utils
// 命名转换相关实用函数

// 大驼峰命名 单词首字母大写 下划线或者中划线链接的字符串转换为首字母大写的字符串
func CamelStr(str string) string {
	return camelConvert(str, true)
}

// 小驼峰命名  第一个单词首字母小写,其他首字母大写
func SmallCamelStr(str string) string {
	return camelConvert(str, false)
}

// 驼峰命名转换
// @author tekintian <tekintian@gmail.com>
func camelConvert(str string, isBig bool) string {
	// 如果有 - 链接,统一将其替换为 _
	if strings.IndexByte(str, '-') != -1 {
		str = strings.Replace(str, "-", "_", -1)
	}
	//处理空格
	if strings.IndexByte(str, ' ') != -1 {
		re, _ := GetRegexp(`(\s{1,})`)
		str = re.ReplaceAllString(str, "_") // 使用正则将多个空额替换为一个下划线
	}
	str = strings.ToLower(str)
	//按下划线切割字符串为单词
	tmp := strings.Split(str, "_")
	for i, w := range tmp {
		// 判断单词第一个字母是否小写字母
		if (isBig || (!isBig && i > 0)) && w[0] >= 'a' && w[0] <= 'z' {
			// 这个地方的 w[0]-32 即将单词的第一个字母转换为小写.
			// 利用ascii码值差来转换, 小写字母的ascii码值比大写字母的ascii码值大32
			tmp[i] = fmt.Sprintf("%v%v", string(w[0]-32), string(w[1:]))
		}
		// 单词首字母非大写字母的情况不需要处理
	}
	return strings.Join(tmp, "") // 将切片拼接为字符串后返回
}

// 即将单词首字母小写, 多个单词使用下划线链接, 如 UserName 转换为 user_name
func SnakeStr(str string) string {
	return toDeliConvert(str, "_")
}

// kebab 所有字母小写,使用中划线 - 链接
func KebabStr(str string) string {
	return toDeliConvert(str, "-")
}

// 全部字母小写,并按照指定分隔符 delimiter来链接字符串
func toDeliConvert(str, delimiter string) string {
	if str == "" {
		return str
	}
	// 如果含有大写字母,将他们转换为小写字母并添加一个前缀 delimiter
	bsStr := []byte(str)
	bbuf := new(bytes.Buffer)
	for i, v := range bsStr {
		// 第2个字符以后得大写字母转换为 delimiter小写字母
		if i > 0 && v >= 'A' && v <= 'Z' {
			fmt.Fprintf(bbuf, "%s%c", delimiter, v+32)
		} else {
			fmt.Fprintf(bbuf, "%c", v)
		}
	}
	str = bbuf.String()
	// 全部转为小写
	str = strings.ToLower(str)
	// 临时分隔符定义, 将他和 delimiter 区分, 目前分隔符 - 或者 _
	var tde byte = '-'
	if delimiter[0] == tde {
		tde = '_'
	}
	if strings.IndexByte(str, tde) != -1 {
		str = strings.Replace(str, string(tde), delimiter, -1)
	}
	return str
}

// 将每个单词的首字母转大写
// str待转换的字符串 默认分隔空格和下划线链接的字符串, 下划线会被转换为空额输出
func UcWords(str string) string {
	if str == "" {
		return str
	}
	sep := " "
	re, _ := GetRegexp(`(_|-|\s+)`) // 处理_ _ 或者空格分隔的字符
	if re.MatchString(str) {
		str = re.ReplaceAllString(str, sep)
	}
	//切割字符串
	tmp := strings.Split(str, sep)
	var buf strings.Builder
	for i, w := range tmp {
		if i == len(tmp)-1 {
			sep = "" //最后一个单词不添加分隔符
		}
		switch {
		case w[0] >= 97 && w[0] <= 122: // 第一个字符的aceii码必须是小写英文字母 97--122 才转
			fmt.Fprintf(&buf, "%v%v%v", string(w[0]-32), string(w[1:]), sep)
		default:
			fmt.Fprintf(&buf, "%v%v", w, sep)
		}
	}
	return buf.String()
}

// 首字母大写 Upper case first letter
func UcFirst(str string) string {
	if len(str) < 1 {
		return ""
	}
	rs := []rune(str)
	if rs[0] >= 97 && rs[0] <= 122 {
		rs[0] -= 32
	}
	return string(rs)
}

// 首字母小写 Lower case first letter
func LcFirst(str string) string {
	if len(str) < 1 {
		return ""
	}
	rs := []rune(str)
	if rs[0] >= 'A' && rs[0] <= 'Z' {
		rs[0] += 32
	}
	return string(rs)
}
