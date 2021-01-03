package main

import (
	"testing"
)

func TestDay05Part1(t *testing.T) {
	runDayTests(t, 5, []dayTest{
		{
			input: "aA",
			want:  0,
		},
		{
			input: "abBA",
			want:  0,
		},
		{
			input: "abAB",
			want:  4,
		},
		{
			input: "aabAAB",
			want:  6,
		},
		{
			input: "dabAcCaCBAcCcaDA",
			want:  10,
		},
	})
}

func TestDay05Part2(t *testing.T) {
	runDayTests(t, 5, []dayTest{
		{
			part2: true,
			input: "dabAcCaCBAcCcaDA",
			want:  4,
		},
	})
}
