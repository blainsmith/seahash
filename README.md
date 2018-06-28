# seahash

[![Build Status](https://travis-ci.org/blainsmith/seahash.svg?branch=master)](https://travis-ci.org/blainsmith/seahash)
[![GoDoc](https://godoc.org/github.com/blainsmith/seahash?status.svg)](https://godoc.org/github.com/blainsmith/seahash)
![](https://img.shields.io/badge/license-MIT-blue.svg)

A Go port of the [SeaHash](https://ticki.github.io/blog/seahash-explained/) algorithm.

## Benchmarks

Intel(R) Xeon(R) CPU E3-1505M v6 @ 3.00GHz

```
$ go test -bench .
BenchmarkSum-8     	30000000	        47.2 ns/op
BenchmarkSum64-8   	50000000	        33.0 ns/op
```
