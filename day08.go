package main

import (
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

var _ = declareDay(8, func(part2 bool, inputReader io.Reader) interface{} {
	if part2 {
		return day08Part2(inputReader)
	}
	return day08Part1(inputReader)
})

func day08Part1(inputReader io.Reader) interface{} {
	return day08ParseTree(inputReader).
		postOrderReduce(func(node day08Node) (sum int) {
			for _, metadata := range node.metadata {
				sum += metadata
			}
			for _, child := range node.children {
				sum += child
			}
			return sum
		})
}

func day08Part2(inputReader io.Reader) interface{} {
	return day08ParseTree(inputReader).
		postOrderReduce(func(node day08Node) (sum int) {
			for _, metadata := range node.metadata {
				if len(node.children) == 0 {
					sum += metadata
				} else if metadata-1 < len(node.children) {
					sum += node.children[metadata-1]
				}
			}
			return sum
		})
}

func day08ParseTree(inputReader io.Reader) (tree day08Tree) {
	rawInput, _ := ioutil.ReadAll(inputReader)
	inputFields := strings.Fields(string(rawInput))

	tree = make(day08Tree, len(inputFields))

	for i, field := range inputFields {
		tree[i], _ = strconv.Atoi(field)
	}

	return tree
}

type day08Tree []int

func (t day08Tree) postOrderReduce(fn func(node day08Node) int) int {
	var stack []day08Node
	var i int

	for {
		node := day08Node{
			children: make([]int, 0, t[i]),
			metadata: make([]int, 0, t[i+1]),
		}

		i += 2

		for len(node.children) == cap(node.children) {
			metadataEnd := i + cap(node.metadata)

			for ; i < metadataEnd; i++ {
				node.metadata = append(node.metadata, t[i])
			}

			result := fn(node)

			if len(stack) == 0 {
				return result
			}

			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			node.children = append(node.children, result)
		}

		stack = append(stack, node)
	}
}

type day08Node struct {
	children []int
	metadata []int
}
