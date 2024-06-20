package animation

import (
	"fmt"
	"image/color"
	"math"

	"github.com/AislingHeanue/aisling-codes/wasm-demo/maths"
	"github.com/AislingHeanue/aisling-codes/wasm-demo/model"
	"github.com/gowebapi/webapi/core/jsconv"
	"github.com/gowebapi/webapi/graphics/webgl"
)

type CubeCube struct {
	cubes             [][][]*maths.Cube
	shapes            []*maths.Cube
	totalSide         float64
	side              float64
	gap               float64
	sideWithGap       float64
	dimension         int
	origin            *maths.Vector
	perspectiveMatrix *maths.Mat4
	bufferSet         *model.BufferSet
	bufferStale       bool
}

var _ model.Animator = &CubeCube{}

func (cc *CubeCube) Init(c *model.GameContext) {
	cc.dimension = c.Dimension
	cc.totalSide = 0.5
	cc.gap = 0.07
	cc.side = cc.totalSide / ((1+cc.gap)*float64(c.Dimension) - cc.gap)
	cc.sideWithGap = cc.side + cc.gap*cc.side
	cc.origin = &maths.Vector{0, 0, 0}
	cc.shapes = []*maths.Cube{}
	// fmt.Println("somewhere in init")
	// fmt.Println("dimension =")
	// fmt.Println(c.Dimension)
	cc.cubes = make([][][]*maths.Cube, cc.dimension)
	for x := 0; x < cc.dimension; x++ {
		cc.cubes[x] = make([][]*maths.Cube, cc.dimension)
		for y := 0; y < cc.dimension; y++ {
			cc.cubes[x][y] = make([]*maths.Cube, cc.dimension)
			for z := 0; z < cc.dimension; z++ {
				cubeOrigin := cc.getCentre(x, y, z)
				colours, external := cubeColours(x, y, z, cc.dimension)
				if external {
					// fmt.Println("making a cube")
					cc.cubes[x][y][z] = maths.NewCubeWithColours(*cubeOrigin, cc.side, colours)
					cc.shapes = append(cc.shapes, cc.cubes[x][y][z])
				}
			}
		}
	}
	cc.perspectiveMatrix = maths.PerspectiveMatrix(
		math.Pi/3,
		float32(c.Width/c.Height),
		-2,
		6,
	)
	// fmt.Println("innit")
	cc.bufferStale = true
	// gl := c.Gl
	// program := c.Program

	// c.Window.RequestAnimationFrame(htmlcommon.FrameRequestCallbackToJS(wrapAnimator(gl, program, c, cc.Render)))
}

// func wrapAnimator(gl *webgl.RenderingContext, p *webgl.Program, c *model.GameContext, f model.RenderFunc) func(float64) {
// 	return func(time float64) {
// 		c.T = time / 1000 // milliseconds to seconds
// 		f(gl, p, c)
// 		c.Window.RequestAnimationFrame(htmlcommon.FrameRequestCallbackToJS(wrapAnimator(gl, p, c, c.Animator.Render)))
// 	}

// }

func (cc *CubeCube) createBuffers(gl *webgl.RenderingContext) {
	sg := maths.GroupBuffersFromCubes(cc.cubes)
	// fmt.Printlsn(cc.cubes)
	// everything below this line could easily be a second function
	vertices := jsconv.Float32ToJs(sg.VerticesArray)
	vertexBuffer := gl.CreateBuffer()
	gl.BindBuffer(webgl.ARRAY_BUFFER, vertexBuffer)
	gl.BufferData2(webgl.ARRAY_BUFFER, webgl.UnionFromJS(vertices), webgl.STATIC_DRAW)
	gl.BindBuffer(webgl.ARRAY_BUFFER, &webgl.Buffer{})

	indices := jsconv.UInt16ToJs(sg.IndicesArray)
	indexBuffer := gl.CreateBuffer()
	gl.BindBuffer(webgl.ELEMENT_ARRAY_BUFFER, indexBuffer)
	gl.BufferData2(webgl.ELEMENT_ARRAY_BUFFER, webgl.UnionFromJS(indices), webgl.STATIC_DRAW)
	gl.BindBuffer(webgl.ELEMENT_ARRAY_BUFFER, &webgl.Buffer{})

	colours := jsconv.Float32ToJs(sg.ColourArray)
	colorBuffer := gl.CreateBuffer()
	gl.BindBuffer(webgl.ARRAY_BUFFER, colorBuffer)
	gl.BufferData2(webgl.ARRAY_BUFFER, webgl.UnionFromJS(colours), webgl.STATIC_DRAW)
	gl.BindBuffer(webgl.ARRAY_BUFFER, &webgl.Buffer{})

	fmt.Println("made some buffers")

	cc.bufferSet = &model.BufferSet{
		Vertices: vertexBuffer,
		Indices:  indexBuffer,
		Colours:  colorBuffer,
		VCount:   sg.VCount,
		ICount:   sg.ICount,
		CCount:   sg.CCount,
	}
	// fmt.Println("for real")
	cc.bufferStale = false

}

