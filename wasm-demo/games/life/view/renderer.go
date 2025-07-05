package view

import (
	_ "embed"
	"fmt"
	"math"

	"github.com/gowebapi/webapi/core/jsconv"
	"github.com/gowebapi/webapi/graphics/webgl"
	webapicanvas "github.com/gowebapi/webapi/html/canvas"

	"github.com/AislingHeanue/aisling-codes/wasm-demo/canvas"
)

func (lg *LifeGame) zoomY(c *canvas.GameContext) float32 {
	return lg.Zoom * c.Height / float32(lg.CellHeight)
}

func (lg *LifeGame) zoomX(c *canvas.GameContext) float32 {
	return lg.Zoom * c.Width / float32(lg.CellWidth)
}

func (lg *LifeGame) zoom(c *canvas.GameContext) float32 {
	return min(lg.zoomX(c), lg.zoomY(c))
}

func (lg *LifeGame) swapTextures() {
	lg.writeFrameBuffer, lg.readFrameBuffer = lg.readFrameBuffer, lg.writeFrameBuffer
	lg.readTexture, lg.writeTexture = lg.writeTexture, lg.readTexture
}

func (lg *LifeGame) getDeadColour() (float32, float32, float32) {
	return float32(0.5 * (1 + math.Sin(2*math.Pi*float64(lg.t)/float64(lg.ColourPeriod)))),
		float32(0.5 * (1 + math.Sin((2*math.Pi/3)+2*math.Pi*float64(lg.t)/float64(lg.ColourPeriod)))),
		float32(0.5 * (1 + math.Sin((4*math.Pi/3)+2*math.Pi*float64(lg.t)/float64(lg.ColourPeriod))))
}

