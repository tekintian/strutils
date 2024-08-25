package strutils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// 生成指定位数的随机数字字符串
func GenRandIntStr(width uint) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.New(rand.NewSource(time.Now().UnixNano()))

	var sb strings.Builder
	for i := 0; i < int(width); i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

// 生成指定位数的随机字符
func GenRandStr(width uint) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.New(rand.NewSource(time.Now().UnixNano()))
	l := len(letterRunes)
	b := make([]rune, width)
	for i := range b {
		b[i] = letterRunes[rand.Intn(l)]
	}
	return string(b)
}
