package main

import (
	. "adventofcode/utils"
	"embed"
	"fmt"
	"sort"
	"strings"
)

//go:embed input.txt
var f embed.FS

type elementCount struct {
	Element string
	Count   int
}

func Part1() Any {
	template, rules := getInput()

	// Initialize first counts
	counter := createPairCounter(rules)
	for _, p := range getPairs(template) {
		counter[p]++
	}

	for i := 0; i < 9; i++ {
		counter = step(counter, rules)
	}

	// Calculate final counts
	elementCounts := map[string]int{}
	for p, c := range counter {
		if c > 0 {
			element := rules[p]
			elementCounts[string(p[0])] += c
			elementCounts[element] += c
		}
	}
	elementCounts[string(template[len(template)-1])] += 1

	min, max := minMaxElementCounts(elementCounts)

	return max - min
}

func Part2() Any {
	template, rules := getInput()

	// Initialize first counts
	counter := createPairCounter(rules)
	for _, p := range getPairs(template) {
		counter[p]++
	}

	for i := 0; i < 39; i++ {
		counter = step(counter, rules)
	}

	// Calculate final counts
	elementCounts := map[string]int{}
	for p, c := range counter {
		if c > 0 {
			element := rules[p]
			elementCounts[string(p[0])] += c
			elementCounts[element] += c
		}
	}
	elementCounts[string(template[len(template)-1])] += 1

	min, max := minMaxElementCounts(elementCounts)

	return max - min
}

func step(counter map[string]int, rules map[string]string) map[string]int {
	pairCounter := createPairCounter(rules)

	for p, c := range counter {
		if c > 0 {
			element := rules[p]
			leftExpansion := string(p[0]) + element
			rightExpansion := element + string(p[1])
			pairCounter[leftExpansion] += c
			pairCounter[rightExpansion] += c
		}
	}

	return pairCounter
}

func createPairCounter(rules map[string]string) map[string]int {
	pairCounter := make(map[string]int)
	for pair := range rules {
		pairCounter[pair] = 0
	}
	return pairCounter
}

func minMaxElementCounts(counter map[string]int) (int, int) {
	elementCounts := []elementCount{}
	for element, count := range counter {
		elementCounts = append(elementCounts, elementCount{Element: element, Count: count})
	}

	sort.Slice(elementCounts, func(i, j int) bool {
		return elementCounts[i].Count < elementCounts[j].Count
	})

	return elementCounts[0].Count, elementCounts[len(elementCounts)-1].Count
}

func getPairs(t string) []string {
	pairs := []string{}
	for i := 0; i < len(t)-1; i++ {
		pairs = append(pairs, string(t[i])+string(t[i+1]))
	}
	return pairs
}

func getInput() (string, map[string]string) {
	lines, _ := ReadLines(f, "input.txt")

	polymerTemplate := strings.Trim(lines[0], " ")
	pairInsertionRules := make(map[string]string)

	for _, line := range lines[2:] {
		var left, right string
		fmt.Sscanf(line, "%s -> %s", &left, &right)
		pairInsertionRules[left] = right
	}

	return polymerTemplate, pairInsertionRules
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 14: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 14: Part 2: = %+v\n", part2Solution)
}
