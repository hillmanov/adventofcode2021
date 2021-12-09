package main

import (
	. "adventofcode/utils"
	"embed"
	"fmt"
	"sort"
	"strconv"
)

//go:embed input.txt
var f embed.FS

type basin map[point]bool

func (b basin) Merge(t basin) {
	for k, v := range t {
		b[k] = v
	}
}

func (b basin) HasPoint(p point) bool {
	if _, ok := b[p]; ok {
		return true
	}
	return false
}

type point struct {
	Col    int
	Row    int
	Height int
}

func (p point) RiskLevel() int {
	return p.Height + 1
}

func Part1() Any {
	heightMap := getInput()

	lowPoints := getLowPoints(heightMap)
	sumOfRiskLevels := 0
	for _, p := range lowPoints {
		sumOfRiskLevels += p.RiskLevel()
	}

	return sumOfRiskLevels
}

func Part2() Any {
	heightMap := getInput()
	lowPoints := getLowPoints(heightMap)

	basins := []basin{}
	for _, lowPoint := range lowPoints {
		b := basin{}
		findBasin(b, heightMap, lowPoint)
		basins = append(basins, b)
	}

	basinSizes := []int{}
	for _, b := range basins {
		basinSizes = append(basinSizes, len(b))
	}
	sort.Ints(basinSizes)

	product := 1
	for _, size := range basinSizes[len(basinSizes)-3:] {
		product *= size
	}

	return product
}

func findBasin(b basin, heightMap [][]int, p point) basin {
	b[p] = true

	// Left
	if p.Col > 0 && heightMap[p.Row][p.Col-1] < 9 {
		point := point{Col: p.Col - 1, Row: p.Row, Height: heightMap[p.Row][p.Col-1]}
		if !b.HasPoint(point) {
			b.Merge(findBasin(b, heightMap, point))
		}
	}

	// Right
	if p.Col < len(heightMap[p.Row])-1 && heightMap[p.Row][p.Col+1] < 9 {
		point := point{Col: p.Col + 1, Row: p.Row, Height: heightMap[p.Row][p.Col+1]}
		if !b.HasPoint(point) {
			b.Merge(findBasin(b, heightMap, point))
		}
	}

	// Up
	if p.Row > 0 && heightMap[p.Row-1][p.Col] < 9 {
		point := point{Col: p.Col, Row: p.Row - 1, Height: heightMap[p.Row-1][p.Col]}
		if !b.HasPoint(point) {
			b.Merge(findBasin(b, heightMap, point))
		}
	}

	// Down
	if p.Row < len(heightMap)-1 && heightMap[p.Row+1][p.Col] < 9 {
		point := point{Col: p.Col, Row: p.Row + 1, Height: heightMap[p.Row+1][p.Col]}
		if !b.HasPoint(point) {
			b.Merge(findBasin(b, heightMap, point))
		}
	}

	return b
}

func getLowPoints(heightMap [][]int) []point {
	lowPoints := []point{}

	for row := range heightMap {
		for col := range heightMap[row] {
			adjecentValues := []int{}

			// Left
			if col > 0 {
				adjecentValues = append(adjecentValues, heightMap[row][col-1])
			}

			// Right
			if col < len(heightMap[row])-1 {
				adjecentValues = append(adjecentValues, heightMap[row][col+1])
			}

			// Up
			if row > 0 {
				adjecentValues = append(adjecentValues, heightMap[row-1][col])
			}

			// Down
			if row < len(heightMap)-1 {
				adjecentValues = append(adjecentValues, heightMap[row+1][col])
			}

			isLowPoint := true
			for _, adadjecentValue := range adjecentValues {
				isLowPoint = isLowPoint && (heightMap[row][col] < adadjecentValue)
			}
			if isLowPoint {
				lowPoints = append(lowPoints, point{Col: col, Row: row, Height: heightMap[row][col]})
			}
		}
	}

	return lowPoints
}

func getInput() [][]int {
	heightMap := [][]int{}
	lines, _ := ReadLines(f, "input.txt")

	for _, line := range lines {
		lineHeights := []int{}
		for _, s := range line {
			digit, _ := strconv.Atoi(string(s))
			lineHeights = append(lineHeights, digit)
		}
		heightMap = append(heightMap, lineHeights)
	}

	return heightMap
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 09: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 09: Part 2: = %+v\n", part2Solution)
}
