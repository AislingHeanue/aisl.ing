package model

import (
	"image/color"

	"github.com/AislingHeanue/aisling-codes/wasm-demo/common"
)

type turnInfo struct {
	xSelector         cubeSelector
	ySelector         cubeSelector
	zSelector         cubeSelector
	reverse           bool
	axis              Axis
	allowedConcurrent []Face
}

type Face string

const (
	u Face = "u"
	d Face = "d"
	r Face = "r"
	l Face = "l"
	f Face = "f"
	b Face = "b"
	m Face = "m"
	e Face = "e"
	s Face = "s"
	x Face = "x"
	y Face = "y"
	z Face = "z"
)

var turnMap = map[Face]turnInfo{
	u: {EVERY, FINAL, EVERY, true, Y, []Face{d, e}},
	d: {EVERY, FIRST, EVERY, false, Y, []Face{e, u}},
	r: {FINAL, EVERY, EVERY, true, X, []Face{l, m}},
	l: {FIRST, EVERY, EVERY, false, X, []Face{m, r}},
	f: {EVERY, EVERY, FIRST, false, Z, []Face{s, b}},
	b: {EVERY, EVERY, FINAL, true, Z, []Face{f, s}},
	m: {INNER, EVERY, EVERY, false, X, []Face{l, r}},
	e: {EVERY, INNER, EVERY, false, Y, []Face{d, u}},
	s: {EVERY, EVERY, INNER, false, Z, []Face{b, f}},
	x: {EVERY, EVERY, EVERY, true, X, []Face{}},
	y: {EVERY, EVERY, EVERY, true, Y, []Face{}},
	z: {EVERY, EVERY, EVERY, false, Z, []Face{}},
}

type RubiksCube struct {
	Data      [][][]Cube
	Dimension int
}

func NewRubiksCube(d int, origin Point, sideLength float32, totalSideLength float32, sideLengthWithGap float32) *RubiksCube {
	data := make([][][]Cube, d)
	for i := range data {
		data[i] = make([][]Cube, d)
		for j := range data[i] {
			data[i][j] = make([]Cube, d)
		}
	}
	for x := range d {
		for y := range d {
			for z := range d {
				cubeOrigin := origin.
					Subtract(Point{totalSideLength / 2, totalSideLength / 2, totalSideLength / 2}).
					Add(Point{sideLengthWithGap * float32(x), sideLengthWithGap * float32(y), sideLengthWithGap * float32(z)}).
					Add(Point{sideLength / 2, sideLength / 2, sideLength / 2})
				colours, external := cubeColours(x, y, z, d)

				if external {
					data[x][y][z] = NewCubeWithColours(cubeOrigin, sideLength, colours)
				}
			}
		}
	}

	return &RubiksCube{Dimension: d, Data: data}
}

func cubeColours(x, y, z, dimension int) ([]color.RGBA, bool) {
	colours := []color.RGBA{BLACK, BLACK, BLACK, BLACK, BLACK, BLACK}
	external := false
	if y == dimension-1 {
		external = true
		colours[0] = DefaultColours[0]
	}
	if z == dimension-1 {
		external = true
		colours[1] = DefaultColours[1]
	}
	if x == 0 {
		external = true
		colours[2] = DefaultColours[2]
	}
	if z == 0 {
		external = true
		colours[3] = DefaultColours[3]
	}
	if x == dimension-1 {
		external = true
		colours[4] = DefaultColours[4]
	}
	if y == 0 {
		external = true
		colours[5] = DefaultColours[5]
	}

	return colours, external
}

func (r RubiksCube) Copy() RubiksCube {
	data := make([][][]Cube, r.Dimension)
	for i := range r.Dimension {
		data[i] = make([][]Cube, r.Dimension)
		for j := range r.Dimension {
			data[i][j] = make([]Cube, r.Dimension)
			for k := range r.Dimension {
				data[i][j][k] = r.Data[i][j][k]
			}
		}
	}

	return RubiksCube{Dimension: r.Dimension, Data: data}
}

func (r RubiksCube) FlattenBuffers() []common.DrawShape {
	cubes := []common.DrawShape{}
	for x := range r.Dimension {
		for y := range r.Dimension {
			for z := range r.Dimension {
				if r.isExternalCube(x, y, z) {
					cubes = append(cubes, r.Data[x][y][z].GetBuffers())
				}
			}
		}
	}

	return cubes
}

type cubeSelector int

const (
	FIRST cubeSelector = iota
	INNER
	FINAL
	EVERY
)

