package main

import (
	. "adventofcode/utils"
	"embed"
	"fmt"
	"math"
	"sort"
)

//go:embed input.txt
var f embed.FS

type node struct {
	Col      int
	Row      int
	Distance int
}

func Part1() Any {
	cavern := getInput()
	lowestCost := findLowestRiskTotal(cavern)
	return lowestCost
}

func Part2() Any {
	cavern := getInput()

	bigCavern := make([][]int, len(cavern)*5)
	for i := range bigCavern {
		bigCavern[i] = make([]int, len(cavern)*5)
	}

	for row := 0; row < len(bigCavern); row++ {
		for col := 0; col < len(bigCavern); col++ {
			value := cavern[row%len(cavern)][col%len(cavern)]

			if row >= len(cavern) {
				value += int(math.Floor(float64(row) / float64(len(cavern))))
			}
			if value > 9 {
				value = value % 9
			}
			if col >= len(cavern) {
				value += int(math.Floor(float64(col) / float64(len(cavern))))
			}
			if value > 9 {
				value = value % 9
			}

			bigCavern[row][col] = value
		}
	}

	lowestCost := findLowestRiskTotal(bigCavern)
	return lowestCost
}

func findLowestRiskTotal(cavern [][]int) int {
	// Initialize distance grid
	distances := make([][]int, len(cavern))

	// Weird optimization to fill the distances with values a bit quicker
	for i := range distances {
		distances[i] = make([]int, len(cavern[i]))
		distances[i][0] = math.MaxInt
		for k := 1; k < len(distances[i]); k *= 2 {
			copy(distances[i][k:], distances[i][:k])
		}
	}

	// Make it easier to navigate in every direction
	dRow := []int{-1, 0, 1, 0}
	dCol := []int{0, 1, 0, -1}

	// Start at top left
	shortestPath := []node{}
	shortestPath = append(shortestPath, node{0, 0, 0})

	distances[0][0] = cavern[0][0]

	for len(shortestPath) != 0 {
		currentNode := shortestPath[0]
		shortestPath = shortestPath[1:]

		for i := 0; i < 4; i++ {
			col := currentNode.Col + dCol[i]
			row := currentNode.Row + dRow[i]

			if !(row >= 0 && row < len(cavern) && col >= 0 && col < len(cavern)) { // It is a square cavern, so we can use the len(cavern) for the col
				continue
			}

			if distances[row][col] > distances[currentNode.Row][currentNode.Col]+cavern[row][col] {
				distances[row][col] = distances[currentNode.Row][currentNode.Col] + cavern[row][col]
				shortestPath = append(shortestPath, node{col, row, distances[row][col]})
			}
		}

		// Only sort if we have to.
		if len(shortestPath) > 1 && shortestPath[len(shortestPath)-1].Distance > shortestPath[len(shortestPath)-2].Distance {
			sort.Slice(shortestPath, func(a, b int) bool {
				return shortestPath[a].Distance < shortestPath[b].Distance
			})
		}
	}

	return distances[len(cavern)-1][len(cavern)-1] - cavern[0][0]
}

func getInput() [][]int {
	lines, _ := ReadLines(f, "input.txt")

	cavern := [][]int{}

	for _, line := range lines {
		row := []int{}
		for _, d := range line {
			row = append(row, ParseInt(string(d)))
		}
		cavern = append(cavern, row)
	}

	return cavern
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 15: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 15: Part 2: = %+v\n", part2Solution)
}