func (cc CubeCube) getCentre(x, y, z int) *maths.Vector {
	return cc.origin.
		Subtract(maths.Vector{cc.totalSide / 2, cc.totalSide / 2, cc.totalSide / 2}).
		Add(maths.Vector{cc.sideWithGap * float64(x), cc.sideWithGap * float64(y), cc.sideWithGap * float64(z)}).
		Add(maths.Vector{cc.side / 2, cc.side / 2, cc.side / 2})
}

func (cc *CubeCube) CreateShaders(gl *webgl.RenderingContext, c *model.GameContext) *webgl.Program {
	// assume we can pass in angleX and angleY.
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

	return program
}

// func (cc *CubeCube) RefreshCoords(c *model.GameContext) {
// 	// fmt.Println(5)
// 	// cc.Sg.CalculatePaths(c)
// 	cc.Sg.Redraw = true
// }

func cubeColours(x, y, z, dimension int) ([]color.RGBA, bool) {
	// White Orange Green Red Blue Yellow
	colours := []color.RGBA{maths.BLACK, maths.BLACK, maths.BLACK, maths.BLACK, maths.BLACK, maths.BLACK}
	external := false
	if y == dimension-1 {
		external = true
		colours[0] = maths.DefaultColours[0]
	}
	if z == dimension-1 {
		external = true
		colours[1] = maths.DefaultColours[1]
	}
	if x == 0 {
		external = true
		colours[2] = maths.DefaultColours[2]
	}
	if z == 0 {
		external = true
		colours[3] = maths.DefaultColours[3]
	}
	if x == dimension-1 {
		external = true
		colours[4] = maths.DefaultColours[4]
	}
	if y == 0 {
		external = true
		colours[5] = maths.DefaultColours[5]
	}

	return colours, external
}

// func (cc *CubeCube) IsRedrawRequired() bool {
// 	if cc.Sg.Redraw {
// 		cc.Sg.Redraw = false

// 		return true
// 	}

// 	return false
// }

func (cc *CubeCube) Render(gl *webgl.RenderingContext, program *webgl.Program, c *model.GameContext) {
	if cc.bufferStale {
		cc.createBuffers(gl)
	}

	gl.ClearColor(0.9, 0.9, 0.9, 1.0)
	gl.ClearDepth(1.0) // clear all objects
	gl.Enable(webgl.DEPTH_TEST)
	gl.DepthFunc(webgl.LEQUAL)

	gl.Clear(webgl.COLOR_BUFFER_BIT)

	// fmt.Println(buffers.ICount)
	// associate shader to buffers
	gl.BindBuffer(webgl.ARRAY_BUFFER, cc.bufferSet.Vertices)
	vPosition := gl.GetAttribLocation(program, "aVertexPosition")
	// point the program to the vertex buffer object we've bound
	gl.VertexAttribPointer(uint(vPosition), 3, webgl.FLOAT, false, 0, 0)
	gl.EnableVertexAttribArray(uint(vPosition))

	gl.BindBuffer(webgl.ARRAY_BUFFER, cc.bufferSet.Colours)
	vColour := gl.GetAttribLocation(program, "aVertexColour")
	gl.VertexAttribPointer(uint(vColour), 4, webgl.FLOAT, false, 0, 0)
	gl.EnableVertexAttribArray(uint(vColour))

	matrix := maths.
		I4().
		Rotate(math.Pi/4, maths.Y).
		Rotate(float32(c.AngleY), maths.Y).
		Rotate(float32(c.AngleX), maths.X).
		Rotate(-math.Pi/5, maths.X)
	matrixLoc := gl.GetUniformLocation(program, "modelView")
	gl.UniformMatrix4fv(matrixLoc, false, matrix.ToJS())

	perspectiveMatrixLoc := gl.GetUniformLocation(program, "perspectiveMatrix")
	gl.UniformMatrix4fv(perspectiveMatrixLoc, false, cc.perspectiveMatrix.ToJS())

	gl.BindBuffer(webgl.ELEMENT_ARRAY_BUFFER, cc.bufferSet.Indices)

	gl.Viewport(0, 0, int(c.Width), int(c.Height))

	gl.DrawElements(webgl.TRIANGLES, cc.bufferSet.ICount, webgl.UNSIGNED_SHORT, 0)
}
