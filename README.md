# Teapot

short and stout golang logger

## Installation

```
$ go get github.com/trevatk/teapot
```

## Usage

```go
package main

func main() {
    log := New()
	log.Debug("hello world")
}
```

## Benchmark 

```
goos: linux
goarch: amd64
pkg: github.com/trevatk/teapot
cpu: Intel(R) Core(TM) i5-10300H CPU @ 2.50GHz
BenchmarkLogger_InfoHotPath-8   	 3322754	       358.0 ns/op	      48 B/op	       1 allocs/op
```