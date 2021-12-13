package main

import (
	. "adventofcode/utils"
	"embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var f embed.FS

type cave string

func Part1() Any {
	caves, connections := getInput()

	return nil
}

func Part2() Any {
	return nil
}

func getInput() ([]cave, map[cave][]cave) {
	lines, _ := ReadLines(f, "input.txt")

	caveConnections := map[cave][]cave{}
	for _, line := range lines {
		var left, right cave
		parts := strings.Split(line, "-")
		left = cave(parts[0])
		right = cave(parts[1])

		if _, ok := caveConnections[left]; !ok {
			caveConnections[left] = []cave{}
		}
		caveConnections[left] = append(caveConnections[left], right)
	}

	caves := []cave{}
	for c := range caveConnections {
		caves = append(caves, c)
	}

	return caves, caveConnections
}

func contains(path []cave, c cave) bool {
	for _, pc := range path {
		if pc == c {
			return true
		}
	}
	return false
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 12: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 12: Part 2: = %+v\n", part2Solution)
}
