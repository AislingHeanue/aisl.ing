package view

import (
	_ "embed"
	"fmt"
	"math"
	"syscall/js"

	"github.com/AislingHeanue/aisling-codes/wasm-demo/canvas"
	cubeController "github.com/AislingHeanue/aisling-codes/wasm-demo/games/rubiks/controller"
	"github.com/AislingHeanue/aisling-codes/wasm-demo/games/rubiks/model"
	"github.com/gowebapi/webapi/core/jsconv"
	"github.com/gowebapi/webapi/graphics/webgl"
)

type CubeRenderer struct {
	TotalSideLength   float32
	SideLength        float32
	GapProportion     float32
	SideLengthWithGap float32
	Origin            model.Point
	Dimension         int
	TurnSeconds       float32

	bufferSet *BufferSet
	program   *webgl.Program

	VertexSource   string
	FragmentSource string

	*model.CubeCubeContext
}

type BufferSet struct {
	Vertices *webgl.Buffer
	Indices  *webgl.Buffer
	Colours  *webgl.Buffer
	VCount   int
	ICount   int
	CCount   int
}

var _ canvas.Animator = &CubeRenderer{}

func (cc *CubeRenderer) Init(c *canvas.GameContext) {
	if cc.program == nil {
		cc.createShaders(c)
	}
}

func (cc CubeRenderer) Dimensions() (int, int) {
	return 0, 0
}

func (cc CubeRenderer) InitListeners(c *canvas.GameContext) {
	cubeController.InitListeners(c, cc.CubeCubeContext)
}

func (cc *CubeRenderer) Render(c *canvas.GameContext) {
	gl := c.GL
	program := cc.program
	animationInProgress := cc.AnimationHandler.Tick(c.IntervalT)
	if cc.bufferSet == nil || animationInProgress {
		cc.createBuffers(c)
	}

	gl.ClearColor(0.9, 0.9, 0.9, 1.0)
	gl.ClearDepth(1.0) // clear all objects
	gl.Enable(webgl.DEPTH_TEST)
	gl.DepthFunc(webgl.LEQUAL)
	gl.Clear(webgl.COLOR_BUFFER_BIT)

	// point the program to the vertex buffer object we've bound
	cc.bindAttribute(c, cc.bufferSet.Vertices, "aVertexPosition", 3)
	cc.bindAttribute(c, cc.bufferSet.Colours, "aVertexColour", 4)

	modelView := model.I4().
		Rotate(math.Pi/4, model.Y).
		Rotate(float32(cc.AngleY), model.Y).
		Rotate(float32(cc.AngleX), model.X).
		Rotate(-math.Pi/5, model.X).
		ToJS()
	matrixLoc := gl.GetUniformLocation(program, "modelView")
	gl.UniformMatrix4fv(matrixLoc, false, modelView)

	perspectiveMatrix := model.PerspectiveMatrix(
		math.Pi/3,
		float32(c.Width/c.Height),
		-2,
		6,
	).ToJS()
	perspectiveMatrixLoc := gl.GetUniformLocation(program, "perspectiveMatrix")
	gl.UniformMatrix4fv(perspectiveMatrixLoc, false, perspectiveMatrix)

	gl.BindBuffer(webgl.ELEMENT_ARRAY_BUFFER, cc.bufferSet.Indices)
	gl.DrawElements(webgl.TRIANGLES, cc.bufferSet.ICount, webgl.UNSIGNED_SHORT, 0)
}

func (cc *CubeRenderer) createShaders(c *canvas.GameContext) {
	gl := c.GL

	vShader := gl.CreateShader(webgl.VERTEX_SHADER)
	gl.ShaderSource(vShader, cc.VertexSource)
	gl.CompileShader(vShader)
	if !gl.GetShaderParameter(vShader, webgl.COMPILE_STATUS).Bool() {
		fmt.Printf("ERROR V: %v", gl.GetShaderInfoLog(vShader))
	}

	fShader := gl.CreateShader(webgl.FRAGMENT_SHADER)
	gl.ShaderSource(fShader, cc.FragmentSource)
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

	gl.UseProgram(program)

	cc.program = program
}

func (cc *CubeRenderer) createBuffers(c *canvas.GameContext) {
	shapeGroup := model.GetBuffers(cc.AnimationHandler, cc.Origin)

	vertices := jsconv.Float32ToJs(shapeGroup.VerticesArray)
	vertexBuffer := bindToBuffer(c, webgl.ARRAY_BUFFER, vertices)

	indices := jsconv.UInt16ToJs(shapeGroup.IndicesArray)
	indexBuffer := bindToBuffer(c, webgl.ELEMENT_ARRAY_BUFFER, indices)

	colours := jsconv.Float32ToJs(shapeGroup.ColourArray)
	colourBuffer := bindToBuffer(c, webgl.ARRAY_BUFFER, colours)

	cc.bufferSet = &BufferSet{
		Vertices: vertexBuffer,
		Indices:  indexBuffer,
		Colours:  colourBuffer,
		VCount:   shapeGroup.VCount,
		ICount:   shapeGroup.ICount,
		CCount:   shapeGroup.CCount,
	}
}

func (cc *CubeRenderer) bindAttribute(c *canvas.GameContext, buffer *webgl.Buffer, name string, stride int) {
	c.GL.BindBuffer(webgl.ARRAY_BUFFER, buffer)
	vPosition := c.GL.GetAttribLocation(cc.program, name)
	c.GL.VertexAttribPointer(uint(vPosition), stride, webgl.FLOAT, false, 0, 0)
	c.GL.EnableVertexAttribArray(uint(vPosition))
	c.GL.BindBuffer(webgl.ARRAY_BUFFER, &webgl.Buffer{})
}

func bindToBuffer(c *canvas.GameContext, target uint, data js.Value) *webgl.Buffer {
	buffer := c.GL.CreateBuffer()
	c.GL.BindBuffer(target, buffer)
	c.GL.BufferData2(target, webgl.UnionFromJS(data), webgl.STATIC_DRAW)
	c.GL.BindBuffer(target, &webgl.Buffer{})

	return buffer
}
