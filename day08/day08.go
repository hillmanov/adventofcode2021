package main

import (
	. "adventofcode/utils"
	"embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var f embed.FS

type entry struct {
	SignalPatterns [10]string
	OutputValues   [4]string
}

func Part1() Any {
	entries := getInput()

	return nil
}

func Part2() Any {
	return nil
}

func getInput() []entry {
	lines, _ := ReadLines(f, "input.txt")

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
