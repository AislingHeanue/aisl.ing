package maths

import (
	"fmt"
	"math"
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
}

type RubiksCube map[string]*Cube

func (r RubiksCube) GroupBuffers() DrawShape {
	d := []DrawShape{}
	for _, v := range r {
		if v != nil {
			d = append(d, v.GetBuffers())
		}

	}

	return GroupBuffers(d)
}

// var _ *CubeController = RubiksCube{}

func (r *RubiksCube) U(reverse bool) {
	cubes := r.getSubset(func(c Coordinate) bool {
		return c.Y() == 2
	})
	fmt.Println("u")

	for _, cube := range cubes {
		cube.AngleY = mod(cube.AngleY-math.Pi/2, 2*math.Pi)
		x, z := cube.CentrePoint.X(), cube.CentrePoint.Z()
		cube.CentrePoint[0] = z
		cube.CentrePoint[2] = -x
	}
}

// func (r *RubiksCube) D(reverse bool) {
// 	panic("not implemented") // TODO: Implement
// }

// func (r *RubiksCube) F(reverse bool) {
// 	panic("not implemented") // TODO: Implement
// }

// func (r *RubiksCube) B(reverse bool) {
// 	panic("not implemented") // TODO: Implement
// }

// func (r *RubiksCube) L(reverse bool) {
// 	panic("not implemented") // TODO: Implement
// }

// func (r *RubiksCube) R(reverse bool) {
// 	panic("not implemented") // TODO: Implement
// }

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

func (r RubiksCube) getSubset(conditions ...func(c Coordinate) bool) []*Cube {
	cubes := []*Cube{}
	for k, v := range r {
		if len(parseCoord(k)) != 3 {
			panic(fmt.Sprintf("parsed coordinate %q has a length not equal to 3", k))
		}
		pass := true
		for _, condition := range conditions {
			if !condition(parseCoord(k)) {
				pass = false
			}
		}
		if pass {
			cubes = append(cubes, v)
		}
	}

	return cubes
}

func mod(in float32, max float32) float32 {
	return in - float32(math.Floor(float64(in)/float64(max)))*max
}
