package common

import (
	_ "embed"
	"fmt"
	"syscall/js"

	"github.com/gowebapi/webapi/core/jsconv"
	"github.com/gowebapi/webapi/graphics/webgl"
	webapicanvas "github.com/gowebapi/webapi/html/canvas"
)

//go:embed shaders/display.vert
var vDisplayShaderSource string

//go:embed shaders/display.frag
var fDisplayShaderSource string

type ShaderGame struct {
	GameInfo
	mainProgram    *webgl.Program
	displayProgram *webgl.Program

	vCount int

	// for 2D stuff
	writeBuffer      *webgl.Buffer
	readBuffer       *webgl.Buffer
	writeFrameBuffer *webgl.Framebuffer
	readFrameBuffer  *webgl.Framebuffer
	writeTexture     *webgl.Texture
	readTexture      *webgl.Texture
	readDepthBuffer  *webgl.Renderbuffer
	writeDepthBuffer *webgl.Renderbuffer

	// for 3D stuff
	bufferSet *BufferSet

	cumulativeIntervalT float32
}

type BufferSet struct {
	Vertices *webgl.Buffer
	Indices  *webgl.Buffer
	Colours  *webgl.Buffer
	Edges    *webgl.Buffer
	VCount   int
	ICount   int
	CCount   int
	ECount   int
}

type DrawShape struct {
	VerticesArray []float32
	IndicesArray  []uint16
	ColourArray   []float32
	EdgeIndices   []uint16
	VCount        int
	ICount        int
	CCount        int
	ECount        int
}

type GameInfo interface {
	AttachAttributes(c *GameContext, program *webgl.Program, writeBuffer, readBuffer *webgl.Buffer, samplerTexture *webgl.Texture)
	PreSetup(c *GameContext)
	PostSetup(c *GameContext)
	InitListeners(c *GameContext)
	GetTps() float32
	GetVertexSource() string
	GetFragmentSource() string
	SetParent(parent *ShaderGame)
	Tick(c *GameContext) bool // true = buffers need to be remade
	GetDrawShape(c *GameContext) DrawShape
	CanRunBetweenFrames() bool // decides whether the program (shader code) is allowed to iterate more than once per frame.
	SkipThisFrame(c *GameContext) bool
}

var _ Animator = &ShaderGame{}

func (g *ShaderGame) Init(c *GameContext) {
	g.createShaders(c, g.GetVertexSource(), g.GetFragmentSource())
	g.PreSetup(c)
	// for game of life, buffers must exist before calling init
	g.CreateBuffers(c)

	if g.zoom(c) <= 0 {
		panic("I refuse to create an infinite loop no thank you")
	}
	c.DX = 0 // default offset position of the grid is 0,0
	c.DY = 0
	if c.Is3D {
		c.GL.Enable(webgl.DEPTH_TEST)
		c.GL.DepthFunc(webgl.LEQUAL)
	}

	g.SetParent(g)
	g.PostSetup(c)
}

func (g *ShaderGame) InitListeners(c *GameContext) {
	g.GameInfo.InitListeners(c)
}

func (g *ShaderGame) RefreshBuffers(c *GameContext) {
	g.CreateBuffers(c)
}

func (g *ShaderGame) zoomY(c *GameContext) float32 {
	return c.Zoom * c.Height / float32(c.PixelsHeight)
}

func (g *ShaderGame) zoomX(c *GameContext) float32 {
	return c.Zoom * c.Width / float32(c.PixelsWidth)
}

func (g *ShaderGame) zoom(c *GameContext) float32 {
	return min(g.zoomX(c), g.zoomY(c))
}

func (g *ShaderGame) swapTextures() {
	g.writeFrameBuffer, g.readFrameBuffer = g.readFrameBuffer, g.writeFrameBuffer
	g.readTexture, g.writeTexture = g.writeTexture, g.readTexture
}

