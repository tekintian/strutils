package strutils

// 字符串转换数字相关函数
// 官方已经提供了一系列的 strconv.xxx函数了,为何还要有这个?
// 因为这里提供的是增强版本的转换, 可转换字符串中包含的数字,小数或者逗号分隔的数字
// 至于数字number 转字符串 推荐使用 fmt.Sprintf("%v",number) 这个内置函数即可
// @author tekintian <tekintian@gmail.com>

import (
	"bytes"
	"io"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// 从字符串中匹配数字 返回数字对应的字符串
// @TODO 这个地方可以使用参数泛型来进行约束,直接返回对应的类型,不过需要go版本大于 1.18, 这里暂时返回字符串
// @author tekintian <tekintian@gmail.com>
func parseNumFromStr(str string) string {
	str = strings.TrimSpace(str) // 先修剪字符串左右的空格
	if str == "" {
		return ""
	}
	reg, _ := GetRegexp(`([\d,]+(\.\d+)?)`) // 可匹配整数,小数或者带有逗号分隔的数字
	ns := reg.FindString(str)               // 这里如果有多个数字,则只匹配第1个
	if strings.IndexByte(ns, ',') >= 0 {
		ns = strings.ReplaceAll(ns, ",", "") // 去除字符串中包含的逗号
	}
	return ns // 返回数字字符串,调用者需要什么类型自己转换即可
}

// 转换字符串到int64
func parseStrToInt64(str string) int64 {
	number := parseNumFromStr(str)
	num := strings.Split(number, ".")[0] // 处理小数情况, 需要删除小数部分 否则ParseInt会报错
	n, _ := strconv.ParseInt(num, 10, 64)
	return n
}

// 字符串转uint
func Str2Uint(str string) uint {
	return uint(parseStrToInt64(str))
}

// 字符串转uint切片
func Str2UintSlice(str string) []uint {
	ss := make([]uint, 0)
	s := strings.TrimSpace(str)
	if s == "" {
		return ss
	}
	strSs := strings.Split(s, ",")
	for _, v := range strSs {
		ss = append(ss, Str2Uint(v))
	}
	return ss
}

// 字符串转int
func Str2Int(str string) int {
	return int(parseStrToInt64(str))
}

// 字符串转int64
func Str2Int64(str string) int64 {
	return parseStrToInt64(str)
}

// 字符串转int切片
// 这里先将字符串按照, 逗号切割为字符串切片,然后在进行匹配和转换
func Str2IntSlice(str string) []int {
	ss := make([]int, 0)
	s := strings.TrimSpace(str)
	if s == "" {
		return ss
	}
	strSs := strings.Split(s, ",")
	for _, v := range strSs {
		ss = append(ss, Str2Int(v))
	}
	return ss
}

// 字符串转int64切片
// 这里先将字符串按照, 逗号切割为字符串切片,然后在进行匹配和转换
func Str2Int64Slice(str string) []int64 {
	ss := make([]int64, 0)
	s := strings.TrimSpace(str)
	if s == "" {
		return ss
	}
	strSs := strings.Split(s, ",")
	for _, v := range strSs {
		ss = append(ss, Str2Int64(v))
	}
	return ss
}

// 字符串转 float32
func Str2Float32(str string) float32 {
	num, err := strconv.ParseFloat(parseNumFromStr(str), 32)
	if err != nil {
		num = 0
	}
	return float32(num)
}

// 字符串转 float64
func Str2Float64(str string) float64 {
	num, err := strconv.ParseFloat(parseNumFromStr(str), 64)
	if err != nil {
		num = 0
	}
	return num
}

// gbk to utf8 encoding conversion
func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := io.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

// utf8 to gbk encoding conversion
func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := io.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

// 字符串编码 gbk到utf8转换
func StrGbkToUtf8(str string) (string, error) {
	data, err := GbkToUtf8([]byte(str))
	return string(data), err
}

// 字符串编码 utf8到gbk转换
func StrUtf8ToGbk(str string) (string, error) {
	data, err := Utf8ToGbk([]byte(str))
	return string(data), err
}
