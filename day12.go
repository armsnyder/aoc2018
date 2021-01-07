package main

import (
	"bufio"
	"io"
	"math/big"
	"strings"
)

var _ = declareDay(12, func(part2 bool, inputReader io.Reader) interface{} {
	if part2 {
		return day12(inputReader, 50000000000)
	}
	return day12(inputReader, 20)
})

func day12(inputReader io.Reader, n int) interface{} {
	pots, rules := day12Parse(inputReader)
	memoIndex := make(map[string]int)
	var memo []day12Pots

	for i := 0; i < n; i++ {
		hash := pots.hash()

		if foundIndex, ok := memoIndex[hash]; ok {
			deltaIndex := i - foundIndex
			growthsToGo := n - i + deltaIndex
			growthCyclesToGo := growthsToGo / deltaIndex
			growthsRemainder := growthsToGo % deltaIndex
			deltaOffset := pots.offset - memo[foundIndex].offset
			foundPotRemainder := memo[foundIndex+growthsRemainder]
			pots = foundPotRemainder
			pots.offset += deltaOffset * growthCyclesToGo
			break
		}

		memoIndex[hash] = i
		memo = append(memo, pots)
		pots = pots.grow(rules)
	}

	return pots.sum()
}

type day12Rules [32]bool

func (r *day12Rules) set(input int, output bool) {
	r[input] = output
}

func (r *day12Rules) check(input int) bool {
	return r[input]
}

type day12Pots struct {
	offset int
	pots   *big.Int
}

func (p *day12Pots) set(index int, value bool) {
	if p.pots == nil {
		p.pots = big.NewInt(0)
	}

	var bit uint
	if value {
		bit = 1
	}

	p.pots = p.pots.SetBit(p.pots, index, bit)
}

func (p day12Pots) hash() string {
	return string(p.pots.Bytes())
}

func (p day12Pots) grow(rules day12Rules) (result day12Pots) {
	const edge = 2
	result.pots = big.NewInt(0)
	result.offset = p.offset - edge
	bitLen := p.pots.BitLen()
	for potIndex := -edge; potIndex <= bitLen+edge; potIndex++ {
		var lookup int
		for li := 0; li < 5; li++ {
			ppi := potIndex - edge + li
			if ppi >= 0 && p.pots.Bit(ppi) == 1 {
				lookup |= 1 << li
			}
		}
		if rules.check(lookup) {
			result.pots.SetBit(result.pots, potIndex+2, 1)
		}
	}
	trailingZeroes := result.pots.TrailingZeroBits()
	result.pots.Rsh(result.pots, trailingZeroes)
	result.offset += int(trailingZeroes)
	return result
}

func (p *day12Pots) sum() int {
	var result int

	bitLen := p.pots.BitLen()
	for i := 0; i < bitLen; i++ {
		if p.pots.Bit(i) == 1 {
			result += p.offset + i
		}
	}

	return result
}

func day12Parse(inputReader io.Reader) (pots day12Pots, rules day12Rules) {
	scanner := bufio.NewScanner(inputReader)
	scanner.Scan()
	initialState := strings.TrimPrefix(scanner.Text(), "initial state: ")
	for i, ch := range initialState {
		pots.set(i, ch == '#')
	}
	scanner.Scan()
	for scanner.Scan() {
		split := strings.SplitN(scanner.Text(), " => ", 2)
		if split[1] == "#" {
			var input int
			for i, ch := range split[0] {
				if ch == '#' {
					input |= 1 << i
				}
			}
			rules.set(input, true)
		}
	}
	return pots, rules
}