func (g *ShaderGame) createShaders(c *GameContext, vertexSource, fragmentSource string) {
	gl := c.GL

	vShader := gl.CreateShader(webgl.VERTEX_SHADER)
	gl.ShaderSource(vShader, vertexSource)
	gl.CompileShader(vShader)
	if !gl.GetShaderParameter(vShader, webgl.COMPILE_STATUS).Bool() {
		fmt.Printf("Error in vertex shader: %v\n", *gl.GetShaderInfoLog(vShader))
	}

	fShader := gl.CreateShader(webgl.FRAGMENT_SHADER)
	gl.ShaderSource(fShader, fragmentSource)
	gl.CompileShader(fShader)
	if !gl.GetShaderParameter(fShader, webgl.COMPILE_STATUS).Bool() {
		fmt.Printf("Error in fragment shader: %v\n", *gl.GetShaderInfoLog(fShader))
	}

	g.mainProgram = gl.CreateProgram()
	gl.AttachShader(g.mainProgram, vShader)
	gl.AttachShader(g.mainProgram, fShader)
	gl.LinkProgram(g.mainProgram)
	if !gl.GetProgramParameter(g.mainProgram, webgl.LINK_STATUS).Bool() {
		fmt.Printf("Error in linking program: %v\n", *gl.GetProgramInfoLog(g.mainProgram))
	}

	vDisplayShader := gl.CreateShader(webgl.VERTEX_SHADER)
	gl.ShaderSource(vDisplayShader, vDisplayShaderSource)
	gl.CompileShader(vDisplayShader)
	if !gl.GetShaderParameter(vShader, webgl.COMPILE_STATUS).Bool() {
		fmt.Printf("Error in display.vert: %v\n", *gl.GetShaderInfoLog(vDisplayShader))
	}

	fDisplayShader := gl.CreateShader(webgl.FRAGMENT_SHADER)
	gl.ShaderSource(fDisplayShader, fDisplayShaderSource)
	gl.CompileShader(fDisplayShader)
	if !gl.GetShaderParameter(fDisplayShader, webgl.COMPILE_STATUS).Bool() {
		fmt.Printf("Error in display.frag: %v\n", *gl.GetShaderInfoLog(fDisplayShader))
	}

	g.displayProgram = gl.CreateProgram()
	gl.AttachShader(g.displayProgram, vDisplayShader)
	gl.AttachShader(g.displayProgram, fDisplayShader)
	gl.LinkProgram(g.displayProgram)
	if !gl.GetProgramParameter(g.displayProgram, webgl.LINK_STATUS).Bool() {
		fmt.Printf("Error in linking display %v\n", *gl.GetProgramInfoLog(g.displayProgram))
	}
}

func (g *ShaderGame) CreateBuffers(c *GameContext) {
	if c.Is3D {
		drawShape := g.GetDrawShape(c)

		vertices := jsconv.Float32ToJs(drawShape.VerticesArray)
		vertexBuffer := bindToBuffer(c, webgl.ARRAY_BUFFER, vertices)

		indices := jsconv.UInt16ToJs(drawShape.IndicesArray)
		indexBuffer := bindToBuffer(c, webgl.ELEMENT_ARRAY_BUFFER, indices)

		colours := jsconv.Float32ToJs(drawShape.ColourArray)
		colourBuffer := bindToBuffer(c, webgl.ARRAY_BUFFER, colours)

		edges := jsconv.UInt16ToJs(drawShape.EdgeIndices)
		edgeBuffer := bindToBuffer(c, webgl.ELEMENT_ARRAY_BUFFER, edges)

		g.bufferSet = &BufferSet{
			Vertices: vertexBuffer,
			Indices:  indexBuffer,
			Colours:  colourBuffer,
			Edges:    edgeBuffer,
			VCount:   drawShape.VCount,
			ICount:   drawShape.ICount,
			CCount:   drawShape.CCount,
			ECount:   drawShape.ECount,
		}
	}
	// this is a fullscreen quad.
	// this is what's drawn to (ie the coordinates of the framebuffer to draw to)
	writeArray := []float32{
		-1.0, -1.0,
		1.0, -1.0,
		-1.0, 1.0,
		-1.0, 1.0,
		1.0, -1.0,
		1.0, 1.0,
	}
	g.writeBuffer = createVertexBuffer(c, writeArray)
	g.vCount = 6

	// this is another fullscreen quad
	// this one is for reading the full contents of the previous frame.
	readArray := []float32{
		0.0, 0.0,
		1.0, 0.0,
		0.0, 1.0,
		0.0, 1.0,
		1.0, 0.0,
		1.0, 1.0,
	}
	g.readBuffer = createVertexBuffer(c, readArray)

	// create the textures to store the contents of previous frames.
	g.writeTexture = createTexture(c, c.PixelsWidth, c.PixelsHeight)
	g.readTexture = createTexture(c, c.PixelsWidth, c.PixelsHeight)

	g.writeFrameBuffer = createFramebuffer(c, g.writeTexture)
	g.readFrameBuffer = createFramebuffer(c, g.readTexture)

	g.writeDepthBuffer = createDepthBuffer(c, g.writeFrameBuffer, c.PixelsWidth, c.PixelsHeight)
	g.readDepthBuffer = createDepthBuffer(c, g.readFrameBuffer, c.PixelsWidth, c.PixelsHeight)
}

