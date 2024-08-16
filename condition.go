package strutils

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/url"
	"strings"
	"unicode/utf8"
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
func IsValidUrl(urlStr string, schemes ...string) bool {
	scheme := "https?"
	if len(schemes) > 0 {
		scheme = schemes[0]
	}
	ssurl, _ := url.PathUnescape(urlStr)
	ssurl = strings.TrimSpace(ssurl)
	regex, _ := GetRegexp(fmt.Sprintf(`^(%s:\/\/)?([\w\./&^#_!-=+$@~*?]+)`, scheme))
	return regex.MatchString(ssurl)
}

// IsUrl 字符串是否URL 优先通过url长度,基本规则和协议进行否点判断,最后通过正则进行判断
func IsUrl(str string) bool {
	if str == "" || len(str) <= 3 || utf8.RuneCountInString(str) >= 2083 || strings.HasPrefix(str, ".") {
		return false
	}
	u, err := url.Parse(str)
	//Couldn't even parse the url
	if err != nil {
		return false
	}

	//Invalid host
	if u.Host == "" || strings.HasPrefix(u.Host, ".") || strings.HasSuffix(u.Host, ":") {
		return false
	}

	//No Scheme found
	if u.Scheme == "" {
		return false
	}

	var inScheme bool
	var schemes = []string{"http", "https", "ftp", "tcp", "udp", "ws", "wss", "irc", "rtmp"}
	for _, s := range schemes {
		if u.Scheme == s {
			inScheme = true
			break
		}
	}
	if !inScheme {
		return false
	}

	return IsValidUrl(str, u.Scheme)
}

// 判断字符串是否是ASCII字符串
func IsASCII(str string) bool {
	for _, v := range str {
		// ASCII字符最大127 十六进制 '\u007F'; 大于127的字符都是非ASCII字符
		// latin1 字符最大255 十六进制 '\u00FF'
		if v > 127 {
			return false
		}
	}
	return true
}

// 通配符 * 问号 ? 匹配, 找出给定的输入字符串str是否与pattern字符串模式相匹配。
//
//		IsWmMatching  wildcard and mask pattern matching
//	 str 要进行匹配的输入字符串
//	 pattern 字符串匹配模式
//
//			通配符 星号'*' -> 星号匹配零个或多个字符。
//			问号'?' -> 匹配任何单个字符。
//
// 如果str与pattern相匹配,则返回true, 否则返回false
func IsWmMatching(str string, pattern string) bool {
	rstrs := []rune(str)
	rpats := []rune(pattern)

	lenInput := len(rstrs)
	lenPattern := len(rpats)

	// 创建一个二维矩阵matrix，其中matrix[i][j] 如果输入字符串中的第一个i字符与模式中的第一个j字符匹配，则为真。
	matrix := make([][]bool, lenInput+1)

	for i := range matrix {
		matrix[i] = make([]bool, lenPattern+1)
	}

	matrix[0][0] = true
	for i := 1; i < lenInput; i++ {
		matrix[i][0] = false
	}

	if lenPattern > 0 {
		if rpats[0] == '*' {
			matrix[0][1] = true
		}
	}

	for j := 2; j <= lenPattern; j++ {
		if rpats[j-1] == '*' {
			matrix[0][j] = matrix[0][j-1]
		}

	}
	for i := 1; i <= lenInput; i++ {
		for j := 1; j <= lenPattern; j++ {

			if rpats[j-1] == '*' {
				matrix[i][j] = matrix[i-1][j] || matrix[i][j-1]
			}

			if rpats[j-1] == '?' || rstrs[i-1] == rpats[j-1] {
				matrix[i][j] = matrix[i-1][j-1]
			}
		}
	}
	return matrix[lenInput][lenPattern]
}

// 星号模式匹配* 和问号模式匹配 ,将模式匹配字符串转换为正则后使用正则进行匹配
// 这个的作用和上面的IsWmMatching 是一样的, 只不过这个函数采用的是正则方式进行模式匹配
func IsWmMatchingReg(str, pattern string) bool {
	// 模式匹配符 *, ? 转换为正则表达式, *替换为 (.*?), .需要进行转义为 \. ; 问号?转换为 (.?)
	rp := strings.NewReplacer("*", `(.*?)`, ".", `\.`, "?", `(.?)`)
	reg := rp.Replace(pattern) // 将v转换为正则表达式
	// 如果正则中不包含前后限定符,则添加上
	if !strings.HasPrefix(reg, "^") && !strings.HasSuffix(reg, "$") {
		reg = fmt.Sprintf(`^%s$`, reg) // 在正则中增加前后限定符
	}
	// 正则对象获取
	regex, err := GetRegexp(reg)
	if err != nil {
		return false
	}
	// 正则匹配
	return regex.MatchString(str)
}
// 判断字符串a是否在list切片中
func StrInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}