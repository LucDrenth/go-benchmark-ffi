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
BenchmarkGo-12                  559262828                1.943 ns/op
BenchmarkCGO-12                  40146758                29.46 ns/op
BenchmarkPurego-12                4470415                265.8 ns/op
BenchmarkPuregoLoad-12              52886                22840 ns/op
```
