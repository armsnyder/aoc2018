package main

import (
	"bufio"
	"io"
	"sort"
	"strconv"
	"strings"
	"time"
)

var _ = declareDay(4, func(part2 bool, inputReader io.Reader) interface{} {
	if part2 {
		return day04Part2(inputReader)
	}
	return day04Part1(inputReader)
})

func day04Part1(inputReader io.Reader) interface{} {
	entries := day04ParseEntries(inputReader)
	summaries := day04CreateSummaries(entries)

	sleepiestGuard := 0
	sleepiestTotalMinutesAsleep := 0

	for guard, sleepSummary := range summaries {
		totalMinutesAsleep := sleepSummary.totalMinutesAsleep()
		if totalMinutesAsleep > sleepiestTotalMinutesAsleep {
			sleepiestGuard = guard
			sleepiestTotalMinutesAsleep = totalMinutesAsleep
		}
	}

	sleepiestMinute := summaries[sleepiestGuard].sleepiestMinute()

	return sleepiestGuard * sleepiestMinute
}

func day04Part2(inputReader io.Reader) interface{} {
	entries := day04ParseEntries(inputReader)
	summaries := day04CreateSummaries(entries)

	sleepiestGuard := 0
	sleepiestMinute := 0
	sleepiestFrequency := 0

	for guard, sleepSummary := range summaries {
		minute := sleepSummary.sleepiestMinute()
		frequency := sleepSummary[minute]
		if frequency > sleepiestFrequency {
			sleepiestGuard = guard
			sleepiestMinute = minute
			sleepiestFrequency = frequency
		}
	}

	return sleepiestGuard * sleepiestMinute
}

func day04ParseEntries(inputReader io.Reader) (entries day04Entries) {
	scanner := bufio.NewScanner(inputReader)

	for scanner.Scan() {
		if len(scanner.Bytes()) == 0 {
			continue
		}
		text := scanner.Text()
		var entry day04Entry
		entry.time, _ = time.Parse("2006-01-02 15:04", text[1:17])
		if strings.HasSuffix(text, "shift") {
			entry.beginsShift = true
			entry.guard, _ = strconv.Atoi(text[26 : len(text)-13])
		} else if strings.HasSuffix(text, "asleep") {
			entry.fallsAsleep = true
		} else {
			entry.wakesUp = true
		}
		entries = append(entries, entry)
	}

	sort.Sort(entries)
	return entries
}

func day04CreateSummaries(entries day04Entries) (summaries map[int]day04GuardSummary) {
	summaries = make(map[int]day04GuardSummary)
	guardID := 0
	fellAsleep := 0

	for _, entry := range entries {
		if entry.beginsShift {
			guardID = entry.guard
		} else if entry.fallsAsleep {
			fellAsleep = entry.time.Minute()
		} else {
			wokeUp := entry.time.Minute()
			summary := summaries[guardID]
			for minute := fellAsleep; minute < wokeUp; minute++ {
				summary[minute]++
			}
			summaries[guardID] = summary
		}
	}

	return summaries
}

type day04Entry struct {
	time        time.Time
	guard       int
	fallsAsleep bool
	beginsShift bool
	wakesUp     bool
}

type day04Entries []day04Entry

func (e day04Entries) Len() int {
	return len(e)
}

func (e day04Entries) Less(i, j int) bool {
	return e[i].time.Before(e[j].time)
}

func (e day04Entries) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

type day04GuardSummary [60]int

func (s day04GuardSummary) totalMinutesAsleep() (total int) {
	for _, m := range s {
		total += m
	}
	return total
}

func (s day04GuardSummary) sleepiestMinute() (minute int) {
	sleepiest := 0
	for i, m := range s {
		if m > sleepiest {
			sleepiest = m
			minute = i
		}
	}
	return minute
}
