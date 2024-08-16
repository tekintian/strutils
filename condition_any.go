package strutils

import (
	"reflect"
)

// 使用放射判断 val是否为空(零值) 可判断 val 对象是否为零值,
// 包含时间类型对象的零值
// 是零值返回 true  否则返回 false; 空字符串,number类型 0 , 引用类型 nil,时间类型空值 都返回true
// @author tekintian <tekintian@gmail.com>
func AnyIsBlank(val interface{}) bool {
	rv := reflect.ValueOf(val)
	switch rv.Kind() {
	case reflect.String:
		return rv.Len() == 0
	case reflect.Bool:
		return !rv.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return rv.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return rv.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return rv.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return rv.IsNil()
	case reflect.Struct:
		// 当前的val是一个结构体, 判断是否有IsZero 这个方法. 如果有这个方法,则通过反射进行调用并获取返回值
		// 这个用于断言时间类型的空值, golang里面时间的零值很特殊 非空,非0也非nil
		isZeroMethod := rv.MethodByName("IsZero")
		if isZeroMethod.IsValid() {
			// 直接调用这个 IsZero() 方法
			// 注意 上面的 IsZero函数是没有参数的 这里要传递一个空[]reflect.Value{}对象
			// 这里的 Call参数返回的结果也是一个 []reflect.Value 这个是根据调用的方法的返回参数决定这个切片里面的数据
			rvs := isZeroMethod.Call([]reflect.Value{})
			return len(rvs) > 0 && rvs[0].Bool() // 注意这里的IsZero方法仅返回1个参数, 所以这里的rvs[0]就是这个参数,且是Bool类型
		}
	}
	// 其他情况  获取反射的接口,然后在创建一个对应类型的反射空对象的接口进行深度比较 如果相等则为空,否则非空
	return reflect.DeepEqual(rv.Interface(), reflect.Zero(rv.Type()).Interface())
}
