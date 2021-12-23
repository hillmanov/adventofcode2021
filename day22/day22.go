package main

import (
	. "adventofcode/utils"
	"embed"
	"fmt"
)

//go:embed input.txt
var f embed.FS

type Cuboid struct {
	State        string
	MinX         int
	MaxX         int
	MinY         int
	MaxY         int
	MinZ         int
	MaxZ         int
	Intersectors []Cuboid
	Absorbed     bool
}

type Voxel struct {
	X int
	Y int
	Z int
}

func (c Cuboid) Volume() int {
	return (c.MaxX - c.MinX) * (c.MaxY - c.MinY) * (c.MaxZ - c.MinZ)
}

func (c Cuboid) IsNil() bool {
	return c.Volume() == 0
}

func (A Cuboid) Intersects(B Cuboid) bool {
	return (MinInt(B.MaxX, A.MaxX)-MaxInt(B.MinX, A.MinX) > 0) &&
		(MinInt(B.MaxY, A.MaxY)-MaxInt(B.MinY, A.MinY) > 0) &&
		(MinInt(B.MaxZ, A.MaxZ)-MaxInt(B.MinZ, A.MinZ) > 0)
}

func (A Cuboid) Intersection(B Cuboid) Cuboid {
	if A.Intersects(B) {
		return Cuboid{
			MinX: MaxInt(B.MinX, A.MinX),
			MaxX: MinInt(B.MaxX, A.MaxX),
			MinY: MaxInt(B.MinY, A.MinY),
			MaxY: MinInt(B.MinY, A.MinY),
			MinZ: MaxInt(B.MinZ, A.MinZ),
			MaxZ: MinInt(B.MinZ, A.MinZ),
		}
	}

	return Cuboid{}
}

func (A Cuboid) Split(B Cuboid) []Cuboid {
	splits := []Cuboid{}

	intersection := A.Intersection(B)

	// Check each plane to see if what splits we need

	return splits
}

type State = string

type Core map[Voxel]State

func Part1() Any {
	steps := getInput()

	core := Core{}

	for _, step := range steps {
		if (step.MinX < -50 && step.MaxX < 50) || (step.MinY < -50 && step.MaxY < 50) || (step.MinZ < -50 && step.MaxZ < 50) {
			continue
		}

		if (step.MinX > 50 && step.MaxX > 50) || (step.MinY > -50 && step.MaxY > 50) || (step.MinZ > -50 && step.MaxZ > 50) {
			continue
		}

		for x := step.MinX; x <= step.MaxX; x++ {
			for y := step.MinY; y <= step.MaxY; y++ {
				for z := step.MinZ; z <= step.MaxZ; z++ {
					core[Voxel{x, y, z}] = step.State
				}
			}
		}
	}

	countOn := 0
	for x := -50; x <= 50; x++ {
		for y := -50; y <= 50; y++ {
			for z := -50; z <= 50; z++ {
				if v := core[Voxel{x, y, z}]; v == "on" {
					countOn++
				}
			}
		}
	}

	return countOn
}

func Part2() Any {
	cuboids := getInput()

	// For each cuboid in the list:
	// Check if there is an intersection:
	// If both boxes are ON
	// aCuboid stays the SAME
	// bCuboid SPLITS into up to 3 different sub-cuiboids
	//
	// If aCuboid is ON and bCuboid is OFF
	// aCuboid splits into 0-3 sub cubes

	return nil
}

func getInput() []Cuboid {
	lines, _ := ReadLines(f, "input.txt")

	steps := []Cuboid{}

	for _, line := range lines {
		step := Cuboid{
			Intersectors: []Cuboid{},
		}
		fmt.Sscanf(line, "%s x=%d..%d,y=%d..%d,z=%d..%d", &step.State, &step.MinX, &step.MaxX, &step.MinY, &step.MaxY, &step.MinZ, &step.MaxZ)
		steps = append(steps, step)
	}

	return steps
}

func main() {
	// part1Solution := Part1()
	part2Solution := Part2()

	// fmt.Printf("Day 22: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 22: Part 2: = %+v\n", part2Solution)
}
