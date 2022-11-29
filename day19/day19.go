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
	Label           int
	X               int
	Y               int
	Z               int
	Beacons         []Beacon
	BeaconRotations [][]Beacon
	RotationVectors [][]DistanceVector
	Rotation        int
	Settled         bool
	Vector          Vector
}

type Vector struct {
	X int
	Y int
	Z int
}

func (v Vector) Equals(o Vector) bool {
	return v.X == o.X && v.Y == o.Y && v.Z == o.Z
}

type Beacon struct {
	X int
	Y int
	Z int
}

func (b Beacon) Vector(o Beacon) Vector {
	return Vector{
		X: b.X - o.X,
		Y: b.Y - o.Y,
		Z: b.Z - o.Z,
	}
}

func (s *Scanner) ManhattanDistance(o *Scanner) int {
	return Abs(s.Vector.X-o.Vector.X) + Abs(s.Vector.Y-o.Vector.Y) + Abs(s.Vector.Z-o.Vector.Z)
}

func (b Beacon) Equals(o Beacon) bool {
	return b.X == o.X &&
		b.Y == o.Y &&
		b.Z == o.Z
}

func (b *Beacon) ApplyOffset(v Vector) {
	b.X = b.X + v.X
	b.Y = b.Y + v.Y
	b.Z = b.Z + v.Z
}

type DistanceVector struct {
	Vector Vector
	Origin Beacon
}

func (s *Scanner) InitRotations() {
	s.BeaconRotations = make([][]Beacon, 24)
	for rotation := 0; rotation < 24; rotation++ {
		s.BeaconRotations[rotation] = make([]Beacon, len(s.Beacons))
		for i, beacon := range s.Beacons {
			s.BeaconRotations[rotation][i] = beacon.Rotate(rotation)
		}
	}
}

func (s *Scanner) InitVectors() {
	s.RotationVectors = make([][]DistanceVector, 24)
	for rotation := range s.BeaconRotations {
		s.RotationVectors[rotation] = []DistanceVector{}
		for i := range s.BeaconRotations[rotation] {
			for j := range s.BeaconRotations[rotation] {
				if i == j {
					continue
				}

				s.RotationVectors[rotation] = append(s.RotationVectors[rotation],
					DistanceVector{
						Vector: s.BeaconRotations[rotation][i].Vector(s.BeaconRotations[rotation][j]),
						Origin: s.BeaconRotations[rotation][i],
					},
				)
			}
		}
	}
}

func (s *Scanner) IsSettled() bool {
	return s.Settled
}

func (s *Scanner) SettledVectors() []DistanceVector {
	return s.RotationVectors[s.Rotation]
}

func (s *Scanner) SettledBeacons() []Beacon {
	return s.BeaconRotations[s.Rotation]
}

func (b Beacon) Rotate(rotation int) Beacon {
	switch rotation {
	case 0:
		return Beacon{b.X, b.Y, b.Z}
	case 1:
		return Beacon{b.X, -b.Z, b.Y}
	case 2:
		return Beacon{b.X, -b.Y, -b.Z}
	case 3:
		return Beacon{b.X, b.Z, -b.Y}
	case 4:
		return Beacon{-b.Y, b.X, b.Z}
	case 5:
		return Beacon{b.Z, b.X, b.Y}
	case 6:
		return Beacon{b.Y, b.X, -b.Z}
	case 7:
		return Beacon{-b.Z, b.X, -b.Y}
	case 8:
		return Beacon{-b.X, -b.Y, b.Z}
	case 9:
		return Beacon{-b.X, -b.Z, -b.Y}
	case 10:
		return Beacon{-b.X, b.Y, -b.Z}
	case 11:
		return Beacon{-b.X, b.Z, b.Y}
	case 12:
		return Beacon{b.Y, -b.X, b.Z}
	case 13:
		return Beacon{b.Z, -b.X, -b.Y}
	case 14:
		return Beacon{-b.Y, -b.X, -b.Z}
	case 15:
		return Beacon{-b.Z, -b.X, b.Y}
	case 16:
		return Beacon{-b.Z, b.Y, b.X}
	case 17:
		return Beacon{b.Y, b.Z, b.X}
	case 18:
		return Beacon{b.Z, -b.Y, b.X}
	case 19:
		return Beacon{-b.Y, -b.Z, b.X}
	case 20:
		return Beacon{-b.Z, -b.Y, -b.X}
	case 21:
		return Beacon{-b.Y, b.Z, -b.X}
	case 22:
		return Beacon{b.Z, b.Y, -b.X}
	case 23:
		return Beacon{b.Y, -b.Z, -b.X}
	default:
		panic("Ack!")
	}
}

