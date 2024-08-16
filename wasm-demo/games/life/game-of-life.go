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
var fLifeShaderSource string

//go:embed shaders/display.frag
var fDisplayShaderSource string

//go:embed shaders/death.frag
var fDeathShaderSource string

type LifeGame struct {
	// tickTime     float32 // time in ms before tick can be updated
	// colourSpeed  float32
	lifeProgram         *webgl.Program
	deathProgram        *webgl.Program
	displayProgram      *webgl.Program
	vertexBuffer        *webgl.Buffer
	textureBuffer       *webgl.Buffer
	vCount              int
	writeFrameBuffer    *webgl.Framebuffer
	readFrameBuffer     *webgl.Framebuffer
	readTexture         *webgl.Texture
	writeTexture        *webgl.Texture
	t                   int
	storedPixels        []uint8
	tps                 int
	cumulativeIntervalT float32
	// colourPalette    []color.Color

	// lastFrameTime time.Time
	// thisFrameTime time.Time
	colourPeriodFrames int
	trailLength        int
	boundaryLoop       bool

	// gamePaused bool
}

var _ canvas.Animator = &LifeGame{}

func (lg *LifeGame) Init(c *canvas.GameContext) {
	lg.createShaders(c)
	lg.createBuffers(c)
	lg.t = -1
	lg.colourPeriodFrames = 60
	lg.trailLength = 15
	lg.tps = 5
	lg.boundaryLoop = false
	// TODO: store max device fps in the game context
	// TODO: load patterns
	// TODO: add all the sliders and buttons!
	// TODO: i wonder if i can transition between sizes without losing all information (eg taking a snapshot of the grid)
	// TODO: pause button
	// TODO: draw to the canvas while paused
	// TODO: random button, and infinite growth button, and maybe some other small ones
	// TODO: scroll screen on drag when unpaused and in looping mode
	// TODO: split this file up
	// TODO: make the actual webpage
}

