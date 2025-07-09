package view

import (
	"fmt"

	"github.com/gowebapi/webapi/core/jsconv"
	"github.com/gowebapi/webapi/graphics/webgl"
	webapicanvas "github.com/gowebapi/webapi/html/canvas"

	"github.com/AislingHeanue/aisling-codes/wasm-demo/util"
)

func (lg *LifeGame) zoomY(c *util.GameContext) float32 {
	return c.Zoom * c.Height / float32(c.PixelsHeight)
}

func (lg *LifeGame) zoomX(c *util.GameContext) float32 {
	return c.Zoom * c.Width / float32(c.PixelsWidth)
}

func (lg *LifeGame) zoom(c *util.GameContext) float32 {
	return min(lg.zoomX(c), lg.zoomY(c))
}

func (lg *LifeGame) swapTextures() {
	lg.writeFrameBuffer, lg.readFrameBuffer = lg.readFrameBuffer, lg.writeFrameBuffer
	lg.readTexture, lg.writeTexture = lg.writeTexture, lg.readTexture
}

func (lg *LifeGame) createShaders(c *util.GameContext, vertexSource, fragmentSource string) {
	gl := c.GL

	vShader := gl.CreateShader(webgl.VERTEX_SHADER)
	gl.ShaderSource(vShader, vertexSource)
	gl.CompileShader(vShader)
	if !gl.GetShaderParameter(vShader, webgl.COMPILE_STATUS).Bool() {
		fmt.Printf("Error in life.vert: %v\n", *gl.GetShaderInfoLog(vShader))
	}

	fShader := gl.CreateShader(webgl.FRAGMENT_SHADER)
	gl.ShaderSource(fShader, fragmentSource)
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
}

func (lg *LifeGame) createBuffers(c *util.GameContext) {
	// this is a fullscreen quad.
	// this is what's drawn to (ie the coordinates of the framebuffer to draw to)
	vertexArray := []float32{
		-1.0, -1.0,
		1.0, -1.0,
		-1.0, 1.0,
		-1.0, 1.0,
		1.0, -1.0,
		1.0, 1.0,
	}
	lg.vertexBuffer = createVertexBuffer(c, vertexArray)
	lg.vCount = 6

	// this is another fullscreen quad
	// this one is for reading the full contents of the previous frame.
	textureArray := []float32{
		0.0, 0.0,
		1.0, 0.0,
		0.0, 1.0,
		0.0, 1.0,
		1.0, 0.0,
		1.0, 1.0,
	}
	lg.textureBuffer = createVertexBuffer(c, textureArray)

	// create the textures to store the contents of previous frames.
	lg.readTexture = createTexture(c, c.PixelsWidth, c.PixelsHeight)
	lg.writeTexture = createTexture(c, c.PixelsWidth, c.PixelsHeight)

	// the first framebuffer blits to the second texture and vice versa.
	lg.writeFrameBuffer = createFramebuffer(c, lg.writeTexture)
	lg.readFrameBuffer = createFramebuffer(c, lg.readTexture)
}

func createVertexBuffer(c *util.GameContext, vertexArray []float32) *webgl.Buffer {
	vertices := jsconv.Float32ToJs(vertexArray)
	vertexBuffer := c.GL.CreateBuffer()
	c.GL.BindBuffer(webgl.ARRAY_BUFFER, vertexBuffer)
	c.GL.BufferData2(webgl.ARRAY_BUFFER, webgl.UnionFromJS(vertices), webgl.STATIC_DRAW)
	c.GL.BindBuffer(webgl.ARRAY_BUFFER, &webgl.Buffer{})

	return vertexBuffer
}

func createTexture(c *util.GameContext, width int, height int) *webgl.Texture {
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

func createFramebuffer(c *util.GameContext, texture *webgl.Texture) *webgl.Framebuffer {
	frameBuffer := c.GL.CreateFramebuffer()
	c.GL.BindFramebuffer(webgl.FRAMEBUFFER, frameBuffer)
	c.GL.FramebufferTexture2D(webgl.FRAMEBUFFER, webgl.COLOR_ATTACHMENT0, webgl.TEXTURE_2D, texture, 0)
	c.GL.BindFramebuffer(webgl.FRAMEBUFFER, &webgl.Framebuffer{})

	return frameBuffer
}

func (lg *LifeGame) renderFrame(c *util.GameContext) {
	c.GL.BindFramebuffer(webgl.FRAMEBUFFER, lg.writeFrameBuffer)
	c.GL.ClearColor(0.0, 0.0, 0.0, 1.0)
	c.GL.DrawArrays(webgl.TRIANGLES, 0, lg.vCount)
	c.GL.BindFramebuffer(webgl.FRAMEBUFFER, &webgl.Framebuffer{})

	lg.swapTextures()
}

func (lg *LifeGame) Render(c *util.GameContext) {
	lg.cumulativeIntervalT += c.IntervalT
	for lg.cumulativeIntervalT > (1 / lg.Tps) {
		if !lg.Paused {
			lg.t++
		}
		lg.attachAttributes(c)
		lg.renderFrame(c)

		lg.cumulativeIntervalT -= 1 / lg.Tps
	}
	lg.drawToCanvas(c)
}

func (lg *LifeGame) drawToCanvas(c *util.GameContext) {
	// the + c.Width/2 here makes it so that the 'anchor' point for zooming in and out is at the centre of the canvas
	topLeftDX := c.DX + c.Width/2 - lg.zoom(c)*float32(c.PixelsWidth)/2
	topLeftDY := c.DY + c.Height/2 - lg.zoom(c)*float32(c.PixelsHeight)/2
	// bound check DX and DY and make sure they're within a valid range to be able to draw each part of all the visible canvases.
	for topLeftDX > 0 {
		topLeftDX -= float32(c.PixelsWidth) * lg.zoom(c)
	}
	for topLeftDY > 0 {
		topLeftDY -= float32(c.PixelsHeight) * lg.zoom(c)
	}

	union := webgl.Union{
		Value: jsconv.UInt8ToJs(make([]uint8, c.PixelsHeight*c.PixelsWidth*4)),
	}
	c.GL.BindFramebuffer(webgl.FRAMEBUFFER, lg.readFrameBuffer)
	c.GL.ReadPixels(0, 0, c.PixelsWidth, c.PixelsHeight, webgl.RGBA, webgl.UNSIGNED_BYTE, &union)
	c.GL.BindFramebuffer(webgl.FRAMEBUFFER, &webgl.Framebuffer{})

	imageData := c.ZoomCtx.CreateImageData(c.PixelsWidth, c.PixelsHeight)
	imageData.Data().JSValue().Call("set", union.JSValue())
	c.ZoomCtx.PutImageData(imageData, 0, 0)
	c.DisplayCtx.ClearRect(0, 0, float64(c.Width), float64(c.Height))
	// tile horizontally if one instance of the grid does not cover the canvas
	for currentDx := topLeftDX; currentDx < c.Width; currentDx += float32(c.PixelsWidth) * lg.zoom(c) {
		// and vertically
		for currentDy := topLeftDY; currentDy < c.Height; currentDy += float32(c.PixelsHeight) * lg.zoom(c) {
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
