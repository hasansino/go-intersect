[![Build Status](https://travis-ci.org/hasansino/go-intersect.svg?branch=master)](https://travis-ci.org/hasansino/go-intersect)

# go-intersect
Simple array intersection using hash maps.

# Installation

```
go get -u github.com/hasansino/go-intersect
```

# Usage

```
var (
    a = []int64{1,2,3}
    b = []int64{3,5,6}
)

c := intersect.Int64(a, b)
```

Two methods are available: `String()` and `Int64()`

# Performance

```
goos: linux
goarch: amd64
pkg: github.com/hasansino/go-intersect
BenchmarkInt64-8    	 3070972	       370 ns/op	      64 B/op	       1 allocs/op
BenchmarkString-8   	  829771	      1382 ns/op	     451 B/op	       2 allocs/op
```