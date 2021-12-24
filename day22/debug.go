package main

import "fmt"

func BlenderDump(cuboids []Cuboid) {
	fmt.Println("")
	fmt.Println("collection = bpy.data.collections.new(\"State\")")
	fmt.Println("bpy.context.scene.collection.children.link(collection)")
	for _, c := range cuboids {
		c.BlenderDump()
	}
}

func (c Cuboid) BlenderDump() {
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
