package main

import (
	"testing"
)

func TestDay09Part1(t *testing.T) {
	runDayTests(t, 9, []dayTest{
		{
			input: "9 players; last marble is worth 25 points",
			want:  32,
		},
		{
			input: "10 players; last marble is worth 1618 points",
			want:  8317,
		},
		{
			input: "13 players; last marble is worth 7999 points",
			want:  146373,
		},
		{
			input: "17 players; last marble is worth 1104 points",
			want:  2764,
		},
		{
			input: "21 players; last marble is worth 6111 points",
			want:  54718,
		},
		{
			input: "30 players; last marble is worth 5807 points",
			want:  37305,
		},
	})
}
