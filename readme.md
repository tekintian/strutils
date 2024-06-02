# Go 语言字符串相关实用工具库 go string related utils

本仓库只干一件事, 就是只做 go 语言中的字符串相关的实用函数或者方法， 当然是官方库 https://pkg.go.dev/strings 中未实现的，官方已有的实现我们不重复造轮子！

欢迎大家踊跃 PR，一起完善这个工具库！

如果你喜欢，请帮忙点亮免费小红星，您的小红星就是我们持续更新的动力！

## 目前已完成函数

1. 各种命名转换函数 naming_convert.go
   **包括:**

- CamelStr 大驼峰 单词全部首字母大写 如: UserName

- SmallCamelStr 小驼峰 第一个单词首字母小写,其他大写, 如: userName

- SnakeStr 单词全部小写,使用下划线链接, 如: user_name

- KebabStr 单词全部小写,使用 中划线链接 如: user-name

2. 字符串过滤函数 filters.go

- Html2str html 字符串过滤函数，过滤所有 script,style 内容, 其他所有<\*>中的内容, 将 1 个以上的换行回车空格替换为 1 个空格

3. 字符串操作函数 handle.go

- Substr 截取指定长度的字符串,支持英文,中文和其他多字节字符串

4.  Regexp 编译缓存 cache_regexp.go

- 这个工具就是为了提高高并发和同一个正则多次使用场景下的正则执行效率，因为字符串处理中经常要用到正则
- [正则编译缓存基准测试结果请点击这里查看](docs/regexp_cache_benchmark.md)

## 使用方法

go 版本环境, 最低 1.16, 我们建议你使用官方最新正式版, 这样就可以使用很多高级特性,比如 1.21 版本以上中的 slices 内置包等......

- 安装最新版本依赖

```sh
go get -u github.com/tekintian/go-str-utils
```

- 使用示例

```go
import (
	"fmt"
	strutils "github.com/tekintian/go-str-utils"
)

func main() {
  camelStr := strutils.CamelStr("hello-world")
  fmt.Println(camelStr) // 输出： HelloWorld
}

```

其他使用示例可参考 \*\_test.go 中的测试用例， 我们的每个函数都有测试用例， 部分函数还使用了 Fuzz 模糊测试技术 来确保我们函数的可靠和高效！ 同时也能帮助我们发现一些潜在的风险等......

也可参考 go pkg 文档,地址: https://pkg.go.dev/github.com/tekintian/go-str-utils

## go 编码规范

每个函数需要有对应的测试函数,以确保函数的可靠性, 编码规范请参考 go 官方的 [Effective Go](https://go.dev/doc/effective_go.html)

## go 参考文档

Go 语言内置包 strings 文档 https://pkg.go.dev/strings

Go 程序设计语言规范 https://go.dev/ref/spec

Fuzzing 模糊测试官方使用文档 https://go.dev/security/fuzz
