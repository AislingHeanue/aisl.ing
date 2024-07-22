package rubiks

import (
	"github.com/AislingHeanue/aisling-codes/wasm-demo/maths"
)

type CubeController interface {
	U(bool)
	D(bool)
	F(bool)
	B(bool)
	R(bool)
	L(bool)
	S(bool)
	E(bool)
	M(bool)
	X(bool)
	Y(bool)
	Z(bool)
	RefreshBuffers()
}

type turnInfo struct {
	xSelector cubeSelector
	ySelector cubeSelector
	zSelector cubeSelector
	reverse   bool
	axis      maths.Axis
}

var turnMap = map[string]turnInfo{
	"u": {ALL, LAST, ALL, false, maths.Y},
	"d": {ALL, FIRST, ALL, true, maths.Y},
	"r": {LAST, ALL, ALL, false, maths.X},
	"l": {FIRST, ALL, ALL, true, maths.X},
	"f": {ALL, ALL, FIRST, false, maths.Z},
	"b": {ALL, ALL, LAST, true, maths.Z},
	"m": {MIDDLE, ALL, ALL, true, maths.X},
	"e": {ALL, MIDDLE, ALL, true, maths.Y},
	"s": {ALL, ALL, MIDDLE, false, maths.Z},
	"x": {ALL, ALL, ALL, false, maths.X},
	"y": {ALL, ALL, ALL, false, maths.Y},
	"z": {ALL, ALL, ALL, false, maths.Z},
}

var _ CubeController = &RubiksCube{}

type RubiksCube struct {
	data        [][][]*maths.Cube
	dimension   int
	bufferStale bool
}

