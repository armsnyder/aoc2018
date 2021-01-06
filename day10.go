package main

import (
	"bufio"
	"io"
	"math"
	"strconv"
	"strings"
	"unicode"
)

var _ = declareDay(10, func(part2 bool, inputReader io.Reader) interface{} {
	points := day10Parse(inputReader)
	elapsed := points.minimizeArea()
	if part2 {
		return elapsed
	}
	return points.String()
})

func day10Parse(inputReader io.Reader) (points day10Points) {
	scanner := bufio.NewScanner(inputReader)

	for scanner.Scan() {
		if len(scanner.Bytes()) == 0 {
			continue
		}

		fields := strings.FieldsFunc(scanner.Text(), func(r rune) bool {
			return !(unicode.IsDigit(r) || r == '-')
		})

		var point day10Point

		point.pos.x, _ = strconv.Atoi(fields[0])
		point.pos.y, _ = strconv.Atoi(fields[1])
		point.vel.x, _ = strconv.Atoi(fields[2])
		point.vel.y, _ = strconv.Atoi(fields[3])

		points = append(points, point)
	}

	return points
}

type day10Vector struct{ x, y int }

type day10Point struct{ pos, vel day10Vector }

func (p day10Point) timeOfIntersection(p2 day10Point) int {
	return (p2.pos.x - p.pos.x) / (p.vel.x - p2.vel.x)
}

func (p day10Point) after(t int) day10Point {
	return day10Point{
		pos: day10Vector{
			x: p.pos.x + t*p.vel.x,
			y: p.pos.y + t*p.vel.y,
		},
		vel: p.vel,
	}
}

type day10Points []day10Point

func (p day10Points) minimizeArea() (elapsed int) {
	pi := p[0]
	pj := p[0]
	for j := 1; pi.vel.x == pj.vel.x; j++ {
		pj = p[j]
	}
	t := pi.timeOfIntersection(pj)
	p.tick(t)
	elapsed += t

	smallestArea := p.area()

	hillClimb := func(dt int) {
		for {
			p.tick(dt)
			elapsed += dt
			candidateArea := p.area()
			if candidateArea < smallestArea {
				smallestArea = candidateArea
			} else {
				p.tick(-dt)
				elapsed -= dt
				return
			}
		}
	}

	hillClimb(1)
	hillClimb(-1)

	return elapsed
}

func (p day10Points) tick(t int) {
	for i, pp := range p {
		p[i] = pp.after(t)
	}
}

func (p day10Points) boundary() (xMin, xMax, yMin, yMax int) {
	xMin, yMin = math.MaxInt32, math.MaxInt32
	xMax, yMax = math.MinInt32, math.MinInt32

	for _, pp := range p {
		if pp.pos.x < xMin {
			xMin = pp.pos.x
		}
		if pp.pos.x > xMax {
			xMax = pp.pos.x
		}
		if pp.pos.y < yMin {
			yMin = pp.pos.y
		}
		if pp.pos.y > yMax {
			yMax = pp.pos.y
		}
	}

	xMax += 1
	yMax += 1

	return xMin, xMax, yMin, yMax
}

func (p day10Points) area() int {
	xMin, xMax, yMin, yMax := p.boundary()
	return (xMax - xMin) * (yMax - yMin)
}

func (p day10Points) String() string {
	xMin, xMax, yMin, yMax := p.boundary()

	buf := make([][]byte, yMax-yMin)

	for y := 0; y < yMax-yMin; y++ {
		buf[y] = make([]byte, xMax-xMin)

		for x := 0; x < xMax-xMin; x++ {
			buf[y][x] = ' '
		}
	}

	for _, pp := range p {
		x := pp.pos.x - xMin
		y := pp.pos.y - yMin

		buf[y][x] = '#'
	}

	sb := &strings.Builder{}
	sb.Grow((xMax - xMin + 1) * (yMax - yMin))

	for _, row := range buf {
		sb.WriteByte('\n')
		sb.Write(row)
	}

	return sb.String()
}