func createVertexBuffer(c *GameContext, vertexArray []float32) *webgl.Buffer {
	vertices := jsconv.Float32ToJs(vertexArray)
	vertexBuffer := c.GL.CreateBuffer()
	c.GL.BindBuffer(webgl.ARRAY_BUFFER, vertexBuffer)
	c.GL.BufferData2(webgl.ARRAY_BUFFER, webgl.UnionFromJS(vertices), webgl.STATIC_DRAW)
	c.GL.BindBuffer(webgl.ARRAY_BUFFER, &webgl.Buffer{})

	return vertexBuffer
}

func createTexture(c *GameContext, width int, height int) *webgl.Texture {
	t := c.GL.CreateTexture()
	c.GL.BindTexture(webgl.TEXTURE_2D, t)
	c.GL.TexImage2D(webgl.TEXTURE_2D, 0, int(webgl.RGBA), width, height, 0, webgl.RGBA, webgl.UNSIGNED_BYTE, &webgl.Union{})
	c.GL.TexParameteri(webgl.TEXTURE_2D, webgl.TEXTURE_WRAP_S, int(webgl.CLAMP_TO_EDGE))
	c.GL.TexParameteri(webgl.TEXTURE_2D, webgl.TEXTURE_WRAP_T, int(webgl.CLAMP_TO_EDGE))
	c.GL.TexParameteri(webgl.TEXTURE_2D, webgl.TEXTURE_MIN_FILTER, int(webgl.NEAREST))
	c.GL.TexParameteri(webgl.TEXTURE_2D, webgl.TEXTURE_MAG_FILTER, int(webgl.NEAREST))
	c.GL.BindTexture(webgl.TEXTURE_2D, &webgl.Texture{})

	return t
}

func createFramebuffer(c *GameContext, texture *webgl.Texture) *webgl.Framebuffer {
	frameBuffer := c.GL.CreateFramebuffer()
	c.GL.BindFramebuffer(webgl.FRAMEBUFFER, frameBuffer)
	c.GL.FramebufferTexture2D(webgl.FRAMEBUFFER, webgl.COLOR_ATTACHMENT0, webgl.TEXTURE_2D, texture, 0)
	c.GL.BindFramebuffer(webgl.FRAMEBUFFER, &webgl.Framebuffer{})

	return frameBuffer
}

func createDepthBuffer(c *GameContext, frameBuffer *webgl.Framebuffer, width int, height int) *webgl.Renderbuffer {
	depthBuffer := c.GL.CreateRenderbuffer()
	c.GL.BindRenderbuffer(webgl.RENDERBUFFER, depthBuffer)
	c.GL.BindFramebuffer(webgl.FRAMEBUFFER, frameBuffer)
	c.GL.RenderbufferStorage(webgl.RENDERBUFFER, webgl.DEPTH_COMPONENT16, width, height)
	c.GL.FramebufferRenderbuffer(webgl.FRAMEBUFFER, webgl.DEPTH_ATTACHMENT, webgl.RENDERBUFFER, depthBuffer)
	c.GL.BindRenderbuffer(webgl.RENDERBUFFER, &webgl.Renderbuffer{})
	c.GL.BindFramebuffer(webgl.FRAMEBUFFER, &webgl.Framebuffer{})

	return depthBuffer
}

