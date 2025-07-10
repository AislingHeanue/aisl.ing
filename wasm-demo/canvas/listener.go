package canvas

import (
	"fmt"
	"math"

	"github.com/AislingHeanue/aisling-codes/wasm-demo/common"
	"github.com/gowebapi/webapi/dom/domcore"
)

type CanvasActionHandler struct{}

var _ common.ActionHandler[any, any] = CanvasActionHandler{}

func GetRelativeMousePosition(e *domcore.Event) (float32, float32) {
	relativeX := float32(e.JSValue().Get("offsetX").Float())
	relativeY := float32(e.JSValue().Get("offsetY").Float())
	return float32(relativeX), float32(relativeY)
}

func GetRelativeTouchPosition(c *common.GameContext, e *domcore.Event) (float32, float32) {
	rect := c.RenderingCanvas.JSValue().Call("getBoundingClientRect")
	touch := e.JSValue().Get("touches").Get("0")
	offsetX := touch.Get("clientX").Float() - rect.Get("left").Float()
	offsetY := touch.Get("clientY").Float() - rect.Get("top").Float()
	return float32(offsetX), float32(offsetY)
}

func GetDistanceBetweenTouches(e *domcore.Event) float32 {
	touches := e.JSValue().Get("touches")
	x1 := touches.Get("0").Get("clientX").Float()
	y1 := touches.Get("0").Get("clientY").Float()
	x2 := touches.Get("1").Get("clientX").Float()
	y2 := touches.Get("1").Get("clientY").Float()
	return float32(math.Hypot(float64(x1-x2), float64(y1-y2)))
}

func setZoom(c *common.GameContext, zoom float32) {
	if !c.ZoomEnabled {
		return
	}

	oldZoom := c.Zoom
	// cap max zoom in
	if zoom > oldZoom && zoom > 20 {
		return
	}
	// cap max zoom out
	if zoom < oldZoom && zoom < 0.2 {
		return
	}
	c.Zoom = zoom
	// scale DX and DY so that the 'anchor' of the zoom is always in the centre of the screen.
	c.DX *= zoom / oldZoom
	c.DY *= zoom / oldZoom
}

func (a CanvasActionHandler) Click(c *common.GameContext, context *any, controller any, e *domcore.Event) {
	c.AnchorX, c.AnchorY = GetRelativeMousePosition(e)
	c.AnchorDX = c.DX
	c.AnchorDY = c.DY
	c.MouseDown = true
}

func (a CanvasActionHandler) Drag(c *common.GameContext, context *any, controller any, e *domcore.Event) {
	if c.MouseDown && c.PanningEnabled {
		e.PreventDefault()
		mouseX, mouseY := GetRelativeMousePosition(e)
		c.DX = (c.AnchorDX - (c.AnchorX - mouseX))
		c.DY = (c.AnchorDY - (c.AnchorY - mouseY))
	}
}

func (a CanvasActionHandler) DragTouch(c *common.GameContext, context *any, controller any, e *domcore.Event) {
	if c.MouseDown && c.PanningEnabled {
		mouseX, mouseY := GetRelativeTouchPosition(c, e)
		c.DX = (c.AnchorDX - float32(c.Window.DevicePixelRatio())*(c.AnchorX-mouseX))
		c.DY = (c.AnchorDY - float32(c.Window.DevicePixelRatio())*(c.AnchorY-mouseY))
	}
	if c.Zooming && c.ZoomEnabled {
		distance := GetDistanceBetweenTouches(e)
		setZoom(c, c.AnchorZoom*(distance/c.AnchorPinchDistance))
	}
}

func (a CanvasActionHandler) MouseUp(c *common.GameContext, context *any, controller any, e *domcore.Event) {
	c.MouseDown = false
}

func (a CanvasActionHandler) Touch(c *common.GameContext, context *any, controller any, e *domcore.Event) {
	fmt.Println("touch")
	if e.JSValue().Get("touches").Length() == 2 {
		// start zooming
		c.Zooming = true
		// don't drag and zoom at the same time because it's probably complicated
		c.MouseDown = false
		c.AnchorPinchDistance = GetDistanceBetweenTouches(e)
		c.AnchorZoom = c.Zoom
		e.PreventDefault()
	} else {
		c.AnchorX, c.AnchorY = GetRelativeTouchPosition(c, e)
		c.AnchorDX = c.DX
		c.AnchorDY = c.DY
		c.MouseDown = true
		e.PreventDefault()
	}
}

func (a CanvasActionHandler) TouchUp(c *common.GameContext, context *any, controller any, e *domcore.Event) {
	c.MouseDown = false
	c.Zooming = false
}

func (a CanvasActionHandler) Keyboard(c *common.GameContext, context *any, controller any, e *domcore.Event) {
	switch e.JSValue().Get("key").String() {
	// pause simulation
	case "0":
		setZoom(c, 1)
	// zoom out
	case "-":
		setZoom(c, 9/10.*c.Zoom)
	// zoom in (+)
	case "=":
		setZoom(c, 10/9.*c.Zoom)
	// reset zoom
	// recentre
	case "o":
		c.DX = 0
		c.DY = 0
	}
}

func (a CanvasActionHandler) Resize(c *common.GameContext, context *any, controller any, e *domcore.Event) {
	InitCanvas(c)
	c.Animator.RefreshBuffers(c)
}
