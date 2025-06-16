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
BenchmarkCGO/CGO_addition-12            39746358                29.96 ns/op
BenchmarkCGO/Go_addition-12             630705363               1.901 ns/op
BenchmarkCGO/Purego_addition-12         4639356                 269.6 ns/op
BenchmarkCGO/Purego_load-12             52376                   23067 ns/op
```
