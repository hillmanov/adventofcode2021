package main

import (
	. "adventofcode/utils"
	"embed"
	"fmt"
)

//go:embed input.txt
var f embed.FS

type Cuboid struct {
	Index int
	State string
	MinX  int
	MaxX  int
	MinY  int
	MaxY  int
	MinZ  int
	MaxZ  int
}

func (c Cuboid) Volume() int {
	return (c.MaxX - c.MinX) * (c.MaxY - c.MinY) * (c.MaxZ - c.MinZ)
}

func (A Cuboid) Intersects(B Cuboid) bool {
	return (A.MaxX > B.MinX) &&
		(A.MinX < B.MaxX) &&
		(A.MaxY > B.MinY) &&
		(A.MinY < B.MaxY) &&
		(A.MaxZ > B.MinZ) &&
		(A.MinZ < B.MaxZ)
}

func (A Cuboid) Split(B Cuboid) []Cuboid {
	splits := []Cuboid{}

	// X
	if A.MinX < B.MinX {
		splits = append(splits, Cuboid{
			Index: A.Index,
			State: "on",
			MinX:  A.MinX,
			MaxX:  B.MinX,
			MinY:  A.MinY,
			MaxY:  A.MaxY,
			MinZ:  A.MinZ,
			MaxZ:  A.MaxZ,
		})
	}

	if A.MaxX > B.MaxX {
		splits = append(splits, Cuboid{
			Index: A.Index,
			State: "on",
			MinX:  B.MaxX,
			MaxX:  A.MaxX,
			MinY:  A.MinY,
			MaxY:  A.MaxY,
			MinZ:  A.MinZ,
			MaxZ:  A.MaxZ,
		})
	}

	// Y
	if A.MinY < B.MinY {
		splits = append(splits, Cuboid{
			Index: A.Index,
			State: "on",
			MinX:  MaxInt(A.MinX, B.MinX),
			MaxX:  MinInt(A.MaxX, B.MaxX),
			MinY:  A.MinY,
			MaxY:  B.MinY,
			MinZ:  A.MinZ,
			MaxZ:  A.MaxZ,
		})
	}

	if A.MaxY > B.MaxY {
		splits = append(splits, Cuboid{
			Index: A.Index,
			State: "on",
			MinX:  MaxInt(A.MinX, B.MinX),
			MaxX:  MinInt(A.MaxX, B.MaxX),
			MinY:  B.MaxY,
			MaxY:  A.MaxY,
			MinZ:  A.MinZ,
			MaxZ:  A.MaxZ,
		})
	}

	// Z
	if A.MinZ < B.MinZ {
		splits = append(splits, Cuboid{
			Index: A.Index,
			State: "on",
			MinX:  MaxInt(A.MinX, B.MinX),
			MaxX:  MinInt(A.MaxX, B.MaxX),
			MinY:  MaxInt(A.MinY, B.MinY),
			MaxY:  MinInt(A.MaxY, B.MaxY),
			MinZ:  A.MinZ,
			MaxZ:  B.MinZ,
		})
	}

	if A.MaxZ > B.MaxZ {
		splits = append(splits, Cuboid{
			Index: A.Index,
			State: "on",
			MinX:  MaxInt(A.MinX, B.MinX),
			MaxX:  MinInt(A.MaxX, B.MaxX),
			MinY:  MaxInt(A.MinY, B.MinY),
			MaxY:  MinInt(A.MaxY, B.MaxY),
			MinZ:  B.MaxZ,
			MaxZ:  A.MaxZ,
		})
	}

	return splits
}

func Part1() Any {
	originalCuboidsToProcess := getInput()
	cuboidsToProcess := []Cuboid{}

	initializationCube := Cuboid{
		MinX: -50,
		MaxX: 50,
		MinY: -50,
		MaxY: 50,
		MinZ: -50,
		MaxZ: 50,
	}

	for _, cuboid := range originalCuboidsToProcess {
		if cuboid.Intersects(initializationCube) {
			cuboidsToProcess = append(cuboidsToProcess, cuboid)
		}
	}

	return process(cuboidsToProcess)
}

func Part2() Any {
	cuboidsToProcess := getInput()
	return process(cuboidsToProcess)
}

func process(cuboidsToProcess []Cuboid) int {
	settledCuboids := []Cuboid{}

	for pI := 0; pI < len(cuboidsToProcess); pI++ {
		for sI := 0; sI < len(settledCuboids); sI++ {
			if !cuboidsToProcess[pI].Intersects(settledCuboids[sI]) {
				continue
			}

			settledCuboids[sI].State = "remove"
			settledCuboids = append(settledCuboids, settledCuboids[sI].Split(cuboidsToProcess[pI])...)
		}

		if cuboidsToProcess[pI].State == "on" {
			settledCuboids = append(settledCuboids, cuboidsToProcess[pI])
		}

		refresh := []Cuboid{}
		for _, cuboid := range settledCuboids {
			if cuboid.State != "remove" {
				refresh = append(refresh, cuboid)
			}
		}
		settledCuboids = refresh
		// BlenderDump(settledCuboids)
	}

	countOn := 0
	for _, c := range settledCuboids {
		countOn += c.Volume()
	}
	return countOn

}

func getInput() []Cuboid {
	lines, _ := ReadLines(f, "input.txt")

	cuboids := []Cuboid{}

	for i, line := range lines {
		cuboid := Cuboid{Index: i}
		fmt.Sscanf(line, "%s x=%d..%d,y=%d..%d,z=%d..%d", &cuboid.State, &cuboid.MinX, &cuboid.MaxX, &cuboid.MinY, &cuboid.MaxY, &cuboid.MinZ, &cuboid.MaxZ)
		cuboid.MaxX += 1
		cuboid.MaxY += 1
		cuboid.MaxZ += 1
		cuboids = append(cuboids, cuboid)
	}

	return cuboids
}

func main() {
	// fmt.Println("import bpy")
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 22: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 22: Part 2: = %+v\n", part2Solution)
}
