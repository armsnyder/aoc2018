package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/armsnyder/aoc2018/aocutil"
)

type dayFn func(part2 bool, inputReader io.Reader) interface{}

var days = make(map[int]dayFn)

func declareDay(day int, dayFn dayFn) interface{} {
	days[day] = dayFn
	return nil
}

func main() {
	easternTime, err := time.LoadLocation("EST")
	if err != nil {
		panic(err)
	}

	var (
		day        int
		part2      bool
		skipSubmit bool
		inputFile  string
	)
	flag.BoolVar(&part2, "2", false, "Run part 2")
	flag.BoolVar(&skipSubmit, "s", false, "Skip submitting the answer")
	flag.StringVar(&inputFile, "f", "", "Puzzle input override filepath")
	flag.Parse()
	day, err = strconv.Atoi(flag.Arg(0))
	if err != nil {
		now := time.Now()
		if now.Month() == 12 {
			day = now.In(easternTime).Day()
		} else {
			panic(fmt.Errorf("Could not parse positional argument: day: %w", err))
		}
	}

	dayFn, ok := days[day]
	if !ok {
		fmt.Printf("Generating a stub for day %d\n", day)
		aocutil.GenerateStub(day)
	}

	var input io.ReadCloser
	if inputFile != "" {
		var err error
		input, err = os.Open(inputFile)
		if err != nil {
			panic(err)
		}
	} else {
		input = aocutil.GetInput(day)
	}
	defer input.Close()

	if !ok {
		return
	}

	start := time.Now()
	output := dayFn(part2, input)
	runTime := time.Since(start)

	fmt.Printf("Finished in %s\n", runTime)
	fmt.Println("Answer:", output)

	if !skipSubmit {
		aocutil.Submit(day, part2, output)
	}
}
