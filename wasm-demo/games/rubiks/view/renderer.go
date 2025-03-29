package view

import (
	"fmt"
	"image/color"
	"math"

	"github.com/AislingHeanue/aisling-codes/wasm-demo/canvas"
	cubeController "github.com/AislingHeanue/aisling-codes/wasm-demo/games/rubiks/controller"
	"github.com/AislingHeanue/aisling-codes/wasm-demo/games/rubiks/model"
	"github.com/gowebapi/webapi/core/jsconv"
	"github.com/gowebapi/webapi/graphics/webgl"
)

type CubeRenderer struct {
	cubes             model.RubiksCube
	totalSide         float32
	side              float32
	gap               float32
	sideWithGap       float32
	dimension         int
	origin            *model.Point
	perspectiveMatrix *model.Mat4
	bufferSet         *BufferSet
	sg                DrawShape
	bufferStale       bool
	program           *webgl.Program
	// animationHandler  *model.RubiksAnimationHandler
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
	cc.dimension = cc.CubeDimension
	cc.totalSide = 0.5
	cc.gap = 0.07
	cc.side = cc.totalSide / ((1+cc.gap)*float32(cc.CubeDimension) - cc.gap)
	cc.sideWithGap = cc.side + cc.gap*cc.side
	cc.origin = &model.Point{0, 0, 0}

	cc.cubes = model.NewRubiksCube(cc.CubeDimension)
	for x := 0; x < cc.dimension; x++ {
		for y := 0; y < cc.dimension; y++ {
			for z := 0; z < cc.dimension; z++ {
				cubeOrigin := cc.getCentre(x, y, z)
				colours, external := cubeColours(x, y, z, cc.dimension)
				if external {
					cc.cubes.Data[x][y][z] = model.NewCubeWithColours(cubeOrigin, cc.side, colours)
				}
			}
		}
	}
	cc.perspectiveMatrix = model.PerspectiveMatrix(
		math.Pi/3,
		float32(c.Width/c.Height),
		-2,
		6,
	)

	cc.AnimationHandler = &model.RubiksAnimationHandler{
		RubiksCube: &cc.cubes,
		MaxTicks:   cc.TurnFrames,
	}

	cc.bufferStale = true

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

func (cc CubeRenderer) getCentre(x, y, z int) model.Point {
	return cc.origin.
		Subtract(model.Point{cc.totalSide / 2, cc.totalSide / 2, cc.totalSide / 2}).
		Add(model.Point{cc.sideWithGap * float32(x), cc.sideWithGap * float32(y), cc.sideWithGap * float32(z)}).
		Add(model.Point{cc.side / 2, cc.side / 2, cc.side / 2})
}

func cubeColours(x, y, z, dimension int) ([]color.RGBA, bool) {
	colours := []color.RGBA{model.BLACK, model.BLACK, model.BLACK, model.BLACK, model.BLACK, model.BLACK}
	external := false
	if y == dimension-1 {
		external = true
		colours[0] = model.DefaultColours[0]
	}
	if z == dimension-1 {
		external = true
		colours[1] = model.DefaultColours[1]
	}
	if x == 0 {
		external = true
		colours[2] = model.DefaultColours[2]
	}
	if z == 0 {
		external = true
		colours[3] = model.DefaultColours[3]
	}
	if x == dimension-1 {
		external = true
		colours[4] = model.DefaultColours[4]
	}
	if y == 0 {
		external = true
		colours[5] = model.DefaultColours[5]
	}

	return colours, external
}

func (cc *CubeRenderer) createShaders(c *canvas.GameContext) {
	gl := c.GL

	vsSource := `
	attribute vec3 aVertexPosition;
	attribute vec4 aVertexColour;
	uniform mat4 modelView;
	uniform mat4 perspectiveMatrix;

	varying lowp vec4 vColour;

	void main(void) {
		gl_Position = perspectiveMatrix * modelView * vec4(aVertexPosition,1.0);
		vColour = aVertexColour;
	}
	`
	fsSource := `
	varying lowp vec4 vColour;

	void main(void) {
		gl_FragColor = vColour;
	}
	`
	vShader := gl.CreateShader(webgl.VERTEX_SHADER)
	gl.ShaderSource(vShader, vsSource)
	gl.CompileShader(vShader)
	if !gl.GetShaderParameter(vShader, webgl.COMPILE_STATUS).Bool() {
		fmt.Printf("ERROR V: %v", gl.GetShaderInfoLog(vShader))
	}

	fShader := gl.CreateShader(webgl.FRAGMENT_SHADER)
	gl.ShaderSource(fShader, fsSource)
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

	cc.program = program
}

func (cc *CubeRenderer) Render(c *canvas.GameContext) {
	gl := c.GL
	program := cc.program
	animationStale := cc.AnimationHandler.Tick()
	if cc.bufferStale || animationStale {
		cc.sg = GetBuffersForRubiksAnimator(cc.AnimationHandler, cc.origin)
		cc.createBuffers(gl)
	}

	gl.ClearColor(0.9, 0.9, 0.9, 1.0)
	gl.ClearDepth(1.0) // clear all objects
	gl.Enable(webgl.DEPTH_TEST)
	gl.DepthFunc(webgl.LEQUAL)

	gl.Clear(webgl.COLOR_BUFFER_BIT)

	gl.BindBuffer(webgl.ARRAY_BUFFER, cc.bufferSet.Vertices)
	vPosition := gl.GetAttribLocation(program, "aVertexPosition")
	// point the program to the vertex buffer object we've bound
	gl.VertexAttribPointer(uint(vPosition), 3, webgl.FLOAT, false, 0, 0)
	gl.EnableVertexAttribArray(uint(vPosition))

	gl.BindBuffer(webgl.ARRAY_BUFFER, cc.bufferSet.Colours)
	vColour := gl.GetAttribLocation(program, "aVertexColour")
	gl.VertexAttribPointer(uint(vColour), 4, webgl.FLOAT, false, 0, 0)
	gl.EnableVertexAttribArray(uint(vColour))

	matrix := model.I4().
		Rotate(math.Pi/4, model.Y).
		Rotate(float32(cc.AngleY), model.Y).
		Rotate(float32(cc.AngleX), model.X).
		Rotate(-math.Pi/5, model.X)
	matrixLoc := gl.GetUniformLocation(program, "modelView")
	gl.UniformMatrix4fv(matrixLoc, false, matrix.ToJS())

	perspectiveMatrixLoc := gl.GetUniformLocation(program, "perspectiveMatrix")
	gl.UniformMatrix4fv(perspectiveMatrixLoc, false, cc.perspectiveMatrix.ToJS())

	gl.BindBuffer(webgl.ELEMENT_ARRAY_BUFFER, cc.bufferSet.Indices)

	gl.DrawElements(webgl.TRIANGLES, cc.bufferSet.ICount, webgl.UNSIGNED_SHORT, 0)
}

func (cc *CubeRenderer) createBuffers(gl *webgl.RenderingContext) {
	vertices := jsconv.Float32ToJs(cc.sg.VerticesArray)
	vertexBuffer := gl.CreateBuffer()
	gl.BindBuffer(webgl.ARRAY_BUFFER, vertexBuffer)
	gl.BufferData2(webgl.ARRAY_BUFFER, webgl.UnionFromJS(vertices), webgl.STATIC_DRAW)
	gl.BindBuffer(webgl.ARRAY_BUFFER, &webgl.Buffer{})

	indices := jsconv.UInt16ToJs(cc.sg.IndicesArray)
	indexBuffer := gl.CreateBuffer()
	gl.BindBuffer(webgl.ELEMENT_ARRAY_BUFFER, indexBuffer)
	gl.BufferData2(webgl.ELEMENT_ARRAY_BUFFER, webgl.UnionFromJS(indices), webgl.STATIC_DRAW)
	gl.BindBuffer(webgl.ELEMENT_ARRAY_BUFFER, &webgl.Buffer{})

	colours := jsconv.Float32ToJs(cc.sg.ColourArray)
	colorBuffer := gl.CreateBuffer()
	gl.BindBuffer(webgl.ARRAY_BUFFER, colorBuffer)
	gl.BufferData2(webgl.ARRAY_BUFFER, webgl.UnionFromJS(colours), webgl.STATIC_DRAW)
	gl.BindBuffer(webgl.ARRAY_BUFFER, &webgl.Buffer{})

	// fmt.Println("made some buffers")

	cc.bufferSet = &BufferSet{
		Vertices: vertexBuffer,
		Indices:  indexBuffer,
		Colours:  colorBuffer,
		VCount:   cc.sg.VCount,
		ICount:   cc.sg.ICount,
		CCount:   cc.sg.CCount,
	}
	cc.bufferStale = false
}
