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
Benchmark/Day_04/Part_1-16          2746            462724 ns/op          145463 B/op       1095 allocs/op
Benchmark/Day_04/Part_2-16          2648            435993 ns/op          145460 B/op       1095 allocs/op
Benchmark/Day_05/Part_1-16          5556            210476 ns/op          177536 B/op         25 allocs/op
Benchmark/Day_05/Part_2-16           692           1549092 ns/op          900047 B/op        284 allocs/op
Benchmark/Day_06/Part_1-16           596           2038004 ns/op           59987 B/op        536 allocs/op
Benchmark/Day_06/Part_2-16           954           1154022 ns/op            8164 B/op        111 allocs/op
Benchmark/Day_07/Part_1-16         28242             42553 ns/op            8146 B/op        141 allocs/op
Benchmark/Day_07/Part_2-16         28450             41781 ns/op            8737 B/op        151 allocs/op
Benchmark/Day_08/Part_1-16          2320            452959 ns/op          741626 B/op       3031 allocs/op
Benchmark/Day_08/Part_2-16          2467            492304 ns/op          741626 B/op       3031 allocs/op
Benchmark/Day_09/Part_1-16           356           3250510 ns/op         2199582 B/op      68548 allocs/op
Benchmark/Day_09/Part_2-16             3         361848361 ns/op        219338976 B/op   6854154 allocs/op
Benchmark/Day_10/Part_1-16          9565            111899 ns/op           72416 B/op        632 allocs/op
Benchmark/Day_10/Part_2-16          9871            111663 ns/op           70888 B/op        620 allocs/op
PASS
ok      github.com/armsnyder/aoc2018    20.539s
```
