package model

import "fmt"

type Matrix [3][3]float64

func NewMatrix() Matrix {
	return Matrix{}
}

func (m Matrix) Dot(v Vector) *Vector {
	out := &Vector{}
	for i := range v {
		for j := range m[i] {
			out[j] += m[i][j] * v[i]
		}
	}

	return out
}

func (m Matrix) Scale(c float64) *Matrix {
	out := &Matrix{}
	for i := range m {
		for j := range m[i] {
			out[i][j] += m[i][j] * c
		}
	}

	return out
}

type Vector [3]float64

func (v Vector) Length() float64 {
	total := 0.
	for _, num := range v {
		total += num * num
	}

	return total
}

func (v1 Vector) Dot(v2 Vector) float64 {
	total := 0.
	for i := range v1 {
		total += v1[i] * v2[i]
	}

	return total
}

func (v1 Vector) Add(v2 Vector) *Vector {
	v3 := &Vector{}
	for i := range v1 {
		v3[i] = v1[i] + v2[i]
	}

	return v3
}

func (v1 Vector) Reverse() *Vector {
	v2 := &Vector{}
	for i, num := range v1 {
		v2[i] = -num
	}

	return v2
}

func (v1 Vector) Subtract(v2 Vector) *Vector {
	return v1.Add(*v2.Reverse())
}

func (v Vector) Scale(c float64) *Vector {
	return &Vector{
		c * v[0],
		c * v[1],
		c * v[2],
	}
}

func (p Vector) X() float64 {
	return p[0]
}

func (p Vector) Y() float64 {
	return p[1]
}

func (p Vector) Z() float64 {
	return p[2]
}

func (p Vector) String() string {
	return fmt.Sprintf("(%d %d %d)", int(p[0]), int(p[1]), int(p[2]))
}
