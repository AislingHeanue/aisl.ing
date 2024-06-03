package model

import "image/color"

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
// New requirement, all faces have their lines specified in clockwise order
type Cube Shape

type CubeColour color.Color

var (
	RED    CubeColour = color.RGBA{255, 0, 0, 255}
	GREEN  CubeColour = color.RGBA{0, 255, 0, 255}
	BLUE   CubeColour = color.RGBA{0, 0, 255, 255}
	ORANGE CubeColour = color.RGBA{255, 128, 0, 255}
	YELLOW CubeColour = color.RGBA{255, 255, 0, 255}
	WHITE  CubeColour = color.RGBA{255, 255, 255, 255}
	BLACK  CubeColour = color.RGBA{0, 0, 0, 255}

	DefaultColours = []color.Color{WHITE, ORANGE, GREEN, RED, BLUE, YELLOW}
)

func NewCube(origin Vector, side float64) *Cube {
	return NewCubeWithColours(origin, side, []color.Color{WHITE, ORANGE, GREEN, RED, BLUE, YELLOW})
}

func NewCubeWithColours(origin Vector, side float64, colours []color.Color) *Cube {
	var out Cube
	// originCopy := origin
	out.Points = []*Point{
		{Vector: origin.Add(Vector{-side / 2, -side / 2, -side / 2})},
		{Vector: origin.Add(Vector{side / 2, -side / 2, -side / 2})},
		{Vector: origin.Add(Vector{side / 2, -side / 2, side / 2})},
		{Vector: origin.Add(Vector{-side / 2, -side / 2, side / 2})},
		{Vector: origin.Add(Vector{-side / 2, side / 2, -side / 2})},
		{Vector: origin.Add(Vector{side / 2, side / 2, -side / 2})},
		{Vector: origin.Add(Vector{side / 2, side / 2, side / 2})},
		{Vector: origin.Add(Vector{-side / 2, side / 2, side / 2})},
	}
	out.Faces = []*Face{
		FaceFromPointsWithColour(colours[0], out.Points[7], out.Points[6], out.Points[5], out.Points[4]),
		FaceFromPointsWithColour(colours[1], out.Points[6], out.Points[7], out.Points[3], out.Points[2]),
		FaceFromPointsWithColour(colours[2], out.Points[7], out.Points[4], out.Points[0], out.Points[3]),
		FaceFromPointsWithColour(colours[3], out.Points[4], out.Points[5], out.Points[1], out.Points[0]),
		FaceFromPointsWithColour(colours[4], out.Points[1], out.Points[5], out.Points[6], out.Points[2]),
		FaceFromPointsWithColour(colours[5], out.Points[0], out.Points[1], out.Points[2], out.Points[3]),
	}

	return &out
}

func (c Cube) GetCentre() *Vector {
	return c.Points[0].Vector.Add(*c.Points[6].Vector).Scale(0.5)
}

func (c *Cube) Rotate(anchor Vector, angle float64, axis Axis) {
	shape := Shape(*c)
	(&shape).Rotate(anchor, angle, axis)
	*c = Cube(shape)
}

// func (c *Cube) AddAngle(num float64, axis Axis) {
// 	shape := Shape(*c)
// 	switch axis {
// 	case X:
// 		shape.AngleX += num
// 	case Y:
// 		shape.AngleY += num
// 	}
// 	*c = Cube(shape)
// }
