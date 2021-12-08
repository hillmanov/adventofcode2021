package main

import (
	. "adventofcode/utils"
	"embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var fe embed.FS

var maskVal = map[string]int{
	"a": 1,
	"b": 2,
	"c": 4,
	"d": 8,
	"e": 16,
	"f": 32,
	"g": 64,
}

type pattern string

func (p pattern) Mask() int {
	val := 0x000000
	for _, c := range p {
		val |= maskVal[string(c)]
	}
	return val
}

func (p pattern) String() string {
	return string(p)
}

type entry struct {
	SignalPatterns [10]pattern
	OutputValues   [4]pattern
}

func Part1() Any {
	entries := getInput()

	countOfOnesFoursSevensAndEights := 0
	for _, entry := range entries {
		for _, v := range entry.OutputValues {
			switch len(v) {
			case 2, 3, 4, 7:
				countOfOnesFoursSevensAndEights++
			}
		}
	}

	return countOfOnesFoursSevensAndEights
}

func Part2() Any {
	entries := getInput()

	outputValues := []int{}
	for _, entry := range entries {
		d := decode(entry.SignalPatterns)

		displayValue := ""
		for _, out := range entry.OutputValues {
			displayValue += strconv.Itoa(d[out.Mask()])
		}

		outputValue, _ := strconv.Atoi(displayValue)
		outputValues = append(outputValues, outputValue)
	}

	return SumOf(outputValues)
}

func decode(patterns [10]pattern) map[int]int {
	maskToDigit := map[int]int{}
	maskOf := map[int]int{}

	// 1, 4, 7, 8
	for _, p := range patterns {
		switch len(p) {
		case 2:
			maskToDigit[p.Mask()] = 1
			maskOf[1] = p.Mask()
		case 4:
			maskToDigit[p.Mask()] = 4
			maskOf[4] = p.Mask()
		case 3:
			maskToDigit[p.Mask()] = 7
			maskOf[7] = p.Mask()
		case 7:
			maskToDigit[p.Mask()] = 8
			maskOf[8] = p.Mask()
		}
	}

	for _, p := range patterns {
		switch {
		// 2
		case len(p) == 5 && p.Mask()&(maskOf[4]^maskOf[1]) != maskOf[4]^maskOf[1] && (p.Mask()&maskOf[1]&maskOf[7] != maskOf[1]&maskOf[7]):
			maskToDigit[p.Mask()] = 2

		// 3
		case len(p) == 5 && p.Mask()&maskOf[1]&maskOf[7] == maskOf[1]&maskOf[7]:
			maskToDigit[p.Mask()] = 3

		// 5
		case len(p) == 5 && p.Mask()&(maskOf[4]^maskOf[1]) == maskOf[4]^maskOf[1]:
			maskToDigit[p.Mask()] = 5

		// 6
		case len(p) == 6 && p.Mask()&maskOf[4]&maskOf[7] != maskOf[4]&maskOf[7]:
			maskToDigit[p.Mask()] = 6

		// 9
		case len(p) == 6 && p.Mask()&maskOf[4] == maskOf[4]:
			maskToDigit[p.Mask()] = 9
		}
	}

	// Easier to find 0 just by looking for the one that we haven't found yet
	for _, p := range patterns {
		if _, ok := maskToDigit[p.Mask()]; !ok {
			maskToDigit[p.Mask()] = 0
		}
	}

	return maskToDigit
}

func getInput() []entry {
	lines, _ := ReadLines(fe, "input.txt")

	entries := []entry{}
	for _, line := range lines {
		e := entry{}

		fmt.Fscanf(
			strings.NewReader(line),
			"%s %s %s %s %s %s %s %s %s %s | %s %s %s %s",
			&e.SignalPatterns[0],
			&e.SignalPatterns[1],
			&e.SignalPatterns[2],
			&e.SignalPatterns[3],
			&e.SignalPatterns[4],
			&e.SignalPatterns[5],
			&e.SignalPatterns[6],
			&e.SignalPatterns[7],
			&e.SignalPatterns[8],
			&e.SignalPatterns[9],
			&e.OutputValues[0],
			&e.OutputValues[1],
			&e.OutputValues[2],
			&e.OutputValues[3],
		)
		entries = append(entries, e)
	}

	return entries
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 08: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 08: Part 2: = %+v\n", part2Solution)
}
