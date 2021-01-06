package main

import (
	"testing"
)

func TestDay11Part1(t *testing.T) {
	runDayTests(t, 11, []dayTest{
		{
			input: "18",
			want:  "33,45",
		},
		{
			input: "42",
			want:  "21,61",
		},
	})
}

func TestDay11Part2(t *testing.T) {
	runDayTests(t, 11, []dayTest{
		{
			part2: true,
			input: "18",
			want:  "90,269,16",
		},
		{
			part2: true,
			input: "42",
			want:  "232,251,12",
		},
	})
}
