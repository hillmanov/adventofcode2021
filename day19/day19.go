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
	RotationVectors [][]Vector
	Rotation        int
	Settled         bool
}

type Dist struct {
	X int
	Y int
	Z int
}

func (d Dist) Equals(o Dist) bool {
	return d.X == o.X && d.Y == o.Y && d.Z == o.Z
}

type Beacon struct {
	X int
	Y int
	Z int
}

func (b Beacon) Dist(o Beacon) Dist {
	return Dist{
		X: b.X - o.X,
		Y: b.Y - o.Y,
		Z: b.Z - o.Z,
	}
}

func (b Beacon) Equals(o Beacon) bool {
	return b.X == o.X &&
		b.Y == o.Y &&
		b.Z == o.Z
}

func (b *Beacon) ApplyOffset(x, y, z int) {
	b.X = b.X + x
	b.Y = b.Y + y
	b.Z = b.Z + z
}

type Vector struct {
	Dist   Dist
	Beacon Beacon
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
	s.RotationVectors = make([][]Vector, 24)
	for rotation := range s.BeaconRotations {
		s.RotationVectors[rotation] = []Vector{}
		for i := range s.BeaconRotations[rotation] {
			for j := range s.BeaconRotations[rotation] {
				if i == j {
					continue
				}

				s.RotationVectors[rotation] = append(s.RotationVectors[rotation],
					Vector{
						Dist:   s.BeaconRotations[rotation][i].Dist(s.BeaconRotations[rotation][j]),
						Beacon: s.BeaconRotations[rotation][i],
					},
				)
			}
		}
	}
}

func (s *Scanner) IsSettled() bool {
	return s.Settled
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

func Part1() Any {
	scanners := getInput()

	for i := range scanners {
		scanners[i].InitRotations()
		scanners[i].InitVectors()
	}

	scanners[0].Rotation = 0
	scanners[0].Settled = true

	settledScanners := []*Scanner{scanners[0]}

loop:
	unsettledScanners := getUnsettledScanners(scanners)
	for i := range unsettledScanners {
		unsettledScanner := unsettledScanners[i]
		for j := range settledScanners {
			settledScanner := settledScanners[j]

			for rotationVectorsIndex := range unsettledScanner.RotationVectors {
				intersections := intersection(settledScanner.SettledVectors(), unsettledScanner.RotationVectors[rotationVectorsIndex])
				if len(intersections) >= 12 {
					offsetX, offsetY, offsetZ := getOffset(intersections[0])

					fmt.Printf("rotationVectorsIndex = %+v\n", rotationVectorsIndex)
					fmt.Printf("offsetX, offsetY, offsetZ = %d %d %d\n", offsetX, offsetY, offsetZ)

					unsettledScanner.Rotation = rotationVectorsIndex
					unsettledScanner.Settled = true
					unsettledScanner.Beacons = unsettledScanner.BeaconRotations[rotationVectorsIndex]

					for i := range unsettledScanner.Beacons {
						unsettledScanner.Beacons[i].ApplyOffset(offsetX, offsetY, offsetZ)
					}

					unsettledScanner.InitRotations()
					unsettledScanner.InitVectors()

					settledScanners = append(settledScanners, unsettledScanner)

					goto loop
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

	fmt.Printf("len(uniques) = %+v\n", len(uniques))

	fmt.Println("Done")

	return nil
}

func Part2() Any {
	return nil
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
	currentScanner.Beacons = append(currentScanner.Beacons, beacon)
	scanners = append(scanners, currentScanner)

	return scanners
}

func intersection(origin, target []Vector) [][2]Vector {
	intersecting := [][2]Vector{}
	for _, v1 := range origin {
		for _, v2 := range target {
			if v1.Dist.Equals(v2.Dist) {
				intersecting = append(intersecting, [2]Vector{v1, v2})
			}
		}
	}

	return intersecting
}

func getUnsettledScanners(scanners []*Scanner) []*Scanner {
	unsettled := []*Scanner{}
	for i := range scanners {
		if !scanners[i].IsSettled() {
			unsettled = append(unsettled, scanners[i])
		}
	}
	return unsettled
}

func getOffset(v [2]Vector) (x, y, z int) {
	return v[0].Beacon.X - v[1].Beacon.X,
		v[0].Beacon.Y - v[1].Beacon.Y,
		v[0].Beacon.Z - v[1].Beacon.Z
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 19: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 19: Part 2: = %+v\n", part2Solution)
}
