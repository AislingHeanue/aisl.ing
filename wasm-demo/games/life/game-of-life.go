package life

import (
	_ "embed"
	"fmt"
	"math"

	"github.com/AislingHeanue/aisling-codes/wasm-demo/canvas"
	"github.com/gowebapi/webapi/core/jsconv"
	"github.com/gowebapi/webapi/graphics/webgl"
)

//go:embed shaders/life.vert
var vShaderSource string

//go:embed shaders/life.frag
var fShaderSource string

//go:embed shaders/display.frag
var fDisplayShaderSource string

type LifeGame struct {
	// tickTime     float32 // time in ms before tick can be updated
	// colourSpeed  float32
	program             *webgl.Program
	displayProgram      *webgl.Program
	vertexBuffer        *webgl.Buffer
	textureBuffer       *webgl.Buffer
	vCount              int
	bufferStale         bool
	bufferIndex         int
	frameBuffer         *webgl.Framebuffer
	frameBufferNotInUse *webgl.Framebuffer
	texture             *webgl.Texture
	textureNotInUse     *webgl.Texture
	t                   int
}

var _ canvas.Animator = &LifeGame{}

func (lg *LifeGame) Init(c *canvas.GameContext) {
	lg.createShaders(c)
	lg.t = -1
}

func (lg LifeGame) InitListeners(c *canvas.GameContext) {
	// cubeController.InitListeners(c, cc.CubeCubeContext)
}

func (lg *LifeGame) createShaders(c *canvas.GameContext) {
	gl := c.GL

	vShader := gl.CreateShader(webgl.VERTEX_SHADER)
	gl.ShaderSource(vShader, vShaderSource)
	gl.CompileShader(vShader)
	if !gl.GetShaderParameter(vShader, webgl.COMPILE_STATUS).Bool() {
		fmt.Printf("ERROR V: %v\n", *gl.GetShaderInfoLog(vShader))
	}

	fShader := gl.CreateShader(webgl.FRAGMENT_SHADER)
	gl.ShaderSource(fShader, fShaderSource)
	gl.CompileShader(fShader)
	if !gl.GetShaderParameter(fShader, webgl.COMPILE_STATUS).Bool() {
		fmt.Printf("ERROR F: %v\n", *gl.GetShaderInfoLog(fShader))
	}

	program := gl.CreateProgram()
	gl.AttachShader(program, vShader)
	gl.AttachShader(program, fShader)
	gl.LinkProgram(program)
	if !gl.GetProgramParameter(program, webgl.LINK_STATUS).Bool() {
		fmt.Printf("ERROR L: %v\n", *gl.GetProgramInfoLog(program))
	}
	lg.program = program

	fDisplayShader := gl.CreateShader(webgl.FRAGMENT_SHADER)
	gl.ShaderSource(fDisplayShader, fDisplayShaderSource)
	gl.CompileShader(fDisplayShader)
	if !gl.GetShaderParameter(fDisplayShader, webgl.COMPILE_STATUS).Bool() {
		fmt.Printf("ERROR F2: %v\n", *gl.GetShaderInfoLog(fDisplayShader))
	}

	program2 := gl.CreateProgram()
	gl.AttachShader(program2, vShader)
	gl.AttachShader(program2, fDisplayShader)
	gl.LinkProgram(program2)
	if !gl.GetProgramParameter(program2, webgl.LINK_STATUS).Bool() {
		fmt.Printf("ERROR L: %v\n", *gl.GetProgramInfoLog(program2))
	}
	lg.displayProgram = program2

	lg.bufferStale = true
	lg.bufferIndex = 0
}

