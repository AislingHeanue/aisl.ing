package life

import (
	_ "embed"
	"fmt"
	"math"
	"math/rand/v2"

	"github.com/gowebapi/webapi/core/jsconv"
	"github.com/gowebapi/webapi/graphics/webgl"
	webapicanvas "github.com/gowebapi/webapi/html/canvas"

	"github.com/AislingHeanue/aisling-codes/wasm-demo/canvas"
	"github.com/AislingHeanue/aisling-codes/wasm-demo/games/life/controller"
	"github.com/AislingHeanue/aisling-codes/wasm-demo/games/life/model"
	"github.com/AislingHeanue/aisling-codes/wasm-demo/games/life/parser"
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
	lifeProgram      *webgl.Program
	deathProgram     *webgl.Program
	displayProgram   *webgl.Program
	vertexBuffer     *webgl.Buffer
	textureBuffer    *webgl.Buffer
	vCount           int
	writeFrameBuffer *webgl.Framebuffer
	readFrameBuffer  *webgl.Framebuffer
	readTexture      *webgl.Texture
	writeTexture     *webgl.Texture
	t                int
	// storedPixels     []uint8

	// configurable but not exposed to the frontend
	colourPeriodFrames int
	trailLength        int

	lifeContext *model.LifeContext
}

// loader stuff
// TODO: load patterns + drop-down selection to pick a design from the life wiki (with some curated samples)

// frontend stuff
// TODO: simulation size slider (with a reasonable default, which changes based on selected wiki design) (causes re-init if used manually)
// TODO: zoom slider
// TODO: random button, clear button

// TIME stuff
// TODO: fps and tps controls toggleable

var _ canvas.Animator = &LifeGame{}

func (lg *LifeGame) zoomY(c *canvas.GameContext) float32 {
	return lg.lifeContext.Zoom * c.Height / float32(lg.lifeContext.CellHeight)
}

func (lg *LifeGame) zoomX(c *canvas.GameContext) float32 {
	return lg.lifeContext.Zoom * c.Width / float32(lg.lifeContext.CellWidth)
}

func (lg *LifeGame) zoom(c *canvas.GameContext) float32 {
	return min(lg.zoomX(c), lg.zoomY(c))
}

func (lg *LifeGame) Init(c *canvas.GameContext) {
	lg.createShaders(c)
	lg.createBuffers(c)
	lg.t = -1
	lg.colourPeriodFrames = 60
	lg.trailLength = 50

	if lg.zoom(c) < 0 {
		panic("I refuse to create an infinite loop no thank you")
	}
	lg.lifeContext.DX = 0 // default offset position of the grid is 0,0
	lg.lifeContext.DY = 0
}

func New(lc *model.LifeContext) *LifeGame {
	return &LifeGame{lifeContext: lc}
}

func (lg LifeGame) Dimensions() (int, int) {
	return lg.lifeContext.CellWidth, lg.lifeContext.CellHeight
}

