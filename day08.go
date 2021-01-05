package main

import (
	"bufio"
	"bytes"
	"io"
	"strconv"
)

var _ = declareDay(8, func(part2 bool, inputReader io.Reader) interface{} {
	if part2 {
		return day08Part2(inputReader)
	}
	return day08Part1(inputReader)
})

func day08Part1(inputReader io.Reader) interface{} {
	var stack day08Stack
	var state day08State
	total := 0

	day08ScanNumbers(inputReader, func(v int) {
		switch state {
		case day08ChildHeader:
			stack.push(day08Node{children: v})
			state = day08MetadataHeader

		case day08MetadataHeader:
			node := stack.pop()
			node.metadata = v
			if node.children == 0 {
				state = day08Metadata
			} else {
				node.children--
				state = day08ChildHeader
			}
			stack.push(node)

		case day08Metadata:
			total += v
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
	})

	return total
}

func day08Part2(inputReader io.Reader) interface{} {
	panic("no solution")
}

func day08ScanNumbers(inputReader io.Reader, fn func(v int)) {
	scanner := bufio.NewScanner(inputReader)

	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}

		if i := bytes.IndexAny(data, " \n"); i >= 0 {
			return i + 1, data[0:i], nil
		}

		if atEOF {
			return len(data), data, nil
		}

		return 0, nil, nil
	})

	for scanner.Scan() {
		v, _ := strconv.Atoi(scanner.Text())
		fn(v)
	}
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
