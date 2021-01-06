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
Benchmark/Day_04/Part_1-16          2649            435321 ns/op          145466 B/op       1095 allocs/op
Benchmark/Day_04/Part_2-16          2464            433433 ns/op          145460 B/op       1095 allocs/op
Benchmark/Day_05/Part_1-16          5434            208460 ns/op          177536 B/op         25 allocs/op
Benchmark/Day_05/Part_2-16           708           1640200 ns/op          901055 B/op        284 allocs/op
Benchmark/Day_06/Part_1-16           645           1934420 ns/op           59969 B/op        536 allocs/op
Benchmark/Day_06/Part_2-16          1136           1145500 ns/op            8165 B/op        111 allocs/op
Benchmark/Day_07/Part_1-16         28768             41358 ns/op            8145 B/op        140 allocs/op
Benchmark/Day_07/Part_2-16         28618             41857 ns/op            8737 B/op        151 allocs/op
Benchmark/Day_08/Part_1-16          2617            453896 ns/op          741626 B/op       3031 allocs/op
Benchmark/Day_08/Part_2-16          2478            496348 ns/op          741626 B/op       3031 allocs/op
Benchmark/Day_09/Part_1-16           352           3276305 ns/op         2199583 B/op      68548 allocs/op
Benchmark/Day_09/Part_2-16             3         347363015 ns/op        219338978 B/op   6854154 allocs/op
Benchmark/Day_10/Part_1-16          9976            111311 ns/op           72416 B/op        632 allocs/op
Benchmark/Day_10/Part_2-16         10000            110036 ns/op           70888 B/op        620 allocs/op
Benchmark/Day_11/Part_1-16          2379            520234 ns/op            2080 B/op          5 allocs/op
Benchmark/Day_11/Part_2-16            48          23883337 ns/op            2080 B/op          6 allocs/op
PASS
ok      github.com/armsnyder/aoc2018    22.368s
```