func (lg *LifeGame) createBuffers(c *canvas.GameContext) {
	gl := c.GL
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

	textureArray := []float32{
		0.0, 0.0,
		1.0, 0.0,
		0.0, 1.0,
		0.0, 1.0,
		1.0, 0.0,
		1.0, 1.0,
	}

	textureCorners := jsconv.Float32ToJs(textureArray)
	textureBuffer := gl.CreateBuffer()
	gl.BindBuffer(webgl.ARRAY_BUFFER, textureBuffer)
	gl.BufferData2(webgl.ARRAY_BUFFER, webgl.UnionFromJS(textureCorners), webgl.STATIC_DRAW)
	gl.BindBuffer(webgl.ARRAY_BUFFER, &webgl.Buffer{})
	lg.textureBuffer = textureBuffer

	pixels := setupPixelArray(int(c.Width), int(c.Height))

	lg.texture = gl.CreateTexture()
	gl.BindTexture(webgl.TEXTURE_2D, lg.texture)
	gl.TexImage2D(webgl.TEXTURE_2D, 0, int(webgl.RGBA), int(c.Width), int(c.Height), 0, webgl.RGBA, webgl.UNSIGNED_BYTE, &webgl.Union{})
	gl.TexParameteri(webgl.TEXTURE_2D, webgl.TEXTURE_WRAP_S, int(webgl.CLAMP_TO_EDGE))
	gl.TexParameteri(webgl.TEXTURE_2D, webgl.TEXTURE_WRAP_T, int(webgl.CLAMP_TO_EDGE))
	gl.TexParameteri(webgl.TEXTURE_2D, webgl.TEXTURE_MIN_FILTER, int(webgl.NEAREST))
	gl.TexParameteri(webgl.TEXTURE_2D, webgl.TEXTURE_MAG_FILTER, int(webgl.NEAREST))

	lg.textureNotInUse = gl.CreateTexture()
	gl.BindTexture(webgl.TEXTURE_2D, lg.textureNotInUse)
	gl.TexImage2D(webgl.TEXTURE_2D, 0, int(webgl.RGBA), int(c.Width), int(c.Height), 0, webgl.RGBA, webgl.UNSIGNED_BYTE, webgl.UnionFromJS(jsconv.UInt8ToJs(pixels)))
	gl.TexParameteri(webgl.TEXTURE_2D, webgl.TEXTURE_WRAP_S, int(webgl.CLAMP_TO_EDGE))
	gl.TexParameteri(webgl.TEXTURE_2D, webgl.TEXTURE_WRAP_T, int(webgl.CLAMP_TO_EDGE))
	gl.TexParameteri(webgl.TEXTURE_2D, webgl.TEXTURE_MIN_FILTER, int(webgl.NEAREST))
	gl.TexParameteri(webgl.TEXTURE_2D, webgl.TEXTURE_MAG_FILTER, int(webgl.NEAREST))

	lg.frameBuffer = gl.CreateFramebuffer()
	gl.BindFramebuffer(webgl.FRAMEBUFFER, lg.frameBuffer)
	gl.FramebufferTexture2D(webgl.FRAMEBUFFER, webgl.COLOR_ATTACHMENT0, webgl.TEXTURE_2D, lg.texture, 0)
	lg.frameBufferNotInUse = gl.CreateFramebuffer()
	gl.BindFramebuffer(webgl.FRAMEBUFFER, lg.frameBufferNotInUse)
	gl.FramebufferTexture2D(webgl.FRAMEBUFFER, webgl.COLOR_ATTACHMENT0, webgl.TEXTURE_2D, lg.textureNotInUse, 0)

	gl.BindTexture(webgl.TEXTURE_2D, &webgl.Texture{})
	gl.BindFramebuffer(webgl.FRAMEBUFFER, &webgl.Framebuffer{})

	fmt.Println("made some buffers")

	lg.bufferStale = false
}

func (lg *LifeGame) Render(c *canvas.GameContext) {
	gl := c.GL
	program := lg.program
	if lg.bufferStale {
		lg.createBuffers(c)
	}
	lg.t++
	if math.Mod(float64(lg.t), 1) != 0 {
		return
	}

	if math.Mod(float64(lg.t), 60) == 30 {
		pixels := lg.getPixelsFromTexture(c)
		// pixels[0] = 255
		// pixels[1] = 255
		// pixels[2] = 255
		// pixels[3] = 255
		lg.setPixelsInTexture(c, pixels)
	}

	gl.ClearColor(0.0, 0.0, 1.0, 1.0)
	// gl.ActiveTexture(webgl.TEXTURE0)

	gl.BindBuffer(webgl.ARRAY_BUFFER, lg.vertexBuffer)
	vPosition := gl.GetAttribLocation(program, "a_position")
	// point the program to the vertex buffer object we've bound
	gl.VertexAttribPointer(uint(vPosition), 2, webgl.FLOAT, false, 0, 0)
	gl.EnableVertexAttribArray(uint(vPosition))

	gl.BindBuffer(webgl.ARRAY_BUFFER, lg.textureBuffer)
	tPosition := gl.GetAttribLocation(program, "a_tex_coord")
	// point the program to the vertex buffer object we've bound
	gl.VertexAttribPointer(uint(tPosition), 2, webgl.FLOAT, false, 0, 0)
	gl.EnableVertexAttribArray(uint(tPosition))

	gl.UseProgram(lg.program)
	gl.BindFramebuffer(webgl.FRAMEBUFFER, lg.frameBuffer)

	decayLoc := gl.GetUniformLocation(lg.program, "u_decay")
	gl.Uniform1f(decayLoc, 0.1)

	initialDecayLoc := gl.GetUniformLocation(lg.program, "u_initial_decay")
	gl.Uniform1f(initialDecayLoc, 0.3)

	deadColourLoc := gl.GetUniformLocation(lg.program, "u_new_dead_colour")
	gl.Uniform3f(deadColourLoc, 0, 1, 0)

	gl.BindTexture(webgl.TEXTURE_2D, lg.textureNotInUse)
	samplerLocation := gl.GetUniformLocation(lg.program, "u_sampler")
	gl.Uniform1i(samplerLocation, 0)

	sizeLoc := gl.GetUniformLocation(lg.program, "u_size")
	gl.Uniform2f(sizeLoc, c.Width, c.Height)

	gl.DrawArrays(webgl.TRIANGLES, 0, lg.vCount)
	gl.BindFramebuffer(webgl.FRAMEBUFFER, &webgl.Framebuffer{})

	// THE BLITTING STUFF
	//
	gl.UseProgram(lg.displayProgram)
	program = lg.displayProgram

	gl.BindBuffer(webgl.ARRAY_BUFFER, lg.textureBuffer)
	tPosition = gl.GetAttribLocation(program, "a_tex_coord")
	// point the program to the vertex buffer object we've bound
	gl.VertexAttribPointer(uint(tPosition), 2, webgl.FLOAT, false, 0, 0)
	gl.EnableVertexAttribArray(uint(tPosition))

	gl.BindBuffer(webgl.ARRAY_BUFFER, lg.vertexBuffer)
	vPosition = gl.GetAttribLocation(program, "a_position")
	// point the program to the vertex buffer object we've bound
	gl.VertexAttribPointer(uint(vPosition), 2, webgl.FLOAT, false, 0, 0)
	gl.EnableVertexAttribArray(uint(vPosition))

	gl.BindTexture(webgl.TEXTURE_2D, lg.texture)
	samplerLocation = gl.GetUniformLocation(program, "u_sampler")
	gl.Uniform1i(samplerLocation, 0)

	gl.Viewport(0, 0, int(c.Width), int(c.Height))
	gl.Clear(webgl.COLOR_BUFFER_BIT)
	gl.DrawArrays(webgl.TRIANGLES, 0, lg.vCount)

	gl.BindTexture(webgl.TEXTURE_2D, &webgl.Texture{})
	lg.frameBuffer, lg.frameBufferNotInUse = lg.frameBufferNotInUse, lg.frameBuffer
	lg.texture, lg.textureNotInUse = lg.textureNotInUse, lg.texture
}

