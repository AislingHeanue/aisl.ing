package life

import (
	_ "embed"
	"math"
	"math/rand"

	common "github.com/AislingHeanue/aisling-codes/wasm-demo/common"
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

//go:embed shaders/life.vert
var vShaderSource string

//go:embed shaders/life.frag
var fLifeShaderSource string

func New(lo LifeOptions) *LifeGame {
	return &LifeGame{
		LifeContext: &LifeContext{
			Tps:          lo.Tps,
			Loop:         lo.Loop,
			TrailLength:  lo.TrailLength,
			ColourPeriod: lo.ColourPeriod,
		},
	}
}

type LifeOptions struct {
	Tps          float32
	Loop         bool
	TrailLength  int
	ColourPeriod int
}

type LifeGame struct {
	Parent *common.ShaderGame
	common.DefaultGame
	*LifeContext
}

var _ common.GameInfo = &LifeGame{}

func (lg *LifeGame) PostSetup(c *common.GameContext) {
	lg.Random(c)

	lg.T = -1
}

func (lg *LifeGame) InitListeners(c *common.GameContext) {
	common.RegisterListeners(c, lg.LifeContext, LifeController(lg), LifeActionHandler{})
}

func (lg *LifeGame) GetFragmentSource() string {
	return fLifeShaderSource
}

func (lg *LifeGame) GetVertexSource() string {
	return vShaderSource
}

func (lg *LifeGame) GetTps() float32 {
	return lg.Tps
}

func (lg *LifeGame) SetParent(parent *common.ShaderGame) {
	lg.Parent = parent
}

func (lg *LifeGame) Tick(c *common.GameContext) bool {
	lg.T++

	return false
}

func (lg *LifeGame) CanRunBetweenFrames() bool {
	return true
}

func (lg *LifeGame) AttachAttributes(c *common.GameContext, program *webgl.Program, writeBuffer, readBuffer *webgl.Buffer, samplerTexture *webgl.Texture) {
	gl := c.GL

	gl.BindBuffer(webgl.ARRAY_BUFFER, writeBuffer)
	vPosition := gl.GetAttribLocation(program, "aPosition")
	gl.VertexAttribPointer(uint(vPosition), 2, webgl.FLOAT, false, 0, 0)
	gl.EnableVertexAttribArray(uint(vPosition))

	gl.BindBuffer(webgl.ARRAY_BUFFER, readBuffer)
	tPosition := gl.GetAttribLocation(program, "aTexCoord")
	gl.VertexAttribPointer(uint(tPosition), 2, webgl.FLOAT, false, 0, 0)
	gl.EnableVertexAttribArray(uint(tPosition))

	decayLoc := gl.GetUniformLocation(program, "uDecay")
	gl.Uniform1f(decayLoc, 0.66/float32(lg.TrailLength))

	boundaryLoc := gl.GetUniformLocation(program, "uBoundaryLoop")
	boundaryLoop := 0.
	if lg.Loop {
		boundaryLoop = 1.
	}
	gl.Uniform1f(boundaryLoc, float32(boundaryLoop))

	initialDecayLoc := gl.GetUniformLocation(program, "uInitialDecay")
	gl.Uniform1f(initialDecayLoc, 0.33)

	deadColourLoc := gl.GetUniformLocation(program, "uNewDeadColour")

	r, g, b := lg.getDeadColour()
	gl.Uniform3f(deadColourLoc, r, g, b)

	gl.BindTexture(webgl.TEXTURE_2D, samplerTexture)
	samplerLocation := gl.GetUniformLocation(program, "uSampler")
	gl.Uniform1i(samplerLocation, 0)

	sizeLoc := gl.GetUniformLocation(program, "uSize")
	gl.Uniform2f(sizeLoc, float32(c.PixelsWidth), float32(c.PixelsHeight))

	pausedLoc := gl.GetUniformLocation(program, "uPaused")
	paused := 0.
	if lg.Paused {
		paused = 1.
	}
	gl.Uniform1f(pausedLoc, float32(paused))
}

func (lg *LifeGame) getDeadColour() (float32, float32, float32) {
	return float32(0.5 * (1 + math.Sin(2*math.Pi*float64(lg.T)/float64(lg.ColourPeriod)))),
		float32(0.5 * (1 + math.Sin((2*math.Pi/3)+2*math.Pi*float64(lg.T)/float64(lg.ColourPeriod)))),
		float32(0.5 * (1 + math.Sin((4*math.Pi/3)+2*math.Pi*float64(lg.T)/float64(lg.ColourPeriod))))
}

func emptyArray(width int, height int) [][]bool {
	m := make([][]bool, height)
	for i := range m {
		m[i] = make([]bool, width)
	}

	return m
}

func randomArray(width int, height int) [][]bool {
	m := make([][]bool, height)
	for i := range m {
		m[i] = make([]bool, width)
		for j := range m[i] {
			if rand.Float32() > 0.8 {
				m[i][j] = true
			}
		}
	}

	return m
}
