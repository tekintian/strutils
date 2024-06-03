#regexp 正则编译缓存

为什么要用编译缓存？ 为了在高并发和同一个正则被多次使用时获得更好的执行效率！ 这个就是算法中一个典型的利用内存空间换时间的最佳实践

### 正则缓存基准测试

到底效果如何，只能用数据来说话，下面分别对使用和不使用缓存进行了基准测试。

- 为使用读锁的情况:

```sh
// 使用缓存(不使用读锁)和不使用缓存的基准测试结果对比, 使用缓存(不使用读锁)可提高约6.2倍的正则执行效率
> go test -run=^$ -count=3 -bench=.
goos: darwin
goarch: amd64
pkg: github.com/tekintian/go-str-utils
cpu: Intel(R) Core(TM) i7-4770HQ CPU @ 2.20GHz
BenchmarkCacheRegexp-8            373348              2902 ns/op
BenchmarkCacheRegexp-8            388639              2878 ns/op
BenchmarkCacheRegexp-8            389702              2889 ns/op
BenchmarkNormalRegexp-8            63574             18020 ns/op
BenchmarkNormalRegexp-8            63924             18126 ns/op
BenchmarkNormalRegexp-8            64536             18196 ns/op
PASS
ok      github.com/tekintian/go-str-utils       14.675s
```

- 使用读锁的情况

```sh
// 使用缓存和不使用缓存的基准测试结果对比, 使用缓存可提高约6倍的正则执行效率
➜  go-strutils git:(main) ✗ go test -bench=. -benchtime=10s -count=3
goos: darwin
goarch: amd64
pkg: github.com/tekintian/go-str-utils
cpu: Intel(R) Core(TM) i7-4770HQ CPU @ 2.20GHz
BenchmarkRegexp-8                     30         418491027 ns/op
BenchmarkRegexp-8                     28         423420166 ns/op
BenchmarkRegexp-8                     27         474324839 ns/op
BenchmarkCacheRegexp-8           4344793              2906 ns/op
BenchmarkCacheRegexp-8           3867787              3047 ns/op
BenchmarkCacheRegexp-8           3954852              3067 ns/op
BenchmarkNormalRegexp-8           658766             18057 ns/op
BenchmarkNormalRegexp-8           671868             18469 ns/op
BenchmarkNormalRegexp-8           707373             17390 ns/op
PASS
ok      github.com/tekintian/go-str-utils       122.823s
```

说明:
BenchmarkRegexp-8 这个是模拟 100 万个并发的基准测试
BenchmarkCacheRegexp-8 使用了缓存的
BenchmarkNormalRegexp-8 未使用缓存

更多正则使用请参考官方文档： https://pkg.go.dev/regexp

测试文档： https://pkg.go.dev/testing
