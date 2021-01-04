package main

import (
	"bufio"
	"io"
	"sort"
)

var _ = declareDay(7, func(part2 bool, inputReader io.Reader) interface{} {
	if part2 {
		return day07Part2(inputReader)
	}
	return day07Part1(inputReader)
})

func day07Part1(inputReader io.Reader) interface{} {
	dag := day07ParseDAG(inputReader)
	queue := make(day07SortedBytes, 0)
	for _, leaf := range dag.leaves() {
		queue.add(leaf)
	}
	var result []byte
	for len(queue) > 0 {
		value := queue.popSmallest()
		result = append(result, value)
		newLeaves := dag.removeLeaf(value)
		for _, leaf := range newLeaves {
			queue.add(leaf)
		}
	}
	return string(result)
}

func day07Part2(inputReader io.Reader) interface{} {
	panic("no solution")
}

func day07ParseDAG(inputReader io.Reader) *day07DAG {
	scanner := bufio.NewScanner(inputReader)
	dag := &day07DAG{}

	for scanner.Scan() {
		if len(scanner.Bytes()) == 0 {
			continue
		}

		upstream := scanner.Bytes()[5]
		downstream := scanner.Bytes()[36]

		dag.addEdge(upstream, downstream)
	}

	return dag
}

type day07SortedBytes []byte

func (q *day07SortedBytes) add(value byte) bool {
	insertIndex := sort.Search(len(*q), func(i int) bool {
		return (*q)[i] <= value
	})

	if insertIndex == len(*q) {
		*q = append(*q, value)
		return true
	}

	if (*q)[insertIndex] == value {
		return false
	}

	*q = append(*q, 0)
	copy((*q)[insertIndex+1:], (*q)[insertIndex:])
	(*q)[insertIndex] = value
	return true
}

func (q *day07SortedBytes) remove(value byte) bool {
	insertIndex := sort.Search(len(*q), func(i int) bool {
		return (*q)[i] <= value
	})

	if insertIndex == len(*q) {
		return false
	}

	copy((*q)[insertIndex:], (*q)[insertIndex+1:])
	*q = (*q)[:len(*q)-1]
	return true
}

func (q *day07SortedBytes) popSmallest() byte {
	if len(*q) == 0 {
		return 0
	}

	value := (*q)[len(*q)-1]
	*q = (*q)[:len(*q)-1]
	return value
}

type day07DAG struct {
	upstreamEdges   map[byte]*day07SortedBytes
	downstreamEdges map[byte]*day07SortedBytes
}

func (d *day07DAG) addEdge(upstreamVertex, downstreamVertex byte) {
	if d.upstreamEdges == nil {
		d.upstreamEdges = make(map[byte]*day07SortedBytes)
	}

	if d.downstreamEdges == nil {
		d.downstreamEdges = make(map[byte]*day07SortedBytes)
	}

	if d.upstreamEdges[downstreamVertex] == nil {
		d.upstreamEdges[downstreamVertex] = &day07SortedBytes{}
	}

	if d.downstreamEdges[upstreamVertex] == nil {
		d.downstreamEdges[upstreamVertex] = &day07SortedBytes{}
	}

	d.downstreamEdges[upstreamVertex].add(downstreamVertex)
	d.upstreamEdges[downstreamVertex].add(upstreamVertex)
}

func (d *day07DAG) removeLeaf(vertex byte) (newLeaves []byte) {
	if downstreamVertices := d.downstreamEdges[vertex]; downstreamVertices != nil {
		for _, downstreamVertex := range *downstreamVertices {
			if upstreamVertices := d.upstreamEdges[downstreamVertex]; upstreamVertices != nil {
				if upstreamVertices.remove(vertex) && len(*upstreamVertices) == 0 {
					newLeaves = append(newLeaves, downstreamVertex)
				}
			}
		}
	}

	d.upstreamEdges[vertex] = nil
	d.downstreamEdges[vertex] = nil

	return newLeaves
}

func (d *day07DAG) leaves() (leaves []byte) {
	seen := make(map[byte]bool)

	for _, values := range d.upstreamEdges {
		if values != nil {
			for _, value := range *values {
				seen[value] = true
			}
		}
	}

	for k := range seen {
		if d.upstreamEdges[k] == nil {
			leaves = append(leaves, k)
		}
	}

	return leaves
}
