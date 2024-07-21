package maths

import (
	"fmt"
)

type Point [3]float32

func (p Point) Length() float32 {
	var total float32
	for _, num := range p {
		total += num * num
	}

	return total
}

func (v1 Point) Dot(v2 Point) float32 {
	var total float32
	for i := range v1 {
		total += v1[i] * v2[i]
	}

	return total
}

func (v1 Point) Add(v2 Point) *Point {
	v3 := &Point{}
	for i := range v1 {
		v3[i] = v1[i] + v2[i]
	}

	return v3
}

func (v1 Point) Reverse() *Point {
	v2 := &Point{}
	for i, num := range v1 {
		v2[i] = -num
	}

	return v2
}

func (v1 Point) Subtract(v2 Point) *Point {
	return v1.Add(*v2.Reverse())
}

func (p Point) Scale(c float32) *Point {
	return &Point{
		c * p[0],
		c * p[1],
		c * p[2],
	}
}

func (p Point) X() float32 {
	return p[0]
}

func (p Point) Y() float32 {
	return p[1]
}

func (p Point) Z() float32 {
	return p[2]
}

func (p Point) String() string {
	return fmt.Sprintf("(%d %d %d)", int(p[0]), int(p[1]), int(p[2]))
}

func (v1 Point) Cross(v2 Point) *Point {
	return &Point{
		v1[1]*v2[2] - v1[2]*v2[1],
		v1[2]*v2[0] - v1[0]*v2[2],
		v1[0]*v2[1] - v1[1]*v2[0],
	}
}

func (p Point) GetZOrthogonal(x, y float32) float32 {
	return -(p[0]*x + p[1]*y) / p[2]
}

func (p Point) ToSlice() []float32 {
	return []float32{
		float32(p[0]),
		float32(p[1]),
		float32(p[2]),
	}
}

func (p Point) Rotate(anchor Point, angle float32, axis Axis) *Point {
	displacement := p.Subtract(anchor)
	matrix := getRotation3Matrix(angle, axis)

	return displacement.MatrixDot(matrix).Add(anchor)
}

func (p Point) MatrixDot(m [3][3]float32) *Point {
	out := Point{}
	for k := range []int{0, 1, 2} {
		for i := range []int{0, 1, 2} {
			out[i] += m[i][k] * p[k]
		}
	}

	return &out
}

func getRotation3Matrix(angle float32, axis Axis) [3][3]float32 {
	switch axis {
	case X:
		return [3][3]float32{
			{1, 0, 0},
			{0, cos(angle), sin(angle)},
			{0, -sin(angle), cos(angle)},
		}
	case Y:
		return [3][3]float32{
			{cos(angle), 0, -sin(angle)},
			{0, 1, 0},
			{sin(angle), 0, cos(angle)},
		}
	case Z:
		return [3][3]float32{
			{cos(angle), sin(angle), 0},
			{-sin(angle), cos(angle), 0},
			{0, 0, 1},
		}
	default:
		return [3][3]float32{}
	}
}
