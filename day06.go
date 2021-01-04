package main

import (
	"bufio"
	"io"
	"math"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
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

	mu := &sync.Mutex{}
	coordAreas := make([]int, len(homeCoords))
	infiniteAreaCoords := make([]bool, len(homeCoords))

	wg := &sync.WaitGroup{}

	box.chunk(32, func(chunk day06BoundingBox) {
		wg.Add(1)

		go func() {
			localCoordAreas := make([]int, len(homeCoords))
			localInfiniteAreaCoords := make([]bool, len(homeCoords))
			var minHomeCoordIndices []int

			for y := chunk.minY; y <= chunk.maxY; y++ {
				for x := chunk.minX; x <= chunk.maxX; x++ {
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
							localInfiniteAreaCoords[minHomeCoordIndices[0]] = true
						} else {
							localCoordAreas[minHomeCoordIndices[0]]++
						}
					}
				}
			}

			mu.Lock()
			for i, v := range localCoordAreas {
				coordAreas[i] += v
			}
			for i, v := range localInfiniteAreaCoords {
				if v {
					infiniteAreaCoords[i] = true
				}
			}
			mu.Unlock()

			wg.Done()
		}()
	})

	wg.Wait()

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
	var area int32
	wg := sync.WaitGroup{}

	box.chunk(32, func(chunk day06BoundingBox) {
		wg.Add(1)

		go func() {
			var localArea int32

			for y := chunk.minY; y <= chunk.maxY; y++ {
				for x := chunk.minX; x <= chunk.maxX; x++ {
					totalDistance := 0

					for _, homeCoord := range homeCoords {
						totalDistance += homeCoord.distance(day06Coord{x, y})
						if totalDistance >= limit {
							break
						}
					}

					if totalDistance < limit {
						localArea++
					}
				}
			}

			atomic.AddInt32(&area, localArea)

			wg.Done()
		}()
	})

	wg.Wait()

	return int(area)
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

func (b day06BoundingBox) chunk(size int, fn func(chunk day06BoundingBox)) {
	for minX := b.minX; minX <= b.maxX; minX += size {
		maxX := minX + size - 1
		if maxX > b.maxX {
			maxX = b.maxX
		}

		for minY := b.minY; minY <= b.maxY; minY += size {
			maxY := minY + size - 1
			if maxY > b.maxY {
				maxY = b.maxY
			}

			fn(day06BoundingBox{minX, minY, maxX, maxY})
		}
	}
}
