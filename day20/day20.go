package main

import (
	. "adventofcode/utils"
	"embed"
	"fmt"
	"math"
	"strconv"
)

//go:embed input.txt
var f embed.FS

type Pixel struct {
	Row int
	Col int
}

func (p Pixel) EnhancerCoords() []Pixel {
	return []Pixel{
		{Row: p.Row - 1, Col: p.Col - 1},
		{Row: p.Row - 1, Col: p.Col},
		{Row: p.Row - 1, Col: p.Col + 1},

		{Row: p.Row, Col: p.Col - 1},
		{Row: p.Row, Col: p.Col},
		{Row: p.Row, Col: p.Col + 1},

		{Row: p.Row + 1, Col: p.Col - 1},
		{Row: p.Row + 1, Col: p.Col},
		{Row: p.Row + 1, Col: p.Col + 1},
	}
}

type Image map[Pixel]string

func (image Image) GetDimensions() (minRow, maxRow, minCol, maxCol int) {
	minCol, minRow = math.MaxInt, math.MaxInt
	maxCol, maxRow = math.MinInt, math.MinInt
	for pixel := range image {
		minRow = MinInt(minRow, pixel.Row)
		maxRow = MaxInt(maxRow, pixel.Row)

		minCol = MinInt(minCol, pixel.Col)
		maxCol = MaxInt(maxCol, pixel.Col)
	}

	return minRow, maxRow, minCol, maxCol
}

func (image Image) Render() {
	// Clear screen fast
	fmt.Printf("\033[0;0H")
	minRow, maxRow, minCol, maxCol := image.GetDimensions()

	for row := minRow - 2; row <= maxRow+2; row++ {
		for col := minCol - 2; col <= maxCol+2; col++ {
			fmt.Printf("%s", image[Pixel{Row: row, Col: col}])
		}
		fmt.Println()
	}
}

func (g Image) GetIndexForPixel(pixel Pixel, iea string, step int) int {
	binString := ""
	for _, iP := range pixel.EnhancerCoords() {
		state, ok := g[iP]
		if !ok {
			if step%2 == 0 {
				state = string(iea[len(iea)-1])
			} else {
				state = string(iea[0])
			}
		}

		switch state {
		case "#":
			binString += "1"
		case ".":
			binString += "0"
		}
	}
	return binToInt(binString)
}

func (image Image) Enhance(iea string, step int) Image {
	enhancedImage := Image{}

	// minRow, maxRow, minCol, maxCol := image.GetDimensions()

	for row := 0 - 2 - (step * 2); row <= 100+2+(step*2); row++ {
		for col := 0 - 2 - (step * 2); col <= 100+2+(step*2); col++ {
			pixel := Pixel{Row: row, Col: col}
			enhancedImage[pixel] = string(iea[image.GetIndexForPixel(pixel, iea, step)])
		}
	}

	return enhancedImage
}

func Part1() Any {
	iea, image := getInput()

	for step := 0; step < 2; step++ {
		image = image.Enhance(iea, step)
	}

	litCount := 0
	for _, state := range image {
		if state == "#" {
			litCount++
		}
	}

	return litCount
}

func Part2() Any {
	iea, image := getInput()

	for step := 0; step < 50; step++ {
		image = image.Enhance(iea, step)
	}

	litCount := 0
	for _, state := range image {
		if state == "#" {
			litCount++
		}
	}

	return litCount
}

func getInput() (string, Image) {
	lines, _ := ReadLines(f, "input.txt")

	imageEnchancementAlgorithm := lines[0]
	image := Image{}

	for rowIndex, row := range lines[2:] {
		for colIndex := range row {
			image[Pixel{Row: rowIndex, Col: colIndex}] = string(row[colIndex])
		}
	}

	return imageEnchancementAlgorithm, image
}

func binToInt(bin string) int {
	v, err := strconv.ParseInt(bin, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(v)
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 20: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 20: Part 2: = %+v\n", part2Solution)
}
