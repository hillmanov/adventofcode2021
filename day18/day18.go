package main

import (
	. "adventofcode/utils"
	"embed"
	"encoding/json"
	"fmt"
	"math"
)

//go:embed input.txt
var f embed.FS

type Expression []interface{}

func (e *Expression) Reduce() *Expression {
	nested := 0
	performedAction := false

reducer:
	for {
		nested = 0
		performedAction = false

		for i := 0; i < len(*e); i++ {
			switch {
			case nested >= 4 && e.IsMatchAtIndex(i, "["):
				e.Explode(i + 1) // This is the int inside the bracket
				performedAction = true
				goto reducer
			case e.IsMatchAtIndex(i, "["):
				nested++
			case e.IsMatchAtIndex(i, "]"):
				nested--
			}
		}

		for i := 0; i < len(*e); i++ {
			switch {
			case e.IsIntAtIndex(i) && e.IntValueAtIndex(i) >= 10:
				e.Split(i)
				performedAction = true
				goto reducer
			}
		}

		if !performedAction {
			break
		}
	}
	return e
}

func (e *Expression) Add(addE Expression) *Expression {
	newE := Expression{}
	newE = append(newE, "[")
	newE = append(newE, (*e)...)
	newE = append(newE, ",")
	newE = append(newE, addE...)
	newE = append(newE, "]")
	(*e) = newE
	return e
}

func (e *Expression) Explode(indexOfLeftValue int) {
	indexOfRightValue := e.IndexOfIntRightOfIndex(indexOfLeftValue)

	leftValue := e.IntValueAtIndex(indexOfLeftValue)
	rightValue := e.IntValueAtIndex(indexOfRightValue)

	e.AddToIntLeftOfIndex(indexOfLeftValue, leftValue)
	e.AddToIntRightOfIndex(indexOfRightValue, rightValue)

	leftBraceIndex := e.IndexOfOpeningBracketLeftOfIndex(indexOfLeftValue)
	rightBraceIndex := e.IndexOfClosingBracketRightOfIndex(indexOfRightValue)

	(*e) = append((*e)[:leftBraceIndex], (*e)[rightBraceIndex:]...)
	(*e)[leftBraceIndex] = 0
}

func (e *Expression) Split(index int) {
	value := e.IntValueAtIndex(index)

	leftValue := int(math.Floor(float64(value) / 2))
	rightValue := int(math.Ceil(float64(value) / 2))

	(*e) = append((*e)[:index], append(Expression{"[", leftValue, ",", rightValue, "]"}, (*e)[index+1:]...)...)
}

func (e *Expression) AddToIntLeftOfIndex(index int, value int) {
	if i := e.IndexOfIntLeftOfIndex(index); i > 0 {
		(*e)[i] = int(value + e.IntValueAtIndex(i))
	}
}

func (e *Expression) AddToIntRightOfIndex(index int, value int) {
	if i := e.IndexOfIntRightOfIndex(index); i > 0 {
		(*e)[i] = int(value + e.IntValueAtIndex(i))
	}
}

func (e *Expression) IndexOfIntRightOfIndex(index int) int {
	for i := index + 1; i < len(*e); i++ {
		if e.IsIntAtIndex(i) {
			return i
		}
	}
	return -1
}

func (e *Expression) IndexOfIntLeftOfIndex(index int) int {
	for i := index - 1; i >= 0; i-- {
		if e.IsIntAtIndex(i) {
			return i
		}
	}
	return -1
}

func (e *Expression) IndexOfClosingBracketRightOfIndex(index int) int {
	for i := index; i < len(*e); i++ {
		if e.IsMatchAtIndex(i, "]") {
			return i
		}
	}
	return -1
}

func (e *Expression) IndexOfOpeningBracketLeftOfIndex(index int) int {
	for i := index; i >= 0; i-- {
		if e.IsMatchAtIndex(i, "[") {
			return i
		}
	}
	return -1
}

func (e *Expression) IntValueAtIndex(index int) int {
	if v, ok := (*e)[index].(int); ok {
		return v
	}
	panic("Not an int at index")
}

func (e *Expression) IsIntAtIndex(index int) bool {
	_, ok := (*e)[index].(int)
	return ok
}

func (e *Expression) IsMatchAtIndex(index int, c string) bool {
	v, ok := (*e)[index].(string)
	return ok && v == c
}

func (e *Expression) String() string {
	s := ""
	for i := 0; i < len(*e); i++ {
		s += fmt.Sprintf("%+v", (*e)[i])
	}
	return s
}

func (e Expression) Materialize() []interface{} {
	var materialized []interface{}
	json.Unmarshal([]byte(e.String()), &materialized)
	return materialized
}

func Part1() Any {
	expressions := getInput()

	runningExpression := expressions[0]
	for _, e := range expressions[1:] {
		runningExpression.Add(e)
		runningExpression.Reduce()
	}
	runningExpression.Reduce()

	mag := getMagnitude(runningExpression.Materialize())

	return mag
}

func Part2() Any {
	expressions := getInput()

	maxMag := 0
	for i := 0; i < len(expressions)-1; i++ {
		for j := i + 1; j < len(expressions)-1; j++ {
			if i == j {
				continue
			}

			// Make copies so we don't mangle everything as we reduce
			a := expressions[i][:]
			b := expressions[j][:]

			c := expressions[j][:]
			d := expressions[i][:]

			maxMag = MaxInt(
				maxMag,
				MaxInt(
					getMagnitude(a.Add(b).Reduce().Materialize()),
					getMagnitude(c.Add(d).Reduce().Materialize()),
				),
			)
		}
	}
	return maxMag
}

func getMagnitude(e interface{}) int {
	if v, ok := e.(float64); ok {
		return int(v)
	}

	ee, _ := e.([]interface{})
	return 3*getMagnitude(ee[0]) + 2*getMagnitude(ee[1])
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
