package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"math"
	"runtime"
)

var _ = declareDay(5, func(part2 bool, inputReader io.Reader) interface{} {
	if part2 {
		return day05Part2(inputReader)
	}
	return day05Part1(inputReader)
})

func day05Part1(inputReader io.Reader) interface{} {
	input, _ := ioutil.ReadAll(inputReader)
	input = bytes.TrimSpace(input)

	var polymer []byte

	for _, unit := range input {
		if len(polymer) > 0 && day05EqualOppositeCase(polymer[len(polymer)-1], unit) {
			polymer = polymer[:len(polymer)-1]
		} else {
			polymer = append(polymer, unit)
		}
	}

	return len(polymer)
}

func day05Part2(inputReader io.Reader) interface{} {
	input, _ := ioutil.ReadAll(inputReader)
	input = bytes.TrimSpace(input)

	numWorkers := runtime.NumCPU()
	ignoreChan := make(chan byte)
	resultChan := make(chan int)

	go day05PopulateIgnoreChannel(ignoreChan)

	for i := 0; i < numWorkers; i++ {
		go day05DoShortestPolymerWork(input, ignoreChan, resultChan)
	}

	shortestPolymer := math.MaxInt32

	for i := 0; i < numWorkers; i++ {
		workerResult := <-resultChan
		if workerResult < shortestPolymer {
			shortestPolymer = workerResult
		}
	}

	return shortestPolymer
}

func day05PopulateIgnoreChannel(ignoreChan chan<- byte) {
	for ignore := byte('A'); ignore <= 'Z'; ignore++ {
		ignoreChan <- ignore
	}
	close(ignoreChan)
}

func day05DoShortestPolymerWork(input []byte, ignoreChan <-chan byte, resultChan chan<- int) {
	shortestPolymer := math.MaxInt32
	var polymer []byte

	for ignore := range ignoreChan {
		polymer = polymer[:0]

		for _, unit := range input {
			if day05EqualIgnoreCase(unit, ignore) {
				continue
			}

			if len(polymer) > 0 && day05EqualOppositeCase(polymer[len(polymer)-1], unit) {
				polymer = polymer[:len(polymer)-1]
			} else {
				polymer = append(polymer, unit)
			}
		}

		if len(polymer) < shortestPolymer {
			shortestPolymer = len(polymer)
		}
	}

	resultChan <- shortestPolymer
}

func day05EqualOppositeCase(a, b byte) bool {
	return a&^0x20 == b&^0x20 && a&0x20 != b&0x20
}

func day05EqualIgnoreCase(a, b byte) bool {
	return a&^0x20 == b&^0x20
}
