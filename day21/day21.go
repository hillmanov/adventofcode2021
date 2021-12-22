package main

import (
	. "adventofcode/utils"
	"embed"
	"fmt"
)

//go:embed input.txt
var f embed.FS

type Dice int

func (d *Dice) Roll(times int) int {
	sum := int(*d)
	for i := 0; i < times-1; i++ {
		(*d)++
		if *d > 100 {
			*d = 1
		}
		sum += int(*d)
	}
	(*d)++
	if *d > 100 {
		*d = 1
	}
	return sum
}

type PlayerPosition int

func (pp *PlayerPosition) Move(times int) int {
	for i := 0; i < times; i++ {
		*pp++
		if *pp > 10 {
			*pp = 1
		}
	}
	return int(*pp)
}

func Part1() Any {
	player1Position, player2Position := getInput()
	player1Score, player2Score := 0, 0

	diceRolls := 0
	dice := Dice(1)
	for {
		// Player 1
		move1 := dice.Roll(3)
		diceRolls += 3
		player1Position.Move(move1)
		player1Score += int(player1Position)

		if player1Score >= 1000 {
			return player2Score * diceRolls
		}

		// Player 2
		move2 := dice.Roll(3)
		diceRolls += 3
		player2Position.Move(move2)
		player2Score += int(player2Position)

		if player2Score >= 1000 {
			return player1Score * diceRolls
		}
	}
}

func Part2() Any {
	// From each position on the board, find out the number of turns it take to get to 21 from any given score that is before 21

	for position := 1; position <= 10; position++ {
		for startingScore := 1; startingScore <= 20; startingScore++ {
			for roll := range []int{1, 2, 3} {
				fmt.Printf("roll = %+v\n", roll)
			}
		}
	}
	return nil
}

func increment(currentValue *int, max int) int {
	*currentValue = *currentValue + 1
	if *currentValue > max {
		*currentValue = 1
	}
	return *currentValue
}

func getInput() (player1Start, player2Start PlayerPosition) {
	lines, _ := ReadLines(f, "input.txt")

	fmt.Sscanf(lines[0], "Player 1 starting position: %d", &player1Start)
	fmt.Sscanf(lines[1], "Player 2 starting position: %d", &player2Start)

	return player1Start, player2Start
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 21: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 21: Part 2: = %+v\n", part2Solution)
}