func NewRubiksCube(d int) RubiksCube {
	r := make([][][]*maths.Cube, d)
	for i := range r {
		r[i] = make([][]*maths.Cube, d)
		for j := range r[i] {
			r[i][j] = make([]*maths.Cube, d)
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

func (r RubiksCube) flatten() []*maths.Cube {
	cubes := []*maths.Cube{}
	for x := range r.dimension {
		for y := range r.dimension {
			for z := range r.dimension {
				cubes = append(cubes, r.data[x][y][z])
			}
		}
	}

	return cubes
}

type cubeSelector int

const (
	FIRST cubeSelector = iota
	MIDDLE
	LAST
	ALL
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
	case MIDDLE:
		out := make([]int, r.dimension-2)
		for i := range out {
			out[i] = i + 1
		}

		return out
	case LAST:
		return []int{r.dimension - 1}
	case ALL:
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
		if v != nil {
			d = append(d, GetBuffers(*v))
		}

	}

	return GroupBuffers(d)
}

// var _ *CubeController = RubiksCube{}

func (r *RubiksCube) U(reverse bool) {
	info := turnMap["u"]
	times := 1
	if reverse {
		times = 3
	}
	for range times {
		newR := r.copy()

		for _, coord := range r.getCubeSubset(info.xSelector, info.ySelector, info.zSelector) {
			x := coord[0]
			y := coord[1]
			z := coord[2]
			newR.data[x][y][z] = r.data[r.dimension-1-z][y][x]
			newR.data[x][y][z].RotateColoursY(info.reverse)
		}

		*r = newR
	}
}

func (r *RubiksCube) D(reverse bool) {
	info := turnMap["d"]
	times := 1
	if reverse {
		times = 3
	}
	for range times {
		newR := r.copy()

		for _, coord := range r.getCubeSubset(info.xSelector, info.ySelector, info.zSelector) {
			x := coord[0]
			y := coord[1]
			z := coord[2]
			newR.data[x][y][z] = r.data[z][y][r.dimension-1-x]
			newR.data[x][y][z].RotateColoursY(info.reverse)
		}

		*r = newR
	}
}

func (r *RubiksCube) R(reverse bool) {
	info := turnMap["r"]
	times := 1
	if reverse {
		times = 3
	}
	for range times {
		newR := r.copy()

		for _, coord := range r.getCubeSubset(info.xSelector, info.ySelector, info.zSelector) {
			x := coord[0]
			y := coord[1]
			z := coord[2]
			// fmt.Println(x, y, z)
			newR.data[x][y][z] = r.data[x][z][r.dimension-1-y]
			newR.data[x][y][z].RotateColoursX(info.reverse)
		}

		*r = newR
	}
}

func (r *RubiksCube) L(reverse bool) {
	info := turnMap["l"]
	times := 1
	if reverse {
		times = 3
	}
	for range times {
		newR := r.copy()

		for _, coord := range r.getCubeSubset(info.xSelector, info.ySelector, info.zSelector) {
			x := coord[0]
			y := coord[1]
			z := coord[2]
			newR.data[x][y][z] = r.data[x][r.dimension-1-z][y]
			newR.data[x][y][z].RotateColoursX(info.reverse)
		}

		*r = newR
	}
}

func (r *RubiksCube) F(reverse bool) {
	info := turnMap["f"]
	times := 1
	if reverse {
		times = 3
	}
	for range times {
		newR := r.copy()

		for _, coord := range r.getCubeSubset(info.xSelector, info.ySelector, info.zSelector) {
			x := coord[0]
			y := coord[1]
			z := coord[2]
			newR.data[x][y][z] = r.data[r.dimension-1-y][x][z]
			newR.data[x][y][z].RotateColoursZ(info.reverse)
		}

		*r = newR
	}
}

func (r *RubiksCube) B(reverse bool) {
	info := turnMap["b"]
	times := 1
	if reverse {
		times = 3
	}
	for range times {
		newR := r.copy()

		for _, coord := range r.getCubeSubset(info.xSelector, info.ySelector, info.zSelector) {
			x := coord[0]
			y := coord[1]
			z := coord[2]
			newR.data[x][y][z] = r.data[y][r.dimension-1-x][z]
			newR.data[x][y][z].RotateColoursZ(info.reverse)
		}

		*r = newR
	}
}

func (r *RubiksCube) M(reverse bool) {
	info := turnMap["m"]
	times := 1
	if reverse {
		times = 3
	}
	for range times {
		newR := r.copy()

		for _, coord := range r.getCubeSubset(info.xSelector, info.ySelector, info.zSelector) {
			x := coord[0]
			y := coord[1]
			z := coord[2]
			newR.data[x][y][z] = r.data[x][r.dimension-1-z][y]
			newR.data[x][y][z].RotateColoursX(info.reverse)
		}

		*r = newR
	}
}

func (r *RubiksCube) E(reverse bool) {
	info := turnMap["e"]
	times := 1
	if reverse {
		times = 3
	}
	for range times {
		newR := r.copy()

		for _, coord := range r.getCubeSubset(info.xSelector, info.ySelector, info.zSelector) {
			x := coord[0]
			y := coord[1]
			z := coord[2]
			newR.data[x][y][z] = r.data[z][y][r.dimension-1-x]
			newR.data[x][y][z].RotateColoursY(info.reverse)
		}

		*r = newR
	}
}

func (r *RubiksCube) S(reverse bool) {
	info := turnMap["s"]
	times := 1
	if reverse {
		times = 3
	}
	for range times {
		newR := r.copy()

		for _, coord := range r.getCubeSubset(info.xSelector, info.ySelector, info.zSelector) {
			x := coord[0]
			y := coord[1]
			z := coord[2]
			newR.data[x][y][z] = r.data[r.dimension-1-y][x][z]
			newR.data[x][y][z].RotateColoursZ(false)
		}

		*r = newR
	}
}

func (r *RubiksCube) X(reverse bool) {
	info := turnMap["x"]
	times := 1
	if reverse {
		times = 3
	}
	for range times {
		newR := r.copy()

		for _, coord := range r.getCubeSubset(info.xSelector, info.ySelector, info.zSelector) {
			x := coord[0]
			y := coord[1]
			z := coord[2]
			newR.data[x][y][z] = r.data[x][z][r.dimension-1-y]
			newR.data[x][y][z].RotateColoursX(false)
		}

		*r = newR
	}
}

func (r *RubiksCube) Y(reverse bool) {
	info := turnMap["y"]
	times := 1
	if reverse {
		times = 3
	}
	for range times {
		newR := r.copy()

		for _, coord := range r.getCubeSubset(info.xSelector, info.ySelector, info.zSelector) {
			x := coord[0]
			y := coord[1]
			z := coord[2]
			newR.data[x][y][z] = r.data[r.dimension-1-z][y][x]
			newR.data[x][y][z].RotateColoursY(false)
		}

		*r = newR
	}
}

func (r *RubiksCube) Z(reverse bool) {
	info := turnMap["z"]
	times := 1
	if reverse {
		times = 3
	}
	for range times {
		newR := r.copy()

		for _, coord := range r.getCubeSubset(info.xSelector, info.ySelector, info.zSelector) {
			x := coord[0]
			y := coord[1]
			z := coord[2]
			newR.data[x][y][z] = r.data[r.dimension-1-y][x][z]
			newR.data[x][y][z].RotateColoursZ(false)
		}

		*r = newR
	}
}
