package maths

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
}

type RubiksCube [][][]*Cube

func NewRubiksCube(d int) RubiksCube {
	r := make(RubiksCube, d)
	for i := range r {
		r[i] = make([][]*Cube, d)
		for j := range r[i] {
			r[i][j] = make([]*Cube, d)
		}
	}

	return r
}

func (r RubiksCube) copy() RubiksCube {
	d := len(r)
	newR := NewRubiksCube(d)
	for x := range d {
		for y := range d {
			for z := range d {
				newR[x][y][z] = r[x][y][z]
			}
		}
	}

	return newR
}

func (r RubiksCube) flatten() []*Cube {
	cubes := []*Cube{}
	for x := range r {
		for y := range r[x] {
			for z := range r[y] {
				cubes = append(cubes, r[x][y][z])
			}
		}
	}

	return cubes
}

func (r RubiksCube) GroupBuffers() DrawShape {
	d := []DrawShape{}
	for _, v := range r.flatten() {
		if v != nil {
			d = append(d, v.GetBuffers())
		}

	}

	return GroupBuffers(d)
}

// var _ *CubeController = RubiksCube{}

func (r *RubiksCube) U(reverse bool) {
	newR := r.copy()

	for _, x := range []int{0, 1, 2} {
		for _, y := range []int{2} {
			for _, z := range []int{0, 1, 2} {
				newR[x][y][z] = (*r)[2-z][y][x]
				newR[x][y][z].RotateColoursY(1)
			}
		}
	}

	*r = newR
}

func (r *RubiksCube) D(reverse bool) {
	newR := r.copy()

	for _, x := range []int{0, 1, 2} {
		for _, y := range []int{0} {
			for _, z := range []int{0, 1, 2} {
				newR[x][y][z] = (*r)[z][y][2-x]
				newR[x][y][z].RotateColoursY(3)
			}
		}
	}

	*r = newR
}

func (r *RubiksCube) R(reverse bool) {
	newR := r.copy()

	for _, x := range []int{2} {
		for _, y := range []int{0, 1, 2} {
			for _, z := range []int{0, 1, 2} {
				newR[x][y][z] = (*r)[x][z][2-y]
				newR[x][y][z].RotateColoursX(1)
			}
		}
	}

	*r = newR
}

func (r *RubiksCube) L(reverse bool) {
	newR := r.copy()

	for _, x := range []int{0} {
		for _, y := range []int{0, 1, 2} {
			for _, z := range []int{0, 1, 2} {
				newR[x][y][z] = (*r)[x][2-z][y]
				newR[x][y][z].RotateColoursX(3)
			}
		}
	}

	*r = newR
}

func (r *RubiksCube) F(reverse bool) {
	newR := r.copy()

	for _, x := range []int{0, 1, 2} {
		for _, y := range []int{0, 1, 2} {
			for _, z := range []int{0} {
				newR[x][y][z] = (*r)[2-y][x][z]
				newR[x][y][z].RotateColoursZ(1)
			}
		}
	}

	*r = newR
}

func (r *RubiksCube) B(reverse bool) {
	newR := r.copy()

	for _, x := range []int{0, 1, 2} {
		for _, y := range []int{0, 1, 2} {
			for _, z := range []int{2} {
				newR[x][y][z] = (*r)[y][2-x][z]
				newR[x][y][z].RotateColoursZ(3)
			}
		}
	}

	*r = newR
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
