package life

import (
	"github.com/AislingHeanue/aisling-codes/wasm-demo/common"
	"github.com/gowebapi/webapi/dom/domcore"
)

type LifeContext struct {
	Paused         bool
	SimulationSize int
	Tps            float32
	NeedsRestart   bool // will probably remove this one
	Loop           bool
	T              int
	TrailLength    int
	ColourPeriod   int

	OpenFileName string
}

type LifeController interface {
	Reset(c *common.GameContext)
	Random(c *common.GameContext)
	ResizeBuffers(c *common.GameContext)
	OpenFile(c *common.GameContext, path string)
	OpenRandomFile(c *common.GameContext)
}

type LifeActionHandler struct {
	common.DefaultActionHandler[LifeContext, LifeController]
}

var _ common.ActionHandler[LifeContext, LifeController] = LifeActionHandler{}

func (l LifeActionHandler) Keyboard(c *common.GameContext, context *LifeContext, controller LifeController, e *domcore.Event) {
	switch e.JSValue().Get("key").String() {
	// pause simulation
	case " ":
		context.Paused = !context.Paused
		e.PreventDefault()
	// set simulation size to 210
	case "b":
		c.PixelsWidth = 500
		c.PixelsHeight = 500
		controller.ResizeBuffers(c)
		controller.Reset(c)
	case "":
		controller.Reset(c)
	case "r":
		controller.Random(c)
	case "l":
		context.Loop = !context.Loop
	case "p":
		controller.OpenRandomFile(c)
	case "k":
		controller.OpenFile(c, context.OpenFileName)
	}
}
