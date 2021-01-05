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
Benchmark/Day_04/Part_1-16          2326            453995 ns/op          145462 B/op       1095 allocs/op
Benchmark/Day_04/Part_2-16          2443            457525 ns/op          145463 B/op       1095 allocs/op
Benchmark/Day_05/Part_1-16          5296            222592 ns/op          177536 B/op         25 allocs/op
Benchmark/Day_05/Part_2-16           745           1637821 ns/op          900670 B/op        284 allocs/op
Benchmark/Day_06/Part_1-16           595           2027423 ns/op           59952 B/op        536 allocs/op
Benchmark/Day_06/Part_2-16          1018           1177671 ns/op            8166 B/op        111 allocs/op
Benchmark/Day_07/Part_1-16         27547             42569 ns/op            8144 B/op        140 allocs/op
Benchmark/Day_07/Part_2-16         27577             43574 ns/op            8736 B/op        151 allocs/op
Benchmark/Day_08/Part_1-16          2372            456757 ns/op          741626 B/op       3031 allocs/op
Benchmark/Day_08/Part_2-16          2271            501312 ns/op          741626 B/op       3031 allocs/op
Benchmark/Day_09/Part_1-16           354           3395335 ns/op         2199583 B/op      68548 allocs/op
Benchmark/Day_09/Part_2-16             3         375546758 ns/op        219338976 B/op   6854154 allocs/op
PASS
ok      github.com/armsnyder/aoc2018    17.500s
```