func solve() (int, int) {

	scanners := getInput()

	for i := range scanners {
		scanners[i].InitRotations()
		scanners[i].InitVectors()
	}

	scanners[0].Settled = true
	scanners[0].Rotation = 0

	settledScanners := []*Scanner{scanners[0]}
	settledScanner := scanners[0]

	baseScannerIndex := 0
	allRotated := false
	for !allRotated {

		for i := range scanners {
			unsettledScanner := scanners[i]
			if !unsettledScanner.IsSettled() {
				for rotationVectorsIndex := range unsettledScanner.RotationVectors {
					intersections := intersections(settledScanner.SettledVectors(), unsettledScanner.RotationVectors[rotationVectorsIndex])

					if len(intersections) >= 12 {
						vector := calculateVector(intersections[0])

						unsettledScanner.Settled = true
						unsettledScanner.Rotation = rotationVectorsIndex
						unsettledScanner.Beacons = unsettledScanner.BeaconRotations[rotationVectorsIndex]
						unsettledScanner.Vector = vector

						for i := range unsettledScanner.Beacons {
							unsettledScanner.Beacons[i].ApplyOffset(vector)
						}

						unsettledScanner.InitVectors()

						settledScanners = append(settledScanners, unsettledScanner)
						break
					}
				}
			}
		}

		allRotated = true
		for _, s := range scanners {
			allRotated = allRotated && s.IsSettled()
		}

		if !allRotated {
			baseScannerIndex++
			if baseScannerIndex >= len(scanners) {
				baseScannerIndex = 0
			}

			for ; baseScannerIndex < len(scanners); baseScannerIndex++ {
				if scanners[baseScannerIndex].IsSettled() {
					settledScanner = scanners[baseScannerIndex]
					break
				}
			}
		}
	}

	uniques := make(map[Beacon]bool)
	for _, settledScanner := range settledScanners {
		for _, beacon := range settledScanner.SettledBeacons() {
			uniques[beacon] = true
		}
	}

	// find max manhattan distance between each scanner
	maxDistance := 0
	for i := range scanners {
		for j := range scanners {
			if i == j {
				continue
			}

			distance := scanners[i].ManhattanDistance(scanners[j])
			fmt.Printf("Scanner %d and %d are %d apart\n", i, j, distance)
			if distance > maxDistance {
				maxDistance = distance
			}
		}
	}

	return len(uniques), maxDistance
}

func Part1(solution int) Any {
	return solution
}

func Part2(solution int) Any {
	return solution
}

func intersections(origin, target []DistanceVector) [][2]DistanceVector {
	intersecting := [][2]DistanceVector{}
	for _, v1 := range origin {
		for _, v2 := range target {
			if v1.Vector.Equals(v2.Vector) {
				intersecting = append(intersecting, [2]DistanceVector{v1, v2})
			}
		}
	}

	return intersecting
}

func calculateVector(v [2]DistanceVector) Vector {
	return Vector{
		X: v[0].Origin.X - v[1].Origin.X,
		Y: v[0].Origin.Y - v[1].Origin.Y,
		Z: v[0].Origin.Z - v[1].Origin.Z,
	}
}

func main() {
	solution1, solution2 := solve()

	part1Solution := Part1(solution1)
	part2Solution := Part2(solution2)

	fmt.Printf("Day 19: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 19: Part 2: = %+v\n", part2Solution)
}

func getInput() []*Scanner {
	lines, _ := ReadLines(f, "input.txt")

	scanners := []*Scanner{}

	var currentScanner *Scanner
	var beacon Beacon
	for _, line := range lines {
		switch {

		case strings.HasPrefix(line, "--- scanner"):
			currentScanner = &Scanner{
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
	scanners = append(scanners, currentScanner)

	return scanners
}
