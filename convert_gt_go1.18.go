//go:build go1.18
// +build go1.18

// 泛型的字符串相关转换函数
// 这里仅支持go版本大于等于1.18的go运行环境,因为只有go版本大于等于1.18才支持泛型
// @author: tekintian@gmail.com
// @see https://dev.tekin.cn
package strutils


// 用于泛型约束的接口
type iNumber interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | uint | uint8 | uint16 | uint32 | uint64
}

// 泛型字符串转数字, 返回的类型为defVal对应的类型,如果转换失败,则返回默认值
// 支持将字符串转换为go中所有的intx,uintx,floatx的数字类型
func StrToNumber[T iNumber](str string, defVal T) T {
	fval, err := strToFloat64(str)
	if err != nil {
		fval = float64(defVal)
	}
	// 强转为指定的T类型
	return T(fval)
}