func (lg *LifeGame) createShaders(c *canvas.GameContext) {
	gl := c.GL

	vShader := gl.CreateShader(webgl.VERTEX_SHADER)
	gl.ShaderSource(vShader, lg.VertexSource)
	gl.CompileShader(vShader)
	if !gl.GetShaderParameter(vShader, webgl.COMPILE_STATUS).Bool() {
		fmt.Printf("Error in life.vert: %v\n", *gl.GetShaderInfoLog(vShader))
	}

	fShader := gl.CreateShader(webgl.FRAGMENT_SHADER)
	gl.ShaderSource(fShader, lg.LifeFragmentSource)
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
	gl.ShaderSource(fDisplayShader, lg.DisplayFragmentSource)
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
	gl.ShaderSource(fDeathShader, lg.DeathFragmentSource)
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
	vertexArray := []float32{
		-1.0, -1.0,
		1.0, -1.0,
		-1.0, 1.0,
		-1.0, 1.0,
		1.0, -1.0,
		1.0, 1.0,
	}

	vertices := jsconv.Float32ToJs(vertexArray)
	vertexBuffer := c.GL.CreateBuffer()
	c.GL.BindBuffer(webgl.ARRAY_BUFFER, vertexBuffer)
	c.GL.BufferData2(webgl.ARRAY_BUFFER, webgl.UnionFromJS(vertices), webgl.STATIC_DRAW)
	c.GL.BindBuffer(webgl.ARRAY_BUFFER, &webgl.Buffer{})
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
	textureBuffer := c.GL.CreateBuffer()
	c.GL.BindBuffer(webgl.ARRAY_BUFFER, textureBuffer)
	c.GL.BufferData2(webgl.ARRAY_BUFFER, webgl.UnionFromJS(textureCorners), webgl.STATIC_DRAW)
	c.GL.BindBuffer(webgl.ARRAY_BUFFER, &webgl.Buffer{})
	lg.textureBuffer = textureBuffer

	pixels := randomArray(lg.CellWidth, lg.CellHeight)
	pixelsArray := setupPixelArray(pixels)

	lg.readTexture = c.GL.CreateTexture()
	c.GL.BindTexture(webgl.TEXTURE_2D, lg.readTexture)
	c.GL.TexImage2D(webgl.TEXTURE_2D, 0, int(webgl.RGBA), lg.CellWidth, lg.CellHeight, 0, webgl.RGBA, webgl.UNSIGNED_BYTE, &webgl.Union{})
	c.GL.TexParameteri(webgl.TEXTURE_2D, webgl.TEXTURE_WRAP_S, int(webgl.CLAMP_TO_EDGE))
	c.GL.TexParameteri(webgl.TEXTURE_2D, webgl.TEXTURE_WRAP_T, int(webgl.CLAMP_TO_EDGE))
	c.GL.TexParameteri(webgl.TEXTURE_2D, webgl.TEXTURE_MIN_FILTER, int(webgl.NEAREST))
	c.GL.TexParameteri(webgl.TEXTURE_2D, webgl.TEXTURE_MAG_FILTER, int(webgl.NEAREST))

	lg.writeTexture = c.GL.CreateTexture()
	c.GL.BindTexture(webgl.TEXTURE_2D, lg.writeTexture)
	c.GL.TexImage2D(webgl.TEXTURE_2D, 0, int(webgl.RGBA), lg.CellWidth, lg.CellHeight, 0, webgl.RGBA, webgl.UNSIGNED_BYTE, webgl.UnionFromJS(jsconv.UInt8ToJs(pixelsArray)))
	c.GL.TexParameteri(webgl.TEXTURE_2D, webgl.TEXTURE_WRAP_S, int(webgl.CLAMP_TO_EDGE))
	c.GL.TexParameteri(webgl.TEXTURE_2D, webgl.TEXTURE_WRAP_T, int(webgl.CLAMP_TO_EDGE))
	c.GL.TexParameteri(webgl.TEXTURE_2D, webgl.TEXTURE_MIN_FILTER, int(webgl.NEAREST))
	c.GL.TexParameteri(webgl.TEXTURE_2D, webgl.TEXTURE_MAG_FILTER, int(webgl.NEAREST))

	lg.writeFrameBuffer = c.GL.CreateFramebuffer()
	c.GL.BindFramebuffer(webgl.FRAMEBUFFER, lg.writeFrameBuffer)
	c.GL.FramebufferTexture2D(webgl.FRAMEBUFFER, webgl.COLOR_ATTACHMENT0, webgl.TEXTURE_2D, lg.readTexture, 0)
	lg.readFrameBuffer = c.GL.CreateFramebuffer()
	c.GL.BindFramebuffer(webgl.FRAMEBUFFER, lg.readFrameBuffer)
	c.GL.FramebufferTexture2D(webgl.FRAMEBUFFER, webgl.COLOR_ATTACHMENT0, webgl.TEXTURE_2D, lg.writeTexture, 0)

	c.GL.BindTexture(webgl.TEXTURE_2D, &webgl.Texture{})
	c.GL.BindFramebuffer(webgl.FRAMEBUFFER, &webgl.Framebuffer{})
}

func (lg *LifeGame) deathFrame(c *canvas.GameContext) {
	gl := c.GL
	gl.BindFramebuffer(webgl.FRAMEBUFFER, lg.writeFrameBuffer)

	gl.ClearColor(0.0, 0.0, 1.0, 1.0)
	gl.BindBuffer(webgl.ARRAY_BUFFER, lg.vertexBuffer)
	vPosition := gl.GetAttribLocation(lg.deathProgram, "a_position")
	gl.VertexAttribPointer(uint(vPosition), 2, webgl.FLOAT, false, 0, 0)
	gl.EnableVertexAttribArray(uint(vPosition))

	gl.BindBuffer(webgl.ARRAY_BUFFER, lg.textureBuffer)
	tPosition := gl.GetAttribLocation(lg.deathProgram, "a_tex_coord")
	gl.VertexAttribPointer(uint(tPosition), 2, webgl.FLOAT, false, 0, 0)
	gl.EnableVertexAttribArray(uint(tPosition))
	gl.UseProgram(lg.deathProgram)
	decayLoc := gl.GetUniformLocation(lg.deathProgram, "u_decay")
	gl.Uniform1f(decayLoc, 0.66/float32(lg.TrailLength))

	gl.BindTexture(webgl.TEXTURE_2D, lg.writeTexture)
	samplerLocation := gl.GetUniformLocation(lg.deathProgram, "u_sampler")
	gl.Uniform1i(samplerLocation, 0)

	sizeLoc := gl.GetUniformLocation(lg.deathProgram, "u_size")
	gl.Uniform2f(sizeLoc, float32(lg.CellWidth), float32(lg.CellHeight))

	gl.DrawArrays(webgl.TRIANGLES, 0, lg.vCount)
	gl.BindFramebuffer(webgl.FRAMEBUFFER, &webgl.Framebuffer{})
	lg.swapTextures()
}

