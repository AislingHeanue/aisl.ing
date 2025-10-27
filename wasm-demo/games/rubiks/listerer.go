package rubiks

import (
	"strings"

	"github.com/AislingHeanue/aisling-codes/wasm-demo/canvas"
	"github.com/AislingHeanue/aisling-codes/wasm-demo/common"
	"github.com/AislingHeanue/aisling-codes/wasm-demo/games/rubiks/model"
	"github.com/gowebapi/webapi/dom/domcore"
)

type CubeActionHandler struct {
	common.DefaultActionHandler[model.CubeCubeContext, model.CubeController]
}

var _ common.ActionHandler[model.CubeCubeContext, model.CubeController] = CubeActionHandler{}

func (l CubeActionHandler) Click(c *common.GameContext, context *model.CubeCubeContext, controller model.CubeController, e *domcore.Event) {
	context.AnchorAngleX = context.AngleX
	context.AnchorAngleY = context.AngleY
}

func (l CubeActionHandler) Drag(c *common.GameContext, context *model.CubeCubeContext, controller model.CubeController, e *domcore.Event) {
	if c.MouseDown {
		mouseX, mouseY := canvas.GetRelativeMousePosition(e)
		context.AngleX = (context.AnchorAngleX + 5*(c.AnchorY-mouseY)/c.ResolutionScale/c.Width*float32(c.Window.DevicePixelRatio()))
		context.AngleY = (context.AnchorAngleY + 5*(c.AnchorX-mouseX)/c.ResolutionScale/c.Height*float32(c.Window.DevicePixelRatio()))
	}
}

func (l CubeActionHandler) DragTouch(c *common.GameContext, context *model.CubeCubeContext, controller model.CubeController, e *domcore.Event) {
	if c.MouseDown {
		mouseX, mouseY := canvas.GetRelativeTouchPosition(c, e)
		context.AngleX = (context.AnchorAngleX + 5*(c.AnchorY-mouseY)/c.ResolutionScale/c.Width*float32(c.Window.DevicePixelRatio()))
		context.AngleY = (context.AnchorAngleY + 5*(c.AnchorX-mouseX)/c.ResolutionScale/c.Height*float32(c.Window.DevicePixelRatio()))
	}
}

func (l CubeActionHandler) Touch(c *common.GameContext, context *model.CubeCubeContext, controller model.CubeController, e *domcore.Event) {
	context.AnchorAngleX = context.AngleX
	context.AnchorAngleY = context.AngleY
}

func (l CubeActionHandler) Keyboard(c *common.GameContext, context *model.CubeCubeContext, controller model.CubeController, e *domcore.Event) {
	shiftPressed := e.JSValue().Get("shiftKey").Bool()
	prime := ""
	if shiftPressed {
		prime = "'"
	}
	face := strings.ToLower(e.JSValue().Get("key").String())

	controller.QueueEvent(model.Turn(face + prime))
}
