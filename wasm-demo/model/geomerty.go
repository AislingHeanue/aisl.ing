package model

type Axis int

const (
	X Axis = iota
	Y
	Z
)

type Line struct {
	Origin *Vector
	End    *Vector
}

func (l Line) Length() float64 {
	return l.Displacement().Length()
}

func (l Line) Displacement() *Vector {
	return Vector(*l.End).Subtract(Vector(*l.Origin))
}

type Face struct {
	Lines  []*Line
	Points []*Vector
}

func FaceFromLines(lines ...*Line) *Face {
	if len(lines) > 3 {
		return nil // face must have at least 3 points
	}
	out := &Face{Lines: lines}
	pointsFounds := make(map[*Vector]bool)
	for _, line := range lines {
		pointsFounds[line.Origin] = true
		pointsFounds[line.End] = true
	}
	for k := range pointsFounds {
		out.Points = append(out.Points, k)
	}

	return out
}

type Shape struct {
	Faces  []*Face
	Lines  []*Line
	Points []*Vector
}
