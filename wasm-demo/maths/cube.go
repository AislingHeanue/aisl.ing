package maths

import (
	"image/color"
)

// creates a new cube at the specified origin in the default rotation.
// The 'origin' is the cubes lowest x, y and z values.
/*
                        . 10.
                       11  9
                      . 8 .

                       6   7
  7 - 6    Y    Z
 /   /|    ^  /       4   5
4 - 5 |    | /
| 3 | 2    . --> X      . 2 .
|   |/                 3   1
0 - 1                 . 0 .
*/
type Cube Shape

var (
	RED    = color.RGBA{255, 0, 0, 255}
	GREEN  = color.RGBA{0, 255, 0, 255}
	BLUE   = color.RGBA{0, 0, 255, 255}
	ORANGE = color.RGBA{255, 128, 0, 255}
	YELLOW = color.RGBA{255, 255, 0, 255}
	WHITE  = color.RGBA{255, 255, 255, 255}
	BLACK  = color.RGBA{0, 0, 0, 255}

	DefaultColours = []color.RGBA{WHITE, BLUE, ORANGE, GREEN, RED, YELLOW}
)

func NewCube(origin Vector, side float64) *Cube {
	return NewCubeWithColours(origin, side, []color.RGBA{WHITE, ORANGE, GREEN, RED, BLUE, YELLOW})
}

func NewCubeWithColours(origin Vector, side float64, colours []color.RGBA) *Cube {
	var out Cube
	out.Points = []*Point{
		{Vector: origin.Add(Vector{-side / 2, -side / 2, -side / 2})}, // GREEN YELLOW RED
		{Vector: origin.Add(Vector{side / 2, -side / 2, -side / 2})},  // BLUE YELLOW RED
		{Vector: origin.Add(Vector{side / 2, -side / 2, side / 2})},   // BLUE YELLOW ORANGE
		{Vector: origin.Add(Vector{-side / 2, -side / 2, side / 2})},  // GREEN YELLOW ORANGE
		{Vector: origin.Add(Vector{-side / 2, side / 2, -side / 2})},  // GREEN WHITE RED
		{Vector: origin.Add(Vector{side / 2, side / 2, -side / 2})},   // BLUE WHITE RED
		{Vector: origin.Add(Vector{side / 2, side / 2, side / 2})},    // BLUE WHITE ORANGE
		{Vector: origin.Add(Vector{-side / 2, side / 2, side / 2})},   // GREEN WHITE ORANGE
	}
	out.Colours = colours

	return &out
}

func (c Cube) GetCentre() *Vector {
	return c.Points[0].Vector.Add(*c.Points[6].Vector).Scale(0.5)
}

func (c Cube) GetBuffers() DrawShape {
	var out DrawShape
	indexOfVerticesArray := []uint16{
		4, 5, 6, 7, // WHITE
		3, 2, 6, 7, // ORANGE
		0, 3, 7, 4, // GREEN
		1, 0, 4, 5, // RED
		1, 2, 6, 5, // BLUE
		0, 1, 2, 3, // YELLOW
	}

	out.VerticesArray = make([]float32, 72)
	for i, index := range indexOfVerticesArray {
		pointSlice := c.Points[index].toSlice()
		out.VerticesArray[3*i] = pointSlice[0]
		out.VerticesArray[3*i+1] = pointSlice[1]
		out.VerticesArray[3*i+2] = pointSlice[2]
	}

	out.IndicesArray = make([]uint16, 36)
	for j := 0; j < 6; j++ {
		// assume points are connected as 0->1->2->3
		// then we need 0,1,2,0,2,3
		out.IndicesArray[6*j] = uint16(4*j + 0)
		out.IndicesArray[6*j+1] = uint16(4*j + 1)
		out.IndicesArray[6*j+2] = uint16(4*j + 2)
		out.IndicesArray[6*j+3] = uint16(4*j + 0)
		out.IndicesArray[6*j+4] = uint16(4*j + 2)
		out.IndicesArray[6*j+5] = uint16(4*j + 3)

	}
	// fmt.Println(out.IndicesArray)

	// out.IndicesArray = []uint16{
	// 	7., 6., 5., 7., 5., 4., // WHITE   +0
	// 	14, 15, 11, 14, 11, 10, // ORANGE  +8
	// 	23, 20, 16, 23, 16, 19, // GREEN   +16
	// 	12, 13, 9., 12, 9., 8., // RED     +8
	// 	17, 21, 22, 17, 22, 18, // BLUE    +16
	// 	0., 1., 2., 0., 2., 3., // YELLOW  +0
	// }

	outColours := []float32{}
	for _, c := range c.Colours {
		outColours = append(outColours, float32(c.R)/256, float32(c.G)/256, float32(c.B)/256, float32(c.A)/256)
		outColours = append(outColours, float32(c.R)/256, float32(c.G)/256, float32(c.B)/256, float32(c.A)/256)
		outColours = append(outColours, float32(c.R)/256, float32(c.G)/256, float32(c.B)/256, float32(c.A)/256)
		outColours = append(outColours, float32(c.R)/256, float32(c.G)/256, float32(c.B)/256, float32(c.A)/256)
	}
	out.ColourArray = outColours
	// fmt.Println(out.ColoursArray)

	out.VCount = 24
	out.ICount = 36
	out.CCount = 24

	return out
}

func GroupBuffersFromCubes(c [][][]*Cube) DrawShapeGroup {
	d := []DrawShape{}
	for _, c1 := range c {
		for _, c2 := range c1 {
			for _, c3 := range c2 {
				if c3 != nil {
					d = append(d, c3.GetBuffers())
				}
			}
		}
	}

	return GroupBuffers(d)
}
