package rubiks

import (
	_ "embed"
	"math"

	"github.com/AislingHeanue/aisling-codes/wasm-demo/games/common"
	"github.com/AislingHeanue/aisling-codes/wasm-demo/games/rubiks/listener"
	"github.com/AislingHeanue/aisling-codes/wasm-demo/games/rubiks/model"
	"github.com/gowebapi/webapi/graphics/webgl"
)

//go:embed shaders/rubiks.vert
var vertexSource string

//go:embed shaders/rubiks.frag
var fragmentSource string

func New(cco CubeCubeOptions) *CubeRenderer {
	return &CubeRenderer{
		CubeCubeContext: &model.CubeCubeContext{
			AnimationHandler: &model.RubiksAnimationHandler{
				MaxTime: cco.TurnSeconds,
			},
			TotalSideLength: cco.TotalSideLength,
			GapProportion:   cco.GapProportion,
			Origin:          model.Point{0, 0, 0},
			Dimension:       cco.Dimension,
			TurnSeconds:     cco.TurnSeconds,
			Tps:             cco.Tps,
		},
	}
}

type CubeCubeOptions struct {
	TurnSeconds     float32
	Dimension       int
	TotalSideLength float32
	GapProportion   float32
	Tps             float32
}

type CubeRenderer struct {
	Parent *common.ShaderGame
	*model.CubeCubeContext
}

var _ common.GameInfo = &CubeRenderer{}

func (cc *CubeRenderer) PreSetup(c *common.GameContext) {
	sideLength := cc.TotalSideLength / ((1+cc.GapProportion)*float32(cc.Dimension) - cc.GapProportion)
	sideLengthWithGap := sideLength + cc.GapProportion*sideLength

	cc.AnimationHandler.FlushAll()
	cc.AnimationHandler.RubiksCube = model.NewRubiksCube(cc.Dimension, cc.Origin, sideLength, cc.TotalSideLength, sideLengthWithGap)
}

func (cc *CubeRenderer) PostSetup(c *common.GameContext) {
}

func (cc *CubeRenderer) InitListeners(c *common.GameContext) {
	common.RegisterListeners(c, cc.CubeCubeContext, model.CubeController{Context: cc.CubeCubeContext}, listener.CubeActionHandler{})
	listener.RegisterButtons(c, cc.CubeCubeContext)
}

func (cc *CubeRenderer) GetFragmentSource() string {
	return fragmentSource
}

func (cc *CubeRenderer) GetVertexSource() string {
	return vertexSource
}

func (cc *CubeRenderer) GetTps() float32 {
	return cc.Tps
}

func (cc *CubeRenderer) SetParent(parent *common.ShaderGame) {
	cc.Parent = parent
}

func (cc *CubeRenderer) Tick(c *common.GameContext) bool {
	return cc.AnimationHandler.Tick(c.IntervalT)
}

func (cc *CubeRenderer) GetDrawShape(c *common.GameContext) common.DrawShape {
	return model.GetBuffers(cc.AnimationHandler, cc.Origin)
}

func (cc *CubeRenderer) CanRunBetweenFrames() bool {
	return false
}

func (cc *CubeRenderer) AttachAttributes(c *common.GameContext, program *webgl.Program, writeBuffer, readBuffer *webgl.Buffer, samplerTexture *webgl.Texture) {
	gl := c.GL

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
}
