package strutils

// 正则编译缓存加速工具 使用原子缓存,可确保高并发情况下的数据一致性
// 高并发,大量相同正则编译的情况下正则的使用性能可提升约6倍

import (
	"fmt"
	"regexp"
	"sync"
	"sync/atomic"
)

var (
	mu          sync.Mutex   // 写锁
	regexpCache atomic.Value // 原子缓存 使用场景 读多写少的情况
)

type expMap map[string]*regexp.Regexp

func init() {
	regexpCache.Store(make(expMap)) //初始化原子缓存对象 避免Load时的nil
}

// 返回 `pattern` 对应的 *regexp.Regexp 对象, 使用原子级别缓存,绝对线程安全!
func GetRegexp(pattern string) (*regexp.Regexp, error) {
	mu.Lock()
	exp := regexpCache.Load().(expMap)[pattern]
	mu.Unlock()
	if exp != nil {
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
	m1 := regexpCache.Load().(expMap) // 加载旧的值
	m1[pattern] = exp
	regexpCache.Store(m1) // 原子操作

	return exp, nil
}