func (lg *LifeGame) InitListeners(c *canvas.GameContext) {
	controller.InitListeners(c, lg.lifeContext, lg)
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

	pixels := randomArray(lg.lifeContext.CellWidth, lg.lifeContext.CellHeight)
	pixelsArray := setupPixelArray(pixels)

	lg.readTexture = gl.CreateTexture()
	gl.BindTexture(webgl.TEXTURE_2D, lg.readTexture)
	gl.TexImage2D(webgl.TEXTURE_2D, 0, int(webgl.RGBA), lg.lifeContext.CellWidth, lg.lifeContext.CellHeight, 0, webgl.RGBA, webgl.UNSIGNED_BYTE, &webgl.Union{})
	gl.TexParameteri(webgl.TEXTURE_2D, webgl.TEXTURE_WRAP_S, int(webgl.CLAMP_TO_EDGE))
	gl.TexParameteri(webgl.TEXTURE_2D, webgl.TEXTURE_WRAP_T, int(webgl.CLAMP_TO_EDGE))
	gl.TexParameteri(webgl.TEXTURE_2D, webgl.TEXTURE_MIN_FILTER, int(webgl.NEAREST))
	gl.TexParameteri(webgl.TEXTURE_2D, webgl.TEXTURE_MAG_FILTER, int(webgl.NEAREST))

	lg.writeTexture = gl.CreateTexture()
	gl.BindTexture(webgl.TEXTURE_2D, lg.writeTexture)
	gl.TexImage2D(webgl.TEXTURE_2D, 0, int(webgl.RGBA), lg.lifeContext.CellWidth, lg.lifeContext.CellHeight, 0, webgl.RGBA, webgl.UNSIGNED_BYTE, webgl.UnionFromJS(jsconv.UInt8ToJs(pixelsArray)))
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
	// lg.cumulativeIntervalT += c.IntervalT
	// ticksToProcess := lg.cumulativeIntervalT * float32(lg.tps)
	// fmt.Println(int(ticksToProcess))
	// if ticksToProcess < 1 {
	// return
	// } else {
	// lg.cumulativeIntervalT = 0
	// }
	for range lg.lifeContext.Tps {
		if lg.lifeContext.Paused {
			lg.deathFrame(c)
		} else {
			lg.t++
			lg.lifeFrame(c)
		}
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
	gl.Uniform2f(sizeLoc, float32(lg.lifeContext.CellWidth), float32(lg.lifeContext.CellHeight))

	gl.DrawArrays(webgl.TRIANGLES, 0, lg.vCount)
	gl.BindFramebuffer(webgl.FRAMEBUFFER, &webgl.Framebuffer{})
	lg.swapTextures()
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
	if lg.lifeContext.Loop {
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
	gl.Uniform2f(sizeLoc, float32(lg.lifeContext.CellWidth), float32(lg.lifeContext.CellHeight))

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
	ctx := c.ZoomCtx
	showCtx := c.DisplayCtx

	// the + c.Width/2 here makes it so that the 'anchor' point for zooming in and out is at the centre of the canvas
	topLeftDX := lg.lifeContext.DX + c.Width/2 - lg.zoom(c)*float32(lg.lifeContext.CellWidth)/2
	topLeftDY := lg.lifeContext.DY + c.Height/2 - lg.zoom(c)*float32(lg.lifeContext.CellHeight)/2
	// bound check DX and DY and make sure they're within a valid range to be able to draw each part of all the visible canvases.
	for topLeftDX > 0 {
		topLeftDX -= float32(lg.lifeContext.CellWidth) * lg.zoom(c)
	}
	for topLeftDY > 0 {
		topLeftDY -= float32(lg.lifeContext.CellHeight) * lg.zoom(c)
	}

	union := webgl.Union{
		Value: jsconv.UInt8ToJs(make([]uint8, lg.lifeContext.CellHeight*lg.lifeContext.CellWidth*4)),
	}
	gl.BindFramebuffer(webgl.FRAMEBUFFER, lg.readFrameBuffer)
	gl.ReadPixels(0, 0, lg.lifeContext.CellWidth, lg.lifeContext.CellHeight, webgl.RGBA, webgl.UNSIGNED_BYTE, &union)
	gl.BindFramebuffer(webgl.FRAMEBUFFER, &webgl.Framebuffer{})

	imageData := ctx.CreateImageData(lg.lifeContext.CellWidth, lg.lifeContext.CellHeight)
	imageData.Data().JSValue().Call("set", union.JSValue())
	ctx.PutImageData(imageData, 0, 0)
	floatWidth := c.Width   //float32(lg.lifeContext.CellWidth)
	floatHeight := c.Height //float32(lg.lifeContext.CellWidth)

	showCtx.ClearRect(0, 0, float64(floatWidth), float64(floatHeight))
	// tile horizontally if one instance of the grid does not cover the canvas
	for currentDx := topLeftDX; currentDx < c.Width; currentDx += float32(lg.lifeContext.CellWidth) * lg.zoom(c) {
		// and vertically
		for currentDy := topLeftDY; currentDy < c.Height; currentDy += float32(lg.lifeContext.CellHeight) * lg.zoom(c) {
			showCtx.DrawImage3(
				webapicanvas.UnionFromJS(c.ZoomCanvas.JSValue()),
				0, 0, // start coords in grid being captured from
				float64((c.Width-currentDx)/lg.zoom(c)), float64((c.Height-currentDy)/lg.zoom(c)),
				float64(currentDx), float64(currentDy), // start coords in grid being displayed to
				float64(c.Width-currentDx), float64(c.Height-currentDy),
			)
		}
	}
}

// func (lg *LifeGame) getPixelsFromTexture(c *canvas.GameContext) []uint8 {
// 	gl := c.GL
// 	gl.BindFramebuffer(webgl.FRAMEBUFFER, lg.readFrameBuffer)
// 	union := webgl.Union{
// 		Value: jsconv.UInt8ToJs(make([]uint8, lg.lifeContext.CellHeight*lg.lifeContext.CellWidth*4)),
// 	}
// 	gl.ReadPixels(0, 0, lg.lifeContext.CellWidth, lg.lifeContext.CellHeight, webgl.RGBA, webgl.UNSIGNED_BYTE, &union)
// 	// fmt.Println(union.Value.Type().String())
// 	// if lg.t < 60 {
// 	// 	canvas.Log(union.Value)
// 	// }
// 	gl.BindFramebuffer(webgl.FRAMEBUFFER, &webgl.Framebuffer{})
//
// 	return jsconv.JsToUInt8(union.Value)
// }

func (lg *LifeGame) setPixelsInTexture(c *canvas.GameContext, in [][]bool) {
	gl := c.GL
	inArray := setupPixelArray(in)
	gl.BindTexture(webgl.TEXTURE_2D, lg.writeTexture)
	gl.TexImage2D(webgl.TEXTURE_2D, 0, int(webgl.RGBA), lg.lifeContext.CellWidth, lg.lifeContext.CellHeight, 0, webgl.RGBA, webgl.UNSIGNED_BYTE, webgl.UnionFromJS(jsconv.UInt8ToJs(inArray)))
	gl.BindTexture(webgl.TEXTURE_2D, &webgl.Texture{})
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

func setupPixelArray(m [][]bool) []uint8 {
	on := []uint8{255, 255, 255, 255}
	off := []uint8{0, 0, 0, 0}
	out := make([]uint8, 4*len(m)*len(m[0]))
	width := len(m[0])
	for i := range m {
		for j := range m[i] {
			if m[i][j] {
				out[4*(i*width+j)+0] = on[0]
				out[4*(i*width+j)+1] = on[1]
				out[4*(i*width+j)+2] = on[2]
				out[4*(i*width+j)+3] = on[3]
			} else {
				out[4*(i*width+j)+0] = off[0]
				out[4*(i*width+j)+1] = off[1]
				out[4*(i*width+j)+2] = off[2]
				out[4*(i*width+j)+3] = off[3]
			}
		}
	}

	return out
}

func (lg *LifeGame) Reset(c *canvas.GameContext) {
	lg.setPixelsInTexture(c, emptyArray(lg.lifeContext.CellWidth, lg.lifeContext.CellHeight))
	// lg.swapTextures()
	lg.drawToCanvas(c)
}

func (lg *LifeGame) Random(c *canvas.GameContext) {
	lg.setPixelsInTexture(c, randomArray(lg.lifeContext.CellWidth, lg.lifeContext.CellHeight))
	lg.deathFrame(c)
	lg.swapTextures()
	lg.drawToCanvas(c)
}

func (lg *LifeGame) OpenFile(c *canvas.GameContext, path string) {
	newArray := parser.ReadRandomFile() // TODO: should return ParsedStuff
	if newArray == nil {
		return
	}
	lg.lifeContext.CellHeight = len(newArray)
	lg.lifeContext.CellWidth = len(newArray[0])
	lg.lifeContext.Zoom = 3
	lg.lifeContext.DX = 0
	lg.lifeContext.DY = 0
	lg.ResizeBuffers(c)

	lg.setPixelsInTexture(c, newArray)
	fmt.Printf("width: %.2f, height: %.2f, cwidth: %d, cheight: %d, zoom: %.2f\n", c.Width, c.Height, lg.lifeContext.CellWidth, lg.lifeContext.CellHeight, lg.lifeContext.Zoom)
}

func (lg *LifeGame) ResizeBuffers(c *canvas.GameContext) {
	canvas.InitCanvas(c)
	lg.createBuffers(c)
}
