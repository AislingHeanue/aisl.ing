package graphics

import (
	"github.com/gowebapi/webapi"
	"github.com/gowebapi/webapi/dom"
	"github.com/gowebapi/webapi/graphics/webgl"
)

type RenderFunc func(*webgl.RenderingContext, *webgl.Program, *GameContext)

type BufferSet struct {
	Vertices *webgl.Buffer
	Indices  *webgl.Buffer
	Colours  *webgl.Buffer
	VCount   int
	ICount   int
	CCount   int
}

type Animator interface {
	Init(*GameContext)
	InitListeners(*GameContext)
	CreateShaders(*webgl.RenderingContext, *GameContext) *webgl.Program
	Render(*webgl.RenderingContext, *webgl.Program, *GameContext)
}

type GameContext struct {
	Height float32
	Width  float32
	T      float32

	Animator        Animator
	MaxTicks        int
	CvsElement      *dom.Element
	Document        *webapi.Document
	Window          *webapi.Window
	ResolutionScale float32

	Gl      *webgl.RenderingContext
	Program *webgl.Program

	Dimension int

	AngleX       float32
	AngleY       float32
	AnchorX      float32
	AnchorY      float32
	AnchorAngleX float32
	AnchorAngleY float32
	MouseDown    bool
}
