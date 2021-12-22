package main

import (
	. "adventofcode/utils"
	"embed"
	"fmt"
)

//go:embed input.txt
var f embed.FS

type Step struct {
	State  string
	StartX int
	EndX   int
	StartY int
	EndY   int
	StartZ int
	EndZ   int
}

type Cube struct {
	X int
	Y int
	Z int
}

func (s Step) VolumeOfCuboid() int {
	return (s.EndX - s.StartX) * (s.EndY - s.StartY) * (s.EndZ - s.StartZ)
}

func (A Step) IntersectionVolume(B Step) int {
	return MaxInt(MinInt(B.EndX, A.EndX)-MaxInt(B.StartX, A.StartX), 0) *
		MaxInt(MinInt(B.EndY, A.EndY)-MaxInt(B.StartY, A.StartY), 0) *
		MaxInt(MinInt(B.EndZ, A.EndZ)-MaxInt(B.StartZ, A.StartZ), 0)
}

type State = string

type Core map[Cube]State

func Part1() Any {
	steps := getInput()

	core := Core{}

	for _, step := range steps {
		if (step.StartX < -50 && step.EndX < 50) || (step.StartY < -50 && step.EndY < 50) || (step.StartZ < -50 && step.EndZ < 50) {
			continue
		}

		if (step.StartX > 50 && step.EndX > 50) || (step.StartY > -50 && step.EndY > 50) || (step.StartZ > -50 && step.EndZ > 50) {
			continue
		}

		for x := step.StartX; x <= step.EndX; x++ {
			for y := step.StartY; y <= step.EndY; y++ {
				for z := step.StartZ; z <= step.EndZ; z++ {
					core[Cube{x, y, z}] = step.State
				}
			}
		}
	}

	countOn := 0
	for x := -50; x <= 50; x++ {
		for y := -50; y <= 50; y++ {
			for z := -50; z <= 50; z++ {
				if v := core[Cube{x, y, z}]; v == "on" {
					countOn++
				}
			}
		}
	}

	return countOn
}

func Part2() Any {
	step := Step{
		StartX: 0,
		EndX:   5,
		StartY: 0,
		EndY:   5,
		StartZ: 0,
		EndZ:   5,
	}

	return step.VolumeOfCuboid()
}

func getInput() []Step {
	lines, _ := ReadLines(f, "input.txt")

	steps := []Step{}

	for _, line := range lines {
		step := Step{}
		fmt.Sscanf(line, "%s x=%d..%d,y=%d..%d,z=%d..%d", &step.State, &step.StartX, &step.EndX, &step.StartY, &step.EndY, &step.StartZ, &step.EndZ)
		steps = append(steps, step)
	}

	return steps
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 22: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 22: Part 2: = %+v\n", part2Solution)
}
