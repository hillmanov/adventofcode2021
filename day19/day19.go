package main

import (
	. "adventofcode/utils"
	"embed"
	"fmt"
	"math"
	"strings"
)

//go:embed input.txt
var f embed.FS

type Scanner struct {
	Label    int
	Rotation int
	X        int
	Y        int
	Z        int
	Beacons  []Beacon
}

type Beacon struct {
	ID           string
	X            int
	Y            int
	Z            int
	Orientations []Beacon
}

type Vector struct {
	Value   float64
	BeaconA Beacon
	BeaconB Beacon
}

func (b *Beacon) getOrientations() []Beacon {
	if b.Orientations != nil {
		return b.Orientations
	}
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
		rX, rY, rZ := r(*b)

		for i := range dX {
			orientations = append(orientations, Beacon{
				ID: b.ID,
				X:  rX * dX[i],
				Y:  rY * dY[i],
				Z:  rZ * dZ[i],
			})
		}
	}

	b.Orientations = orientations

	return b.Orientations
}

func (s Scanner) Vectors() []Vector {
	vectors := []Vector{}

	for i := 0; i < len(s.Beacons)-1; i++ {
		for j := 0; j < len(s.Beacons)-1; j++ {
			if i == j {
				continue
			}
			vectors = append(vectors, Vector{
				Value: math.Sqrt(
					math.Pow(float64(s.Beacons[i].X)-float64(s.Beacons[j].X), 2) +
						math.Pow(float64(s.Beacons[i].Y)-float64(s.Beacons[j].Y), 2) +
						math.Pow(float64(s.Beacons[i].Z)-float64(s.Beacons[j].Z), 2),
				),
				BeaconA: s.Beacons[i],
				BeaconB: s.Beacons[j],
			})
		}
	}

	return vectors
}

func Part1() Any {
	// scanners := getInput()
	// s := scanners[0]

	// fmt.Printf("s.Vectors() = %+v\n", s.Vectors())

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
	beaconID := 0
	for _, line := range lines {
		switch {

		case strings.HasPrefix(line, "--- scanner"):
			currentScanner = Scanner{
				Rotation: -1, // NO rotation has been set yet
				Beacons:  []Beacon{},
			}
			fmt.Sscanf(line, "--- scanner %d ---", &currentScanner.Label)

		case len(line) == 0:
			scanners = append(scanners, currentScanner)

		default:
			beacon = Beacon{
				ID: fmt.Sprintf("%03d", beaconID),
			}
			fmt.Sscanf(line, "%d,%d,%d", &beacon.X, &beacon.Y, &beacon.Z)
			currentScanner.Beacons = append(currentScanner.Beacons, beacon)
			beaconID++
		}
	}
	currentScanner.Beacons = append(currentScanner.Beacons, beacon)
	scanners = append(scanners, currentScanner)

	return scanners
}

func intersection(origin, target []Vector) [][2]Vector {
	intersecting := [][2]Vector{}
	for _, v1 := range origin {
		for _, v2 := range target {
			if v1.Value == v2.Value {
				intersecting = append(intersecting, [2]Vector{v1, v2})
			}
		}
	}

	return intersecting
}

func getOffset(v [][2]Vector) (x, y, z int) {
	return v[0][0].BeaconA.X - v[0][1].BeaconA.X,
		v[0][0].BeaconA.Y - v[0][1].BeaconA.Y,
		v[0][0].BeaconA.Z - v[0][1].BeaconA.Z
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 19: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 19: Part 2: = %+v\n", part2Solution)
}
