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
Benchmark/Day_04/Part_1-16          2384            441776 ns/op          145463 B/op       1095 allocs/op
Benchmark/Day_04/Part_2-16          2548            456783 ns/op          145463 B/op       1095 allocs/op
Benchmark/Day_05/Part_1-16          5458            205392 ns/op          177536 B/op         25 allocs/op
Benchmark/Day_05/Part_2-16           723           1651305 ns/op          901041 B/op        283 allocs/op
Benchmark/Day_06/Part_1-16           608           1998166 ns/op           59945 B/op        536 allocs/op
Benchmark/Day_06/Part_2-16           996           1143428 ns/op            8164 B/op        111 allocs/op
Benchmark/Day_07/Part_1-16         29428             41023 ns/op            8145 B/op        140 allocs/op
Benchmark/Day_07/Part_2-16         28593             42775 ns/op            8736 B/op        151 allocs/op
Benchmark/Day_08/Part_1-16          2884            412504 ns/op          741627 B/op       3031 allocs/op
Benchmark/Day_08/Part_2-16          2401            453654 ns/op          741627 B/op       3031 allocs/op
PASS
ok      github.com/armsnyder/aoc2018    14.384s
```