func New() *LifeGame {
	return &LifeGame{}
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
		fmt.Printf("Error in life.vert: %v\n", *gl.GetShaderInfoLog(vShader))
	}

	fShader := gl.CreateShader(webgl.FRAGMENT_SHADER)
	gl.ShaderSource(fShader, fLifeShaderSource)
	gl.CompileShader(fShader)
	if !gl.GetShaderParameter(fShader, webgl.COMPILE_STATUS).Bool() {
		fmt.Printf("Error in life.frag: %v\n", *gl.GetShaderInfoLog(fShader))
	}

	lg.lifeProgram = gl.CreateProgram()
	gl.AttachShader(lg.lifeProgram, vShader)
	gl.AttachShader(lg.lifeProgram, fShader)
	gl.LinkProgram(lg.lifeProgram)
	if !gl.GetProgramParameter(lg.lifeProgram, webgl.LINK_STATUS).Bool() {
		fmt.Printf("Error in linking life: %v\n", *gl.GetProgramInfoLog(lg.lifeProgram))
	}

	fDisplayShader := gl.CreateShader(webgl.FRAGMENT_SHADER)
	gl.ShaderSource(fDisplayShader, fDisplayShaderSource)
	gl.CompileShader(fDisplayShader)
	if !gl.GetShaderParameter(fDisplayShader, webgl.COMPILE_STATUS).Bool() {
		fmt.Printf("Error in display.frag: %v\n", *gl.GetShaderInfoLog(fDisplayShader))
	}

	lg.displayProgram = gl.CreateProgram()
	gl.AttachShader(lg.displayProgram, vShader)
	gl.AttachShader(lg.displayProgram, fDisplayShader)
	gl.LinkProgram(lg.displayProgram)
	if !gl.GetProgramParameter(lg.displayProgram, webgl.LINK_STATUS).Bool() {
		fmt.Printf("Error in linking display %v\n", *gl.GetProgramInfoLog(lg.displayProgram))
	}

	fDeathShader := gl.CreateShader(webgl.FRAGMENT_SHADER)
	gl.ShaderSource(fDeathShader, fDeathShaderSource)
	gl.CompileShader(fDeathShader)
	if !gl.GetShaderParameter(fDeathShader, webgl.COMPILE_STATUS).Bool() {
		fmt.Printf("Error in death shader: %v\n", *gl.GetShaderInfoLog(fDeathShader))
	}

	lg.deathProgram = gl.CreateProgram()
	gl.AttachShader(lg.deathProgram, vShader)
	gl.AttachShader(lg.deathProgram, fDeathShader)
	gl.LinkProgram(lg.deathProgram)
	if !gl.GetProgramParameter(lg.deathProgram, webgl.LINK_STATUS).Bool() {
		fmt.Printf("Error in linking death: %v\n", *gl.GetProgramInfoLog(lg.deathProgram))
	}

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

	lg.readTexture = gl.CreateTexture()
	gl.BindTexture(webgl.TEXTURE_2D, lg.readTexture)
	gl.TexImage2D(webgl.TEXTURE_2D, 0, int(webgl.RGBA), int(c.Width), int(c.Height), 0, webgl.RGBA, webgl.UNSIGNED_BYTE, &webgl.Union{})
	gl.TexParameteri(webgl.TEXTURE_2D, webgl.TEXTURE_WRAP_S, int(webgl.CLAMP_TO_EDGE))
	gl.TexParameteri(webgl.TEXTURE_2D, webgl.TEXTURE_WRAP_T, int(webgl.CLAMP_TO_EDGE))
	gl.TexParameteri(webgl.TEXTURE_2D, webgl.TEXTURE_MIN_FILTER, int(webgl.NEAREST))
	gl.TexParameteri(webgl.TEXTURE_2D, webgl.TEXTURE_MAG_FILTER, int(webgl.NEAREST))

	lg.writeTexture = gl.CreateTexture()
	gl.BindTexture(webgl.TEXTURE_2D, lg.writeTexture)
	gl.TexImage2D(webgl.TEXTURE_2D, 0, int(webgl.RGBA), int(c.Width), int(c.Height), 0, webgl.RGBA, webgl.UNSIGNED_BYTE, webgl.UnionFromJS(jsconv.UInt8ToJs(pixels)))
	gl.TexParameteri(webgl.TEXTURE_2D, webgl.TEXTURE_WRAP_S, int(webgl.CLAMP_TO_EDGE))
	gl.TexParameteri(webgl.TEXTURE_2D, webgl.TEXTURE_WRAP_T, int(webgl.CLAMP_TO_EDGE))
	gl.TexParameteri(webgl.TEXTURE_2D, webgl.TEXTURE_MIN_FILTER, int(webgl.NEAREST))
	gl.TexParameteri(webgl.TEXTURE_2D, webgl.TEXTURE_MAG_FILTER, int(webgl.NEAREST))

	lg.writeFrameBuffer = gl.CreateFramebuffer()
	gl.BindFramebuffer(webgl.FRAMEBUFFER, lg.writeFrameBuffer)
	gl.FramebufferTexture2D(webgl.FRAMEBUFFER, webgl.COLOR_ATTACHMENT0, webgl.TEXTURE_2D, lg.readTexture, 0)
	lg.readFrameBuffer = gl.CreateFramebuffer()
	gl.BindFramebuffer(webgl.FRAMEBUFFER, lg.readFrameBuffer)
	gl.FramebufferTexture2D(webgl.FRAMEBUFFER, webgl.COLOR_ATTACHMENT0, webgl.TEXTURE_2D, lg.writeTexture, 0)

	gl.BindTexture(webgl.TEXTURE_2D, &webgl.Texture{})
	gl.BindFramebuffer(webgl.FRAMEBUFFER, &webgl.Framebuffer{})

	fmt.Println("made some buffers")
}

