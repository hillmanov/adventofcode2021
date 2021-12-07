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
	nums := getInput()

	for day := 0; day <= 18; day++ {
		nextNums
	}

	return nil
}

func Part2() Any {
	return nil
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
