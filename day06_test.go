package main

import (
	"strings"
	"testing"
)

func TestDay06Part1(t *testing.T) {
	runDayTests(t, 6, []dayTest{
		{
			input: `
1, 1
1, 6
8, 3
3, 4
5, 5
8, 9
`,
			want: 17,
		},
	})
}

func TestDay06Part2(t *testing.T) {
	input := strings.NewReader(`
1, 1
1, 6
8, 3
3, 4
5, 5
8, 9
`)
	want := 16
	got := day06Part2(input, 32)
	if want != got {
		t.Errorf("got %d; wanted %d", got, want)
	}
}
