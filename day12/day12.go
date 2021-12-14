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

func (c cave) IsSmall() bool {
	return strings.ToLower(string(c)) == string(c)
}

func (c cave) IsBig() bool {
	return !c.IsSmall()
}

type path []cave

func (p path) Visits(c cave) int {
	visits := 0
	for _, pc := range p {
		if pc == c {
			visits++
		}
	}
	return visits
}

func (p path) HasAnySmallCaveTwice() bool {
	counts := map[cave]int{}
	for _, c := range p {
		if c.IsSmall() {
			counts[c]++
			if counts[c] >= 2 {
				return true
			}
		}
	}
	return false
}

func (p path) String() string {
	j := ""
	for _, c := range p {
		j += string(c) + ","
	}
	return strings.Trim(j, ",")
}

func Part1() Any {
	_, connections := getInput()

	startingPath := path{cave("start")}

	fullPaths := map[string]bool{
		startingPath.String(): true,
	}
	fullPaths = explore(connections, startingPath, fullPaths, false)

	i := 0
	for p := range fullPaths {
		if strings.HasSuffix(p, "end") {
			i++
		}
	}

	return i
}

func Part2() Any {
	_, connections := getInput()

	startingPath := path{cave("start")}

	fullPaths := map[string]bool{
		startingPath.String(): true,
	}
	fullPaths = explore(connections, startingPath, fullPaths, true)

	i := 0
	for p := range fullPaths {
		if strings.HasSuffix(p, "end") {
			i++
		}
	}

	return i
}
func explore(connections map[cave][]cave, workingPath path, fullPaths map[string]bool, allowDoubleVisitToSingleCave bool) map[string]bool {
	currentCave := workingPath[len(workingPath)-1]

	if currentCave == "end" {
		fullPaths[workingPath.String()] = true
		return fullPaths
	}

	for _, connectedCave := range connections[currentCave] {
		if connectedCave == "start" {
			continue
		}

		if connectedCave.IsSmall() {
			if allowDoubleVisitToSingleCave {
				if workingPath.Visits(connectedCave) == 1 && !workingPath.HasAnySmallCaveTwice() {
					goto allowExplore
				} else if workingPath.Visits(connectedCave) >= 1 && workingPath.HasAnySmallCaveTwice() {
					continue
				}
			} else {
				if workingPath.Visits(connectedCave) >= 1 {
					continue
				}
			}
		}

	allowExplore:
		exploreTo := make(path, len(workingPath))
		copy(exploreTo, workingPath)
		exploreTo = append(exploreTo, connectedCave)

		for path := range explore(connections, exploreTo, fullPaths, allowDoubleVisitToSingleCave) {
			fullPaths[path] = true
		}
	}

	return fullPaths
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
		if _, ok := caveConnections[right]; !ok {
			caveConnections[right] = []cave{}
		}

		if right != "start" {
			caveConnections[left] = append(caveConnections[left], right)
		}
		if left != "start" {
			caveConnections[right] = append(caveConnections[right], left)
		}
	}

	caves := []cave{}
	for c := range caveConnections {
		caves = append(caves, c)
	}

	return caves, caveConnections
}

func main() {
	part1Solution := Part1()
	fmt.Printf("Day 12: Part 1: = %+v\n", part1Solution)

	part2Solution := Part2()
	fmt.Printf("Day 12: Part 2: = %+v\n", part2Solution)
}
