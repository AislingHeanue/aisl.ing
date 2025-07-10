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
			CentreX: opts.CentreX,
			CentreY: opts.CentreY,
			Zoom:    opts.Zoom,
		},
	}
}

type MandelbrotOptions struct {
	CentreX float64
	CentreY float64
	Zoom    float64
}

type Mandelbrot struct {
	common.DefaultGame
	*MandelbrotContext
}

func (m *Mandelbrot) InitListeners(c *common.GameContext) {
	common.RegisterListeners(c, m.MandelbrotContext, MandelbrotController{}, MandelbrotActionHandler{})
}

// AttachAttributes implements common.GameInfo.
func (m *Mandelbrot) AttachAttributes(c *common.GameContext, program *webgl.Program, writeBuffer *webgl.Buffer, readBuffer *webgl.Buffer, samplerTexture *webgl.Texture) {
	gl := c.GL

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
}

func (m *Mandelbrot) GetFragmentSource() string {
	return fragmentSource
}

func (m *Mandelbrot) GetVertexSource() string {
	return vertexSource
}

var _ common.GameInfo = &Mandelbrot{}
