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
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
Benchmark/Day_04/Part_1-16          2824            418264 ns/op          145455 B/op       1095 allocs/op
Benchmark/Day_04/Part_2-16          2877            424348 ns/op          145451 B/op       1095 allocs/op
Benchmark/Day_05/Part_1-16          5355            215655 ns/op          258944 B/op         34 allocs/op
Benchmark/Day_05/Part_2-16           754           1628480 ns/op          982025 B/op        293 allocs/op
Benchmark/Day_06/Part_1-16           619           2044924 ns/op           59966 B/op        536 allocs/op
Benchmark/Day_06/Part_2-16          1054           1153188 ns/op            8162 B/op        111 allocs/op
Benchmark/Day_07/Part_1-16         28878             40929 ns/op            7751 B/op        140 allocs/op
Benchmark/Day_07/Part_2-16         29137             42224 ns/op            8342 B/op        151 allocs/op
Benchmark/Day_08/Part_1-16          2511            452615 ns/op          763426 B/op       3039 allocs/op
Benchmark/Day_08/Part_2-16          2401            494968 ns/op          763426 B/op       3039 allocs/op
Benchmark/Day_09/Part_1-16           504           2366983 ns/op         1649696 B/op      68547 allocs/op
Benchmark/Day_09/Part_2-16             3         377833151 ns/op        164504248 B/op   6854153 allocs/op
Benchmark/Day_10/Part_1-16          9464            113338 ns/op           72416 B/op        632 allocs/op
Benchmark/Day_10/Part_2-16         10000            109330 ns/op           70888 B/op        620 allocs/op
Benchmark/Day_11/Part_1-16          2342            504162 ns/op             544 B/op          4 allocs/op
Benchmark/Day_11/Part_2-16            51          24198209 ns/op             544 B/op          5 allocs/op
Benchmark/Day_12/Part_1-16         15339             76401 ns/op           10665 B/op        178 allocs/op
Benchmark/Day_12/Part_2-16          2535            483413 ns/op           41778 B/op        714 allocs/op
```
