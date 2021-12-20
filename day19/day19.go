package main

import (
	. "adventofcode/utils"
	"embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var f embed.FS

type Scanner struct {
	Label   int
	X       int
	Y       int
	Z       int
	Beacons []Beacon
}

type Beacon struct {
	X int
	Y int
	Z int
}

func (b Beacon) getOrientations() []Beacon {
	dX := []int{1, -1, 1, 1}
	dY := []int{1, 1, -1, 1}
	dZ := []int{1, 1, 1, -1}

	rotations := []func(b Beacon) (int, int, int){
		func(b Beacon) (x, y, z int) {
			return b.X, b.Y, b.Z
		},

		func(b Beacon) (x, y, z int) {
			return b.X, b.Z, b.Y
		},

		func(b Beacon) (x, y, z int) {
			return b.Y, b.X, b.Z
		},

		func(b Beacon) (x, y, z int) {
			return b.Y, b.Z, b.X
		},

		func(b Beacon) (x, y, z int) {
			return b.Z, b.X, b.Y
		},

		func(b Beacon) (x, y, z int) {
			return b.Z, b.Y, b.X
		},
	}

	orientations := []Beacon{}
	for _, r := range rotations {
		rX, rY, rZ := r(b)

		for i := range dX {
			orientations = append(orientations, Beacon{
				X: rX * dX[i],
				Y: rY * dY[i],
				Z: rZ * dZ[i],
			})
		}
	}

	return orientations
}

func Part1() Any {
	scanners := getInput()
	s := scanners[0]

	o := s.Beacons[0].getOrientations()

	for _, oo := range o {
		fmt.Printf("oo = %+v\n", oo)
	}

	fmt.Printf("o = %+v\n", len(o))

	return nil
}

func Part2() Any {
	return nil
}

func getInput() []Scanner {
	lines, _ := ReadLines(f, "input.txt")

	scanners := []Scanner{}

	var currentScanner Scanner
	var beacon Beacon
	for _, line := range lines {
		switch {

		case strings.HasPrefix(line, "--- scanner"):
			currentScanner = Scanner{
				Beacons: []Beacon{},
			}
			fmt.Sscanf(line, "--- scanner %d ---", &currentScanner.Label)

		case len(line) == 0:
			scanners = append(scanners, currentScanner)

		default:
			beacon = Beacon{}
			fmt.Sscanf(line, "%d,%d,%d", &beacon.X, &beacon.Y, &beacon.Z)
			currentScanner.Beacons = append(currentScanner.Beacons, beacon)
		}
	}
	currentScanner.Beacons = append(currentScanner.Beacons, beacon)
	scanners = append(scanners, currentScanner)

	return scanners
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 19: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 19: Part 2: = %+v\n", part2Solution)
}
