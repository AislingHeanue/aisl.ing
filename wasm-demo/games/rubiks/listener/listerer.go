package listener

import (
	"strings"

	"github.com/AislingHeanue/aisling-codes/wasm-demo/canvas"
	"github.com/AislingHeanue/aisling-codes/wasm-demo/games/rubiks/model"
	"github.com/AislingHeanue/aisling-codes/wasm-demo/util"
	"github.com/gowebapi/webapi/dom/domcore"
)

type CubeActionHandler struct{}

var _ util.ActionHandler[model.CubeCubeContext, model.CubeController] = CubeActionHandler{}

func (l CubeActionHandler) Click(c *util.GameContext, context *model.CubeCubeContext, controller model.CubeController, e *domcore.Event) {
	context.AnchorAngleX = context.AngleX
	context.AnchorAngleY = context.AngleY
}

func (l CubeActionHandler) Drag(c *util.GameContext, context *model.CubeCubeContext, controller model.CubeController, e *domcore.Event) {
	if c.MouseDown {
		mouseX, mouseY := canvas.GetRelativeMousePosition(e)
		context.AngleX = (context.AnchorAngleX + 5*(c.AnchorY-mouseY)/c.ResolutionScale/c.Width)
		context.AngleY = (context.AnchorAngleY + 5*(c.AnchorX-mouseX)/c.ResolutionScale/c.Height)
	}
}

func (l CubeActionHandler) DragTouch(c *util.GameContext, context *model.CubeCubeContext, controller model.CubeController, e *domcore.Event) {
	if c.MouseDown {
		mouseX, mouseY := canvas.GetRelativeTouchPosition(c, e)
		context.AngleX = (context.AnchorAngleX + 5*(c.AnchorY-mouseY)/c.ResolutionScale/c.Width)
		context.AngleY = (context.AnchorAngleY + 5*(c.AnchorX-mouseX)/c.ResolutionScale/c.Height)
	}
}

func (l CubeActionHandler) MouseUp(c *util.GameContext, context *model.CubeCubeContext, controller model.CubeController, e *domcore.Event) {
}

func (l CubeActionHandler) Resize(c *util.GameContext, context *model.CubeCubeContext, controller model.CubeController, e *domcore.Event) {
}

func (l CubeActionHandler) Touch(c *util.GameContext, context *model.CubeCubeContext, controller model.CubeController, e *domcore.Event) {
	context.AnchorAngleX = context.AngleX
	context.AnchorAngleY = context.AngleY
}

func (l CubeActionHandler) TouchUp(c *util.GameContext, context *model.CubeCubeContext, controller model.CubeController, e *domcore.Event) {
}

func (l CubeActionHandler) Keyboard(c *util.GameContext, context *model.CubeCubeContext, controller model.CubeController, e *domcore.Event) {
	shiftPressed := e.JSValue().Get("shiftKey").Bool()
	prime := ""
	if shiftPressed {
		prime = "'"
	}
	face := strings.ToLower(e.JSValue().Get("key").String())

	controller.QueueEvent(model.Turn(face + prime))
}
