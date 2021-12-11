package main

import (
	. "adventofcode/utils"
	"embed"
	"fmt"
	"strconv"
)

//go:embed input.txt
var f embed.FS

type octopus struct {
	E         int
	Flashed   bool
	Neighbors []*octopus
}

func Part1() Any {
	octopi := getInput()

	totalFlashes := 0
	for steps := 0; steps < 100; steps++ {

		// Increase each E by 1
		for row := range octopi {
			for col := range octopi[row] {
				octopi[row][col].E += 1
			}
		}

		// Flash and ripple out until there were no more flashes
		for {
			wasFlash := false
			// Find all that are more than 9
			for row := range octopi {
				for col := range octopi[row] {
					octopus := &octopi[row][col]
					if octopus.E > 9 && !octopus.Flashed {
						octopus.Flashed = true
						totalFlashes++
						wasFlash = true
						for _, n := range octopus.Neighbors {
							if !n.Flashed {
								n.E += 1
							}
						}
					}
				}
			}

			if !wasFlash {
				break
			}
		}

		// Reset all
		for row := range octopi {
			for col := range octopi[row] {
				octopi[row][col].Flashed = false
				if octopi[row][col].E > 9 {
					octopi[row][col].E = 0
				}
			}
		}
	}

	return totalFlashes
}

func Part2() Any {
	octopi := getInput()

	for steps := 0; ; steps++ {
		stepFlashes := 0

		// Increase each E by 1
		for row := range octopi {
			for col := range octopi[row] {
				octopi[row][col].E += 1
			}
		}

		// Flash and ripple out until there were no more flashes
		for {
			wasFlash := false
			// Find all that are more than 9
			for row := range octopi {
				for col := range octopi[row] {
					octopus := &octopi[row][col]
					if octopus.E > 9 && !octopus.Flashed {
						octopus.Flashed = true
						stepFlashes++
						wasFlash = true
						for _, n := range octopus.Neighbors {
							if !n.Flashed {
								n.E += 1
							}
						}
					}
				}
			}

			if !wasFlash {
				break
			}
		}

		// Reset all
		for row := range octopi {
			for col := range octopi[row] {
				octopi[row][col].Flashed = false
				if octopi[row][col].E > 9 {
					octopi[row][col].E = 0
				}
			}
		}

		if stepFlashes == len(octopi)*len(octopi) {
			return steps + 1
		}
	}
}

func getInput() [][]octopus {
	lines, _ := ReadLines(f, "input.txt")

	octopi := [][]octopus{}

	for _, line := range lines {
		rowOctopi := []octopus{}
		for _, c := range line {
			e, _ := strconv.Atoi(string(c))
			rowOctopi = append(rowOctopi, octopus{E: e, Neighbors: []*octopus{}})
		}
		octopi = append(octopi, rowOctopi)
	}

	for row := range octopi {
		for col := range octopi[row] {
			// T
			if row > 0 {
				octopi[row][col].Neighbors = append(octopi[row][col].Neighbors, &octopi[row-1][col])
			}
			// TR
			if row > 0 && col < len(octopi[row])-1 {
				octopi[row][col].Neighbors = append(octopi[row][col].Neighbors, &octopi[row-1][col+1])
			}
			// R
			if col < len(octopi[row])-1 {
				octopi[row][col].Neighbors = append(octopi[row][col].Neighbors, &octopi[row][col+1])
			}
			//BR
			if row < len(octopi)-1 && col < len(octopi[row])-1 {
				octopi[row][col].Neighbors = append(octopi[row][col].Neighbors, &octopi[row+1][col+1])
			}
			//B
			if row < len(octopi)-1 {
				octopi[row][col].Neighbors = append(octopi[row][col].Neighbors, &octopi[row+1][col])
			}
			//BL
			if row < len(octopi)-1 && col > 0 {
				octopi[row][col].Neighbors = append(octopi[row][col].Neighbors, &octopi[row+1][col-1])
			}
			// L
			if col > 0 {
				octopi[row][col].Neighbors = append(octopi[row][col].Neighbors, &octopi[row][col-1])
			}
			// TL
			if row > 0 && col > 0 {
				octopi[row][col].Neighbors = append(octopi[row][col].Neighbors, &octopi[row-1][col-1])
			}
		}
	}

	return octopi
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 11: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 11: Part 2: = %+v\n", part2Solution)
}
