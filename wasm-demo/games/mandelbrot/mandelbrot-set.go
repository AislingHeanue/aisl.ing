package mandelbrot

import (
	_ "embed"

	"github.com/AislingHeanue/aisling-codes/wasm-demo/common"
	"github.com/gowebapi/webapi/graphics/webgl"
)

//go:embed shaders/mandelbrot.vert
var vertexSource string

//go:embed shaders/mandelbrot.frag
var fragmentSource string

func New(opts MandelbrotOptions) *Mandelbrot {
	return &Mandelbrot{
		MandelbrotContext: &MandelbrotContext{
			CentreX:    opts.CentreX,
			CentreY:    opts.CentreY,
			Zoom:       opts.Zoom,
			Iterations: opts.Iterations,
			FpsTarget:  opts.FpsTarget,
		},
	}
}

type MandelbrotOptions struct {
	CentreX    float64
	CentreY    float64
	Zoom       float64
	Iterations int
	FpsTarget  int
}

type Mandelbrot struct {
	common.DefaultGame
	*MandelbrotContext
}

func (m *Mandelbrot) InitListeners(c *common.GameContext) {
	common.RegisterListeners(c, m.MandelbrotContext, MandelbrotController{}, MandelbrotActionHandler{})
}

func (m *Mandelbrot) SkipThisFrame(c *common.GameContext) bool {
	return m.FrameUpToDate
}

// AttachAttributes implements common.GameInfo.
func (m *Mandelbrot) AttachAttributes(c *common.GameContext, program *webgl.Program, writeBuffer *webgl.Buffer, readBuffer *webgl.Buffer, samplerTexture *webgl.Texture) {
	gl := c.GL
	currentFps := 1 / c.IntervalT
	if currentFps < 0.5*float32(m.FpsTarget) && m.Iterations > 400 {
		m.Iterations = max(m.Iterations*4/5, 400)
		// c.Logf("Decreasing max iterations to %d", m.Iterations)
	}
	if currentFps > 1.7*float32(m.FpsTarget) && m.Iterations < 3000 {
		m.Iterations = min(m.Iterations*5/4, 3000)
		// c.Logf("Increasing max iterations to %d", m.Iterations)
	}

	gl.BindBuffer(webgl.ARRAY_BUFFER, writeBuffer)
	vPosition := gl.GetAttribLocation(program, "aPosition")
	gl.VertexAttribPointer(uint(vPosition), 2, webgl.FLOAT, false, 0, 0)
	gl.EnableVertexAttribArray(uint(vPosition))

	gl.BindBuffer(webgl.ARRAY_BUFFER, readBuffer)
	tPosition := gl.GetAttribLocation(program, "aTexCoord")
	gl.VertexAttribPointer(uint(tPosition), 2, webgl.FLOAT, false, 0, 0)
	gl.EnableVertexAttribArray(uint(tPosition))

	zoomLoc := gl.GetUniformLocation(program, "uZoom")
	gl.Uniform1f(zoomLoc, float32(m.Zoom))

	centreLoc := gl.GetUniformLocation(program, "uCentre")
	gl.Uniform2f(centreLoc, float32(m.CentreX), float32(m.CentreY))

	iterLoc := gl.GetUniformLocation(program, "uMaxIterations")
	gl.Uniform1i(iterLoc, m.Iterations)
	m.FrameUpToDate = true
}

func (m *Mandelbrot) GetFragmentSource() string {
	return fragmentSource
}

func (m *Mandelbrot) GetVertexSource() string {
	return vertexSource
}

var _ common.GameInfo = &Mandelbrot{}
