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

// Define a stack
type stack []string

func (s *stack) Push(v string) {
	*s = append(*s, v)
}

func (s *stack) Pop() string {
	if len(*s) == 0 {
		return ""
	}

	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

func (s *stack) Peek() string {
	if len(*s) == 0 {
		return ""
	}

	return (*s)[len(*s)-1]
}

var openerPairs = map[string]string{
	"(": ")",
	"{": "}",
	"[": "]",
	"<": ">",
}

var closerPairs = map[string]string{
	")": "(",
	"}": "{",
	"]": "[",
	">": "<",
}

var corruptedPoints = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}

var balancePoints = map[string]int{
	")": 1,
	"]": 2,
	"}": 3,
	">": 4,
}

////////////////////////////////////////////////////////
func Part1() Any {
	lines := getInput()

	badChars := []string{}

	for _, line := range lines {
		balancer := stack{}
		for _, c := range line {
			if isCloser(c) {
				if balancer.Peek() == closerPairs[c] {
					balancer.Pop()
					continue
				} else {
					badChars = append(badChars, c)
					break
				}
			}
			balancer.Push(c)
		}
	}

	syntaxErrorScore := 0
	for _, c := range badChars {
		syntaxErrorScore += corruptedPoints[c]
	}

	return syntaxErrorScore
}

func Part2() Any {
	lines := getInput()

	incompleteBalancers := []stack{}

	for _, line := range lines {
		balancer := stack{}
		corrupted := false
		for _, c := range line {
			if isCloser(c) {
				if balancer.Peek() == closerPairs[c] {
					balancer.Pop()
					continue
				} else {
					corrupted = true
					break
				}
			}
			balancer.Push(c)
		}
		if !corrupted {
			incompleteBalancers = append(incompleteBalancers, balancer)
		}
	}

	balanceScores := []int{}
	for _, b := range incompleteBalancers {
		var balanceScore int
		for len(b) > 0 {
			c := openerPairs[b.Pop()]
			balanceScore = (5 * balanceScore) + balancePoints[c]
		}
		balanceScores = append(balanceScores, balanceScore)
	}

	sort.Ints(balanceScores)
	middle := balanceScores[len(balanceScores)/2]

	return middle
}

func getInput() [][]string {
	contents, _ := ReadLines(f, "input.txt")

	lines := [][]string{}
	for _, l := range contents {
		lines = append(lines, strings.Split(l, ""))
	}

	return lines

}

func isCloser(c string) bool {
	return closerPairs[c] != ""
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 10: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 10: Part 2: = %+v\n", part2Solution)
}
