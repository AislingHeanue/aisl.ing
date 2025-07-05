package view

import (
	"fmt"

	"github.com/AislingHeanue/aisling-codes/wasm-demo/canvas"
	"github.com/AislingHeanue/aisling-codes/wasm-demo/games/life/controller"
	"github.com/AislingHeanue/aisling-codes/wasm-demo/games/life/model"
	"github.com/gowebapi/webapi/graphics/webgl"
)

// loader stuff
// TODO: load patterns + drop-down selection to pick a design from the life wiki (with some curated samples)

// frontend stuff
// TODO: simulation size slider (with a reasonable default, which changes based on selected wiki design) (causes re-init if used manually)
// TODO: fps slider
// TODO: zoom slider
// TODO: step button
// TODO: random button, clear button

type LifeGame struct {
	lifeProgram    *webgl.Program
	deathProgram   *webgl.Program
	displayProgram *webgl.Program

	vertexBuffer  *webgl.Buffer
	textureBuffer *webgl.Buffer

	vCount int

	writeFrameBuffer *webgl.Framebuffer
	readFrameBuffer  *webgl.Framebuffer
	readTexture      *webgl.Texture
	writeTexture     *webgl.Texture

	t                   int
	cumulativeIntervalT float32

	// configurable but not exposed to the frontend
	TrailLength  int
	ColourPeriod int

	VertexSource          string
	LifeFragmentSource    string
	DeathFragmentSource   string
	DisplayFragmentSource string

	*model.LifeContext
}

var _ canvas.Animator = &LifeGame{}

func (lg *LifeGame) Init(c *canvas.GameContext) {
	lg.createShaders(c)
	lg.createBuffers(c)
	lg.t = -1

	if lg.zoom(c) < 0 {
		panic("I refuse to create an infinite loop no thank you")
	}
	lg.DX = 0 // default offset position of the grid is 0,0
	lg.DY = 0
}

func (lg LifeGame) Dimensions() (int, int) {
	fmt.Println("Dimensions: width:", lg.CellWidth, "height:", lg.CellHeight)
	return lg.CellWidth, lg.CellHeight
}

func (lg *LifeGame) InitListeners(c *canvas.GameContext) {
	controller.InitListeners(c, lg.LifeContext, lg)
}

func (lg *LifeGame) Render(c *canvas.GameContext) {
	lg.cumulativeIntervalT += c.IntervalT
	for lg.cumulativeIntervalT > (1 / lg.Tps) {
		if lg.Paused {
			lg.deathFrame(c)
		} else {
			lg.t++
			lg.lifeFrame(c)
		}
		lg.cumulativeIntervalT -= 1 / lg.Tps
	}
	lg.drawToCanvas(c)
}
