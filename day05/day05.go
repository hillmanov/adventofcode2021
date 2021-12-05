package main

import (
	. "adventofcode/utils"
	"embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var f embed.FS

type point struct {
	X int
	Y int
}

type line struct {
	Start point
	End   point
}

func (l line) IsVertical() bool {
	return l.Start.X == l.End.X && l.Start.Y != l.End.Y
}

func (l line) IsHoritzontal() bool {
	return l.Start.X != l.End.X && l.Start.Y == l.End.Y
}

func (l line) GetPoints() []point {
	points := []point{}
	x, y := l.Start.X, l.Start.Y
	for {
		points = append(points, point{X: x, Y: y})
		if x == l.End.X && y == l.End.Y {
			break
		}

		if x < l.End.X {
			x++
		} else if x > l.End.X {
			x--
		}

		if y < l.End.Y {
			y++
		} else if y > l.End.Y {
			y--
		}
	}

	return points
}

func Part1() Any {
	lines := getInput()

	pointCounts := make(map[point]int)
	for _, line := range lines {
		if line.IsHoritzontal() || line.IsVertical() {
			for _, point := range line.GetPoints() {
				pointCounts[point] = pointCounts[point] + 1
			}
		}
	}

	overlappingCount := 0
	for _, count := range pointCounts {
		if count > 1 {
			overlappingCount++
		}
	}

	return overlappingCount
}

func Part2() Any {
	lines := getInput()

	pointCounts := make(map[point]int)
	for _, line := range lines {
		for _, point := range line.GetPoints() {
			pointCounts[point] = pointCounts[point] + 1
		}
	}

	overlappingCount := 0
	for _, count := range pointCounts {
		if count >= 2 {
			overlappingCount++
		}
	}

	return overlappingCount
}

func getInput() []line {
	inputs, _ := ReadLines(f, "input.txt")

	lines := []line{}
	for _, input := range inputs {
		l := line{}
		fmt.Fscanf(strings.NewReader(input), "%d,%d -> %d,%d", &l.Start.X, &l.Start.Y, &l.End.X, &l.End.Y)
		lines = append(lines, l)
	}

	return lines
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 05: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 05: Part 2: = %+v\n", part2Solution)
}
