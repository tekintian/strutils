package strutils

// 字符串操作相关函数

// 截取指定长度的字符串 支持中文截取
// length 要截取的字符串长度, 此处使用uint16 0--65535长度足够了! 这里避免使用 int 或者 int64 可省去判断负数和少用内存 这样代码更精简更高效
func Substr(str string, length int) string {
	rstr := []rune(str)
	if length > len(rstr) {
		length=len(rstr)
	}
	return string(rstr[:length])
}
