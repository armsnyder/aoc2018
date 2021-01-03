# aoc2018

My [Advent of Code 2018](https://adventofcode.com/2018) solutions. Days 1-3 are missing because I
abandoned this calendar early when I was doing it originally.

Usage:

```
$ go run . [-d day] [-2]
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
Benchmark/Day_04/Part_1-16          2361            457102 ns/op          145465 B/op       1095 allocs/op
Benchmark/Day_04/Part_2-16          2539            458153 ns/op          145462 B/op       1095 allocs/op
Benchmark/Day_05/Part_1-16          5559            215441 ns/op          177536 B/op         25 allocs/op
Benchmark/Day_05/Part_2-16           223           5387807 ns/op          177536 B/op         25 allocs/op
PASS
ok      github.com/armsnyder/aoc2018    5.434s
```
