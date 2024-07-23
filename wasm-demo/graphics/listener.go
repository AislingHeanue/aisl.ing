package graphics

import (
	"fmt"
	"strconv"
	"strings"
	"syscall/js"

	"github.com/gowebapi/webapi/css/cssom/view"
	"github.com/gowebapi/webapi/dom/domcore"
)

type CCListener struct {
	cc *CubeRenderer
	c  *GameContext
}

func (l *CCListener) HandleEvent(e *domcore.Event) {
	shiftPressed := e.JSValue().Get("shiftKey").Bool()
	face := strings.ToLower(e.JSValue().Get("key").String())
	l.cc.animationHandler.AddEvent(face, shiftPressed)
}

func Log(v js.Value) {
	js.Global().Get("console").Call("log", v)
}

func RegisterListeners(c *GameContext) {
	c.Window.AddEventListener("resize", domcore.NewEventListener(&Listener{c, RESIZE}), nil)

	res := c.Document.GetElementById("dimension")
	handleDimension(c, res.JSValue())
	res.AddEventListener("input", domcore.NewEventListener(&Listener{c, DIMENSION_CHANGED}), nil)

	c.CvsElement.AddEventListener("mousedown", domcore.NewEventListener(&Listener{c, CLICK}), nil)
	c.CvsElement.AddEventListener("mousemove", domcore.NewEventListener(&Listener{c, MOUSE_MOVE}), nil)
	c.CvsElement.AddEventListener("mouseup", domcore.NewEventListener(&Listener{c, MOUSE_UP}), nil)
	c.CvsElement.AddEventListener("mouseleave", domcore.NewEventListener(&Listener{c, MOUSE_UP}), nil)

	c.CvsElement.AddEventListener("touchstart", domcore.NewEventListener(&Listener{c, TOUCH}), nil)
	c.CvsElement.AddEventListener("touchmove", domcore.NewEventListener(&Listener{c, TOUCH_MOVE}), nil)
	c.CvsElement.AddEventListener("touchend", domcore.NewEventListener(&Listener{c, TOUCH_UP}), nil)
	c.CvsElement.AddEventListener("touchcancel", domcore.NewEventListener(&Listener{c, TOUCH_UP}), nil)

}

type ListenerKind int

const (
	CLICK ListenerKind = iota
	MOUSE_MOVE
	MOUSE_UP
	RESIZE
	DIMENSION_CHANGED
	TOUCH
	TOUCH_MOVE
	TOUCH_UP
)

type Listener struct {
	c    *GameContext
	kind ListenerKind
}

func (l *Listener) HandleEvent(e *domcore.Event) {
	switch l.kind {
	case CLICK:
		click(l.c, e)
	case MOUSE_MOVE:
		dragCanvas(l.c, e)
	case MOUSE_UP:
		mouseUp(l.c)
	case TOUCH:
		touch(l.c, e)
	case TOUCH_MOVE:
		dragCanvasTouch(l.c, e)
	case TOUCH_UP:
		touchUp(l.c)
	case RESIZE:
		InitCanvas(l.c)
	case DIMENSION_CHANGED:
		handleDimension(l.c, e.SrcElement().JSValue())
	}
}

func click(c *GameContext, e *domcore.Event) {
	c.AnchorX, c.AnchorY = getRelativeMousePosition(c, e.Value_JS)
	c.AnchorAngleX = c.AngleX
	c.AnchorAngleY = c.AngleY
	c.MouseDown = true
}

func touch(c *GameContext, e *domcore.Event) {
	c.AnchorX, c.AnchorY = getRelativeTouchPosition(c, e.Value_JS)
	c.AnchorAngleX = c.AngleX
	c.AnchorAngleY = c.AngleY
	c.MouseDown = true
	lockScroll(c)
}

func dragCanvas(c *GameContext, e *domcore.Event) {
	if c.MouseDown {
		mouseX, mouseY := getRelativeMousePosition(c, e.JSValue())
		c.AngleX = (c.AnchorAngleX + 5*(c.AnchorY-mouseY)/c.ResolutionScale)
		c.AngleY = (c.AnchorAngleY + 5*(c.AnchorX-mouseX)/c.ResolutionScale)
	}
}

func dragCanvasTouch(c *GameContext, e *domcore.Event) {
	if c.MouseDown {
		mouseX, mouseY := getRelativeTouchPosition(c, e.JSValue())
		c.AngleX = (c.AnchorAngleX + 5*(c.AnchorY-mouseY)/c.ResolutionScale)
		c.AngleY = (c.AnchorAngleY + 5*(c.AnchorX-mouseX)/c.ResolutionScale)
	}
}

func mouseUp(c *GameContext) {
	c.MouseDown = false
}

func touchUp(c *GameContext) {
	c.MouseDown = false
	unlockScroll(c)
}

func handleDimension(c *GameContext, value js.Value) {
	i, _ := strconv.Atoi(value.Get("value").String())
	c.Dimension = i
	c.Animator.Init(c)
}

func getRelativeMousePosition(c *GameContext, click js.Value) (float32, float32) {
	relativeX := float32(click.Get("offsetX").Float()) / c.Width
	relativeY := float32(click.Get("offsetY").Float()) / c.Height
	return float32(relativeX), float32(relativeY)
}

func getRelativeTouchPosition(c *GameContext, touch js.Value) (float32, float32) {
	rect := c.CvsElement.JSValue().Call("getBoundingClientRect")
	touchInfo := touch.Get("touches").Get("0") // only care about the first touch point
	offsetX := touchInfo.Get("clientX").Float() - rect.Get("left").Float()
	offsetY := touchInfo.Get("clientY").Float() - rect.Get("top").Float()
	return float32(offsetX) / c.Width, float32(offsetY) / c.Height
}

func lockScroll(c *GameContext) {
	c.ScrollPosition = c.Document.ActiveElement().ScrollTop() //c.Window.PageYOffset()
	c.Document.Body().Style().SetProperty("overflow", "hidden", nil)
	c.Document.Body().Style().SetProperty("position", "fixed", nil)
	c.Document.Body().Style().SetProperty("top", fmt.Sprintf("-%dpx", int(c.ScrollPosition)), nil)
	c.Document.Body().Style().SetProperty("width", "100%", nil)
}

func unlockScroll(c *GameContext) {
	c.Document.Body().Style().SetProperty("overflow", "", nil)
	c.Document.Body().Style().SetProperty("position", "", nil)
	c.Document.Body().Style().SetProperty("top", "", nil)
	c.Document.Body().Style().SetProperty("width", "100%", nil)
	c.Window.ScrollTo(&view.ScrollToOptions{Left: 0, Top: c.ScrollPosition})
}
