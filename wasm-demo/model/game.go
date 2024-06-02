package model

import (
	"image/color"
	"syscall/js"
	"time"

	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/markfarnan/go-canvas/canvas"
)

type AnimationFunc func(gc *draw2dimg.GraphicContext, c *GameContext) bool
type Projector interface {
	DoFirstRotation(v *Vector, anchor Vector)
	GetCoords(v Vector, height float64, width float64, anchor Vector) (float64, float64)
}
type GameContext struct {
	Height float64
	Width  float64
	T      time.Duration
	Colour color.RGBA

	Animation   AnimationFunc
	Projector   Projector
	Cvs         *canvas.Canvas2d
	CvsElement  js.Value
	Document    js.Value
	Window      js.Value
	Fps         float64
	RenderDelay time.Duration

	Cube Shape

	AngleX       float64
	AngleY       float64
	AnchorX      float64
	AnchorY      float64
	AnchorAngleX float64
	AnchorAngleY float64
	MouseDown    bool
}
