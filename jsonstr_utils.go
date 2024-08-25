// JSON 相关工具函数
package strutils

import (
	"bytes"
	"encoding/json"
)

// 利用JSON decoder扫描data数据到 dest
// 注意这里的 data是JSON字符串的[]byte切片
// dest 是用来接收json data的指针对象, 可以是map 或者自定义结构体,
// 注意json中字段的类型需要和接收对象的字段的数据类型一致, 除非是interface{}
// 另外这里json解码后会将int64类型数据转换为json number类型 而不是默认的 float64类型.
func JsonDataScan(data []byte, dest interface{}) error {
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber() // int类型的数据转换为 json number 而非float64
	return decoder.Decode(&dest)
}