func (lg *LifeGame) Render(c *canvas.GameContext) {
	lg.cumulativeIntervalT += c.IntervalT
	// ticksToProcess := lg.cumulativeIntervalT * float32(lg.tps)
	// fmt.Println(int(ticksToProcess))
	// if ticksToProcess < 1 {
	// return
	// } else {
	// lg.cumulativeIntervalT = 0
	// }
	for range lg.tps {
		lg.t++
		lg.cumulativeIntervalT = 0
		if math.Mod(float64(lg.t), 1) != 0 {
			return
		}
		if math.Mod(float64(lg.t), 60) == 59 {
			lg.storedPixels = lg.getPixelsFromTexture(c)
			if math.Mod(float64(lg.t), 60) == 59 {
				lg.storedPixels[3*4*int(c.Width)+8] = 255
				lg.storedPixels[3*4*int(c.Width)+9] = 255
				lg.storedPixels[3*4*int(c.Width)+10] = 255
				lg.storedPixels[3*4*int(c.Width)+11] = 255
				lg.storedPixels[3*4*int(c.Width)+12] = 255
				lg.storedPixels[3*4*int(c.Width)+13] = 255
				lg.storedPixels[3*4*int(c.Width)+14] = 255
				lg.storedPixels[3*4*int(c.Width)+15] = 255
				lg.storedPixels[3*4*int(c.Width)+16] = 255
				lg.storedPixels[3*4*int(c.Width)+17] = 255
				lg.storedPixels[3*4*int(c.Width)+18] = 255
				lg.storedPixels[3*4*int(c.Width)+19] = 255
				lg.storedPixels[2*4*int(c.Width)+16] = 255
				lg.storedPixels[2*4*int(c.Width)+17] = 255
				lg.storedPixels[2*4*int(c.Width)+18] = 255
				lg.storedPixels[2*4*int(c.Width)+19] = 255
				lg.storedPixels[4*int(c.Width)+12] = 255
				lg.storedPixels[4*int(c.Width)+13] = 255
				lg.storedPixels[4*int(c.Width)+14] = 255
				lg.storedPixels[4*int(c.Width)+15] = 255

				lg.storedPixels[0] = 255
				lg.storedPixels[3] = 254
			}

			lg.setPixelsInTexture(c, lg.storedPixels)
			lg.deathFrame(c)
			lg.drawToCanvas(c)

			return
		}

		lg.lifeFrame(c)
	}
	lg.drawToCanvas(c)
	// thisFrameTime := time.Now()
	// duration := thisFrameTime.Sub(lg.lastFrameTime)
	// fps := 1 / duration.Seconds()
	// lg.lastFrameTime = thisFrameTime
	// if math.Mod(float64(lg.t), 60) == 0 {
	// 	fmt.Println(int(fps))
	// }
}

func (lg *LifeGame) deathFrame(c *canvas.GameContext) {
	gl := c.GL
	program := lg.deathProgram
	gl.BindFramebuffer(webgl.FRAMEBUFFER, lg.writeFrameBuffer)

	gl.ClearColor(0.0, 0.0, 1.0, 1.0)
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
	gl.UseProgram(program)
	decayLoc := gl.GetUniformLocation(program, "u_decay")
	gl.Uniform1f(decayLoc, 0.66/float32(lg.trailLength))

	gl.BindTexture(webgl.TEXTURE_2D, lg.writeTexture)
	samplerLocation := gl.GetUniformLocation(program, "u_sampler")
	gl.Uniform1i(samplerLocation, 0)

	sizeLoc := gl.GetUniformLocation(program, "u_size")
	gl.Uniform2f(sizeLoc, c.Width, c.Height)

	gl.DrawArrays(webgl.TRIANGLES, 0, lg.vCount)
	gl.BindFramebuffer(webgl.FRAMEBUFFER, &webgl.Framebuffer{})
}

func (lg *LifeGame) swapTextures() {
	lg.writeFrameBuffer, lg.readFrameBuffer = lg.readFrameBuffer, lg.writeFrameBuffer
	lg.readTexture, lg.writeTexture = lg.writeTexture, lg.readTexture
}

