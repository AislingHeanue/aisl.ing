package model

import (
	"image/color"
	"syscall/js"
	"time"

	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/markfarnan/go-canvas/canvas"
)

type Animator interface {
	Init(*GameContext)
	Render(*draw2dimg.GraphicContext, *GameContext) bool
	IsActive() bool
}

type Projector interface {
	GetCoords(v Vector, height float64, width float64, angleX float64, angleY float64, anchor Vector) *Vector
}

type GameContext struct {
	Height float64
	Width  float64
	T      time.Duration
	Colour color.RGBA

	Animator    Animator
	Projector   Projector
	Cvs         *canvas.Canvas2d
	CvsElement  js.Value
	Document    js.Value
	Window      js.Value
	Fps         float64
	RenderDelay time.Duration

	Cube      *Cube
	Dimension int

	AngleX       float64
	AngleY       float64
	AnchorX      float64
	AnchorY      float64
	AnchorAngleX float64
	AnchorAngleY float64
	MouseDown    bool
}