func (r RubiksCube) getCubeSubset(xSelector, ySelector, zSelector cubeSelector) [][3]int {
	xs := r.getRangeFromSelection(xSelector)
	ys := r.getRangeFromSelection(ySelector)
	zs := r.getRangeFromSelection(zSelector)

	coords := [][3]int{}

	for _, x := range xs {
		for _, y := range ys {
			for _, z := range zs {
				if r.isExternalCube(x, y, z) {
					coords = append(coords, [3]int{x, y, z})
				}
			}
		}
	}

	return coords
}

func (r RubiksCube) getRangeFromSelection(selector cubeSelector) []int {
	switch selector {
	case FIRST:
		return []int{0}
	case INNER:
		out := make([]int, r.Dimension-2)
		for i := range out {
			out[i] = i + 1
		}

		return out
	case FINAL:
		return []int{r.Dimension - 1}
	case EVERY:
		out := make([]int, r.Dimension)
		for i := range out {
			out[i] = i
		}

		return out
	default:
		return []int{}
	}
}

func (r RubiksCube) isExternalCube(x, y, z int) bool {
	xCond := x == 0 || x == r.Dimension-1
	yCond := y == 0 || y == r.Dimension-1
	zCond := z == 0 || z == r.Dimension-1

	return xCond || yCond || zCond
}

func (r RubiksCube) getPermutedCubeColours(x, y, z int, axis Axis, reverse bool) []color.RGBA {
	switch axis {
	case X:
		if reverse {
			return r.Data[x][z][r.Dimension-1-y].Colours
		} else {
			return r.Data[x][r.Dimension-1-z][y].Colours
		}
	case Y:
		if reverse {
			return r.Data[r.Dimension-1-z][y][x].Colours
		} else {
			return r.Data[z][y][r.Dimension-1-x].Colours
		}
	case Z:
		if reverse {
			return r.Data[y][r.Dimension-1-x][z].Colours
		} else {
			return r.Data[r.Dimension-1-y][x][z].Colours
		}
	default:
		return []color.RGBA{}
	}
}

// var _ *CubeController = RubiksCube{}

func (r *RubiksCube) Turn(face Face, reverse bool) {
	info, ok := turnMap[face]
	if !ok {
		return
	}
	if info.reverse {
		reverse = !reverse
	}
	newR := r.Copy()

	for _, coord := range r.getCubeSubset(info.xSelector, info.ySelector, info.zSelector) {
		newR.Data[coord[0]][coord[1]][coord[2]].Colours = r.getPermutedCubeColours(coord[0], coord[1], coord[2], info.axis, reverse)
		newR.Data[coord[0]][coord[1]][coord[2]].RotateColours(reverse, info.axis)
	}

	*r = newR
}

func (c Cube) GetBuffers() common.DrawShape {
	var out common.DrawShape

	out.VerticesArray = make([]float32, 72)
	for i, index := range c.VertexArrayIndices {
		pointSlice := c.Points[index].ToSlice()
		out.VerticesArray[3*i] = pointSlice[0]
		out.VerticesArray[3*i+1] = pointSlice[1]
		out.VerticesArray[3*i+2] = pointSlice[2]
	}

	out.IndicesArray = make([]uint16, 36)
	for j := range 6 {
		// assume points are connected as 0->1->2->3
		// then we need 0,1,2,0,2,3
		out.IndicesArray[6*j] = uint16(4*j + 0)
		out.IndicesArray[6*j+1] = uint16(4*j + 1)
		out.IndicesArray[6*j+2] = uint16(4*j + 2)
		out.IndicesArray[6*j+3] = uint16(4*j + 0)
		out.IndicesArray[6*j+4] = uint16(4*j + 2)
		out.IndicesArray[6*j+5] = uint16(4*j + 3)

	}

	outColours := []float32{}
	for _, c := range c.Colours {
		outColours = append(outColours, float32(c.R)/256, float32(c.G)/256, float32(c.B)/256, float32(c.A)/256)
		outColours = append(outColours, float32(c.R)/256, float32(c.G)/256, float32(c.B)/256, float32(c.A)/256)
		outColours = append(outColours, float32(c.R)/256, float32(c.G)/256, float32(c.B)/256, float32(c.A)/256)
		outColours = append(outColours, float32(c.R)/256, float32(c.G)/256, float32(c.B)/256, float32(c.A)/256)
	}
	out.ColourArray = outColours

	out.VCount = 24
	out.ICount = 36
	out.CCount = 24

	return out
}
