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

type path []cave

func hasVisitedAnySmallCaveTwice(visits map[cave]int) bool {
	for _, c := range visits {
		if c >= 2 {
			return true
		}
	}
	return false
}

func Part1() Any {
	caveConnections := getInput()

	visited := map[cave]bool{}
	currentPath := path{}
	paths := []path{}

	explore(caveConnections, cave("start"), currentPath, visited, &paths)
	return len(paths)
}

func Part2() Any {
	caveConnections := getInput()

	// Different - we count the visits, not just if we visited
	visited := map[cave]int{}
	currentPath := path{}
	paths := []path{}

	explore2(caveConnections, cave("start"), currentPath, visited, &paths)
	return len(paths)
}

func explore(caveConnections map[cave][]cave, currentCave cave, currentPath path, visited map[cave]bool, paths *[]path) {
	if visited[currentCave] {
		return
	}

	if currentCave.IsSmall() {
		visited[currentCave] = true
	}
	currentPath = append(currentPath, currentCave)

	if currentCave == "end" {
		*paths = append(*paths, currentPath)
		visited["end"] = false
		return
	}
	for _, c := range caveConnections[currentCave] {
		explore(caveConnections, c, currentPath, visited, paths)
	}
	visited[currentCave] = false
}

func explore2(caveConnections map[cave][]cave, currentCave cave, currentPath path, visited map[cave]int, paths *[]path) {
	if visited[currentCave] >= 1 && hasVisitedAnySmallCaveTwice(visited) {
		return
	}

	if currentCave.IsSmall() {
		visited[currentCave]++
	}

	currentPath = append(currentPath, currentCave)

	if currentCave == "end" {
		*paths = append(*paths, currentPath)
		visited["end"] = 0
		return
	}
	for _, c := range caveConnections[currentCave] {
		explore2(caveConnections, c, currentPath, visited, paths)
	}

	if currentCave.IsSmall() {
		visited[currentCave]--
	}
}

func getInput() map[cave][]cave {
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

	return caveConnections
}

func main() {
	part1Solution := Part1()
	fmt.Printf("Day 12: Part 1: = %+v\n", part1Solution)

	part2Solution := Part2()
	fmt.Printf("Day 12: Part 2: = %+v\n", part2Solution)
}
