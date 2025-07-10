package mandelbrot

import (
	"github.com/AislingHeanue/aisling-codes/wasm-demo/common"
)

type MandelbrotActionHandler struct {
	common.DefaultActionHandler[MandelbrotContext, MandelbrotController]
}

var _ common.ActionHandler[MandelbrotContext, MandelbrotController] = MandelbrotActionHandler{}

type MandelbrotContext struct {
	AnchorCentreX float64
	AnchorCentreY float64
	CentreX       float64
	CentreY       float64
	Zoom          float64
}
type MandelbrotController struct{}

// func (l MandelbrotActionHandler) Click(c *common.GameContext, context *MandelbrotContext, controller MandelbrotController, e *domcore.Event) {
// 	context.AnchorAngleX = context.AngleX
// 	context.AnchorAngleY = context.AngleY
// }
//
// func (l CubeActionHandler) Drag(c *common.GameContext, context *MandelbrotContext, controller MandelbrotController, e *domcore.Event) {
// 	if c.MouseDown {
// 		mouseX, mouseY := canvas.GetRelativeMousePosition(e)
// 		context.AngleX = (context.AnchorAngleX + 5*(c.AnchorY-mouseY)/c.ResolutionScale/c.Width)
// 		context.AngleY = (context.AnchorAngleY + 5*(c.AnchorX-mouseX)/c.ResolutionScale/c.Height)
// 	}
// }
//
// func (l CubeActionHandler) DragTouch(c *common.GameContext, context *MandelbrotContext, controller MandelbrotController, e *domcore.Event) {
// 	if c.MouseDown {
// 		mouseX, mouseY := canvas.GetRelativeTouchPosition(c, e)
// 		context.AngleX = (context.AnchorAngleX + 5*(c.AnchorY-mouseY)/c.ResolutionScale/c.Width)
// 		context.AngleY = (context.AnchorAngleY + 5*(c.AnchorX-mouseX)/c.ResolutionScale/c.Height)
// 	}
// }
//
// func (l CubeActionHandler) Touch(c *common.GameContext, context *MandelbrotContext, controller MandelbrotController, e *domcore.Event) {
// 	context.AnchorAngleX = context.AngleX
// 	context.AnchorAngleY = context.AngleY
// }
//
// func (l CubeActionHandler) Keyboard(c *common.GameContext, context *MandelbrotContext, controller MandelbrotController, e *domcore.Event) {
// 	shiftPressed := e.JSValue().Get("shiftKey").Bool()
// 	prime := ""
// 	if shiftPressed {
// 		prime = "'"
// 	}
// 	face := strings.ToLower(e.JSValue().Get("key").String())
//
// 	controller.QueueEvent(model.Turn(face + prime))
// }
