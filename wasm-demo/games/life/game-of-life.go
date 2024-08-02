package life

import (
	"fmt"

	"github.com/AislingHeanue/aisling-codes/wasm-demo/canvas"
	"github.com/gowebapi/webapi/core/jsconv"
	"github.com/gowebapi/webapi/graphics/webgl"
)

type LifeGame struct {
	// tickTime     float32 // time in ms before tick can be updated
	// colourSpeed  float32
	program      *webgl.Program
	vertexBuffer *webgl.Buffer
	vCount       int
	bufferStale  bool
}

var _ canvas.Animator = &LifeGame{}

func (lg *LifeGame) Init(c *canvas.GameContext) {
	lg.createShaders(c)
}

func (lg LifeGame) InitListeners(c *canvas.GameContext) {
	// cubeController.InitListeners(c, cc.CubeCubeContext)
}

func (lg *LifeGame) createShaders(c *canvas.GameContext) {
	gl := c.GL
	vShaderSource := `
attribute vec2 v_position;
void main() {
	// vertex shader knows where it is, because it knows where it isn't
	gl_Position = vec4(v_position, 0.0, 1.0);
}`

	fShaderSource := `
precision lowp float;
uniform vec2 u_resolution;
void main() {
	gl_FragColor =  vec4(gl_FragCoord.x/10.0,1.0-gl_FragCoord.y/10.0, 0.5, 1.0);
}`

	vShader := gl.CreateShader(webgl.VERTEX_SHADER)
	gl.ShaderSource(vShader, vShaderSource)
	gl.CompileShader(vShader)
	if !gl.GetShaderParameter(vShader, webgl.COMPILE_STATUS).Bool() {
		fmt.Printf("ERROR V: %v", gl.GetShaderInfoLog(vShader))
	}

	fShader := gl.CreateShader(webgl.FRAGMENT_SHADER)
	gl.ShaderSource(fShader, fShaderSource)
	gl.CompileShader(fShader)
	if !gl.GetShaderParameter(fShader, webgl.COMPILE_STATUS).Bool() {
		fmt.Printf("ERROR F: %v", gl.GetShaderInfoLog(fShader))
	}

	program := gl.CreateProgram()
	gl.AttachShader(program, vShader)
	gl.AttachShader(program, fShader)
	gl.LinkProgram(program)
	if !gl.GetProgramParameter(program, webgl.LINK_STATUS).Bool() {
		fmt.Printf("ERROR L: %v", gl.GetProgramInfoLog(program))
	}

	fmt.Println("made some shaders")
	gl.UseProgram(program)

	lg.program = program
	lg.bufferStale = true
}

func (lg *LifeGame) Render(c *canvas.GameContext) {
	gl := c.GL
	program := lg.program
	if lg.bufferStale {
		lg.createBuffers(gl)
	}
	gl.ClearColor(0.0, 0.0, 0.0, 1.0)

	gl.Clear(webgl.COLOR_BUFFER_BIT)

	gl.BindBuffer(webgl.ARRAY_BUFFER, lg.vertexBuffer)
	vPosition := gl.GetAttribLocation(program, "v_position")
	// point the program to the vertex buffer object we've bound
	gl.VertexAttribPointer(uint(vPosition), 2, webgl.FLOAT, false, 0, 0)
	gl.EnableVertexAttribArray(uint(vPosition))

	resLocation := gl.GetUniformLocation(lg.program, "u_resolution")
	gl.Uniform2f(resLocation, c.Width, c.Height)

	gl.DrawArrays(webgl.TRIANGLES, 0, lg.vCount)
}

func (lg *LifeGame) createBuffers(gl *webgl.RenderingContext) {
	vertexArray := []float32{
		-1.0, -1.0,
		1.0, -1.0,
		-1.0, 1.0,
		-1.0, 1.0,
		1.0, -1.0,
		1.0, 1.0,
	}

	vertices := jsconv.Float32ToJs(vertexArray)
	vertexBuffer := gl.CreateBuffer()
	gl.BindBuffer(webgl.ARRAY_BUFFER, vertexBuffer)
	gl.BufferData2(webgl.ARRAY_BUFFER, webgl.UnionFromJS(vertices), webgl.STATIC_DRAW)
	gl.BindBuffer(webgl.ARRAY_BUFFER, &webgl.Buffer{})
	lg.vertexBuffer = vertexBuffer
	lg.vCount = 6

	fmt.Println("made some buffers")

	lg.bufferStale = false
}
