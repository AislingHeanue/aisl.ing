package model

import (
	"math"
)

func (sg *ShapeGroup) Rotate(anchor Vector, angle float64, axis Axis) {
	for _, s := range sg.Shapes {
		s.Rotate(anchor, angle, axis)
	}
}

// func (sg *ShapeGroup) ResetRotation(anchor Vector) {
// 	for _, s := range sg.Shapes {
// 		s.Rotate(anchor, -s.AngleX, X)
// 		s.Rotate(anchor, -s.AngleY, Y)
// 	}
// }

// for each point
// construct a line from anchor to point
// rotate 2 of the coordinate of the line through an angle, anticlockwise
func (s *Shape) Rotate(anchor Vector, angle float64, axis Axis) {
	// switch axis {
	// case X:
	// 	s.AngleX += angle
	// case Y:
	// 	s.AngleY += angle
	// }
	for _, p := range s.Points {
		p.Rotate(anchor, angle, axis)
	}
	// rotate the centre point of each face in the shape
	for _, f := range s.Faces {
		f.Centre.Rotate(anchor, angle, axis)
	}
}

func (p *Point) Rotate(anchor Vector, angle float64, axis Axis) {
	p.Vector.Rotate(anchor, angle, axis)
}

func (p *Vector) Rotate(anchor Vector, angle float64, axis Axis) {
	displacement := p.Subtract(anchor)
	matrix := getRotationMatrix(angle, axis)

	*p = *matrix.Dot(*displacement).Add(anchor)
}

func getRotationMatrix(angle float64, axis Axis) Matrix {
	switch axis {
	case X:
		return Matrix{
			{1, 0, 0},
			{0, math.Cos(angle), math.Sin(angle)},
			{0, -math.Sin(angle), math.Cos(angle)},
		}
	case Y:
		return Matrix{
			{math.Cos(angle), 0, -math.Sin(angle)},
			{0, 1, 0},
			{math.Sin(angle), 0, math.Cos(angle)},
		}
	case Z:
		return Matrix{
			{math.Cos(angle), math.Sin(angle), 0},
			{-math.Sin(angle), math.Cos(angle), 0},
			{0, 0, 1},
		}
	default:
		return Matrix{}
	}
}
