package main

import (
	. "adventofcode/utils"
	"embed"
	"fmt"
	"math"
	"sort"
	"strings"
)

//go:embed input.txt
var f embed.FS

func Part1() Any {
	positions := getInput()
	sort.Ints(positions)
	bestPosition := positions[len(positions)/2] // Median!

	fuel := 0
	for _, position := range positions {
		fuel += abs(position - bestPosition)
	}
	return fuel
}

func Part2() Any {
	positions := getInput()
	sort.Ints(positions)

	bestFuelCost := math.MaxInt
	for position := range positions {
		checkFuelCost := totalFuelCostForPosition(positions, position)
		if checkFuelCost < bestFuelCost {
			bestFuelCost = checkFuelCost
		}
	}

	return bestFuelCost
}

func getInput() []int {
	contents, _ := ReadContents(f, "input.txt")
	parts := strings.Split(strings.Trim(contents, "\n"), ",")

	positions := []int{}
	for _, n := range parts {
		positions = append(positions, ParseInt(n))
	}
	return positions
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func totalFuelCostForPosition(positions []int, position int) int {
	cost := 0
	for _, p := range positions {
		cost += fuelCostAtDistance(abs(p - position))
	}
	return cost
}

func fuelCostAtDistance(dist int) int {
	return ((dist * dist) + dist) / 2
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 07: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 07: Part 2: = %+v\n", part2Solution)
}
