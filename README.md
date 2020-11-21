# Go benchmarktests

This is just a playground repo to become more familiar with Golang benchmarks.

Examples:
```bash
$ go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/c0rby/go-benchmarktests
BenchmarkFibIterative-4         29249113            35.4 ns/op
BenchmarkFibRecusrive-4              136       7747940 ns/op
BenchmarkFibMemoization-4         157252          6715 ns/op
PASS
ok      github.com/c0rby/go-benchmarktests  4.152s
``` 

With memory benchmarks:
```bash
goos: linux
goarch: amd64
pkg: github.com/c0rby/go-benchmarktests
BenchmarkFibIterative-4         34703954            30.5 ns/op         0 B/op          0 allocs/op
BenchmarkFibRecusrive-4              160       7348830 ns/op           0 B/op          0 allocs/op
BenchmarkFibMemoization-4         190088          5878 ns/op        2428 B/op          8 allocs/op
PASS
ok      github.com/c0rby/go-benchmarktests  5.179s
```
