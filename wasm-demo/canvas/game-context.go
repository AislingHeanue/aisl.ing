package canvas

import (
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
	// returns CellHeight, CellWidth
	Dimensions() (int, int)
}

type GameContext struct {
	Animator Animator
	// the canvas actually being displayed to
	// optional intermediate canvas to facilitate zooming
	ZoomCanvas *dom.Element
	// the canvas with the webgl stuff
	RenderingCanvas *dom.Element
	Document        *webapi.Document
	Window          *webapi.Window
	ResolutionScale float32

	Square     bool
	Height     float32
	Width      float32
	T          float32
	IntervalT  float32
	GL         *webgl.RenderingContext
	ZoomCtx    *canvas.CanvasRenderingContext2D
	DisplayCtx *canvas.CanvasRenderingContext2D
}
