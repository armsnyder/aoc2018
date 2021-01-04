# aoc2018

My [Advent of Code 2018](https://adventofcode.com/2018) solutions. Days 1-3 are missing because I
abandoned this calendar early when I was doing it originally.

Usage:

```
$ go run . [day] [-2]
```

Requires a `session.txt` file containing a session token, for pulling puzzle input and submitting answers.
(Inputs and answers are cached.)

## Benchmarks

I thought it would be fun to share performance [benchmarks](https://golang.org/pkg/testing/#hdr-Benchmarks)
for each of my puzzle solutions, since I write benchmarks anyway to help guide my optimizations.
I don't always optimize for the best possible time if I think it impacts code readability.
Benchmarks use the real puzzle input, which is preloaded in memory.

```
$ go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/armsnyder/aoc2018
Benchmark/Day_04/Part_1-16          2385            425543 ns/op          145461 B/op       1095 allocs/op
Benchmark/Day_04/Part_2-16          2698            426258 ns/op          145462 B/op       1095 allocs/op
Benchmark/Day_05/Part_1-16          5826            203503 ns/op          177536 B/op         25 allocs/op
Benchmark/Day_05/Part_2-16           720           1570948 ns/op          900835 B/op        284 allocs/op
Benchmark/Day_06/Part_1-16           649           1973529 ns/op           59961 B/op        536 allocs/op
Benchmark/Day_06/Part_2-16           994           1143134 ns/op            8160 B/op        111 allocs/op
Benchmark/Day_07/Part_1-16         29248             39906 ns/op            8145 B/op        140 allocs/op
Benchmark/Day_07/Part_2-16         29246             40674 ns/op            8736 B/op        151 allocs/op
PASS
ok      github.com/armsnyder/aoc2018    11.888s
```
