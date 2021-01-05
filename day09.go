package main

import (
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

var _ = declareDay(9, func(part2 bool, inputReader io.Reader) interface{} {
	players, lastMarble := day09Parse(inputReader)
	if part2 {
		lastMarble *= 100
	}

	cur := &day09Marble{}
	cur.cw = cur
	cur.ccw = cur

	scores := make([]int, players)
	turn := 0

	for marble := 1; marble <= lastMarble; marble++ {
		if marble%23 == 0 {
			scores[turn] += marble
			removed := cur.getMarbleCCW(7).remove()
			scores[turn] += removed.value
			cur = removed.cw
		} else {
			cur = cur.cw.addMarbleCW(marble)
		}

		turn++
		turn %= players
	}

	highScore := 0

	for _, score := range scores {
		if score > highScore {
			highScore = score
		}
	}

	return highScore
})

func day09Parse(inputReader io.Reader) (players, lastMarble int) {
	b, _ := ioutil.ReadAll(inputReader)
	text := string(b)

	i := strings.IndexByte(text, ' ')
	players, _ = strconv.Atoi(text[:i])

	j := strings.LastIndexByte(text, ' ')
	i = strings.LastIndexByte(text[:j], ' ')
	lastMarble, _ = strconv.Atoi(text[i+1 : j])

	return players, lastMarble
}

type day09Marble struct {
	value int
	cw    *day09Marble
	ccw   *day09Marble
}

func (m *day09Marble) addMarbleCW(value int) *day09Marble {
	displacedMarble := m.cw
	marble := &day09Marble{value: value}
	marble.cw = displacedMarble
	marble.ccw = m
	m.cw = marble
	displacedMarble.ccw = marble
	return marble
}

func (m *day09Marble) remove() *day09Marble {
	m.cw.ccw = m.ccw
	m.ccw.cw = m.cw
	return m
}

func (m *day09Marble) getMarbleCCW(n int) *day09Marble {
	cur := m
	for i := 0; i < n; i++ {
		cur = cur.ccw
	}
	return cur
}
