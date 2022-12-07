package main

import (
	. "adventofcode/utils"
	"embed"
	"fmt"
)

//go:embed input.txt
var f embed.FS

func Part1() Any {
	return nil
}

func Part2() Any {
	return nil
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 23: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 23: Part 2: = %+v\n", part2Solution)
}

// What I need to do is make the entire WORLD a Pather. Not each node. I need to be able to calculate the cost of every possible movement at any step.

// When finding paths, keep the following rules in mind!
// Amphipods will never stop on the space immediately outside any room.
// They can move into that space so long as they immediately continue moving.
// (Specifically, this refers to the four open spaces in the hallway that are directly above an amphipod starting position.)

// Amphipods will never move from the hallway into a room unless that room is their destination room and that room contains no amphipods which do not also have that room as their own destination.
// If an amphipod's starting room is not its destination room, it can stay in that room until it leaves the room.
// (For example, an Amber amphipod will not move from the hallway into the right three rooms, and will only move into the leftmost room if that room is empty or if it only contains other Amber amphipods.)

// Once an amphipod stops moving in the hallway, it will stay in that spot until it can move into a room.
// (That is, once any amphipod starts moving, any other amphipods currently in the hallway are locked in place and will not move again until they can move fully into a room.)

// Do I want to try to represent the world like this:
// AABBCCDD...........
// And do the math? Would be a lot faster...
// Could create quick references to specific positions:
/*
0 - HomeA1
1 - homeA2
2 - homeB1
3 - homeB2
4 - ...
5
6
7
8 - hallway1
9 - entryway1
10
*/
