# seahash

[![Build Status](https://travis-ci.org/blainsmith/seahash.svg?branch=master)](https://travis-ci.org/blainsmith/seahash)
[![GoDoc](https://godoc.org/github.com/blainsmith/seahash?status.svg)](https://godoc.org/github.com/blainsmith/seahash)
![](https://img.shields.io/badge/license-MIT-blue.svg)

A Go port of the [SeaHash](https://ticki.github.io/blog/seahash-explained/) algorithm.

## Benchmarks
```
$ go test -bench . -benchmem
BenchmarkHash-4      	30000000	        41.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkDiffuse-4   	2000000000	         0.29 ns/op	       0 B/op	       0 allocs/op
```
