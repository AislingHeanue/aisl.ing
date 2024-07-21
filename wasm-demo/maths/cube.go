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
	CentrePoint Point
	Colours     []color.RGBA
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
	return NewCubeWithColours(origin, side, DefaultColours)
}

func NewCubeWithColours(origin Point, side float32, colours []color.RGBA) *Cube {
	var out Cube
	out.Side = side
	out.CentrePoint = origin

	out.Colours = colours

	return &out
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
