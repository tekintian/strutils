package strutils

import (
	"log"
	"regexp"
	"testing"
)

func TestRegex(t *testing.T) {
	l, cnt := concurentDemo()
	t.Logf("并发数:%d, 缓存命中数:%d ", l, cnt)
}

// 并发测试
func concurentDemo() (int, int) {
	pattern := `(\d{2,})`
	b := []byte("123abc 3456654 , asfasdfd 9305345034.999")
	var l int = 1000000
	ch := make(chan int, l)
	for i := 0; i < l; i++ {
		// 开启协程执行
		go func() {
			reg, err := GetRegexp(pattern)
			if err != nil {
				log.Fatal(err)
				return
			}
			rb := reg.FindAll(b, 1)
			if len(rb[0]) > 0 {
				ch <- 1
			} else {
				ch <- 0
			}
		}()
	}

	var cnt int = 0
	for i := 0; i < l; i++ {
		if a := <-ch; a == 1 {
			cnt++
		}
	}
	return l, cnt
}
func BenchmarkRegexp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concurentDemo()
		//b.Logf("并发数:%d, 缓存命中数:%d ",l,cnt)
	}
}

var regStr string = `<(script|style)>([\S\s]+?)<(/|/\s+)(script|style)>`
var txt string = `
<title>go生成指定范围的随机数 - Search</title><meta content="text/html; charset=utf-8" http-equiv="content-type" /><meta name="referrer" content="origin-when-cross-origin" /><meta property="og:description" content="Intelligent search from Bing makes it easier to quickly find what you’re looking for and rewards you." /><meta property="og:site_name" content="Bing" /><meta property="og:title" content="go生成指定范围的随机数 - Bing" /><meta property="og:url" content="https://cn.bing.com:9943/search?q=go生成指定范围的随机数" /><meta property="fb:app_id" content="3732605936979161" /><meta property="og:image" content="http://www.bing.com/sa/simg/facebook_sharing_5.png" /><meta property="og:type" content="website" /><meta property="og:image:width" content="600" /><meta property="og:image:height" content="315" /><link href="/search?format=rss&amp;q=go%E7%94%9F%E6%88%90%E6%8C%87%E5%AE%9A%E8%8C%83%E5%9B%B4%E7%9A%84%E9%9A%8F%E6%9C%BA%E6%95%B0&amp;qs=UT&amp;pq=go%E7%94%9F%E6%88%90%E6%8C%87%E5%AE%9A&amp;sc=8-6&amp;cvid=FBE1BFE52F564F30A659263053049CB8&amp;FORM=QBRE&amp;sp=1&amp;lq=0" data-orighref="" rel="alternate" title="XML" type="text/xml" /><link href="/search?format=rss&amp;q=go%E7%94%9F%E6%88%90%E6%8C%87%E5%AE%9A%E8%8C%83%E5%9B%B4%E7%9A%84%E9%9A%8F%E6%9C%BA%E6%95%B0&amp;qs=UT&amp;pq=go%E7%94%9F%E6%88%90%E6%8C%87%E5%AE%9A&amp;sc=8-6&amp;cvid=FBE1BFE52F564F30A659263053049CB8&amp;FORM=QBRE&amp;sp=1&amp;lq=0" data-orighref="" rel="alternate" title="RSS" type="application/rss+xml" /><link href="/sa/simg/favicon-trans-bg-blue-mg.ico" data-orighref="" rel="icon" /><link rel="preconnect" href="https://r.bing.com" /><link rel="preconnect" href="https://r.bing.com" crossorigin/><link rel="dns-prefetch" href="https://r.bing.com" /><link rel="dns-prefetch" href="https://r.bing.com" crossorigin/><link rel="preconnect" href="https://th.bing.com" /><link rel="preconnect" href="https://th.bing.com" crossorigin/><link rel="dns-prefetch" href="https://th.bing.com" /><link rel="dns-prefetch" href="https://th.bing.com" crossorigin/><script type="text/javascript" nonce="MugYWcN70DTy0G6Qqqu0oT2jEHFvdkh8KmMhLt3NjhA=">//<![CDATA[
_G={Region:"CN",Lang:"en-US",ST:(typeof si_ST!=='undefined'?si_ST:new Date),Mkt:"en-US",RevIpCC:"cn",RTL:false,Ver:"22",IG:"CAA84B55CCB44D61A78ABA8FB65BAFAC",EventID:"665bfe5d4cf44a38a79a2a9913e7eb67",V:"web",P:"SERP",DA:"HKGE01",SUIH:"NPVC-F9thG3JmeMFxUAocw",adc:"b_ad",EF:{cookss:1,bmcov:1,crossdomainfix:1,bmasynctrigger:1,bmasynctrigger3:1,getslctspt:1,newtabsloppyclick:1,chevroncheckmousemove:1,sharepreview:1,shareoutimage:1,sharefixreadnum:1,sharepreviewthumbnailid:1,chatskip2content:1},gpUrl:"\/fd\/ls\/GLinkPing.aspx?" }; _G.lsUrl="/fd/ls/l?IG="+_G.IG ;curUrl="https:\/\/cn.bing.com\/search";function si_T(a){ if(document.images){_G.GPImg=new Image;_G.GPImg.src=_G.gpUrl+'IG='+_G.IG+'&'+a;}return true;}_G.NTT="600000";_G.CTT="3000";_G.BNFN="Default";_G.LG="160";_G.FilterFlareInterval=5;;
//]]></script>
`

// 采用自定义缓存调用regexp
func BenchmarkCacheRegexp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reg, _ := GetRegexp(regStr)
		reg.ReplaceAllString(txt, "")
	}
}

// 普通模式调用regexp
func BenchmarkNormalRegexp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reg, _ := regexp.Compile(regStr)
		reg.ReplaceAllString(txt, "")
	}
}

/*
// 使用缓存和不使用缓存的基准测试结果对比, 使用缓存可提高6倍的正则执行效率
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

*/
