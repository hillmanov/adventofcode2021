package main

import (
	. "adventofcode/utils"
	"fmt"
)

func Part1() Any {
	return 5
}

func Part2() Any {
	return "Hello World"
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 04: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 04: Part 2: = %+v\n", part2Solution)
}
