package controller

import (
	"github.com/gowebapi/webapi/dom/domcore"

	"github.com/AislingHeanue/aisling-codes/wasm-demo/util"
)

type LifeContext struct {
	Paused         bool
	SimulationSize int
	Tps            float32
	NeedsRestart   bool // will probably remove this one
	Loop           bool

	OpenFileName string
}

type LifeController interface {
	Reset(c *util.GameContext)
	Random(c *util.GameContext)
	ResizeBuffers(c *util.GameContext)
	OpenFile(c *util.GameContext, path string)
	OpenRandomFile(c *util.GameContext)
}

type LifeActionHandler struct{}

var _ util.ActionHandler[LifeContext, LifeController] = LifeActionHandler{}

func (l LifeActionHandler) Click(c *util.GameContext, context *LifeContext, controller LifeController, e *domcore.Event) {
}

func (l LifeActionHandler) Drag(c *util.GameContext, context *LifeContext, controller LifeController, e *domcore.Event) {
}

func (l LifeActionHandler) DragTouch(c *util.GameContext, context *LifeContext, controller LifeController, e *domcore.Event) {
}

func (l LifeActionHandler) MouseUp(c *util.GameContext, context *LifeContext, controller LifeController, e *domcore.Event) {
}

func (l LifeActionHandler) Resize(c *util.GameContext, context *LifeContext, controller LifeController, e *domcore.Event) {
}

func (l LifeActionHandler) Touch(c *util.GameContext, context *LifeContext, controller LifeController, e *domcore.Event) {
}

func (l LifeActionHandler) TouchUp(c *util.GameContext, context *LifeContext, controller LifeController, e *domcore.Event) {
}

func (l LifeActionHandler) Keyboard(c *util.GameContext, context *LifeContext, controller LifeController, e *domcore.Event) {
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