func (lg *LifeGame) lifeFrame(c *canvas.GameContext) {
	gl := c.GL
	program := lg.lifeProgram
	// lg.swapTextures()
	gl.BindFramebuffer(webgl.FRAMEBUFFER, lg.writeFrameBuffer)

	gl.ClearColor(0.0, 0.0, 1.0, 1.0)
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
	gl.UseProgram(program)

	decayLoc := gl.GetUniformLocation(program, "u_decay")
	gl.Uniform1f(decayLoc, 0.66/float32(lg.trailLength))

	boundaryLoc := gl.GetUniformLocation(program, "u_boundary_loop")
	boundaryLoop := 0.
	if lg.boundaryLoop {
		boundaryLoop = 1.
	}
	gl.Uniform1f(boundaryLoc, float32(boundaryLoop))

	initialDecayLoc := gl.GetUniformLocation(program, "u_initial_decay")
	gl.Uniform1f(initialDecayLoc, 0.33)

	deadColourLoc := gl.GetUniformLocation(program, "u_new_dead_colour")

	r, g, b := lg.getDeadColour()
	gl.Uniform3f(deadColourLoc, r, g, b)

	gl.BindTexture(webgl.TEXTURE_2D, lg.writeTexture)
	samplerLocation := gl.GetUniformLocation(program, "u_sampler")
	gl.Uniform1i(samplerLocation, 0)

	sizeLoc := gl.GetUniformLocation(program, "u_size")
	gl.Uniform2f(sizeLoc, c.Width, c.Height)

	gl.DrawArrays(webgl.TRIANGLES, 0, lg.vCount)
	gl.BindFramebuffer(webgl.FRAMEBUFFER, &webgl.Framebuffer{})
	lg.swapTextures()
}

func (lg *LifeGame) getDeadColour() (float32, float32, float32) {
	return float32(0.5 * (1 + math.Sin(2*math.Pi*float64(lg.t)/float64(lg.colourPeriodFrames)))),
		float32(0.5 * (1 + math.Sin((2*math.Pi/3)+2*math.Pi*float64(lg.t)/float64(lg.colourPeriodFrames)))),
		float32(0.5 * (1 + math.Sin((4*math.Pi/3)+2*math.Pi*float64(lg.t)/float64(lg.colourPeriodFrames))))
}

func (lg *LifeGame) drawToCanvas(c *canvas.GameContext) {
	gl := c.GL

	// THE BLITTING STUFF
	//
	// lg.swapTextures()
	gl.UseProgram(lg.displayProgram)
	program := lg.displayProgram

	gl.BindBuffer(webgl.ARRAY_BUFFER, lg.textureBuffer)
	tPosition := gl.GetAttribLocation(program, "a_tex_coord")
	// point the program to the vertex buffer object we've bound
	gl.VertexAttribPointer(uint(tPosition), 2, webgl.FLOAT, false, 0, 0)
	gl.EnableVertexAttribArray(uint(tPosition))

	gl.BindBuffer(webgl.ARRAY_BUFFER, lg.vertexBuffer)
	vPosition := gl.GetAttribLocation(program, "a_position")
	// point the program to the vertex buffer object we've bound
	gl.VertexAttribPointer(uint(vPosition), 2, webgl.FLOAT, false, 0, 0)
	gl.EnableVertexAttribArray(uint(vPosition))

	gl.BindTexture(webgl.TEXTURE_2D, lg.readTexture)
	samplerLocation := gl.GetUniformLocation(program, "u_sampler")
	gl.Uniform1i(samplerLocation, 0)

	gl.Viewport(0, 0, int(c.Width), int(c.Height))
	gl.Clear(webgl.COLOR_BUFFER_BIT)
	gl.DrawArrays(webgl.TRIANGLES, 0, lg.vCount)

	gl.BindTexture(webgl.TEXTURE_2D, &webgl.Texture{})
	// lg.swapTextures()
}

func (lg *LifeGame) getPixelsFromTexture(c *canvas.GameContext) []uint8 {
	gl := c.GL
	gl.BindFramebuffer(webgl.FRAMEBUFFER, lg.readFrameBuffer)
	gl.BindTexture(webgl.TEXTURE_2D, lg.readTexture)
	union := webgl.Union{
		Value: jsconv.UInt8ToJs(make([]uint8, int(c.Height*c.Width*4))),
	}
	gl.ReadPixels(0, 0, int(c.Width), int(c.Height), webgl.RGBA, webgl.UNSIGNED_BYTE, &union)
	// fmt.Println(union.Value.Type().String())
	// if lg.t < 60 {
	// 	canvas.Log(union.Value)
	// }
	gl.BindTexture(webgl.TEXTURE_2D, &webgl.Texture{})
	gl.BindFramebuffer(webgl.FRAMEBUFFER, &webgl.Framebuffer{})

	return jsconv.JsToUInt8(union.Value)
}

func (lg *LifeGame) setPixelsInTexture(c *canvas.GameContext, in []uint8) {
	gl := c.GL
	gl.BindTexture(webgl.TEXTURE_2D, lg.writeTexture)
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
