package main

import (
	. "adventofcode/utils"
	"embed"
	"fmt"
)

//go:embed input.txt
var f embed.FS

func Part1() Any {
	depths := getInput()
	increases := 0

	for i := range depths {
		if i == 0 {
			continue
		}
		if depths[i] > depths[i-1] {
			increases++
		}
	}
	return increases
}

func Part2() Any {
	depths := getInput()
	windowIncreases := 0

	for i := range depths {
		if i == 0 || i == 1 || i == 2 {
			continue
		}
		currentWindow := depths[i-2 : i+1]
		previousWindow := depths[i-3 : i]
		if SumOf(currentWindow) > SumOf(previousWindow) {
			windowIncreases++
		}
	}
	return windowIncreases
}

func getInput() []int {
	depths, err := ReadInts(f, "input.txt")
	if err != nil {
		panic(err)
	}
	return depths
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 01: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 01: Part 2: = %+v\n", part2Solution)
}
