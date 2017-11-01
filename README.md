# seahash

[![Build Status](https://travis-ci.org/blainsmith/seahash.svg?branch=master)](https://travis-ci.org/blainsmith/seahash)
[![Coverage Status](https://coveralls.io/repos/github/blainsmith/seahash/badge.svg?branch=master)](https://coveralls.io/github/blainsmith/seahash?branch=master)
[![GoDoc](https://godoc.org/github.com/blainsmith/seahash?status.svg)](https://godoc.org/github.com/blainsmith/seahash)
![](https://img.shields.io/badge/license-MIT-blue.svg)

A Go port of the [SeaHash](https://ticki.github.io/blog/seahash-explained/) algorithm.

## Benchmarks

On Intel(R) Core(TM) i5-6500 CPU @ 3.20GHz, with go 1.9.1.

```
$ go test -bench . -benchmem
BenchmarkSum-4     	20000000	        63.5 ns/op
BenchmarkSum64-4   	30000000	        43.1 ns/op
```
