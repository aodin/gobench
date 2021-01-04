# gobench
Benchmarking Go

```sh
go test -bench=BenchmarkRandStringBytesRmndr -benchmem
goos: darwin
goarch: arm64
pkg: github.com/aodin/gobench
BenchmarkRandStringBytesRmndr-8   	 5611254	       178.7 ns/op	      32 B/op	       2 allocs/op
PASS
ok  	github.com/aodin/gobench	1.609s
```
