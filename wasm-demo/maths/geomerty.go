package maths

import (
	"fmt"
	"image/color"
)

type Axis int

type Point struct {
	Vector *Vector
}

func (p Point) toSlice() []float32 {
	return []float32{
		float32(p.Vector[0]),
		float32(p.Vector[1]),
		float32(p.Vector[2]),
	}
}

const (
	X Axis = iota
	Y
	Z
)

type Face struct {
	Vertices []*Point
	Centre   *Point
	Colour   color.Color
}

type Shape struct {
	Points  []*Point
	Colours []color.RGBA
}

type DrawShape struct {
	VerticesArray []float32
	IndicesArray  []uint16
	ColourArray   []float32
	VCount        int
	ICount        int
	CCount        int
}

type DrawShapeGroup struct {
	VerticesArray []float32
	IndicesArray  []uint16
	ColourArray   []float32
	VCount        int
	ICount        int
	CCount        int
}

func GroupBuffers(s []DrawShape) DrawShapeGroup {
	var out DrawShapeGroup
	out.VerticesArray = []float32{}
	out.IndicesArray = []uint16{}
	out.ColourArray = []float32{}
	vCountOffset := 0
	iCountOffset := 0
	fmt.Println("sup")
	for _, shape := range s {
		// fmt.Println("yup")
		// out.Shapes = append(out.Shapes, shape)
		// out.Points = append(out.Points, shape.Points...)

		out.VerticesArray = append(out.VerticesArray, shape.VerticesArray...)

		for _, i := range shape.IndicesArray {
			out.IndicesArray = append(out.IndicesArray, i+uint16(vCountOffset))
		}
		out.ColourArray = append(out.ColourArray, shape.ColourArray...)

		vCountOffset += shape.VCount
		iCountOffset += shape.ICount
	}
	fmt.Println("bup")
	out.VCount = vCountOffset
	out.ICount = iCountOffset

	return out
}
