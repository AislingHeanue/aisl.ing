package controller

import (
	"syscall/js"

	"github.com/gowebapi/webapi"
	"github.com/gowebapi/webapi/dom"
	"github.com/gowebapi/webapi/graphics/webgl"
)

type RenderFunc func(*webgl.RenderingContext, *webgl.Program, *GameContext)

type Animator interface {
	Init(*GameContext)
	CreateShaders(*webgl.RenderingContext, *GameContext) *webgl.Program
	Render(*webgl.RenderingContext, *webgl.Program, *GameContext)
	QueueEvent(string, bool) bool
	Shuffle()
}

type GameContext struct {
	Height float32
	Width  float32
	T      float32

	Animator        Animator
	TurnFrames      int
	CvsElement      *dom.Element
	Document        *webapi.Document
	Window          *webapi.Window
	ResolutionScale float32

	CubeDimension int

	AngleX       float32
	AngleY       float32
	AnchorX      float32
	AnchorY      float32
	AnchorAngleX float32
	AnchorAngleY float32
	MouseDown    bool
}

func Log(v js.Value) {
	js.Global().Get("console").Call("log", v)
}