func (lg *LifeGame) lifeFrame(c *canvas.GameContext) {
	gl := c.GL
	gl.BindFramebuffer(webgl.FRAMEBUFFER, lg.writeFrameBuffer)

	gl.ClearColor(0.0, 0.0, 1.0, 1.0)
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

	gl.BindTexture(webgl.TEXTURE_2D, lg.writeTexture)
	samplerLocation := gl.GetUniformLocation(lg.lifeProgram, "u_sampler")
	gl.Uniform1i(samplerLocation, 0)

	sizeLoc := gl.GetUniformLocation(lg.lifeProgram, "u_size")
	gl.Uniform2f(sizeLoc, float32(lg.CellWidth), float32(lg.CellHeight))

	gl.DrawArrays(webgl.TRIANGLES, 0, lg.vCount)
	gl.BindFramebuffer(webgl.FRAMEBUFFER, &webgl.Framebuffer{})
	lg.swapTextures()
}

func (lg *LifeGame) drawToCanvas(c *canvas.GameContext) {
	// the + c.Width/2 here makes it so that the 'anchor' point for zooming in and out is at the centre of the canvas
	topLeftDX := lg.DX + c.Width/2 - lg.zoom(c)*float32(lg.CellWidth)/2
	topLeftDY := lg.DY + c.Height/2 - lg.zoom(c)*float32(lg.CellHeight)/2
	// bound check DX and DY and make sure they're within a valid range to be able to draw each part of all the visible canvases.
	for topLeftDX > 0 {
		topLeftDX -= float32(lg.CellWidth) * lg.zoom(c)
	}
	for topLeftDY > 0 {
		topLeftDY -= float32(lg.CellHeight) * lg.zoom(c)
	}

	union := webgl.Union{
		Value: jsconv.UInt8ToJs(make([]uint8, lg.CellHeight*lg.CellWidth*4)),
	}
	c.GL.BindFramebuffer(webgl.FRAMEBUFFER, lg.readFrameBuffer)
	c.GL.ReadPixels(0, 0, lg.CellWidth, lg.CellHeight, webgl.RGBA, webgl.UNSIGNED_BYTE, &union)
	c.GL.BindFramebuffer(webgl.FRAMEBUFFER, &webgl.Framebuffer{})

	imageData := c.ZoomCtx.CreateImageData(lg.CellWidth, lg.CellHeight)
	imageData.Data().JSValue().Call("set", union.JSValue())
	c.ZoomCtx.PutImageData(imageData, 0, 0)
	floatWidth := c.Width   //float32(lg.lifeContext.CellWidth)
	floatHeight := c.Height //float32(lg.lifeContext.CellWidth)

	c.DisplayCtx.ClearRect(0, 0, float64(floatWidth), float64(floatHeight))
	// tile horizontally if one instance of the grid does not cover the canvas
	for currentDx := topLeftDX; currentDx < c.Width; currentDx += float32(lg.CellWidth) * lg.zoom(c) {
		// and vertically
		for currentDy := topLeftDY; currentDy < c.Height; currentDy += float32(lg.CellHeight) * lg.zoom(c) {
			c.DisplayCtx.DrawImage3(
				webapicanvas.UnionFromJS(c.ZoomCanvas.JSValue()),
				0, 0, // start coords in grid being captured from
				float64((c.Width-currentDx)/lg.zoom(c)), float64((c.Height-currentDy)/lg.zoom(c)),
				float64(currentDx), float64(currentDy), // start coords in grid being displayed to
				float64(c.Width-currentDx), float64(c.Height-currentDy),
			)
		}
	}
}
