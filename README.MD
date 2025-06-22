# About
The benchmarks compare the performance of native Go, cgo, and [purego](https://github.com/ebitengine/purego) for two different scenarios: a simple function call and a real-world call to an external API (Apple's Metal framework). This is done with Go 1.24.

## Simple addition
In [/simple_addition](./simple_addition/), we compare functions that add two numbers and returns the result.

| Method	          | Time (ns/op) | Overhead vs. Go | Overhead vs. cgo |
| ------------------- | ------------ | --------------- | ---------------- |
| Go (native)	      | 1.997	     | -	           | -                |
| cgo	              | 29.69	     | +27.7 ns	       | -                |
| purego (syscall)	  | 89.82	     | +87.8 ns	       | +60.1 ns         |
| purego (registered) |	259.5	     | +257.5 ns	   | +229.8 ns        |

**Key takeaways:**
- A cgo call has a fixed overhead of approximately 28 ns over a native Go call.
- purego's syscall mechanism adds another 60 ns of overhead on top of cgo.
- purego's registered function is significantly slower, likely due to its use of Go's reflection package.

_Note: The one-time cost of loading the dynamic library with purego (BenchmarkAdditionPuregoLoad-8) was measured at ~20 µs (microseconds), which is insignificant for most applications._

## Calling the Metal API
In [/metal_device](./metal_device/), we compare functions that call the [Metal](https://developer.apple.com/documentation/metal?language=objc) API to get a reference to a MTLDevice. This is a real-world scenario, as getting an MTLDevice is the first step to doing any GPU related work on an Apple device. The times here are much higher, but we are interested in the _difference_ in overhead, not the times of the calls itself.

| Method 	          | Time (ns/op) | Overhead vs. cgo |
| ------------------- | ------------ | ---------------- |
| cgo                 | 30001        | -                |
| purego (syscall)    |	30215	     | +214 ns          |
| purego (registered) | 30466	     | +465 ns          |

**Key takeaways:**
- The overhead of purego's syscall compared to cgo is 214 ns, which is substantially higher than the 60 ns overhead seen in the simple addition test. <br />
The reason for this increased overhead is not immediately clear but could be related to factors like data marshalling complexity or the size of the framework being called. Further benchmarks for performance-critical Metal calls, like submitting commands to a command queue, would be valuable to see if this overhead trend continues.
- The additional overhead of using a registered function versus a direct syscall in purego is +251 ns. This is a higher but comparable overhead to the +170 ns seen in the simple addition test, suggesting the reflection-based cost is relatively consistent.

_Note: The one-time cost of loading the Metal framework with purego was measured at ~1.3 µs (microseconds), which is also negligible._

# Running the benchmarks yourself
The benchmarks in this repository are tailored for MacOS with ARM64 architecture:
- [/simple_addition](./simple_addition/) contains `dylib` that is compiled from the arm64 architecture on MacOS, using `make compile-c`. If you want to run this on a different machine, you'll need to compile [add.c](./simple_addition/c/add.c) yourself.
- [/metal_device](./metal_device/) makes calls to Metal, Apple's graphics API. This is only available on Mac devices.

To run the benchmarks yourself, run `go test -bench=. ./...`.

# Test results
```
goos: darwin
goarch: arm64
cpu: Apple M2

pkg: github.com/LucDrenth/go-benchmark-ffi/metal_device
BenchmarkGetMTLDeviceCgo-8                          30001 ns/op
BenchmarkGetMTLDevicePurego/registered_func-8       30466 ns/op
BenchmarkGetMTLDevicePurego/sys_call-8              30215 ns/op
BenchmarkGetMTLDevicePuregoLoad-8                    1269 ns/op

pkg: github.com/LucDrenth/go-benchmark-ffi/simple_addition
BenchmarkAdditionGo-8                               1.997 ns/op
BenchmarkAdditionCgo-8                              29.69 ns/op
BenchmarkAdditionPurego/registered_func-8           259.5 ns/op
BenchmarkAdditionPurego/sys_call-8                  89.82 ns/op
BenchmarkAdditionPuregoLoad-8                       19343 ns/op
```
