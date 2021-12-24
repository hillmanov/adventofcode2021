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
	Side  string
	State string
	MinX  int
	MaxX  int
	MinY  int
	MaxY  int
	MinZ  int
	MaxZ  int
}

type Voxel struct {
	X int
	Y int
	Z int
}

func BlenderDump(cuboids []Cuboid, collection string) {
	fmt.Println("")
	fmt.Printf("collection = bpy.data.collections.new(\"%s\")\n", collection)
	fmt.Println("bpy.context.scene.collection.children.link(collection)")
	for _, c := range cuboids {
		c.Blender()
	}
}

func (c Cuboid) Blender() {
	var x, y, z float64
	for i := 0; i < c.Index; i++ {
		x += .05
		y += .3
		z += .7
		if x > 3 {
			x = 0
		}
		if y > 3 {
			y = 0
		}
		if z > 3 {
			z = 0
		}
	}

	fmt.Printf(`
cubeMesh = bpy.data.meshes.new("cube%d")
color = bpy.data.materials.new("CubeColor")
color.diffuse_color = ( %.1f, %.1f, %.1f, 0.9 )
cubeMesh.materials.append(color)
cubeObj = bpy.data.objects.new("cube%d", cubeMesh)
cubeObj.location = bpy.context.scene.cursor.location
collection.objects.link(cubeObj)
cubeMesh.from_pydata([(%d, %d, %d), (%d, %d, %d), (%d, %d, %d), (%d, %d, %d), (%d, %d, %d), (%d, %d, %d), (%d, %d, %d), (%d, %d, %d)],[],[(0,1,2,3), (4,5,6,7), (0,4,5,1), (1,5,6,2), (2,6,7,3), (3,7,4,0)])
cubeMesh.update(calc_edges=True)

	`, c.Index, float64(x), float64(y), float64(z),
		c.Index,
		c.MinX, c.MaxY, c.MinZ,
		c.MaxX, c.MaxY, c.MinZ,
		c.MaxX, c.MinY, c.MinZ,
		c.MinX, c.MinY, c.MinZ,

		c.MinX, c.MaxY, c.MaxZ,
		c.MaxX, c.MaxY, c.MaxZ,
		c.MaxX, c.MinY, c.MaxZ,
		c.MinX, c.MinY, c.MaxZ,
	)

}

func (c Cuboid) Volume() int {
	return (c.MaxX - c.MinX) * (c.MaxY - c.MinY) * (c.MaxZ - c.MinZ)
}

func (c Cuboid) IsNil() bool {
	return c.Volume() == 0
}

func (A Cuboid) Contains(B Cuboid) bool {
	return A.MinX <= B.MinX && A.MaxX >= B.MaxX &&
		A.MinY <= B.MinY && A.MaxY >= B.MaxY &&
		A.MinZ <= B.MinZ && A.MaxZ >= B.MaxZ
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
			Side:  "Left",
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
			Side:  "Right",
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
			Side:  "Face",
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
			Side:  "Back",
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
			Side:  "Bottom",
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
			Side:  "Top",
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
	newCuboids := getInput()
	fmt.Println("import bpy")

	cuboids := []Cuboid{}

	for nI := 0; nI < len(newCuboids); nI++ {
		for cI := 0; cI < len(cuboids); cI++ {
			if newCuboids[nI].Intersects(cuboids[cI]) {

				switch {
				case newCuboids[nI].State == "on":
					if newCuboids[nI].Contains(cuboids[cI]) {
						cuboids[cI].State = "remove"
					} else if cuboids[cI].Contains(newCuboids[nI]) {
						newCuboids[nI].State = "remove"
					} else {
						cuboids[cI].State = "remove"
						cuboids = append(cuboids, cuboids[cI].Split(newCuboids[nI])...)
					}

				case newCuboids[nI].State == "off": // Only "on" cuboids are added to the list, but I check here just to help me keep things straight
					cuboids[cI].State = "remove"
					cuboids = append(cuboids, cuboids[cI].Split(newCuboids[nI])...)
				}
			}
		}

		if newCuboids[nI].State == "on" {
			cuboids = append(cuboids, newCuboids[nI])
		}

		// Get rid of all the cuboids marked to remove
		refresh := []Cuboid{}
		for _, cuboid := range cuboids {
			if cuboid.State == "remove" {
				continue
			}
			refresh = append(refresh, cuboid)
		}
		cuboids = refresh
		BlenderDump(cuboids, "State")
	}

	countOn := 0
	for _, c := range cuboids {
		countOn += c.Volume()
	}

	return countOn
}

func getInput() []Cuboid {
	lines, _ := ReadLines(f, "input.txt")

	steps := []Cuboid{}

	for i, line := range lines {
		step := Cuboid{Index: i}
		fmt.Sscanf(line, "%s x=%d..%d,y=%d..%d,z=%d..%d", &step.State, &step.MinX, &step.MaxX, &step.MinY, &step.MaxY, &step.MinZ, &step.MaxZ)
		step.MaxX += 1
		step.MaxY += 1
		step.MaxZ += 1
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
