package model

import (
	"image/color"
	"syscall/js"
	"time"

	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/markfarnan/go-canvas/canvas"
)

type AnimationFunc func(gc *draw2dimg.GraphicContext, c GameContext) bool

type GameContext struct {
	Height float64
	Width  float64
	T      time.Duration
	Colour color.RGBA

	Cvs         *canvas.Canvas2d
	Document    js.Value
	Window      js.Value
	Fps         float64
	RenderDelay time.Duration
}