func (lg *LifeGame) getPixelsFromTexture(c *canvas.GameContext) []uint8 {
	gl := c.GL
	gl.BindTexture(webgl.TEXTURE_2D, lg.texture)
	union := webgl.Union{
		Value: jsconv.UInt8ToJs(make([]uint8, int(c.Height*c.Width*4))),
	}
	gl.ReadPixels(0, 0, int(c.Width), int(c.Height), webgl.RGBA, webgl.UNSIGNED_BYTE, &union)
	// fmt.Println(union.Value.Type().String())
	if lg.t < 60 {
		canvas.Log(union.Value)
	}
	gl.BindTexture(webgl.TEXTURE_2D, &webgl.Texture{})

	return jsconv.JsToUInt8(union.Value)
}

func (lg *LifeGame) setPixelsInTexture(c *canvas.GameContext, in []uint8) {
	gl := c.GL
	gl.BindTexture(webgl.TEXTURE_2D, lg.textureNotInUse)
	gl.TexImage2D(webgl.TEXTURE_2D, 0, int(webgl.RGBA), int(c.Width), int(c.Height), 0, webgl.RGBA, webgl.UNSIGNED_BYTE, webgl.UnionFromJS(jsconv.UInt8ToJs(in)))
	gl.BindTexture(webgl.TEXTURE_2D, &webgl.Texture{})
}

func setupPixelArray(width int, height int) []uint8 {
	if width == 5 && height == 5 {
		return []uint8{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		}
	}

	on := []uint8{255, 255, 255, 255}
	off := []uint8{0, 0, 0, 0}
	out := make([]uint8, 4*height*width)
	m := make([][]bool, height)
	for i := range m {
		m[i] = make([]bool, width)
	}
	// matrix is upside down because webgl reasons
	//
	// .....
	// ..*..
	// ...*.
	// .***.
	// ..... etc.
	// m[height-1-1][2] = true
	// m[height-2-1][3] = true
	// m[height-3-1][1] = true
	// m[height-3-1][2] = true
	// m[height-3-1][3] = true

	// ......*.
	// ....*.**
	// ....*.*.
	// ....*...
	// ..*.....
	// *.*.....
	midWidth := width/2 - 5
	midHeight := height/2 - 3
	m[midHeight][midWidth] = true
	m[midHeight][midWidth+2] = true
	m[midHeight+1][midWidth+2] = true
	m[midHeight+2][midWidth+4] = true
	m[midHeight+3][midWidth+4] = true
	m[midHeight+4][midWidth+4] = true
	m[midHeight+3][midWidth+6] = true
	m[midHeight+4][midWidth+6] = true
	m[midHeight+4][midWidth+7] = true
	m[midHeight+5][midWidth+6] = true

	for i := range m {
		for j := range m[i] {
			if m[i][j] {
				out[4*(i*width+j)+0] = on[0]
				out[4*(i*width+j)+1] = on[0]
				out[4*(i*width+j)+2] = on[0]
				out[4*(i*width+j)+3] = on[0]
			} else {
				out[4*(i*width+j)+0] = off[0]
				out[4*(i*width+j)+1] = off[0]
				out[4*(i*width+j)+2] = off[0]
				out[4*(i*width+j)+3] = off[0]
			}
		}
	}

	return out
}
