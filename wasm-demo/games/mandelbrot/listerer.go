package mandelbrot

import (
	"math"

	"github.com/AislingHeanue/aisling-codes/wasm-demo/canvas"
	"github.com/AislingHeanue/aisling-codes/wasm-demo/common"
	"github.com/gowebapi/webapi/dom/domcore"
)

type MandelbrotActionHandler struct {
	common.DefaultActionHandler[MandelbrotContext, MandelbrotController]
}

var _ common.ActionHandler[MandelbrotContext, MandelbrotController] = MandelbrotActionHandler{}

type MandelbrotContext struct {
	AnchorCentreX float64
	AnchorCentreY float64
	AnchorZoom    float64
	CentreX       float64
	CentreY       float64
	Zoom          float64
}
type MandelbrotController struct{}

func (a MandelbrotActionHandler) Wheel(c *common.GameContext, context *MandelbrotContext, controller MandelbrotController, e *domcore.Event) {
	e.PreventDefault()
	deltaY := e.JSValue().Get("deltaY").Float()
	if deltaY < 0 {
		mouseX, mouseY := canvas.GetRelativeMousePosition(e)
		canvasCentreX, canvasCentreY := c.Width/2/c.ResolutionScale, c.Height/2/c.ResolutionScale

		context.CentreX = context.CentreX - 2.5*c.Window.DevicePixelRatio()*float64((canvasCentreX-mouseX)*c.ResolutionScale/c.Width)/context.Zoom
		context.CentreY = context.CentreY - 2.5*c.Window.DevicePixelRatio()*float64((canvasCentreY-mouseY)*c.ResolutionScale/c.Height)/context.Zoom

		setZoom(context, math.Pow(1.1, -deltaY/180)*context.Zoom)

		context.CentreX = context.CentreX + 2.5*c.Window.DevicePixelRatio()*float64((canvasCentreX-mouseX)*c.ResolutionScale/c.Width)/context.Zoom
		context.CentreY = context.CentreY + 2.5*c.Window.DevicePixelRatio()*float64((canvasCentreY-mouseY)*c.ResolutionScale/c.Height)/context.Zoom
		c.Log(context.Zoom)
	} else {
		setZoom(context, math.Pow(1.1, -deltaY/180)*context.Zoom)
	}
}

func setZoom(context *MandelbrotContext, zoom float64) {
	oldZoom := context.Zoom

	// cap max zoom in
	if zoom > oldZoom && zoom > 100000 {
		return
	}
	// cap max zoom out
	if zoom < oldZoom && zoom < 0.2 {
		return
	}
	context.Zoom = zoom
}

func (a MandelbrotActionHandler) Click(c *common.GameContext, context *MandelbrotContext, controller MandelbrotController, e *domcore.Event) {
	context.AnchorCentreX = context.CentreX
	context.AnchorCentreY = context.CentreY
}

func (a MandelbrotActionHandler) Drag(c *common.GameContext, context *MandelbrotContext, controller MandelbrotController, e *domcore.Event) {
	if c.MouseDown {
		mouseX, mouseY := canvas.GetRelativeMousePosition(e)
		context.CentreX = context.AnchorCentreX + 2.5*c.Window.DevicePixelRatio()*float64((c.AnchorX-mouseX)*c.ResolutionScale/c.Width)/context.Zoom
		context.CentreY = context.AnchorCentreY + 2.5*c.Window.DevicePixelRatio()*float64((c.AnchorY-mouseY)*c.ResolutionScale/c.Height)/context.Zoom
	}
}

func (a MandelbrotActionHandler) DragTouch(c *common.GameContext, context *MandelbrotContext, controller MandelbrotController, e *domcore.Event) {
	if c.MouseDown {
		mouseX, mouseY := canvas.GetRelativeTouchPosition(c, e)
		context.CentreX = context.AnchorCentreX + c.Window.DevicePixelRatio()*float64((c.AnchorX-mouseX)*c.ResolutionScale/c.Width)/context.Zoom
		context.CentreY = context.AnchorCentreY + c.Window.DevicePixelRatio()*float64((c.AnchorY-mouseY)*c.ResolutionScale/c.Height)/context.Zoom
	}

	if c.Zooming {
		distance := canvas.GetDistanceBetweenTouches(e)
		setZoom(context, context.AnchorZoom*float64(distance/c.AnchorPinchDistance))
	}
}

func (a MandelbrotActionHandler) Touch(c *common.GameContext, context *MandelbrotContext, controller MandelbrotController, e *domcore.Event) {
	context.AnchorCentreX = context.CentreX
	context.AnchorCentreY = context.CentreY

	if e.JSValue().Get("touches").Length() == 2 {
		context.AnchorZoom = context.Zoom
		e.PreventDefault()
	} else {
		c.AnchorX, c.AnchorY = canvas.GetRelativeTouchPosition(c, e)
		c.AnchorDX = c.DX
		c.AnchorDY = c.DY
		c.MouseDown = true
		e.PreventDefault()
	}
}

func (a MandelbrotActionHandler) Keyboard(c *common.GameContext, context *MandelbrotContext, controller MandelbrotController, e *domcore.Event) {
	switch e.JSValue().Get("key").String() {
	// reset zoom
	case "0":
		setZoom(context, 1)
	// zoom out
	case "-":
		setZoom(context, 9/10.*context.Zoom)
	// zoom in (+)
	case "=":
		setZoom(context, 10/9.*context.Zoom)
	}
}
