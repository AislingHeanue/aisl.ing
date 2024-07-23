package rubiks

import (
	"image/color"

	"github.com/AislingHeanue/aisling-codes/wasm-demo/maths"
)

type CubeController interface {
	Turn(string, bool)
	RefreshBuffers()
}

type turnInfo struct {
	xSelector         cubeSelector
	ySelector         cubeSelector
	zSelector         cubeSelector
	reverse           bool
	axis              maths.Axis
	allowedConcurrent []string
}

var turnMap = map[string]turnInfo{
	"u": {EVERY, FINAL, EVERY, true, maths.Y, []string{"d", "e"}},
	"d": {EVERY, FIRST, EVERY, false, maths.Y, []string{"d", "u"}},
	"r": {FINAL, EVERY, EVERY, true, maths.X, []string{"u", "e"}},
	"l": {FIRST, EVERY, EVERY, false, maths.X, []string{}},
	"f": {EVERY, EVERY, FIRST, false, maths.Z, []string{"m", "l"}},
	"b": {EVERY, EVERY, FINAL, true, maths.Z, []string{"l", "r"}},
	"m": {INNER, EVERY, EVERY, false, maths.X, []string{"r", "m"}},
	"e": {EVERY, INNER, EVERY, false, maths.Y, []string{}},
	"s": {EVERY, EVERY, INNER, false, maths.Z, []string{"s", "b"}},
	"x": {EVERY, EVERY, EVERY, true, maths.X, []string{"b", "f"}},
	"y": {EVERY, EVERY, EVERY, true, maths.Y, []string{"f", "s"}},
	"z": {EVERY, EVERY, EVERY, false, maths.Z, []string{}},
}

var _ CubeController = &RubiksCube{}

type RubiksCube struct {
	data        [][][]maths.Cube
	dimension   int
	bufferStale bool
}

func NewRubiksCube(d int) RubiksCube {
	r := make([][][]maths.Cube, d)
	for i := range r {
		r[i] = make([][]maths.Cube, d)
		for j := range r[i] {
			r[i][j] = make([]maths.Cube, d)
		}
	}

	return RubiksCube{r, d, true}
}

func (r RubiksCube) copy() RubiksCube {
	newR := NewRubiksCube(r.dimension)
	for x := range r.dimension {
		for y := range r.dimension {
			for z := range r.dimension {
				newR.data[x][y][z] = r.data[x][y][z]
			}
		}
	}

	return newR
}

func (r RubiksCube) flatten() []maths.Cube {
	cubes := []maths.Cube{}
	for x := range r.dimension {
		for y := range r.dimension {
			for z := range r.dimension {
				if r.isExternalCube(x, y, z) {
					cubes = append(cubes, r.data[x][y][z])
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
		out := make([]int, r.dimension-2)
		for i := range out {
			out[i] = i + 1
		}

		return out
	case FINAL:
		return []int{r.dimension - 1}
	case EVERY:
		out := make([]int, r.dimension)
		for i := range out {
			out[i] = i
		}

		return out
	default:
		return []int{}
	}
}

func (r RubiksCube) isExternalCube(x, y, z int) bool {
	xCond := x == 0 || x == r.dimension-1
	yCond := y == 0 || y == r.dimension-1
	zCond := z == 0 || z == r.dimension-1

	return xCond || yCond || zCond
}

func (r *RubiksCube) RefreshBuffers() {
	r.bufferStale = true
}

func (r RubiksCube) GroupBuffers() DrawShape {
	d := []DrawShape{}
	for _, v := range r.flatten() {
		d = append(d, GetBuffers(v))
	}

	return GroupBuffers(d)
}

func (r RubiksCube) getPermutedCubeColours(x, y, z int, axis maths.Axis, reverse bool) []color.RGBA {
	switch axis {
	case maths.X:
		if reverse {
			return r.data[x][z][r.dimension-1-y].Colours
		} else {
			return r.data[x][r.dimension-1-z][y].Colours
		}
	case maths.Y:
		if reverse {
			return r.data[r.dimension-1-z][y][x].Colours
		} else {
			return r.data[z][y][r.dimension-1-x].Colours
		}
	case maths.Z:
		if reverse {
			return r.data[y][r.dimension-1-x][z].Colours
		} else {
			return r.data[r.dimension-1-y][x][z].Colours
		}
	default:
		return []color.RGBA{}
	}
}

// var _ *CubeController = RubiksCube{}

func (r *RubiksCube) Turn(face string, reverse bool) {
	info, ok := turnMap[face]
	if !ok {
		return
	}
	times := 1
	if reverse {
		times = 3
	}
	for range times {
		newR := r.copy()

		for _, coord := range r.getCubeSubset(info.xSelector, info.ySelector, info.zSelector) {
			newR.data[coord[0]][coord[1]][coord[2]].Colours = r.getPermutedCubeColours(coord[0], coord[1], coord[2], info.axis, info.reverse)
			newR.data[coord[0]][coord[1]][coord[2]].RotateColours(info.reverse, info.axis)
		}

		*r = newR
	}
}
