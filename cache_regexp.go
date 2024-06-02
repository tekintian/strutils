package strutils

// 正则编译缓存加速工具
// 高并发,大量相同正则编译的情况下正则的使用性能可提升1倍以上

import (
	"fmt"
	"regexp"
	"sync"
)

var (
	regexMu  = sync.RWMutex{}                  // 读写锁
	regexMap = make(map[string]*regexp.Regexp) // 用来缓存正则编译对象的map
)

// 返回 `pattern` 对应的 *regexp.Regexp 对象, 使用内存缓存, 线程安全.
func GetRegexp(pattern string) (*regexp.Regexp, error) {
	exp, _, err := getExp(pattern)
	return exp, err
}

// ct 这个是缓存命中次数
func getExp(pattern string) (exp *regexp.Regexp, ct int, err error) {
	// regexMu.RLock()
	exp, ok := regexMap[pattern]
	// regexMu.RUnlock()
	if ok {
		ct++
		return
	}
	// 不存在缓存,编译正则然后写入map缓存
	if exp, err = regexp.Compile(pattern); err != nil {
		err = fmt.Errorf(`regexp.Compile failed for pattern "%s" errors: %v`, pattern, err)
		return
	}
	// 使用写入锁缓存正则编译对象
	regexMu.Lock() // 写锁
	regexMap[pattern] = exp
	regexMu.Unlock()
	return
}
