package main

import (
	"testing"
)

func TestDay08Part1(t *testing.T) {
	runDayTests(t, 8, []dayTest{
		{
			input: "2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2",
			want:  138,
		},
	})
}

func TestDay08Part2(t *testing.T) {
	runDayTests(t, 8, []dayTest{
		{
			part2: true,
			input: "2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2",
			want:  66,
		},
	})
}
