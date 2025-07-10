package controller

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

type LifeActionHandler struct{}

var _ common.ActionHandler[LifeContext, LifeController] = LifeActionHandler{}

func (l LifeActionHandler) Click(c *common.GameContext, context *LifeContext, controller LifeController, e *domcore.Event) {
}

func (l LifeActionHandler) Drag(c *common.GameContext, context *LifeContext, controller LifeController, e *domcore.Event) {
}

func (l LifeActionHandler) DragTouch(c *common.GameContext, context *LifeContext, controller LifeController, e *domcore.Event) {
}

func (l LifeActionHandler) MouseUp(c *common.GameContext, context *LifeContext, controller LifeController, e *domcore.Event) {
}

func (l LifeActionHandler) Resize(c *common.GameContext, context *LifeContext, controller LifeController, e *domcore.Event) {
}

func (l LifeActionHandler) Touch(c *common.GameContext, context *LifeContext, controller LifeController, e *domcore.Event) {
}

func (l LifeActionHandler) TouchUp(c *common.GameContext, context *LifeContext, controller LifeController, e *domcore.Event) {
}

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

