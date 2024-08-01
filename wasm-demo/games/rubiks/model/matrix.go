package model

import (
	"github.com/gowebapi/webapi/core/jsconv"
	"github.com/gowebapi/webapi/graphics/webgl"
)

type Mat4 [4][4]float32

func I4() Mat4 {
	return Mat4{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
}

// This just, does not work for me, so I'll use a halfway solution that looks good
// https://developer.mozilla.org/en-US/docs/Web/API/WebGL_API/WebGL_model_view_projection#perspective_projection_matrix
func PerspectiveMatrix(fov float32, ratio float32, nearZ float32, farZ float32) *Mat4 {
	f := 1 / Tan(fov/2)
	r := 1 / (farZ - nearZ)
	return &Mat4{
		{f / ratio, 0, 0, 0},
		{0, f, 0, 0},
		{0, 0, r * (nearZ + farZ), 1},
		{0, 0, 0, 1},
	}
}

func (m Mat4) Scale(c float32) *Mat4 {
	out := Mat4{}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			out[i][j] = c * m[i][j]
		}
	}

	return &out
}

func (m Mat4) Rotate(angle float32, axis Axis) *Mat4 {
	return m.MatrixDot(getRotationMat4(angle, axis))
}

func (m1 Mat4) MatrixDot(m2 Mat4) *Mat4 {
	out := &Mat4{}
	for i := range []int{0, 1, 2, 3} {
		for j := range []int{0, 1, 2, 3} {
			for k := range []int{0, 1, 2, 3} {
				out[i][j] += m1[i][k] * m2[k][j]
			}
		}
	}

	return out
}

func (m Mat4) ToJS() *webgl.Union {
	flatMatrix := make([]float32, 16)
	for i := range []int{0, 1, 2, 3} {
		for j := range []int{0, 1, 2, 3} {
			flatMatrix[i*4+j] = m[i][j]
		}
	}

	return webgl.UnionFromJS(jsconv.Float32ToJs(flatMatrix))
}

func getRotationMat4(angle float32, axis Axis) Mat4 {
	switch axis {
	case X:
		return Mat4{
			{1, 0, 0, 0},
			{0, Cos(angle), Sin(angle), 0},
			{0, -Sin(angle), Cos(angle), 0},
			{0, 0, 0, 1},
		}
	case Y:
		return Mat4{
			{Cos(angle), 0, -Sin(angle), 0},
			{0, 1, 0, 0},
			{Sin(angle), 0, Cos(angle), 0},
			{0, 0, 0, 1},
		}
	case Z:
		return Mat4{
			{Cos(angle), Sin(angle), 0, 0},
			{-Sin(angle), Cos(angle), 0, 0},
			{0, 0, 1, 0},
			{0, 0, 0, 1},
		}
	default:
		return Mat4{}
	}
}
