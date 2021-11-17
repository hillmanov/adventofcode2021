package main

import (
	. "adventofcode/utils"
	"fmt"
	"time"
)

func Part1() Any {
	return 17
}

func Part2() Any {
	time.Sleep(1 * time.Second)
	return "It worked!"
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 05: Part 1: = %+v\\n", part1Solution)
	fmt.Printf("Day 05: Part 2: = %+v\\n", part2Solution)
}
