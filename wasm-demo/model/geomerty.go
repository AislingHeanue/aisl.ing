package model

import "image/color"

type Axis int

// var largeNumber float64 = 30000

type Point struct {
	Vector              *Vector
	RenderedCoordinates *Vector
}

const (
	X Axis = iota
	Y
	Z
	LocalX
	LocalY
	LocalZ
)

// type Line struct {
// 	Origin *Point
// 	End    *Point
// }

// func (l Line) Length() float64 {
// 	return l.Displacement().Length()
// }

// func (l Line) Displacement() *Vector {
// 	return l.End.Vector.Subtract(*l.Origin.Vector)
// }

// func (l Line) RenderedDisplaceMent() *Vector {
// 	return l.End.RenderedCoordinates.Subtract(*l.Origin.RenderedCoordinates)
// }

// func (l Line) PointIsLeft(x, y float64) bool {
// 	ab := l.RenderedDisplaceMent()
// 	// if ab[0] < 0 {
// 	// 	ab = ab.Reverse() // want x to be positive
// 	// }
// 	newVector := Vector{x, y, 0}.Subtract(*l.Origin.RenderedCoordinates)

// 	return (ab[0]*newVector[1] - ab[1]*newVector[0]) < 0
// }

// func (l *Line) String() string {
// 	return fmt.Sprintf("%v -> %v\n", *l.Origin, *l.End)
// }

type Face struct {
	Vertices []*Point
	Centre   *Point
	Colour   color.Color
}

func (f Face) GetRenderedNormal() *Vector {
	return f.Vertices[2].RenderedCoordinates.Subtract(*f.Vertices[0].RenderedCoordinates).Cross(*f.Vertices[1].RenderedCoordinates.Subtract(*f.Vertices[0].RenderedCoordinates))
}

func FaceFromPoints(points ...*Point) *Face {
	centre := &Point{Vector: &Vector{0, 0, 0}}
	for _, p := range points {
		centre.Vector = centre.Vector.Add(*p.Vector.Scale(1 / float64(len(points))))
	}
	return &Face{Vertices: points, Centre: centre}
}

func FaceFromPointsWithColour(colour color.Color, points ...*Point) *Face {
	face := FaceFromPoints(points...)
	face.Colour = colour

	return face
}

// func FaceFromLines(lines ...*Line) *Face {
// 	if len(lines) < 3 {
// 		return nil // face must have at least 3 points
// 	}
// 	out := &Face{}
// 	pointsFounds := make(map[*Point]bool)
// 	for _, line := range lines {
// 		pointsFounds[line.Origin] = true
// 		pointsFounds[line.End] = true
// 	}
// 	for k := range pointsFounds {
// 		out.Points = append(out.Points, k)
// 	}

// 	return out
// }

// // only works on squares for now. I wanted this to be much smarter
// func (f *Face) CheckContains(x, y float64) bool {
// 	results := []int{0, 0, 0, 0}
// 	for i, line := range f.Lines {
// 		if line.PointIsLeft(x, y) {
// 			results[i] = 1
// 		}
// 	}
// 	return slices.Equal(results, []int{0, 0, 1, 1}) ||
// 		slices.Equal(results, []int{0, 1, 1, 0}) ||
// 		slices.Equal(results, []int{1, 1, 0, 0}) ||
// 		slices.Equal(results, []int{1, 0, 0, 1})
// }

// // assumes the first 2 lines of f are not parallel
// func (f *Face) GetZ(x, y float64) float64 {
// 	return (*f.Lines[0].Displacement()).
// 		Cross(*f.Lines[1].Displacement()).
// 		GetZOrthogonal(
// 			x-f.Lines[0].Origin.RenderedCoordinates[0],
// 			y-f.Lines[0].Origin.RenderedCoordinates[1],
// 		) + f.Lines[0].Origin.RenderedCoordinates[2]
// }

type Shape struct {
	Faces []*Face
	// Lines  []*Line
	Points []*Point
	// AngleX float64
	// AngleY float64
	// AngleZ float64
}

type ShapeGroup struct {
	Shapes []*Shape
	Faces  []*Face
	// // Lines  []*Line
	// Points []*Point
}

func NewShapeGroup(s []*Shape) ShapeGroup {
	var out ShapeGroup
	for _, shape := range s {
		out.Shapes = append(out.Shapes, shape)
		// // out.Lines = append(out.Lines, shape.Lines...)
		out.Faces = append(out.Faces, shape.Faces...)
		// out.Points = append(out.Points, shape.Points...)
	}

	return out
}
