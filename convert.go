package strutils

// 字符串转换数字相关函数
// 官方已经提供了一系列的 strconv.xxx函数了,为何还要有这个?
// 因为这里提供的是增强版本的转换, 可转换字符串中包含的数字,小数或者逗号分隔的数字
// 至于数字number 转字符串 推荐使用 fmt.Sprintf("%v",number) 这个内置函数即可
// @author tekintian <tekintian@gmail.com>

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"
	"time"

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
// str 要转换的字符串  如: "123,456"
// seps 可选参数, 字符串中的切割分隔符,默认 逗号, 可以自行根据要切割转换的字符串指定
// 返回按照指定的sep分隔符切割转换后的 uint切片
func Str2UintSlice(str string, seps ...string) []uint {
	ss := make([]uint, 0)
	s := strings.TrimSpace(str)
	if s == "" {
		return ss
	}
	sep := "," // 默认分隔符 逗号
	if len(seps) > 0 {
		sep = seps[0]
	}
	strSs := strings.Split(s, sep)
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
// str 要转换的字符串  如: "123,456"
// seps 可选参数, 字符串中的切割分隔符,默认 逗号, 可以自行根据要切割转换的字符串指定
// 返回按照指定的sep分隔符切割转换后的 int切片
func Str2IntSlice(str string, seps ...string) []int {
	ss := make([]int, 0)
	s := strings.TrimSpace(str)
	if s == "" {
		return ss
	}
	sep := "," // 默认分隔符 逗号
	if len(seps) > 0 {
		sep = seps[0]
	}
	strSs := strings.Split(s, sep)
	for _, v := range strSs {
		ss = append(ss, Str2Int(v))
	}
	return ss
}

// 字符串转int64切片
// str 要转换的字符串  如: "123,456"
// seps 可选参数, 字符串中的切割分隔符,默认 逗号, 可以自行根据要切割转换的字符串指定
// 返回按照指定的sep分隔符切割转换后的 int64切片
func Str2Int64Slice(str string, seps ...string) []int64 {
	ss := make([]int64, 0)
	s := strings.TrimSpace(str)
	if s == "" {
		return ss
	}
	sep := "," // 默认分隔符 逗号
	if len(seps) > 0 {
		sep = seps[0]
	}
	strSs := strings.Split(s, sep)
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

// AnyToStr 返回any类型的数据的字符串
// nil数据返回空字符串,
// 数组,切片 返回以空格分隔的值, map类型的数据返回以空格分隔的 k:v ,
// 如果value是Time对象,则默认使用 time.RFC3339 格式化时间返回
// 实现了Stringer接口的结构体调用 String返回字符串, 字符串类型的直接返回字符串
// 其他类型全部使用 fmt.Sprintf("%v",value) 返回字符串
func AnyToStr(value interface{}) string {
	if value == nil {
		return ""
	}
	// 先使用 type switch来判断数据类型
	switch vt := value.(type) {
	case string:
		return vt
	case []byte:
		return string(vt)
	}
	// 使用 reflect反射方式 处理字符串
	retStr := ""
	// Indirect方法兼容指针或者值
	vrt := reflect.Indirect(reflect.ValueOf(value))
	switch vrt.Kind() {
	case reflect.Struct:
		// 如果value是一个时间对象  这里因为time.Time基本上都是直接使用对象而非指针,所以这里只考虑非指针Time
		if tt, ok := value.(time.Time); ok {
			if tt.IsZero() { // 时间对象里面的零值
				return ""
			}
			return tt.Format(time.RFC3339) // 默认使用 time.RFC3339 格式化时间后返回
		} else if f, ok := value.(fmt.Stringer); ok {
			// 如果value实现了Stringer接口,则调用接口中的String()方法返回数据
			return f.String()
		}
	case reflect.Map:
		vmap := value.(map[string]interface{})
		sb := strings.Builder{}
		for k, v := range vmap {
			sb.WriteString(fmt.Sprintf("%s:%v ", k, v))
		}
		str := sb.String()
		retStr = str[:len(str)-1] // 删除最后一个空格
	// 注意 这里后面的 int,uint,float序列的最终效果和 fmt.Sprintf("%v", value) 的效果是一样的
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(vrt.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(vrt.Uint(), 10)
	case reflect.Float64:
		return strconv.FormatFloat(vrt.Float(), 'g', -1, 64)
	case reflect.Float32:
		return strconv.FormatFloat(vrt.Float(), 'g', -1, 32)
	case reflect.Bool:
		return strconv.FormatBool(vrt.Bool())
	case reflect.Slice, reflect.Array:
		retStr = fmt.Sprintf("%v", vrt.Interface())
		retStr = strings.Trim(retStr, "[]") // 删除数组和切片数据中的[]
	// 其他类型全部使用 %v 来转换了
	default:
		if vrt.IsValid() {
			// 使用%v 格式化为字符串返回
			retStr = fmt.Sprintf("%v", vrt.Interface())
		}
	}
	if retStr == "" {
		retStr = fmt.Sprintf("%v", value)
	}
	return retStr
}

// 使用正则从字符串中解析数字, 支持小数和使用了,分隔的数字的匹配
// 如  123,456 获取的结果为 123456;  abc12.88def   返回 12.88
func ParseNumberStr(str string) string {
	regex, _ := GetRegexp(`([\d\.,]+)`)
	dstr := regex.FindString(str)
	if strings.Contains(dstr, ",") {
		dstr = strings.ReplaceAll(dstr, ",", "") // 删除数字中的分隔符 ,
	}
	return dstr
}

// 使用正则匹配字符串中的数字并转换为 float64
// 支持使用了,分隔的数字的匹配, 如 123,456.99  获取的结果为 123456.99
func strToFloat64(str string) (float64, error) {
	dstr := ParseNumberStr(str)
	// 注意这里因为 ParseInt 无法解析小数(会返回0), 所以转换为其他类型也是使用ParseFloat解析为float64后再强转为int序列的类型
	f64Val, err := strconv.ParseFloat(dstr, 64)
	if err != nil {
		return 0, err
	}
	return f64Val, nil
}

// 字符串到 float64 转换
func StrToFloat64(str string) float64 {
	fval, err := strToFloat64(str)
	if err != nil {
		return 0
	}
	return fval
}

// 字符串到 int 转换
func StrToInt(str string) int {
	fval := StrToFloat64(str)
	return int(fval)
}

// 字符串到 int64 转换
func StrToInt64(str string) int64 {
	fval := StrToFloat64(str)
	return int64(fval)
}

// 字符串到 uint 转换
func StrToUint(str string) uint {
	fval := StrToFloat64(str)
	return uint(fval)
}

// 字符串到 uint64 转换
func StrToUint64(str string) uint64 {
	fval := StrToFloat64(str)
	return uint64(fval)
}

// JsonStrToStruct 反序列化字符串并赋值给对应结构体
func JsonStrToStruct(m string, dst interface{}) error {
	err := json.Unmarshal([]byte(m), dst)
	if err != nil {
		return err
	}
	return nil
}

// TimeToStr 转换 时间字符串/时间戳/时间对象 到字符串
// tval 要转换的时间对象,  时间戳(支持秒和毫秒), 时间字符串(注意,如果时间格式非默认的格式,需要指定时间格式)
// layouts 可选的时间格式
//
//		默认输出字符串格式 "2006-01-02T15:04:05Z07:00",
//		默认tval字符串对应的时间格式 "2006-01-02 15:04:05"
//
//	layouts可以传递多个时间格式参数,
//		第一个为最终返回的字符串格式,默认"2006-01-02T15:04:05Z07:00"
//		第二个为源格式(tval字符串对应的时间格式),默认"2006-01-02 15:04:05",仅tval为字符串时有效,
//			还会自动尝试格式 time.RFC3339, 2006年01月02日15:04:05
//
// 时间字符串
// @author tekintian <tekintian@gmail.com>
// @see https://dev.tekin.cn
func TimeToStr(tval interface{}, layouts ...string) string {
	// 默认时间格式,
	toLayout := time.RFC3339   // 默认转换后的字符串对应的时间格式 "2006-01-02T15:04:05Z07:00"
	srcLayout := time.DateTime // 默认tval对应的时间格式 "2006-01-02 15:04:05"
	switch len(layouts) {
	case 1:
		if layouts[0] != "" {
			toLayout = layouts[0]
		}
	case 2:
		if layouts[0] != "" {
			toLayout = layouts[0]
		}
		if layouts[1] != "" {
			srcLayout = layouts[1]
		}
	}
	// 根据不同类型选择对应的解析方式
	switch v := tval.(type) {
	case time.Time: // 时间对象
		return v.Format(toLayout)
	case int: // 时间戳解析
		return TsToStr(int64(v), toLayout)
	case uint:
		return TsToStr(int64(v), toLayout)
	case int64:
		return TsToStr(v, toLayout)
	case uint64:
		return TsToStr(int64(v), toLayout)
	case int32:
		return TsToStr(int64(v), toLayout)
	case uint32:
		return TsToStr(int64(v), toLayout)
	case string: // 字符串解析
		// 如果是字符串类型，将其解析为时间对象
		if t, err := time.Parse(srcLayout, v); err == nil {
			return t.Format(toLayout)
		}
		// 指定的时间格式解析失败,使用常用格式再次尝试
		if t, err := time.Parse(time.RFC3339, v); err == nil {
			return t.Format(toLayout)
		}
		// 使用这个 2006年01月02日15:04:05 格式尝试解析
		if t, err := time.Parse("2006年01月02日15:04:05", v); err == nil {
			return t.Format(toLayout)
		}
		return "" // 解析失败,返回空字符串
	default:
		return ""
	}
}

// 格式化时间戳到字符串
// ts 时间戳 支持秒 10位, 毫秒 13位
// layout 时间格式 如果为空,则默认使用 time.RFC3339 作为格式 "2006-01-02T15:04:05Z07:00"
// 格式后的时间字符串
func TsToStr(ts int64, layout string) string {
	var t time.Time
	if ts < 9999999999 { // 当前的时间戳是秒 秒的时间戳最长10位, 毫秒的时间戳长度为13位
		t = time.Unix(ts, 0)
	} else {
		t = time.UnixMilli(ts)
	}
	// 格式化时间字符串
	if layout == "" {
		layout = time.RFC3339
	}
	return t.Format(layout)
}

// NumberFormat 以千位分隔符方式格式化一个数字.
// decimal为要保留的小数位数,point为小数点显示的字符,thousand为千位分隔符显示的字符.
// 有效数值是长度(包括小数点)为17位之内的数值,最后一位会四舍五入.
func NumberFormat(number float64, decimal uint8, point, thousand string) string {
	neg := false
	if number < 0 {
		number = -number
		neg = true
	}
	dec := int(decimal)
	// Will round off
	str := fmt.Sprintf("%."+strconv.Itoa(dec)+"F", number)
	prefix, suffix := "", ""
	if dec > 0 {
		prefix = str[:len(str)-(dec+1)]
		suffix = str[len(str)-dec:]
	} else {
		prefix = str
	}
	sep := []byte(thousand)
	n, l1, l2 := 0, len(prefix), len(sep)
	// thousands sep num
	c := (l1 - 1) / 3
	tmp := make([]byte, l2*c+l1)
	pos := len(tmp) - 1
	for i := l1 - 1; i >= 0; i, n, pos = i-1, n+1, pos-1 {
		if l2 > 0 && n > 0 && n%3 == 0 {
			for j := range sep {
				tmp[pos] = sep[l2-j-1]
				pos--
			}
		}
		tmp[pos] = prefix[i]
	}
	s := string(tmp)
	if dec > 0 {
		s += point + suffix
	}
	if neg {
		s = "-" + s
	}

	return s
}

// 解析字符串切片 到 map[string][]string 这个的应用场景通常是对权限验证字符串进行切割和转换
//
//		strs 要进行切割转换的字符串切片 , 如: []string{"/api/login:GET,POST"}
//		options 字符串切割可选项,可以传3个可选项参数:
//		 	第一个是字符串的key和值的分隔符 默认 分号 :
//			第二个是第一个分割后的值的分隔符, 默认 逗号 ,
//	 	第二个是分割后的key/val值的trim 字符, 默认 空格
//
//	 返回 map[string][]string  如: mps["/api/login"]=[]string{"GET","POST"}
//
// @author tekintian <tekintian@gmail.com>
func ParseStrsToMapStrStrSlice(strs []string, options ...string) map[string][]string {
	mps := make(map[string][]string) // 定义要返回的map对象
	var sepKv, sepV, cutStr string   // 默认移除的字符串 空白字符
	// 可选参数赋值
	switch len(options) {
	case 0: // 默认的字符串切割规则 即没有options参数的情况
		sepKv = ":"
		sepV = ","
		cutStr = " "
	case 1:
		sepKv = options[0]
	case 2:
		sepKv = options[0]
		sepV = options[1]
	case 3:
		sepKv = options[0]
		sepV = options[1]
		cutStr = options[2]
	}

	for _, str := range strs {
		kvals := strings.Split(str, sepKv) // 这里使用 sepKv 分字符串进行分割为 key 和vals
		if len(kvals) != 2 {
			continue // 跳过不符合规则的内容
		}
		key := strings.Trim(kvals[0], cutStr)           // 第一部分为map的key
		vals := StrSplitAndTrim(kvals[1], sepV, cutStr) // 这个是vals的值,使用 sepV对其进行分割并去除 cutStr
		mps[key] = vals
	}
	return mps
}
