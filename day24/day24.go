package main

import (
	. "adventofcode/utils"
	"embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var f embed.FS

func Part1() Any {
	program := getInput()

	r := process([]int{10}, program)
	fmt.Printf("r = %+v\n", r)

	// for i := 99999999999999; i > 11111111111111; i-- {
	// 	if containsZero(i) {
	// 		continue
	// 	}
	// 	input := intToSlice(i, nil)
	// 	r := process(input, program)
	// 	fmt.Printf("r = %+v\n", r)
	// 	if r["z"] == 0 {
	// 		fmt.Println("DONE", i)
	// 		return i
	// 	}
	// }
	return -1
}

func Part2() Any {
	return nil
}

func process(input []int, program []string) map[string]int {

	fmt.Printf("input = %+v\n", input)
	fmt.Printf("len(input) = %+v\n", len(input))

	vars := make(map[string]int)

	for _, line := range program {
		parts := strings.Split(line, " ")

		instruction := string(parts[0])
		variable := string(parts[1])

		switch instruction {
		case "inp":
			vars[variable] = input[0]
			input = input[1:]
		case "add":
			arg := string(parts[2])
			switch kind, val := varOrInt(arg); {
			case kind == "string":
				vars[variable] = vars[variable] + vars[val.(string)]
			case kind == "int":
				vars[variable] = vars[variable] + val.(int)
			}
		case "mul":
			arg := string(parts[2])
			switch kind, val := varOrInt(arg); {
			case kind == "string":
				vars[variable] = vars[variable] * vars[val.(string)]
			case kind == "int":
				vars[variable] = vars[variable] * val.(int)
			}
		case "div":
			arg := string(parts[2])
			switch kind, val := varOrInt(arg); {
			case kind == "string":
				vars[variable] = vars[variable] / vars[val.(string)]
			case kind == "int":
				vars[variable] = vars[variable] / val.(int)
			}
		case "mod":
			arg := string(parts[2])
			switch kind, val := varOrInt(arg); {
			case kind == "string":
				vars[variable] = vars[variable] % vars[val.(string)]
			case kind == "int":
				vars[variable] = vars[variable] % val.(int)
			}
		case "eql":
			arg := string(parts[2])
			switch kind, val := varOrInt(arg); {
			case kind == "string":
				if vars[variable] == vars[val.(string)] {
					vars[variable] = 1
				} else {
					vars[variable] = 0
				}
			}
		}
	}

	return vars
}

func getInput() []string {
	lines, _ := ReadLines(f, "input.txt")
	return lines
}

func intToSlice(num int, sequence []int) []int {
	if sequence == nil {
		sequence = []int{}
	}
	if num != 0 {
		i := num % 10
		sequence = append([]int{i}, sequence...)
		return intToSlice(num/10, sequence)
	}
	return sequence
}

func containsZero(num int) bool {
	if num == 0 {
		return true
	}

	for num > 0 {
		if num%10 == 0 {
			return true
		}
		num /= 10
	}
	return false
}

func varOrInt(v string) (string, interface{}) {
	if intVal, err := strconv.Atoi(v); err == nil {
		return "int", intVal
	}
	return "string", v
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 24: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 24: Part 2: = %+v\n", part2Solution)
}
