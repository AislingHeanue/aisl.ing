package view

import (
	"math"

	"github.com/AislingHeanue/aisling-codes/wasm-demo/games/life/controller"
	"github.com/AislingHeanue/aisling-codes/wasm-demo/util"
	"github.com/gowebapi/webapi/graphics/webgl"
)

// loader stuff
// TODO: load patterns + drop-down selection to pick a design from the life wiki (with some curated samples)
// TODO: make it so that the loader pattern name and comment shows up in a description box next to the app

// frontend stuff
// TODO: simulation size slider (with a reasonable default, which changes based on selected wiki design) (causes re-init if used manually)
// TODO: fps slider
// TODO: zoom slider
// TODO: step button
// TODO: random button, clear button

type LifeGame struct {
	lifeProgram    *webgl.Program
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
	DisplayFragmentSource string

	*controller.LifeContext
}

var _ util.Animator = &LifeGame{}

func (lg *LifeGame) Init(c *util.GameContext) {
	lg.createShaders(c, lg.VertexSource, lg.LifeFragmentSource)
	lg.createBuffers(c)
	lg.Random(c)

	lg.t = -1

	if lg.zoom(c) < 0 {
		panic("I refuse to create an infinite loop no thank you")
	}
	c.DX = 0 // default offset position of the grid is 0,0
	c.DY = 0
}

func (lg *LifeGame) InitListeners(c *util.GameContext) {
	util.RegisterListeners(c, lg.LifeContext, controller.LifeController(lg), controller.LifeActionHandler{})
}

func (lg *LifeGame) getDeadColour() (float32, float32, float32) {
	return float32(0.5 * (1 + math.Sin(2*math.Pi*float64(lg.t)/float64(lg.ColourPeriod)))),
		float32(0.5 * (1 + math.Sin((2*math.Pi/3)+2*math.Pi*float64(lg.t)/float64(lg.ColourPeriod)))),
		float32(0.5 * (1 + math.Sin((4*math.Pi/3)+2*math.Pi*float64(lg.t)/float64(lg.ColourPeriod))))
}

func (lg *LifeGame) attachAttributes(c *util.GameContext) {
	gl := c.GL

	gl.BindBuffer(webgl.ARRAY_BUFFER, lg.vertexBuffer)
	vPosition := gl.GetAttribLocation(lg.lifeProgram, "a_position")
	gl.VertexAttribPointer(uint(vPosition), 2, webgl.FLOAT, false, 0, 0)
	gl.EnableVertexAttribArray(uint(vPosition))

	gl.BindBuffer(webgl.ARRAY_BUFFER, lg.textureBuffer)
	tPosition := gl.GetAttribLocation(lg.lifeProgram, "a_tex_coord")
	gl.VertexAttribPointer(uint(tPosition), 2, webgl.FLOAT, false, 0, 0)
	gl.EnableVertexAttribArray(uint(tPosition))
	gl.UseProgram(lg.lifeProgram)

	decayLoc := gl.GetUniformLocation(lg.lifeProgram, "u_decay")
	gl.Uniform1f(decayLoc, 0.66/float32(lg.TrailLength))

	boundaryLoc := gl.GetUniformLocation(lg.lifeProgram, "u_boundary_loop")
	boundaryLoop := 0.
	if lg.Loop {
		boundaryLoop = 1.
	}
	gl.Uniform1f(boundaryLoc, float32(boundaryLoop))

	initialDecayLoc := gl.GetUniformLocation(lg.lifeProgram, "u_initial_decay")
	gl.Uniform1f(initialDecayLoc, 0.33)

	deadColourLoc := gl.GetUniformLocation(lg.lifeProgram, "u_new_dead_colour")

	r, g, b := lg.getDeadColour()
	gl.Uniform3f(deadColourLoc, r, g, b)

	gl.BindTexture(webgl.TEXTURE_2D, lg.readTexture)
	samplerLocation := gl.GetUniformLocation(lg.lifeProgram, "u_sampler")
	gl.Uniform1i(samplerLocation, 0)

	sizeLoc := gl.GetUniformLocation(lg.lifeProgram, "u_size")
	gl.Uniform2f(sizeLoc, float32(c.PixelsWidth), float32(c.PixelsHeight))

	pausedLoc := gl.GetUniformLocation(lg.lifeProgram, "u_paused")
	paused := 0.
	if lg.Paused {
		paused = 1.
	}
	gl.Uniform1f(pausedLoc, float32(paused))
}
