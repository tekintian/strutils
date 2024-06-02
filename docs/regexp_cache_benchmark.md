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
> go test -bench=. -benchtime=10s -count=10
goos: darwin
goarch: amd64
pkg: github.com/tekintian/go-str-utils
cpu: Intel(R) Core(TM) i7-4770HQ CPU @ 2.20GHz
BenchmarkCacheRegexp-8           3936446              2876 ns/op
BenchmarkCacheRegexp-8           4215283              3142 ns/op
BenchmarkCacheRegexp-8           3976014              3090 ns/op
BenchmarkCacheRegexp-8           3832299              3063 ns/op
BenchmarkCacheRegexp-8           3947518              3007 ns/op
BenchmarkCacheRegexp-8           3930441              3120 ns/op
BenchmarkCacheRegexp-8           3868693              3089 ns/op
BenchmarkCacheRegexp-8           3757822              3020 ns/op
BenchmarkCacheRegexp-8           3982918              2917 ns/op
BenchmarkCacheRegexp-8           3633404              3065 ns/op
BenchmarkNormalRegexp-8           648982             17346 ns/op
BenchmarkNormalRegexp-8           683965             17628 ns/op
BenchmarkNormalRegexp-8           684678             17825 ns/op
BenchmarkNormalRegexp-8           646743             17986 ns/op
BenchmarkNormalRegexp-8           651927             18032 ns/op
BenchmarkNormalRegexp-8           629230             17448 ns/op
BenchmarkNormalRegexp-8           557157             18398 ns/op
BenchmarkNormalRegexp-8           633378             18138 ns/op
BenchmarkNormalRegexp-8           669038             18004 ns/op
BenchmarkNormalRegexp-8           681964             17552 ns/op
PASS
ok      github.com/tekintian/go-str-utils       272.780s
```

更多正则使用请参考官方文档： https://pkg.go.dev/regexp

测试文档： https://pkg.go.dev/testing





