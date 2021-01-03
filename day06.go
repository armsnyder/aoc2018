package main

import (
	"bufio"
	"io"
	"math"
	"strconv"
	"strings"
)

var _ = declareDay(6, func(part2 bool, inputReader io.Reader) interface{} {
	if part2 {
		return day06Part2(inputReader, 10000)
	}
	return day06Part1(inputReader)
})

func day06Part1(inputReader io.Reader) interface{} {
	homeCoords := day06ParseCoords(inputReader)
	box := day06GetBoundingBox(homeCoords)
	coordAreas := make([]int, len(homeCoords))
	infiniteAreaCoords := make([]bool, len(homeCoords))
	var minHomeCoordIndices []int

	for y := box.minY; y <= box.maxY; y++ {
		for x := box.minX; x <= box.maxX; x++ {
			minDistance := math.MaxInt32

			for i, homeCoord := range homeCoords {
				dist := homeCoord.distance(day06Coord{x, y})
				if dist < minDistance {
					minDistance = dist
					minHomeCoordIndices = append(minHomeCoordIndices[:0], i)
				} else if dist == minDistance {
					minHomeCoordIndices = append(minHomeCoordIndices, i)
				}
			}

			if len(minHomeCoordIndices) == 1 {
				if box.atBounds(day06Coord{x, y}) {
					infiniteAreaCoords[minHomeCoordIndices[0]] = true
				} else {
					coordAreas[minHomeCoordIndices[0]]++
				}
			}
		}
	}

	maxArea := math.MinInt32

	for i, area := range coordAreas {
		if infiniteAreaCoords[i] {
			continue
		}

		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}

func day06Part2(inputReader io.Reader, limit int) interface{} {
	homeCoords := day06ParseCoords(inputReader)
	box := day06GetBoundingBox(homeCoords)
	area := 0

	for y := box.minY; y <= box.maxY; y++ {
		for x := box.minX; x <= box.maxX; x++ {
			totalDistance := 0

			for _, homeCoord := range homeCoords {
				totalDistance += homeCoord.distance(day06Coord{x, y})
				if totalDistance >= limit {
					break
				}
			}

			if totalDistance < limit {
				area++
			}
		}
	}

	return area
}

func day06ParseCoords(inputReader io.Reader) (coords []day06Coord) {
	scanner := bufio.NewScanner(inputReader)

	for scanner.Scan() {
		if len(scanner.Bytes()) == 0 {
			continue
		}

		var coord day06Coord
		split := strings.SplitN(scanner.Text(), ", ", 2)
		coord.x, _ = strconv.Atoi(split[0])
		coord.y, _ = strconv.Atoi(split[1])

		coords = append(coords, coord)
	}

	return coords
}

func day06GetBoundingBox(coords []day06Coord) (boundingBox day06BoundingBox) {
	boundingBox.minX, boundingBox.minY = math.MaxInt32, math.MaxInt32
	boundingBox.maxX, boundingBox.maxY = math.MinInt32, math.MinInt32

	for _, coord := range coords {
		if coord.x < boundingBox.minX {
			boundingBox.minX = coord.x
		}
		if coord.y < boundingBox.minY {
			boundingBox.minY = coord.y
		}
		if coord.x > boundingBox.maxX {
			boundingBox.maxX = coord.x
		}
		if coord.y > boundingBox.maxY {
			boundingBox.maxY = coord.y
		}
	}

	return boundingBox
}

type day06Coord struct{ x, y int }

func (c day06Coord) distance(c2 day06Coord) int {
	dx := c.x - c2.x
	if dx < 0 {
		dx = -dx
	}

	dy := c.y - c2.y
	if dy < 0 {
		dy = -dy
	}

	return dx + dy
}

type day06BoundingBox struct{ minX, minY, maxX, maxY int }

func (b day06BoundingBox) atBounds(coord day06Coord) bool {
	return coord.x == b.minX || coord.y == b.minY || coord.x == b.maxX || coord.y == b.maxY
}
