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

	for day := 0; day < 80; day++ {
		nextFishCycles := []int{}
		newFishCycles := []int{}
		for _, n := range fishCycles {
			if n == 0 {
				nextFishCycles = append(nextFishCycles, 6)
				newFishCycles = append(newFishCycles, 8)
			} else {
				nextFishCycles = append(nextFishCycles, n-1)
			}
		}
		nextFishCycles = append(nextFishCycles, newFishCycles...)
		fishCycles = nextFishCycles
	}

	return len(fishCycles)
}

func Part2() Any {
	nums := getInput()

	dayCycleCounts := [9]int{}
	for _, n := range nums {
		dayCycleCounts[n]++
	}

	for day := 0; day < 256; day++ {
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

	nums := []int{}
	for _, num := range parts {
		nums = append(nums, ParseInt(num))
	}
	return nums
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 06: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 06: Part 2: = %+v\n", part2Solution)
}
