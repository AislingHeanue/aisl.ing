package maths

import (
	"image/color"
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

func NewRubiksCube(d int) RubiksCube {
	r := make([][][]Cube, d)
	for i := range r {
		r[i] = make([][]Cube, d)
		for j := range r[i] {
			r[i][j] = make([]Cube, d)
		}
	}

	return RubiksCube{r, d}
}

func (r RubiksCube) Copy() RubiksCube {
	newR := NewRubiksCube(r.Dimension)
	for x := range r.Dimension {
		for y := range r.Dimension {
			for z := range r.Dimension {
				newR.Data[x][y][z] = r.Data[x][y][z]
			}
		}
	}

	return newR
}

func (r RubiksCube) Flatten() []Cube {
	cubes := []Cube{}
	for x := range r.Dimension {
		for y := range r.Dimension {
			for z := range r.Dimension {
				if r.isExternalCube(x, y, z) {
					cubes = append(cubes, r.Data[x][y][z])
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
