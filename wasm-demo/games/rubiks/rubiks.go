package rubiks

import (
	"github.com/AislingHeanue/aisling-codes/wasm-demo/games/rubiks/model"
	"github.com/AislingHeanue/aisling-codes/wasm-demo/games/rubiks/view"
)

func New(cco CubeCubeOptions) *view.CubeRenderer {
	return &view.CubeRenderer{CubeCubeContext: &model.CubeCubeContext{
		TurnFrames:    cco.TurnFrames,
		CubeDimension: cco.Dimension,
	}}
}

type CubeCubeOptions struct {
	TurnFrames int
	Dimension  int
}
