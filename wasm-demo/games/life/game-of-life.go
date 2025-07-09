package life

import (
	_ "embed"

	"github.com/AislingHeanue/aisling-codes/wasm-demo/games/life/controller"
	"github.com/AislingHeanue/aisling-codes/wasm-demo/games/life/view"
)

//go:embed shaders/life.vert
var vShaderSource string

//go:embed shaders/life.frag
var fLifeShaderSource string

//go:embed shaders/display.frag
var fDisplayShaderSource string

func New(lo LifeOptions) *view.LifeGame {
	return &view.LifeGame{
		LifeContext: &controller.LifeContext{
			Tps:  lo.Tps,
			Loop: lo.Loop,
		},
		VertexSource:          vShaderSource,
		LifeFragmentSource:    fLifeShaderSource,
		DisplayFragmentSource: fDisplayShaderSource,
		TrailLength:           lo.TrailLength,
		ColourPeriod:          lo.ColourPeriod,
	}
}

type LifeOptions struct {
	Tps          float32
	Loop         bool
	TrailLength  int
	ColourPeriod int
}
