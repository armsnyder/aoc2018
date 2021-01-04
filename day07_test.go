package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestDay07Part1(t *testing.T) {
	runDayTests(t, 7, []dayTest{
		{
			input: `
Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.
`,
			want: "CABDFE",
		},
	})
}

func TestDay07Part2(t *testing.T) {
	input := strings.NewReader(`
Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.
`)
	got := day07Part2(input, 2, 0)
	want := 15
	if got != want {
		t.Errorf("got %d; wanted %d", got, want)
	}
}

func Test_day07SortedBytes(t *testing.T) {
	tests := []struct {
		push    string
		wantPop string
	}{
		{
			push:    "ABC",
			wantPop: "ABC",
		},
		{
			push:    "CBA",
			wantPop: "ABC",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s->%s", tt.push, tt.wantPop), func(t *testing.T) {
			queue := make(day07SortedBytes, 0)
			for i := 0; i < len(tt.push); i++ {
				queue.add(tt.push[i])
			}
			var popped []byte
			for i := 0; i < len(tt.wantPop); i++ {
				popped = append(popped, queue.popSmallest())
			}
			got := string(popped)
			if tt.wantPop != got {
				t.Errorf("wanted %q; got %q", tt.wantPop, got)
			}
		})
	}
}

func Test_day07DAG(t *testing.T) {
	tests := []struct {
		edges         [][2]byte
		wantLeaves    []byte
		removeVertex  byte
		wantNewLeaves []byte
	}{
		{
			edges:         [][2]byte{{'A', 'B'}},
			wantLeaves:    []byte{'A'},
			removeVertex:  'A',
			wantNewLeaves: []byte{'B'},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			dag := &day07DAG{}
			for _, edge := range tt.edges {
				dag.addEdge(edge[0], edge[1])
			}
			gotLeaves := dag.leaves()
			if !reflect.DeepEqual(gotLeaves, tt.wantLeaves) {
				t.Errorf("wanted initial leaves: %value; got: %value", tt.wantLeaves, gotLeaves)
			}
			gotNewLeaves := dag.removeLeaf(tt.removeVertex)
			if !reflect.DeepEqual(gotNewLeaves, tt.wantNewLeaves) {
				t.Errorf("wanted new leaves: %value; got: %value", tt.wantNewLeaves, gotNewLeaves)
			}
		})
	}
}
