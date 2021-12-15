package main

import (
	. "adventofcode/utils"
	"embed"
	"fmt"
)

type command struct {
	Direction string
	Units     int
}

//go:embed input.txt
var f embed.FS

func Part1() Any {
	commands := getInput()

	horizontalPosition := 0
	depth := 0

	for _, command := range commands {
		switch command.Direction {
		case "forward":
			horizontalPosition += command.Units
		case "down":
			depth += command.Units
		case "up":
			depth -= command.Units
		default:
			panic("Unsupport direction")
		}
	}

	return horizontalPosition * depth
}

func Part2() Any {
	commands := getInput()

	aim := 0
	horizontalPosition := 0
	depth := 0

	for _, command := range commands {
		switch command.Direction {
		case "forward":
			horizontalPosition += command.Units
			depth += aim * command.Units
		case "down":
			aim += command.Units
		case "up":
			aim -= command.Units
		default:
			panic("Unsupport direction")
		}
	}

	return horizontalPosition * depth
}

func getInput() []command {
	lines, err := ReadLines(f, "input.txt")
	if err != nil {
		panic(err)
	}

	commands := []command{}
	for _, line := range lines {
		c := command{}
		fmt.Sscanf(line, "%s %d", &c.Direction, &c.Units)
		commands = append(commands, c)
	}

	return commands
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 02: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 02: Part 2: = %+v\n", part2Solution)
}
