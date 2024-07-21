package rubiks

import (
	"github.com/AislingHeanue/aisling-codes/wasm-demo/maths"
)

type CubeController interface {
	U(bool)
	D(bool)
	F(bool)
	B(bool)
	L(bool)
	R(bool)
	X(bool)
	Y(bool)
	Z(bool)
	S(bool)
	E(bool)
	M(bool)
	RefreshBuffers()
}

var _ CubeController = &RubiksCube{}

type RubiksCube struct {
	data        [][][]*maths.Cube
	dimension   int
	bufferStale bool
	// animationHandler animationHandler
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
	times := 1
	if reverse {
		times = 3
	}
	for range times {
		newR := r.copy()

		for x := range r.dimension {
			y := r.dimension - 1
			for z := range r.dimension {
				newR.data[x][y][z] = r.data[r.dimension-1-z][y][x]
				newR.data[x][y][z].RotateColoursY(false)
			}
		}

		*r = newR
	}
}

func (r *RubiksCube) D(reverse bool) {
	times := 1
	if reverse {
		times = 3
	}
	for range times {
		newR := r.copy()

		for x := range r.dimension {
			y := 0
			for z := range r.dimension {
				newR.data[x][y][z] = r.data[z][y][r.dimension-1-x]
				newR.data[x][y][z].RotateColoursY(true)
			}
		}

		*r = newR
	}
}

func (r *RubiksCube) R(reverse bool) {
	times := 1
	if reverse {
		times = 3
	}
	for range times {
		newR := r.copy()

		x := r.dimension - 1
		for y := range r.dimension {
			for z := range r.dimension {
				newR.data[x][y][z] = r.data[x][z][r.dimension-1-y]
				newR.data[x][y][z].RotateColoursX(false)
			}
		}

		*r = newR
	}
}

func (r *RubiksCube) L(reverse bool) {
	times := 1
	if reverse {
		times = 3
	}
	for range times {
		newR := r.copy()

		x := 0
		for y := range r.dimension {
			for z := range r.dimension {
				newR.data[x][y][z] = r.data[x][r.dimension-1-z][y]
				newR.data[x][y][z].RotateColoursX(true)
			}
		}

		*r = newR
	}
}

func (r *RubiksCube) F(reverse bool) {
	times := 1
	if reverse {
		times = 3
	}
	for range times {
		newR := r.copy()

		for x := range r.dimension {
			for y := range r.dimension {
				z := 0
				newR.data[x][y][z] = r.data[r.dimension-1-y][x][z]
				newR.data[x][y][z].RotateColoursZ(false)
			}
		}

		*r = newR
	}
}

func (r *RubiksCube) B(reverse bool) {
	times := 1
	if reverse {
		times = 3
	}
	for range times {
		newR := r.copy()

		for x := range r.dimension {
			for y := range r.dimension {
				z := r.dimension - 1
				newR.data[x][y][z] = r.data[y][r.dimension-1-x][z]
				newR.data[x][y][z].RotateColoursZ(true)
			}
		}

		*r = newR
	}
}

func (r *RubiksCube) M(reverse bool) {
	times := 1
	if reverse {
		times = 3
	}
	for range times {
		newR := r.copy()

		for x := range r.dimension - 2 {
			for y := range r.dimension {
				for z := range r.dimension {
					if r.isExternalCube(x+1, y, z) {
						newR.data[x+1][y][z] = r.data[x+1][r.dimension-1-z][y]
						newR.data[x+1][y][z].RotateColoursX(true)
					}
				}
			}
		}

		*r = newR
	}
}

func (r *RubiksCube) E(reverse bool) {
	times := 1
	if reverse {
		times = 3
	}
	for range times {
		newR := r.copy()

		for x := range r.dimension {
			for y := range r.dimension - 2 {
				for z := range r.dimension {
					if r.isExternalCube(x, y+1, z) {
						newR.data[x][y+1][z] = r.data[z][y+1][r.dimension-1-x]
						newR.data[x][y+1][z].RotateColoursY(true)
					}
				}
			}
		}

		*r = newR
	}
}

func (r *RubiksCube) S(reverse bool) {
	times := 1
	if reverse {
		times = 3
	}
	for range times {
		newR := r.copy()

		for x := range r.dimension {
			for y := range r.dimension {
				for z := range r.dimension - 2 {
					if r.isExternalCube(x, y, z+1) {
						newR.data[x][y][z+1] = r.data[r.dimension-1-y][x][z+1]
						newR.data[x][y][z+1].RotateColoursZ(false)
					}
				}
			}
		}

		*r = newR
	}
}

func (r *RubiksCube) X(reverse bool) {
	times := 1
	if reverse {
		times = 3
	}
	for range times {
		newR := r.copy()

		for x := range r.dimension {
			for y := range r.dimension {
				for z := range r.dimension {
					if r.isExternalCube(x, y, z) {
						newR.data[x][y][z] = r.data[x][z][r.dimension-1-y]
						newR.data[x][y][z].RotateColoursX(false)
					}
				}
			}
		}

		*r = newR
	}
}

func (r *RubiksCube) Y(reverse bool) {
	times := 1
	if reverse {
		times = 3
	}
	for range times {
		newR := r.copy()

		for x := range r.dimension {
			for y := range r.dimension {
				for z := range r.dimension {
					if r.isExternalCube(x, y, z) {
						newR.data[x][y][z] = r.data[r.dimension-1-z][y][x]
						newR.data[x][y][z].RotateColoursY(false)
					}
				}
			}
		}

		*r = newR
	}
}

func (r *RubiksCube) Z(reverse bool) {
	times := 1
	if reverse {
		times = 3
	}
	for range times {
		newR := r.copy()

		for x := range r.dimension {
			for y := range r.dimension {
				for z := range r.dimension {
					if r.isExternalCube(x, y, z) {
						newR.data[x][y][z] = r.data[r.dimension-1-y][x][z]
						newR.data[x][y][z].RotateColoursZ(false)
					}
				}
			}
		}

		*r = newR
	}
}

// func (r *RubiksCube) X(reverse bool) {
// 	panic("not implemented") // TODO: Implement
// }

// func (r *RubiksCube) Y(reverse bool) {
// 	panic("not implemented") // TODO: Implement
// }

// func (r *RubiksCube) Z(reverse bool) {
// 	panic("not implemented") // TODO: Implement
// }

// func (r *RubiksCube) S(reverse bool) {
// 	panic("not implemented") // TODO: Implement
// }

// func (r *RubiksCube) E(reverse bool) {
// 	panic("not implemented") // TODO: Implement
// }

// func (r *RubiksCube) M(reverse bool) {
// 	panic("not implemented") // TODO: Implement
// }
