package main

import (
	. "adventofcode/utils"
	"embed"
	"fmt"
	"regexp"
	"strings"
)

//go:embed input.txt
var f embed.FS

const (
	boardSize = 5
)

type square struct {
	Num    int
	Marked bool
}

type board struct {
	Squares [boardSize][boardSize]*square
}

func (b *board) MarkNum(n int) {
	for i := range b.Squares {
		for j := range b.Squares[i] {
			if b.Squares[i][j].Num == n {
				b.Squares[i][j].Marked = true
			}
		}
	}
}

func (b *board) HasWin() bool {
	rowWin := true
	colWin := true

	for i := range b.Squares {
		for j := range b.Squares[i] {
			rowWin = rowWin && b.Squares[i][j].Marked
			colWin = colWin && b.Squares[j][i].Marked
		}
		if rowWin || colWin {
			return true
		}
		rowWin = true
		colWin = true
	}

	return false
}

func (b *board) Score() int {
	score := 0

	for i := range b.Squares {
		for j := range b.Squares[i] {
			if !b.Squares[i][j].Marked {
				score = score + b.Squares[i][j].Num
			}
		}
	}
	return score
}

func Part1() Any {
	callNums, boards := getInput()

	for _, callNum := range callNums {
		for _, board := range boards {
			board.MarkNum(callNum)
			if board.HasWin() {
				return callNum * board.Score()
			}
		}
	}

	return nil
}

func Part2() Any {
	callNums, boards := getInput()
	boardsToSkip := make(map[int]bool)

	for _, callNum := range callNums {
		for boardIndex, board := range boards {
			if boardsToSkip[boardIndex] {
				continue
			}
			board.MarkNum(callNum)
			if board.HasWin() {
				boardsToSkip[boardIndex] = true
			}

			if len(boardsToSkip) == len(boards) {
				return callNum * board.Score()
			}
		}
	}

	return nil
}

func getInput() ([]int, []board) {
	lines, _ := ReadLines(f, "input.txt")

	var callNums []int
	for _, val := range strings.Split(lines[0], ",") {
		callNums = append(callNums, ParseInt(val))
	}

	blockLineSplitter := regexp.MustCompile("\\d+")

	var boards []board
	for _, block := range strings.Split(strings.Join(lines[1:], "\n"), "\n\n") {
		var b board
		for i, blockLine := range strings.Split(strings.Trim(block, "\n"), "\n") {
			for j, num := range blockLineSplitter.FindAllStringSubmatch(blockLine, -1) {
				b.Squares[i][j] = &square{Num: ParseInt(num[0])}
			}
		}
		boards = append(boards, b)
	}

	return callNums, boards
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 04: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 04: Part 2: = %+v\n", part2Solution)
}
