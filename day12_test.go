package main

import (
	"testing"
)

func TestDay12Part1(t *testing.T) {
	runDayTests(t, 12, []dayTest{
		{
			input: `initial state: #..#.#..##......###...###

...## => #
..#.. => #
.#... => #
.#.#. => #
.#.## => #
.##.. => #
.#### => #
#.#.# => #
#.### => #
##.#. => #
##.## => #
###.. => #
###.# => #
####. => #
`,
			want: 325,
		},
		{
			input: `initial state: #

.#... => #
..#.. => #
.##.. => #
`,
			want: 10,
		},
	})
}

func TestDay12Part2(t *testing.T) {
	runDayTests(t, 12, []dayTest{
		{
			part2: true,
			input: `initial state: #

.#... => #
..#.. => #
.##.. => #
`,
			want: 25000000000,
		},
	})
}
