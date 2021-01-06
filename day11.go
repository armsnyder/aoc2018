package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"math"
	"strconv"
)

var _ = declareDay(11, func(part2 bool, inputReader io.Reader) interface{} {
	if part2 {
		return day11Part2(inputReader)
	}
	return day11Part1(inputReader)
})

func day11Part1(inputReader io.Reader) interface{} {
	var grid day11Grid
	grid.loadSerialNumber(inputReader)
	grid.calculatePowerLevels()
	x, y, _ := grid.locateHighestPowerSquare(3)
	return strconv.Itoa(x) + "," + strconv.Itoa(y)
}

func day11Part2(inputReader io.Reader) interface{} {
	var grid day11Grid
	grid.loadSerialNumber(inputReader)
	grid.calculatePowerLevels()
	var maxPowerX, maxPowerY, maxPowerSize, maxPower int
	maxPower = math.MinInt32
	for size := 1; size <= 300; size++ {
		x, y, power := grid.locateHighestPowerSquare(size)
		if power > maxPower {
			maxPowerX, maxPowerY, maxPowerSize, maxPower = x, y, size, power
		}
	}
	return strconv.Itoa(maxPowerX) + "," + strconv.Itoa(maxPowerY) + "," + strconv.Itoa(maxPowerSize)
}

type day11Grid struct {
	serialNumber int
	power        [300][300]int8
}

func (g *day11Grid) loadSerialNumber(inputReader io.Reader) {
	raw, _ := ioutil.ReadAll(inputReader)
	g.serialNumber, _ = strconv.Atoi(string(bytes.TrimSpace(raw)))
}

func (g *day11Grid) calculatePowerLevels() {
	for x := 0; x < 300; x++ {
		rackID := x + 11
		for y := 0; y < 300; y++ {
			g.power[y][x] = int8(((rackID*(y+1)+g.serialNumber)*rackID/100)%10 - 5)
		}
	}
}

func (g *day11Grid) locateHighestPowerSquare(size int) (maxPowerX, maxPowerY, maxPower int) {
	maxPower = math.MinInt32

	for x1 := 0; x1 <= 300-size; x1++ {
		for y1 := 0; y1 <= 300-size; y1++ {
			power := 0

			for x2 := x1; x2 < x1+size; x2++ {
				for y2 := y1; y2 < y1+size; y2++ {
					power += int(g.power[y2][x2])
				}
			}

			if power > maxPower {
				maxPowerX, maxPowerY, maxPower = x1, y1, power
			}
		}
	}

	return maxPowerX + 1, maxPowerY + 1, maxPower
}
