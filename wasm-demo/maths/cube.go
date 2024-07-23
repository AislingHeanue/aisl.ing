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
	// maybe this field can be moved?
	Colours            []color.RGBA
	VertexArrayIndices []uint16
	Points             []Point
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

func NewCube(origin Point, side float32) Cube {
	return NewCubeWithColours(origin, side, DefaultColours)
}

func NewCubeWithColours(origin Point, side float32, colours []color.RGBA) Cube {
	var out Cube
	out.Side = side

	out.Colours = make([]color.RGBA, len(colours))
	copy(out.Colours, colours)

	out.VertexArrayIndices = []uint16{
		4, 5, 6, 7, // WHITE
		3, 2, 6, 7, // ORANGE
		0, 3, 7, 4, // GREEN
		1, 0, 4, 5, // RED
		1, 2, 6, 5, // BLUE
		0, 1, 2, 3, // YELLOW
	}

	out.Points = []Point{
		origin.Add(Point{-side / 2, -side / 2, -side / 2}), // GREEN YELLOW RED
		origin.Add(Point{side / 2, -side / 2, -side / 2}),  // BLUE YELLOW RED
		origin.Add(Point{side / 2, -side / 2, side / 2}),   // BLUE YELLOW ORANGE
		origin.Add(Point{-side / 2, -side / 2, side / 2}),  // GREEN YELLOW ORANGE
		origin.Add(Point{-side / 2, side / 2, -side / 2}),  // GREEN WHITE RED
		origin.Add(Point{side / 2, side / 2, -side / 2}),   // BLUE WHITE RED
		origin.Add(Point{side / 2, side / 2, side / 2}),    // BLUE WHITE ORANGE
		origin.Add(Point{-side / 2, side / 2, side / 2}),   // GREEN WHITE ORANGE
	}

	return out
}

func (c *Cube) RotateColoursX(flip bool) {
	times := 1
	if flip {
		times = 3
	}
	for i := 0; i < times; i++ {
		c.Colours[0], c.Colours[1], c.Colours[5], c.Colours[3] =
			c.Colours[3], c.Colours[0], c.Colours[1], c.Colours[5]
	}
}

func (c *Cube) RotateColoursY(flip bool) {
	times := 1
	if flip {
		times = 3
	}
	for i := 0; i < times; i++ {
		c.Colours[1], c.Colours[2], c.Colours[3], c.Colours[4] =
			c.Colours[2], c.Colours[3], c.Colours[4], c.Colours[1]
	}
}

func (c *Cube) RotateColoursZ(flip bool) {
	times := 1
	if flip {
		times = 3
	}
	for i := 0; i < times; i++ {
		c.Colours[0], c.Colours[4], c.Colours[5], c.Colours[2] =
			c.Colours[2], c.Colours[0], c.Colours[4], c.Colours[5]
	}
}

func (c Cube) Rotate(anchor Point, angle float32, axis Axis) Cube {
	cubeCopy := c.Copy()
	for i, p := range c.Points {
		cubeCopy.Points[i] = p.Rotate(anchor, angle, axis)
	}

	return cubeCopy
}

func (c Cube) Copy() Cube {
	side := c.Side
	vertexIndices := c.VertexArrayIndices
	colours := make([]color.RGBA, len(c.Colours))
	points := make([]Point, len(c.Points))

	copy(colours, c.Colours)
	copy(points, c.Points)

	return Cube{
		Side:               side,
		VertexArrayIndices: vertexIndices,
		Colours:            colours,
		Points:             points,
	}
}
