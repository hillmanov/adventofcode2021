package main

import (
	. "adventofcode/utils"
	"embed"
	"fmt"
)

//go:embed input.txt
var f embed.FS

type point struct {
	Row int
	Col int
}

type fold struct {
	Axis  string
	Value int
}

func Part1() Any {
	paper, folds := getInput()

	paper = doFold(paper, folds[0])

	dotCount := 0
	for row := range paper {
		for col := range paper[row] {
			dotCount += paper[row][col]
		}
	}

	return dotCount
}

func Part2() Any {
	paper, folds := getInput()

	for _, f := range folds {
		paper = doFold(paper, f)
	}

	for row := range paper {
		for col := range paper[row] {
			value := " "
			if paper[row][col] == 1 {
				value = "#"
			}
			fmt.Printf("%s", value)
		}
		fmt.Println("")
	}

	return "See output"
}

func doFold(paper [][]int, f fold) [][]int {
	switch f.Axis {
	case "x":
		for row := range paper {
			for col := range paper[row] {
				paper[row][col] = paper[row][col] | paper[row][len(paper[row])-1-col]
			}
			paper[row] = paper[row][:f.Value]
		}

	case "y":
		for row := range paper {
			for col := range paper[row] {
				paper[row][col] = paper[row][col] | paper[len(paper)-1-row][col]
			}
		}
		paper = paper[:f.Value]
	}

	return paper
}

func getInput() ([][]int, []fold) {
	lines, _ := ReadLines(f, "input.txt")

	maxCol := 0
	maxRow := 0

	points := []point{}
	folds := []fold{}

	parseCoords := true
	parseFolds := false
	for _, line := range lines {

		if len(line) == 0 {
			parseCoords = false
			parseFolds = true
			continue
		}

		if parseCoords {
			p := point{}
			fmt.Sscanf(line, "%d,%d", &p.Col, &p.Row)
			maxCol = MaxInt(maxCol, p.Col)
			maxRow = MaxInt(maxRow, p.Row)
			points = append(points, p)
		} else if parseFolds {
			fld := fold{}
			fmt.Sscanf(line, "fold along %1s=%d", &fld.Axis, &fld.Value)
			folds = append(folds, fld)
		}
	}

	grid := make([][]int, maxRow+1)
	for row := range grid {
		grid[row] = make([]int, maxCol+1)
	}

	for _, p := range points {
		grid[p.Row][p.Col] = 1
	}

	return grid, folds
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 13: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 13: Part 2: = %+v\n", part2Solution)
}
