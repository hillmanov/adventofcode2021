package main

import (
	. "adventofcode/utils"
	"embed"
	"fmt"
	"math"
	"strconv"
)

//go:embed input.txt
var f embed.FS

func Part1() Any {
	diagnoticOutputs, bitCount := getInput()
	countsOfOneByPosition := getCountsOfOneByPosition(diagnoticOutputs, bitCount)

	binString := ""
	for _, count := range countsOfOneByPosition {
		if count > len(diagnoticOutputs)/2 {
			binString += "1"
		} else {
			binString += "0"
		}
	}

	gammaRate, _ := strconv.ParseInt(binString, 2, 64)
	epsilonRate := gammaRate ^ (int64(math.Pow(2, float64(bitCount))) - 1)

	powerConsumption := gammaRate * epsilonRate

	return powerConsumption
}

func Part2() Any {
	diagnoticOutputs, bitCount := getInput()

	oxygenSystemRating := whittle(diagnoticOutputs, 0, bitCount, false)
	co2ScrubberRating := whittle(diagnoticOutputs, 0, bitCount, true)

	return oxygenSystemRating * co2ScrubberRating
}

func whittle(currentNums []int, position, bitCount int, invert bool) int {
	for len(currentNums) > 1 {
		filteredNums := filter(currentNums, position, bitCount, invert)
		return whittle(filteredNums, position+1, bitCount, invert)
	}
	return currentNums[0]
}

func filter(nums []int, position, bitCount int, invert bool) []int {
	countsOfOne := getCountsOfOneByPosition(nums, bitCount)

	var filterFor int
	if float64(countsOfOne[position]) >= float64(len(nums))/2 {
		filterFor = 1
	} else {
		filterFor = 0
	}
	if invert {
		filterFor = filterFor ^ 1
	}

	filtered := []int{}
	for _, num := range nums {
		mask := 1 << (bitCount - position - 1)
		switch filterFor {
		case 1:
			if num&(mask) == mask {
				filtered = append(filtered, num)
			}
		case 0:
			if num&(mask) != mask {
				filtered = append(filtered, num)
			}
		}
	}
	return filtered
}

func getCountsOfOneByPosition(diagnosticOutputs []int, bitCount int) []int {
	counts := make([]int, bitCount)
	for _, line := range diagnosticOutputs {
		for i := 0; i < bitCount; i++ {
			if line&(1<<(bitCount-i-1)) != 0 {
				counts[i]++
			}
		}
	}
	return counts
}

func getInput() ([]int, int) {
	lines, err := ReadLines(f, "input.txt")
	if err != nil {
		panic(err)
	}

	nums := []int{}
	var bits int
	for _, v := range lines {
		bits = len(v)
		num, _ := strconv.ParseInt(v, 2, 64)
		nums = append(nums, int(num))
	}

	return nums, bits
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 03: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 03: Part 2: = %+v\n", part2Solution)
}
