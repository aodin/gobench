# gobench
Benchmarking Go

    go test -bench=BenchmarkRandStringBytesRmndr -benchmem


#### 2020 MacBook Pro M1

```sh
goos: darwin
goarch: arm64
pkg: github.com/aodin/gobench
BenchmarkRandStringBytesRmndr-8   	 5611254	       178.7 ns/op	      32 B/op	       2 allocs/op
PASS
ok  	github.com/aodin/gobench	1.609s
```

#### Mid 2014 MacBook Pro 2.6 GHz Dual-Core Intel Core i5

```sh
goos: darwin
goarch: amd64
pkg: github.com/aodin/gobench
BenchmarkRandStringBytesRmndr-4   	 4203007	       276 ns/op	      32 B/op	       2 allocs/op
PASS
ok  	github.com/aodin/gobench	1.658s
```
