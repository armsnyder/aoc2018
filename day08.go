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
	rawInput, _ := ioutil.ReadAll(inputReader)
	inputFields := strings.Fields(string(rawInput))

	var stack day08Stack
	var state day08State

	metadataTotal := 0

	for _, field := range inputFields {
		inputValue, _ := strconv.Atoi(field)

		switch state {
		case day08ChildHeader:
			stack.push(day08Node{children: inputValue})
			state = day08MetadataHeader

		case day08MetadataHeader:
			node := stack.pop()
			node.metadata = inputValue
			if node.children == 0 {
				state = day08Metadata
			} else {
				node.children--
				state = day08ChildHeader
			}
			stack.push(node)

		case day08Metadata:
			metadataTotal += inputValue
			node := stack.pop()
			node.metadata--
			if node.metadata > 0 {
				stack.push(node)
			} else if stack.peek().children > 0 {
				node = stack.pop()
				node.children--
				stack.push(node)
				state = day08ChildHeader
			}
		}
	}

	return metadataTotal
}

func day08Part2(inputReader io.Reader) interface{} {
	panic("no solution")
}

type day08State int

const (
	day08ChildHeader day08State = iota
	day08MetadataHeader
	day08Metadata
)

type day08Node struct {
	children int
	metadata int
}

type day08Stack []day08Node

func (s *day08Stack) push(node day08Node) {
	*s = append(*s, node)
}

func (s *day08Stack) pop() day08Node {
	if len(*s) == 0 {
		return day08Node{}
	}
	node := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return node
}

func (s *day08Stack) peek() day08Node {
	if len(*s) == 0 {
		return day08Node{}
	}
	return (*s)[len(*s)-1]
}
