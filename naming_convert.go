/*
naming conventions utils
命名转换相关实用函数
@author tekintian@gmail.com
*/
package strutils

import (
	"bytes"
	"fmt"
	"strings"
)

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

// 大驼峰命名 单词首字母大写 下划线或者中划线链接的字符串转换为首字母大写的字符串
func CamelStr(str string) string {
	return camelConvert(str, true)
}

// 大驼峰 CamelStr的别名
func CaseCamel(str string) string {
	return camelConvert(str, true)
}

// 小驼峰命名  第一个单词首字母小写,其他首字母大写
func SmallCamelStr(str string) string {
	return camelConvert(str, false)
}

// 小驼峰 SmallCamelStr的别名
func CaseCamelLower(str string) string {
	return camelConvert(str, false)
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

// 将字符串s中所有单词的第一个字母转换为大写
// @author tekintian <tekintian@gmail.com>
func Capitalize(s string) string {
	return UcWords(s)
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

// 转换字符串中单词的首字母大小写
//
//		s 待转换的字符串
//		sep 单词分隔符 如果指定了分隔符且不为空同时字符串中包含指定分隔符,则返回的字符串中的单词将会带上这个分隔符,否则分隔符全部会被设置为空
//	    isUpper 是否转换为大写 true 是, false 否(转换为小写)
//
// 使用示例:
//
//	ConvertWrodsFirstUpperLower("hello word","",true) // HelloWorld
//	ConvertWrodsFirstUpperLower("Hello Word"," ",false) // hello world
//
// 返回转换后的字符串
func ConvertWrodsFirstUpperLower(s, sep string, isUpper bool) string {
	// 定义切割字符串的正则
	regexp := `(\s+|\n|\r|\t|\f|\v|_|-|\b)`
	// 如果sep不为空,且字符串中包含用户提供的分隔符,则将分隔符放入到正则中
	if sep != "" && strings.Contains(s, sep) {
		regexp = fmt.Sprintf(`(%s|\s+|\n|\r|\t|\f|\v|_|-|\b)`, sep)
	} else {
		sep = "" // 其他情况将分隔符设置为空
	}
	re, err := GetRegexp(regexp)
	if err != nil {
		return s
	}
	ss := re.Split(s, -1) // 按照上面的正则切割字符串
	var sb strings.Builder
	sb.Grow(len(ss)) // 指定容量为切割后的切片个数
	for i, v := range ss {
		if v == "" {
			continue
		}
		r0 := []rune(v)[0]
		// 如果单词第一个rune是小写或者大写字母  大写字母 65-90  小写字母97-122
		if r0 <= 122 && ((r0 >= 'A' && r0 <= 'Z') || (r0 >= 'a' && r0 <= 'z')) {
			wr := false
			if isUpper && 'a' <= r0 && r0 <= 'z' {
				wr = true
				r0 -= 'a' - 'A' // 转换为大写 小写字母比大写字母的ascii码大32  注意这里的转换大小写必须的前提
			} else if !isUpper && 'A' <= r0 && r0 <= 'Z' {
				wr = true
				r0 += 'a' - 'A' // 转换为小写
			}
			if wr {
				sb.WriteRune(r0)                      // 将转换后的单词第一个rune写入缓存
				sb.WriteString(string([]rune(v)[1:])) // 写入剩余的rune写入到缓存
			} else {
				sb.WriteString(v)
			}
		} else { // 非小写或者大写字母,直接原样写入
			sb.WriteString(v)
		}
		if sep != "" && i < len(ss)-1 {
			sb.WriteString(sep) // 写入单词分隔符
		}
	}
	if sb.Len() > 0 {
		return sb.String()
	}
	return s // 根据单词分隔符切割后的字符为空,原样返回
}

// 将字符串中所有单词的第一个字母转换为大写
func Title(s string) string {
	if len(s) == 0 {
		return s
	}
	return ConvertWrodsFirstUpperLower(s, " ", true)
}

// 将字符串中所有单词的第一个字母转换为小写
func UnTitle(s string) string {
	if len(s) == 0 {
		return s
	}
	return ConvertWrodsFirstUpperLower(s, " ", false)
}
