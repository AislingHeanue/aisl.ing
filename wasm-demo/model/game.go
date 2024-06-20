package model

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
	// CreateBuffers(*webgl.RenderingContext, *GameContext)
	CreateShaders(*webgl.RenderingContext, *GameContext) *webgl.Program
	Render(*webgl.RenderingContext, *webgl.Program, *GameContext)
	// RefreshCoords(*GameContext)
	// IsActive() bool
	// IsRedrawRequired() bool
}

type GameContext struct {
	Height float64
	Width  float64
	T      float64

	Animator        Animator
	CvsElement      *dom.Element
	Document        *webapi.Document
	Window          *webapi.Window
	ResolutionScale float64

	Gl      *webgl.RenderingContext
	Program *webgl.Program

	Dimension int

	AngleX       float64
	AngleY       float64
	AnchorX      float64
	AnchorY      float64
	AnchorAngleX float64
	AnchorAngleY float64
	MouseDown    bool
}
