package common

import (
	"fmt"
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
	RefreshBuffers(*GameContext)
}

type GameContext struct {
	Animator Animator
	Game     GameInfo
	// the canvas actually being displayed to
	// optional intermediate canvas to facilitate zooming
	ZoomCanvas *dom.Element
	// the canvas with the webgl stuff
	RenderingCanvas *dom.Element
	Document        *webapi.Document
	Window          *webapi.Window

	ResolutionScale float32
	SmoothImage     bool
	AutoSizePixels  bool
	Square          bool
	Height          float32
	Width           float32
	T               float32
	IntervalT       float32
	GL              *webgl.RenderingContext
	ZoomCtx         *canvas.CanvasRenderingContext2D
	DisplayCtx      *canvas.CanvasRenderingContext2D
	ZoomEnabled     bool
	PanningEnabled  bool
	Is3D            bool
	EdgeWidth       float32

	DX                  float32
	DY                  float32
	AnchorDX            float32
	AnchorDY            float32
	AnchorX             float32
	AnchorY             float32
	Zoom                float32
	MouseDown           bool
	Zooming             bool
	AnchorZoom          float32
	AnchorPinchDistance float32
	PixelsHeight        int
	PixelsWidth         int

	IsInitialised bool
}

func (c *GameContext) Log(value any) {
	js.Global().Get("console").Call("log", js.ValueOf(value))
}

func (c *GameContext) Logf(format string, args ...any) {
	js.Global().Get("console").Call("log", js.ValueOf(fmt.Sprintf(format+"\n", args...)))
}
