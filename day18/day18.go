package main

import (
	. "adventofcode/utils"
	"embed"
	"fmt"
)

//go:embed input.txt
var f embed.FS

type Expression []interface{}

func (e *Expression) Reduce() {

}

func (e *Expression) Explode(index int) {

}

func (e *Expression) Split(index int) {

}

func (e *Expression) AddToIntLeftOfIndex(index int, value int) {
	i := e.IndexOfIntLeftOfIndex(index)
	v, ok := (*e)[i].(int)
	if !ok {
		panic("not an int at index " + string(i))
	}
	v += value
	(*e)[i] = v
}

func (e *Expression) AddToIntRightOfIndex(index int, value int) {
	i := e.IndexOfIntRightOfIndex(index)
	v, ok := (*e)[i].(int)
	if !ok {
		panic("not an int at index " + string(i))
	}
	v += value
	(*e)[i] = v
}

func (e *Expression) IndexOfIntRightOfIndex(index int) int {
	for i := index + 1; i < len(*e); i++ {
		if _, ok := (*e)[i].(int); ok {
			return i
		}
	}
	return -1
}

func (e *Expression) IndexOfIntLeftOfIndex(index int) int {
	for i := index - 1; i >= 0; i-- {
		if _, ok := (*e)[i].(int); ok {
			return i
		}
	}
	return -1
}

func (e *Expression) IndexOfClosingBracketRightOfIndex(index int) int {
	for i := index; i < len(*e); i++ {
		if c, ok := (*e)[i].(string); ok && c == "]" {
			return i
		}
	}
	return -1
}

func (e *Expression) IndexOfOpeningBracketLeftOfIndex(index int) int {
	for i := index; i >= 0; i-- {
		if c, ok := (*e)[i].(string); ok && c == "]" {
			return i
		}
	}
	return -1
}

func (e *Expression) ClearSurroundingBracketsAtIndex(index int) {
	lIndex, rIndex := e.IndexOfOpeningBracketLeftOfIndex(index), e.IndexOfClosingBracketRightOfIndex(index)
	if lIndex > 0 {
		(*e)[lIndex] = ""
	}
	if rIndex > 0 {
		(*e)[rIndex] = ""
	}
}

func Part1() Any {
	expressions := getInput()
	fmt.Printf("expressions = %+v\n", expressions)

	e := expressions[0]
	i := e.IndexOfIntRightOfIndex(0)

	e.AddToIntRightOfIndex(i, 5)

	fmt.Printf("e = %+v\n", e)

	return nil
}

func Part2() Any {
	return nil
}

func getInput() []Expression {
	lines, _ := ReadLines(f, "input.txt")

	expressions := []Expression{}
	for _, line := range lines {
		expression := Expression{}
		for _, c := range line {
			switch string(c) {
			case "[", "]", ",":
				expression = append(expression, string(c))
			case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
				expression = append(expression, ParseInt(string(c)))
			}
		}
		expressions = append(expressions, expression)
	}

	return expressions
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 18: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 18: Part 2: = %+v\n", part2Solution)
}
