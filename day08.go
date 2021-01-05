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
	var stack []day08Node
	var state day08State
	total := 0

	day08ScanNumbers(inputReader, func(v int) {
		switch state {
		case day08ChildHeader:
			stack = append(stack, day08Node{children: v})
			state = day08MetadataHeader
		case day08MetadataHeader:
			stack[len(stack)-1].metadata = v
			if stack[len(stack)-1].children == 0 {
				state = day08Metadata
			} else {
				stack[len(stack)-1].children--
				state = day08ChildHeader
			}
		case day08Metadata:
			total += v
			stack[len(stack)-1].metadata--
			if stack[len(stack)-1].metadata == 0 {
				stack = stack[:len(stack)-1]
				if len(stack) > 0 && stack[len(stack)-1].children > 0 {
					stack[len(stack)-1].children--
					state = day08ChildHeader
				}
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
