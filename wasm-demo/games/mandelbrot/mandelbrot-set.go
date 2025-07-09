package mandelbrot

import (
	_ "embed"

	"github.com/AislingHeanue/aisling-codes/wasm-demo/games/common"
	"github.com/gowebapi/webapi/graphics/webgl"
)

//go:embed shaders/mandelbrot.vert
var vertexSource string

//go:embed shaders/mandelbrot.frag
var fragmentSource string

func New(cco MandelbrotOptions) *Mandelbrot {
	return &Mandelbrot{}
}

type MandelbrotOptions struct {
}

type Mandelbrot struct {
	common.DefaultGame
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
}

func (m *Mandelbrot) GetFragmentSource() string {
	return fragmentSource
}

func (m *Mandelbrot) GetVertexSource() string {
	return vertexSource
}

var _ common.GameInfo = &Mandelbrot{}
