package main

import (
	. "adventofcode/utils"
	"embed"
	"fmt"
)

//go:embed input.txt
var f embed.FS

func Part1() Any {
	return nil
}

func Part2() Any {
	return nil
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

  fmt.Printf("Day 23: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 23: Part 2: = %+v\n", part2Solution)
}
