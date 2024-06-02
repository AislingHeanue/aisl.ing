package animation

import (
	"math"

	"github.com/AislingHeanue/aisling-codes/wasm-demo/model"
)

// DA RULES FOR PROJECTIONS
// Take in 3 coordinates and give an x and a y to draw in the grid.
// in gc notation, x,y = 0,0 is the top left corner. All of these projections
// should make it so that 0,0,0 is the bottom left corner (or analogous to it).
// For this reason, projection functions are given the height of the canvas, so
// they can flip the coordinate themselves

type OrthographicProjector struct{}

var _ model.Projector = OrthographicProjector{}

func (o OrthographicProjector) DoFirstRotation(_ *model.Vector, _ model.Vector) {}

func (o OrthographicProjector) GetCoords(v model.Vector, height float64, width float64, anchor model.Vector) (float64, float64) {
	x := v.X() + 1/math.Sqrt2*v.Z()
	y := v.Y() + 1/math.Sqrt2*v.Z()
	return x, height - y
}

type IsometricProjector struct{}

var _ model.Projector = IsometricProjector{}

func (i IsometricProjector) DoFirstRotation(v *model.Vector, anchor model.Vector) {
	v.Rotate(anchor, -math.Pi/4, model.Y)

}

func (i IsometricProjector) GetCoords(v model.Vector, height float64, width float64, anchor model.Vector) (float64, float64) {
	v.Rotate(anchor, -math.Asin(math.Tan(math.Pi/6)), model.X)
	return v[0], height - v[1]
}

type PerspectiveProjector struct{}

var _ model.Projector = PerspectiveProjector{}

func (i PerspectiveProjector) DoFirstRotation(v *model.Vector, anchor model.Vector) {
	v.Rotate(anchor, -math.Pi/4, model.Y)

}

func (i PerspectiveProjector) GetCoords(v model.Vector, height float64, width float64, anchor model.Vector) (float64, float64) {
	v.Rotate(anchor, -math.Asin(math.Tan(math.Pi/6)), model.X)
	vanishingPoint := model.Vector{width / 2, height / 2, 3 * height}

	if v[2] < vanishingPoint[2] {
		v[0] += (vanishingPoint[0] - v[0]) * (v[2] / vanishingPoint[2])
		v[1] += (vanishingPoint[1] - v[1]) * (v[2] / vanishingPoint[2])
	} else {
		return height / 2, height / 2
	}
	return v[0], height - v[1]
}
