# About
Compare Go FFI overhead

# Running the benchmarks yourself
This repo contains a `dylib` that is compiled from the arm64 architecture using `make compile-c`. If you want to run this on a different architecture, you'll need to compile [add.c](./c/add.c) yourself.

To run the benchmarks, run `go test ./... -bench=.`.

# Test results
```
goos: darwin
goarch: arm64
pkg: github.com/LucDrenth/go-benchmark-ffi
cpu: Apple M2 Pro
BenchmarkGo-12                                  586882443               1.931 ns/op
BenchmarkCGO-12                                  41127932               29.52 ns/op
BenchmarkPurego/registered_func-12                4373442               261.9 ns/op
BenchmarkPurego/sys_call-12                      13341868               89.52 ns/op
BenchmarkPuregoLoad-12                              56335               20762 ns/op
```
