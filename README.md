# seahash

[![Build Status](https://travis-ci.org/blainsmith/seahash.svg?branch=master)](https://travis-ci.org/blainsmith/seahash)
[![GoDoc](https://godoc.org/github.com/blainsmith/seahash?status.svg)](https://godoc.org/github.com/blainsmith/seahash)
![](https://img.shields.io/badge/license-MIT-blue.svg)

A Go port of the [SeaHash](https://ticki.github.io/blog/seahash-explained/) algorithm.

## Benchmarks

Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz running Linux 5.3.0

```
BenchmarkSeahash64-8           	55457641	        20.8 ns/op	       0 B/op	       0 allocs/op

BenchmarkFnvHash32-8           	40879219	        30.2 ns/op	       4 B/op	       1 allocs/op
BenchmarkFnvHash64-8           	34940972	        30.6 ns/op	       8 B/op	       1 allocs/op

BenchmarkFarmHashHash32-8      	61766998	        19.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkFarmHashHash64-8      	100000000	        11.6 ns/op	       0 B/op	       0 allocs/op

BenchmarkHuichenMurmur-8       	81137569	        13.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkReuseeMurmur-8        	21982393	        52.7 ns/op	       4 B/op	       1 allocs/op
BenchmarkZhangMurmur-8         	91060291	        14.1 ns/op	       0 B/op	       0 allocs/op

BenchmarkDgryskiSpooky32-8     	44230748	        27.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkDgryskiSpooky64-8     	48947028	        23.9 ns/op	       0 B/op	       0 allocs/op

BenchmarkDgryskiStatdx64-8     	100000000	        11.5 ns/op	       0 B/op	       0 allocs/op

BenchmarkHashlandSpooky32-8   	40680852	        29.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashlandSpooky64-8   	43915500	        28.0 ns/op	       0 B/op	       0 allocs/op

BenchmarkCreachadairCity32-8   	39462910	        28.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkCreachadairCity64-8   	80871393	        15.4 ns/op	       0 B/op	       0 allocs/op
```