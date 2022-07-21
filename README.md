# Polyglot-Go Benchmarks

This repository contains a series of benchmarks for the Go implementation of Polyglot.

## Polyglot Benchmark

This benchmark is designed to be as close to a 1:1 comparison with protobuf and other serialization/deserialization frameworks as possible.
In order to facilitate this, all the generated code makes use of the same `benchmark.proto` file.

We've also kept the benchmarking code as similar as possible.

### Running the Benchmarks

Running the benchmarks is as simple as using the built-in `go test` command as follows:

```shell
$ go test ./... -bench=. -v
```

This will run the entire benchmark suite and for each serialization framework and return the average amount of time it takes to serialize or deserialize **1024** messages in a row.

This means that for the following example protobuf benchmark:

```shell
pkg: github.com/loopholelabs/polyglot-go-benchmarks/protobuf
BenchmarkEncode
BenchmarkEncode-10                  1198            905053 ns/op          376838 B/op      13312 allocs/op
BenchmarkDecode
BenchmarkDecode-10                   732           1630963 ns/op         1785871 B/op      44032 allocs/op
BenchmarkDecodeValidate
BenchmarkDecodeValidate-10           520           2318041 ns/op         2032424 B/op      54274 allocs/op
```

The `BenchmarkEncode` test took `905053 ns` for **1024 messages** - meaning it took `905053 ns / 1024 = 883.840820313 ns` per message. Similarly, it took
`1630963 ns / 1024 = 1592.73730469 ns` per message to decode in the `BenchmarkDecode` test.

Here's the same benchmarks but for Polyglot

```shell
pkg: github.com/loopholelabs/polyglot-go-benchmarks/polyglot
BenchmarkEncode
BenchmarkEncode-10                          3650            328514 ns/op               1 B/op          0 allocs/op
BenchmarkDecode
BenchmarkDecode-10                          2388            486361 ns/op          598215 B/op      17408 allocs/op
BenchmarkDecodeWithValidate
BenchmarkDecodeWithValidate-10              1008           1175060 ns/op          844389 B/op      27649 allocs/op
```

The `BenchmarkEncode` test took `328514 ns` for **1024 messages** - meaning it took `328514 ns / 1024 = 320.814453125 ns` per message. Similarly, it took
`486361 ns / 1024 = 474.961914063 ns` per message to decode in the `BenchmarkDecode` test.
