package rubiks

import (
	_ "embed"
	"github.com/AislingHeanue/aisling-codes/wasm-demo/games/rubiks/model"
	"github.com/AislingHeanue/aisling-codes/wasm-demo/games/rubiks/view"
)

//go:embed shaders/rubiks.vert
var vertexSource string

//go:embed shaders/rubiks.frag
var fragmentSource string

func New(cco CubeCubeOptions) *view.CubeRenderer {
	sideLength := cco.TotalSideLength / ((1+cco.GapProportion)*float32(cco.Dimension) - cco.GapProportion)
	sideLengthWithGap := sideLength + cco.GapProportion*sideLength
	origin := model.Point{0, 0, 0}
	cubes := model.NewRubiksCube(cco.Dimension, origin, sideLength, cco.TotalSideLength, sideLengthWithGap)

	return &view.CubeRenderer{
		CubeCubeContext: &model.CubeCubeContext{
			AnimationHandler: &model.RubiksAnimationHandler{
				RubiksCube: &cubes,
				MaxTime:    cco.TurnSeconds,
			},
		},
		VertexSource:   vertexSource,
		FragmentSource: fragmentSource,

		TotalSideLength:   cco.TotalSideLength,
		GapProportion:     cco.GapProportion,
		SideLength:        sideLength,
		SideLengthWithGap: sideLengthWithGap,
		Origin:            model.Point{0, 0, 0},
		Dimension:         cco.Dimension,
		TurnSeconds:       cco.TurnSeconds,
	}
}

type CubeCubeOptions struct {
	TurnSeconds     float32
	Dimension       int
	TotalSideLength float32
	GapProportion   float32
}
