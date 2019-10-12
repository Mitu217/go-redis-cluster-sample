# go-redis-cluster-sample

## Test

```bash
# Connect 127.0.0.1:7000
redis-cli -c -h localhost -p 7000

localhost:7000> set hoge fuga
localhost:7000> get hoge
"fuga"

# Connect 127.0.0.1:7001
redis-cli -c -h localhost -p 7001

localhost:7001> get hoge
-> Redirected to slot [1525] located at 127.0.0.1:7000
"fuga"
```

## Bench

```bash
% make bench
go test -bench . --benchmem
goos: darwin
goarch: amd64
pkg: github.com/Mitu217/go-redis-cluster-sample
BenchmarkSerialGetSet/radix-8               1000           1693677 ns/op             298 B/op              8 allocs/op
BenchmarkSerialGetSet/go-redis-8            1000           1145892 ns/op             358 B/op             10 allocs/op
BenchmarkSerialGetSetLargeArgs/radix-8              1000           1807767 ns/op      13560 B/op          13 allocs/op
BenchmarkSerialGetSetLargeArgs/go-redis-8           1000           1410732 ns/op      13943 B/op          10 allocs/op
BenchmarkParallelGetSet/radix-8                    30000             56843 ns/op         77 B/op           4 allocs/op
BenchmarkParallelGetSet/goredis-8               --- FAIL: BenchmarkParallelGetSet/goredis-8
    bench_test.go:97: read tcp 127.0.0.1:53881->127.0.0.1:7000: read: connection reset by peer
    bench_test.go:97: read tcp 127.0.0.1:53883->127.0.0.1:7000: read: connection reset by peer
    bench_test.go:97: read tcp 127.0.0.1:53882->127.0.0.1:7000: read: connection reset by peer
    bench_test.go:97: read tcp 127.0.0.1:53886->127.0.0.1:7000: read: connection reset by peer
    bench_test.go:97: read tcp 127.0.0.1:53890->127.0.0.1:7000: read: connection reset by peer
    bench_test.go:97: read tcp 127.0.0.1:53891->127.0.0.1:7000: read: connection reset by peer
    bench_test.go:97: read tcp 127.0.0.1:53880->127.0.0.1:7000: read: connection reset by peer
--- FAIL: BenchmarkParallelGetSet
FAIL
exit status 1
FAIL    github.com/Mitu217/go-redis-cluster-sample      9.056s
make: *** [bench] Error 1
```
