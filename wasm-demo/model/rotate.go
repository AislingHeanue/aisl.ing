package model

import (
	"math"
)

// for each point
// construct a line from anchor to point
// rotate 2 of the coordinate of the line through an angle, anticlockwise
func (s *Shape) Rotate(anchor Vector, angle float64, axis Axis) {
	for _, p := range s.Points {
		p.Rotate(anchor, angle, axis)
	}
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
