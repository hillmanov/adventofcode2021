package main

import (
	. "adventofcode/utils"
	"embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var f embed.FS

func Part1() Any {
	fishCycles := getInput()
	return simulate(fishCycles, 80)
}

func Part2() Any {
	fishCycles := getInput()
	return simulate(fishCycles, 256)
}

func simulate(fishCycles []int, days int) int {
	dayCycleCounts := [9]int{}
	for _, n := range fishCycles {
		dayCycleCounts[n]++
	}

	for day := 0; day < days; day++ {
		nextDayCycleCounts := [9]int{}
		createAndResetCount := dayCycleCounts[0]
		for i := 0; i < 8; i++ {
			nextDayCycleCounts[i] = dayCycleCounts[i+1]
		}
		nextDayCycleCounts[6] = nextDayCycleCounts[6] + createAndResetCount
		nextDayCycleCounts[8] = nextDayCycleCounts[8] + createAndResetCount
		dayCycleCounts = nextDayCycleCounts
	}

	sum := 0
	for _, n := range dayCycleCounts {
		sum += n
	}
	return sum
}

func getInput() []int {
	contents, _ := ReadContents(f, "input.txt")
	parts := strings.Split(strings.Trim(contents, "\n"), ",")

	fishCycles := []int{}
	for _, num := range parts {
		fishCycles = append(fishCycles, ParseInt(num))
	}
	return fishCycles
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 06: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 06: Part 2: = %+v\n", part2Solution)
}
