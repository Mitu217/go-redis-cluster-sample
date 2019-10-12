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
BenchmarkSerialGetSet/radix-8               1000           1639788 ns/op             227 B/op          8 allocs/op
BenchmarkSerialGetSet/go-redis-8            2000           1048931 ns/op             284 B/op          9 allocs/op
BenchmarkSerialGetSetLargeArgs/radix-8              1000           1799455 ns/op           13087 B/op         12 allocs/op
BenchmarkSerialGetSetLargeArgs/go-redis-8           1000           1338310 ns/op           13852 B/op          9 allocs/op
BenchmarkParallelGetSet/radix-8                    20000             56267 ns/op              72 B/op          4 allocs/op
BenchmarkParallelGetSet/goredis-8                  10000            192767 ns/op             330 B/op          9 allocs/op
PASS
ok      github.com/Mitu217/go-redis-cluster-sample      12.324s
```
