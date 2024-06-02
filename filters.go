package strutils

import "strings"

// 字符串过滤相关函数

// Html过滤 过滤所有script,style内容, 其他所有<*>中的内容, 将1个以上的换行回车空格替换为1个空格
// @author: tekintian <tekintian@gmail.com>
func Html2str(htmlStr string) string {
	// script 或者 style 标签内容删除 包括换行的内容 ([\S\s]+?) 非贪婪匹配换行
	re, _ := GetRegexp(`<(script|style)>([\S\s]+?)<(/|/\s+)(script|style)>`)
	htmlStr = re.ReplaceAllString(htmlStr, "")
	// 内联标签尖括号内的内容删除
	re, _ = GetRegexp(`(<[^>]+>)`)
	htmlStr = re.ReplaceAllString(htmlStr, "")
	// 多个换行回车空格替换为一个空格
	re, _ = GetRegexp(`(\n|\t|\r|\s|\\n|\\t|\\r){1,}`) // 将所有1个以上的换行符号和空格替换为1个空格
	htmlStr = re.ReplaceAllString(htmlStr, " ")
	return strings.Trim(htmlStr, " ") // 删除左右空格
}
