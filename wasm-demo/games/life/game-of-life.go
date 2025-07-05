package life

import (
	_ "embed"

	"github.com/AislingHeanue/aisling-codes/wasm-demo/games/life/model"
	"github.com/AislingHeanue/aisling-codes/wasm-demo/games/life/view"
)

//go:embed shaders/life.vert
var vShaderSource string

//go:embed shaders/life.frag
var fLifeShaderSource string

//go:embed shaders/display.frag
var fDisplayShaderSource string

//go:embed shaders/death.frag
var fDeathShaderSource string

func New(lo LifeOptions) *view.LifeGame {
	return &view.LifeGame{
		LifeContext: &model.LifeContext{
			CellHeight: lo.CellHeight,
			CellWidth:  lo.CellWidth,
			Zoom:       lo.Zoom,
			Tps:        lo.Tps,
			Loop:       lo.Loop,
		},
		VertexSource:          vShaderSource,
		LifeFragmentSource:    fLifeShaderSource,
		DeathFragmentSource:   fDeathShaderSource,
		DisplayFragmentSource: fDisplayShaderSource,
		TrailLength:           lo.TrailLength,
		ColourPeriod:          lo.ColourPeriod,
	}
}

type LifeOptions struct {
	CellHeight   int
	CellWidth    int
	Zoom         float32
	Tps          float32
	Loop         bool
	TrailLength  int
	ColourPeriod int
}

