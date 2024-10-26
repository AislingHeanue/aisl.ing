package canvas

import (
	"syscall/js"

	"github.com/gowebapi/webapi"
	"github.com/gowebapi/webapi/dom"
	"github.com/gowebapi/webapi/graphics/webgl"
	"github.com/gowebapi/webapi/html/canvas"
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
	SecondaryCanvas *dom.Element
	ZoomCanvas      *dom.Element
	Document        *webapi.Document
	Window          *webapi.Window
	ResolutionScale float32

	CellHeight int
	CellWidth  int

	Height     float32
	Width      float32
	T          float32
	IntervalT  float32
	GL         *webgl.RenderingContext
	ZoomCtx    *canvas.CanvasRenderingContext2D
	DisplayCtx *canvas.CanvasRenderingContext2D
}

func Log(v js.Value) {
	js.Global().Get("console").Call("log", v)
}