func bindToBuffer(c *GameContext, target uint, data js.Value) *webgl.Buffer {
	buffer := c.GL.CreateBuffer()
	c.GL.BindBuffer(target, buffer)
	c.GL.BufferData2(target, webgl.UnionFromJS(data), webgl.STATIC_DRAW)
	c.GL.BindBuffer(target, &webgl.Buffer{})

	return buffer
}

func (g *ShaderGame) Render(c *GameContext) {
	c.GL.UseProgram(g.mainProgram)
	g.cumulativeIntervalT += c.IntervalT
	tps := g.GetTps()
	buffersStale := false
	for g.cumulativeIntervalT > (1 / tps) {
		if g.Tick(c) {
			buffersStale = true
		}
		if g.CanRunBetweenFrames() {
			// game of life runs on in-between frames, since it may run multiple physics frames in one draw frame
			g.runProgram(c, buffersStale)
			buffersStale = false
		}
		g.cumulativeIntervalT -= 1. / tps
	}
	if !g.CanRunBetweenFrames() {
		g.runProgram(c, buffersStale)
	}
	g.drawToCanvas(c)
}

func (g *ShaderGame) runProgram(c *GameContext, buffersStale bool) {
	if buffersStale {
		g.CreateBuffers(c)
	}
	if !g.SkipThisFrame(c) {
		g.AttachAttributes(c, g.mainProgram, g.writeBuffer, g.readBuffer, g.readTexture)
		g.renderFrame(c)
	}
}

func (g *ShaderGame) renderFrame(c *GameContext) {
	if c.Is3D {
		c.GL.BindFramebuffer(webgl.FRAMEBUFFER, g.writeFrameBuffer)
		c.GL.ClearColor(0.9, 0.9, 0.9, 1.0)
		c.GL.ClearDepth(1.0) // clear all objects
		c.GL.Clear(webgl.COLOR_BUFFER_BIT)
		c.GL.Clear(webgl.DEPTH_BUFFER_BIT)

		// point the program to the vertex buffer object we've bound
		g.bindArrayBuffer(c, g.mainProgram, g.bufferSet.Vertices, "aVertexPosition", 3)
		g.bindArrayBuffer(c, g.mainProgram, g.bufferSet.Colours, "aVertexColour", 4)

		useUniformLoc := c.GL.GetUniformLocation(g.mainProgram, "useUniformColour")
		c.GL.Uniform1f(useUniformLoc, 0.)

		c.GL.BindBuffer(webgl.ELEMENT_ARRAY_BUFFER, g.bufferSet.Indices)
		c.GL.DrawElements(webgl.TRIANGLES, g.bufferSet.ICount, webgl.UNSIGNED_SHORT, 0)
		if c.EdgeWidth != 0 {
			c.GL.LineWidth(c.EdgeWidth) // thickness of the border
			c.GL.BindBuffer(webgl.ELEMENT_ARRAY_BUFFER, g.bufferSet.Edges)

			colourLoc := c.GL.GetUniformLocation(g.mainProgram, "uColour")
			c.GL.Uniform4f(colourLoc, 0., 0., 0., 1.)

			useUniformLoc := c.GL.GetUniformLocation(g.mainProgram, "useUniformColour")
			c.GL.Uniform1f(useUniformLoc, 1.)

			c.GL.DrawElements(webgl.LINES, g.bufferSet.ECount, webgl.UNSIGNED_SHORT, 0)
		}
		c.GL.BindFramebuffer(webgl.FRAMEBUFFER, &webgl.Framebuffer{})
	} else {
		c.GL.BindFramebuffer(webgl.FRAMEBUFFER, g.writeFrameBuffer)
		c.GL.ClearColor(0.0, 0.0, 0.0, 1.0)
		c.GL.DrawArrays(webgl.TRIANGLES, 0, g.vCount)
		c.GL.BindFramebuffer(webgl.FRAMEBUFFER, &webgl.Framebuffer{})
	}

	g.swapTextures()
}

