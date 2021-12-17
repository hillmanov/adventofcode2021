package main

import (
	. "adventofcode/utils"
	"embed"
	"fmt"
	"math"
)

//go:embed input.txt
var f embed.FS

type target struct {
	MinX int
	MaxX int
	MinY int
	MaxY int
}

type probe struct {
	X         int
	Y         int
	VelocityX int
	VelocityY int
}

func (p *probe) Step() {
	p.X += p.VelocityX
	p.Y += p.VelocityY

	if p.VelocityX > 0 {
		p.VelocityX--
	} else if p.VelocityX < 0 {
		p.VelocityX++
	}

	p.VelocityY--
}

func Part1() Any {
	t := getInput()

	maxY := math.MinInt
	for velocityX := 0; velocityX <= t.MaxX; velocityX++ {
		for velocityY := t.MinY; velocityY <= -t.MinY; velocityY++ {
			p := probe{
				X:         0,
				Y:         0,
				VelocityX: velocityX,
				VelocityY: velocityY,
			}
			attemptMaxY := math.MinInt
			for p.X < t.MaxX && p.Y >= t.MinY {
				p.Step()
				attemptMaxY = MaxInt(attemptMaxY, p.Y)
				if probeInTargetArea(p, t) {
					maxY = MaxInt(maxY, attemptMaxY)
					break
				}
			}
		}
	}

	return maxY
}

func Part2() Any {
	t := getInput()

	count := 0
	for velocityX := 0; velocityX <= t.MaxX; velocityX++ {
		for velocityY := t.MinY; velocityY <= -t.MinY; velocityY++ {
			p := probe{
				X:         0,
				Y:         0,
				VelocityX: velocityX,
				VelocityY: velocityY,
			}
			for p.X < t.MaxX && p.Y >= t.MinY {
				p.Step()
				if probeInTargetArea(p, t) {
					count++
					break
				}
			}
		}
	}

	return count
}

func probeInTargetArea(p probe, t target) bool {
	return p.X >= t.MinX && p.X <= t.MaxX && p.Y >= t.MinY && p.Y <= t.MaxY
}

func getInput() target {
	contents, _ := ReadContents(f, "input.txt")
	t := target{}
	fmt.Sscanf(contents, "target area: x=%d..%d, y=%d..%d", &t.MinX, &t.MaxX, &t.MinY, &t.MaxY)
	return t
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 17: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 17: Part 2: = %+v\n", part2Solution)
}
