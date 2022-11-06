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
	Label           int
	X               int
	Y               int
	Z               int
	Beacons         []Beacon
	BeaconRotations [][]Beacon
	RotationVectors [][]Vector
	Rotation        int
}

type Beacon struct {
	X int
	Y int
	Z int
}

type Vector struct {
	Value   float64
	BeaconA Beacon
	BeaconB Beacon
}

func (s *Scanner) InitRotations() {
	s.BeaconRotations = make([][]Beacon, 24)
	for i := 0; i < 24; i++ {
		s.BeaconRotations[i] = make([]Beacon, len(s.Beacons))
		for j, b := range s.Beacons {
			s.BeaconRotations[i][j] = b.Rotate(i)
		}
	}
}

func (s *Scanner) InitVectors() {
	s.RotationVectors = make([][]Vector, 24)
	for rotation := range s.BeaconRotations {
		s.RotationVectors[rotation] = make([]Vector, len(s.Beacons)*len(s.Beacons))
		for i := range s.BeaconRotations[rotation] {
			for j := range s.BeaconRotations[rotation] {
				if i == j {
					continue
				}

				s.RotationVectors[rotation] = append(s.RotationVectors[rotation],
					Vector{
						Value: math.Sqrt(
							math.Pow(float64(s.BeaconRotations[rotation][i].X)-float64(s.BeaconRotations[rotation][j].X), 2) +
								math.Pow(float64(s.BeaconRotations[rotation][i].Y)-float64(s.BeaconRotations[rotation][j].Y), 2) +
								math.Pow(float64(s.BeaconRotations[rotation][i].Z)-float64(s.BeaconRotations[rotation][j].Z), 2),
						),
						BeaconA: s.BeaconRotations[rotation][i],
						BeaconB: s.BeaconRotations[rotation][j],
					},
				)
			}
		}
	}
}

func (s *Scanner) IsSettled() bool {
	return s.Rotation >= 0
}

func (s *Scanner) SettledVectors() []Vector {
	return s.RotationVectors[s.Rotation]
}

func (s *Scanner) SettledBeacons() []Beacon {
	return s.BeaconRotations[s.Rotation]
}

func (b Beacon) Rotate(rotation int) Beacon {
	switch rotation {
	case 0:
		return Beacon{X: b.X, Y: b.Y, Z: b.Z}
	case 1:
		return Beacon{X: -b.X, Y: b.Y, Z: -b.Z}
	case 2:
		return Beacon{X: b.Y, Y: -b.X, Z: b.Z}
	case 3:
		return Beacon{X: -b.Y, Y: b.X, Z: b.Z}
	case 4:
		return Beacon{X: b.Z, Y: b.Y, Z: -b.X}
	case 5:
		return Beacon{X: -b.Z, Y: b.Y, Z: b.X}
	case 6:
		return Beacon{X: b.X, Y: -b.Z, Z: b.Y}
	case 7:
		return Beacon{X: -b.X, Y: b.Z, Z: b.Y}
	case 8:
		return Beacon{X: b.Y, Y: -b.Z, Z: -b.X}
	case 9:
		return Beacon{X: -b.Y, Y: -b.Z, Z: b.X}
	case 10:
		return Beacon{X: b.Z, Y: b.X, Z: b.Y}
	case 11:
		return Beacon{X: -b.Z, Y: -b.X, Z: b.Y}
	case 12:
		return Beacon{X: b.X, Y: -b.Y, Z: -b.Z}
	case 13:
		return Beacon{X: -b.X, Y: -b.Y, Z: b.Z}
	case 14:
		return Beacon{X: b.Y, Y: b.X, Z: -b.Z}
	case 15:
		return Beacon{X: -b.Y, Y: -b.X, Z: -b.Z}
	case 16:
		return Beacon{X: b.Z, Y: -b.Y, Z: b.X}
	case 17:
		return Beacon{X: -b.Z, Y: -b.Y, Z: -b.X}
	case 18:
		return Beacon{X: b.X, Y: b.Z, Z: -b.Y}
	case 19:
		return Beacon{X: -b.X, Y: -b.Z, Z: -b.Y}
	case 20:
		return Beacon{X: b.Y, Y: b.Z, Z: b.X}
	case 21:
		return Beacon{X: -b.Y, Y: b.Z, Z: -b.X}
	case 22:
		return Beacon{X: b.Z, Y: -b.X, Z: -b.Y}
	case 23:
		return Beacon{X: -b.Z, Y: b.X, Z: -b.Y}
	default:
		panic("Unsupported rotate value")
	}
}

func Part1() Any {
	scanners := getInput()

	for i := range scanners {
		scanners[i].InitRotations()
		scanners[i].InitVectors()
	}

	homeScanner := scanners[0]
	homeScanner.Rotation = 0

	for i := range scanners {
		scanner := scanners[i]
		if scanner.IsSettled() || i == 0 {
			continue
		}

		for rotationVectorsIndex := range scanner.RotationVectors {
			if len((intersection(homeScanner.SettledVectors(), scanner.RotationVectors[rotationVectorsIndex]))) >= 12 {
				scanner.Rotation = rotationVectorsIndex
				fmt.Println("Oh baby!")
				fmt.Printf("i = %+v\n", i)
				fmt.Printf("rotationVectorsIndex = %+v\n", rotationVectorsIndex)
				break
			}
		}

	}

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
	for _, line := range lines {
		switch {

		case strings.HasPrefix(line, "--- scanner"):
			currentScanner = Scanner{
				Beacons:  []Beacon{},
				Rotation: -1,
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

// Go through the 24 rotations
// Go through each scanner.
// If rotation has NOT been found, need to find it's rotation that matches another scanner that HAS been found.
// If it founds a match of twelve, need to store which rotation matched up (0-23) (you can now delete all other rotations if you want...)
// GO through and recalculatue the vectors on the scanne you just found using the translation (any match can give the translation) just on the rotation that was found
// Find all scanners for the current scanner that match. If there are no matches, then go to the next scanner that has a rotation.
// Once all scanner have a rotation:
// Once you do that, you can go through the rotated and tranlated beacons, using the x,y,z values as the unique identifier to find all the unique beacons.
