package strutils

// 正则编译缓存加速工具
// 高并发,大量相同正则编译的情况下正则的使用性能可提升1倍以上

import (
	"fmt"
	"regexp"
	"sync"
)

var (
	mu       sync.RWMutex                      // 读写锁
	expCache = make(map[string]*regexp.Regexp) // 初始化expMap容器
)

// 返回 `pattern` 对应的 *regexp.Regexp 对象, 使用内存缓存, 线程安全.
func GetRegexp(pattern string) (*regexp.Regexp, error) {
	mu.RLock()
	exp, ok := expCache[pattern]
	mu.RUnlock()
	if ok {
		return exp, nil
	}
	return storeExp(pattern)
}

func storeExp(pattern string) (*regexp.Regexp, error) {
	// 不存在缓存,编译正则然后写入map缓存
	mu.Lock()         // 加锁
	defer mu.Unlock() //退出时释放锁
	exp, err := regexp.Compile(pattern)
	if err != nil {
		return nil, fmt.Errorf(`regexp.Compile failed for pattern "%s" errors: %v`, pattern, err)
	}
	expCache[pattern] = exp
	return exp, nil
}
