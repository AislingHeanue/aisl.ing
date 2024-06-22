package maths

import (
	"image/color"
)

type Axis int

const (
	X Axis = iota
	Y
	Z
)

type Cube struct {
	Side float32
	// Points           []*Point
	CentrePoint Point
	// FinalCentrePoint Point
	Colours []color.RGBA
	// These angles are the angular displacement of the cube from its original position
	// They are performed in the order Y, X
	// These rotations are applied when DrawCube is used.
	AngleX float32
	AngleY float32
	// FinalAngleX float64
	// FinalAngleY float64
}

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

func NewCube(origin Point, side float32) *Cube {
	return NewCubeWithColours(origin, side, []color.RGBA{WHITE, ORANGE, GREEN, RED, BLUE, YELLOW})
}

func NewCubeWithColours(origin Point, side float32, colours []color.RGBA) *Cube {
	var out Cube
	out.Side = side
	out.CentrePoint = origin

	// out.Points = []*Point{
	// 	origin.Add(Point{-side / 2, -side / 2, -side / 2}), // GREEN YELLOW RED
	// 	origin.Add(Point{side / 2, -side / 2, -side / 2}),  // BLUE YELLOW RED
	// 	origin.Add(Point{side / 2, -side / 2, side / 2}),   // BLUE YELLOW ORANGE
	// 	origin.Add(Point{-side / 2, -side / 2, side / 2}),  // GREEN YELLOW ORANGE
	// 	origin.Add(Point{-side / 2, side / 2, -side / 2}),  // GREEN WHITE RED
	// 	origin.Add(Point{side / 2, side / 2, -side / 2}),   // BLUE WHITE RED
	// 	origin.Add(Point{side / 2, side / 2, side / 2}),    // BLUE WHITE ORANGE
	// 	origin.Add(Point{-side / 2, side / 2, side / 2}),   // GREEN WHITE ORANGE
	// }
	out.Colours = colours

	return &out
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

	points := []*Point{
		c.CentrePoint.Add(Point{-c.Side / 2, -c.Side / 2, -c.Side / 2}), // GREEN YELLOW RED
		c.CentrePoint.Add(Point{c.Side / 2, -c.Side / 2, -c.Side / 2}),  // BLUE YELLOW RED
		c.CentrePoint.Add(Point{c.Side / 2, -c.Side / 2, c.Side / 2}),   // BLUE YELLOW ORANGE
		c.CentrePoint.Add(Point{-c.Side / 2, -c.Side / 2, c.Side / 2}),  // GREEN YELLOW ORANGE
		c.CentrePoint.Add(Point{-c.Side / 2, c.Side / 2, -c.Side / 2}),  // GREEN WHITE RED
		c.CentrePoint.Add(Point{c.Side / 2, c.Side / 2, -c.Side / 2}),   // BLUE WHITE RED
		c.CentrePoint.Add(Point{c.Side / 2, c.Side / 2, c.Side / 2}),    // BLUE WHITE ORANGE
		c.CentrePoint.Add(Point{-c.Side / 2, c.Side / 2, c.Side / 2}),   // GREEN WHITE ORANGE
	}
	for i, p := range points {
		points[i] = p.Rotate(c.CentrePoint, c.AngleY, Y)
	}

	out.VerticesArray = make([]float32, 72)
	for i, index := range indexOfVerticesArray {
		pointSlice := points[index].toSlice()
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

	// fmt.Println(len(c.Points))
	// fmt.Println(len(out.VerticesArray) / 3)
	// fmt.Println(len(out.IndicesArray))
	// fmt.Println(len(out.ColourArray) / 4)

	out.VCount = 24
	out.ICount = 36
	out.CCount = 24

	return out
}