func (g *ShaderGame) bindArrayBuffer(c *GameContext, program *webgl.Program, buffer *webgl.Buffer, name string, stride int) {
	c.GL.BindBuffer(webgl.ARRAY_BUFFER, buffer)
	attributeLocation := c.GL.GetAttribLocation(program, name)
	c.GL.VertexAttribPointer(uint(attributeLocation), stride, webgl.FLOAT, false, 0, 0)
	c.GL.EnableVertexAttribArray(uint(attributeLocation))
	c.GL.BindBuffer(webgl.ARRAY_BUFFER, &webgl.Buffer{})
}

func (g *ShaderGame) drawToCanvas(c *GameContext) {
	// the + c.Width/2 here makes it so that the 'anchor' point for zooming in and out is at the centre of the canvas
	topLeftDX := c.DX + c.Width/2 - g.zoom(c)*float32(c.PixelsWidth)/2
	bottomLeftDY := c.DY + c.Height/2 - g.zoom(c)*float32(c.PixelsHeight)/2
	// bound check DX and DY and make sure they're within a valid range to be able to draw each part of all the visible canvases.

	for topLeftDX > 0 {
		topLeftDX -= float32(c.PixelsWidth) * g.zoom(c)
	}
	for bottomLeftDY > 0 {
		bottomLeftDY -= float32(c.PixelsHeight) * g.zoom(c)
	}

	union := webgl.Union{
		Value: jsconv.UInt8ToJs(make([]uint8, c.PixelsHeight*c.PixelsWidth*4)),
	}
	c.GL.BindFramebuffer(webgl.FRAMEBUFFER, g.readFrameBuffer)
	c.GL.ReadPixels(0, 0, c.PixelsWidth, c.PixelsHeight, webgl.RGBA, webgl.UNSIGNED_BYTE, &union)
	c.GL.BindFramebuffer(webgl.FRAMEBUFFER, &webgl.Framebuffer{})

	imageData := c.DisplayCtx.CreateImageData(c.PixelsWidth, c.PixelsHeight)
	imageData.Data().JSValue().Call("set", union.JSValue())
	c.ZoomCtx.PutImageData(imageData, 0, 0)

	c.DisplayCtx.ClearRect(0, 0, float64(c.Width), float64(c.Height))
	// tile horizontally if one instance of the grid does not cover the canvas
	for currentDx := topLeftDX; currentDx < c.Width; currentDx += float32(c.PixelsWidth) * g.zoom(c) {
		// and vertically
		for currentDy := bottomLeftDY; currentDy < c.Height; currentDy += float32(c.PixelsHeight) * g.zoom(c) {
			c.DisplayCtx.DrawImage3(
				webapicanvas.UnionFromJS(c.ZoomCanvas.JSValue()),
				0, 0, // start coords in grid being captured from
				float64((c.Width-currentDx)/g.zoom(c)), float64((c.Height-currentDy)/g.zoom(c)),
				float64(currentDx), float64(currentDy), // start coords in grid being displayed to
				float64(c.Width-currentDx), float64(c.Height-currentDy),
			)
		}
	}
}

func (g *ShaderGame) SetPixelsInTexture(c *GameContext, in [][]bool) {
	c.GL.BindTexture(webgl.TEXTURE_2D, g.writeTexture)
	c.GL.TexImage2D(webgl.TEXTURE_2D, 0, int(webgl.RGBA), c.PixelsWidth, c.PixelsHeight, 0, webgl.RGBA, webgl.UNSIGNED_BYTE, webgl.UnionFromJS(jsconv.UInt8ToJs(setupPixelArray(in))))
	c.GL.BindTexture(webgl.TEXTURE_2D, &webgl.Texture{})

	g.swapTextures()
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
