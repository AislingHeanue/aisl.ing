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
	InitListeners(*GameContext)
	Render(*GameContext)
}

type GameContext struct {
	Animator        Animator
	CvsElement      *dom.Element
	Document        *webapi.Document
	Window          *webapi.Window
	ResolutionScale float32
	Height          float32
	Width           float32
	T               float32
	GL              *webgl.RenderingContext
	Program         *webgl.Program
}

func Log(v js.Value) {
	js.Global().Get("console").Call("log", v)
}
