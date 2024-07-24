package controller

import (
	"strconv"
	"strings"
	"syscall/js"

	"github.com/gowebapi/webapi/dom/domcore"
)

func RegisterListeners(c *GameContext) {
	c.Window.AddEventListener("resize", domcore.NewEventListener(&CanvasListener{c, RESIZE}), nil)

	res := c.Document.GetElementById("dimension")
	handleDimension(c, res.JSValue())
	res.AddEventListener("input", domcore.NewEventListener(&CanvasListener{c, DIMENSION_CHANGED}), nil)

	c.CvsElement.AddEventListener("mousedown", domcore.NewEventListener(&CanvasListener{c, CLICK}), nil)
	c.CvsElement.AddEventListener("mousemove", domcore.NewEventListener(&CanvasListener{c, MOUSE_MOVE}), nil)
	c.CvsElement.AddEventListener("mouseup", domcore.NewEventListener(&CanvasListener{c, MOUSE_UP}), nil)
	c.CvsElement.AddEventListener("mouseleave", domcore.NewEventListener(&CanvasListener{c, MOUSE_UP}), nil)

	c.CvsElement.AddEventListener("touchstart", domcore.NewEventListener(&CanvasListener{c, TOUCH}), nil)
	c.CvsElement.AddEventListener("touchmove", domcore.NewEventListener(&CanvasListener{c, TOUCH_MOVE}), nil)
	c.CvsElement.AddEventListener("touchend", domcore.NewEventListener(&CanvasListener{c, TOUCH_UP}), nil)
	c.CvsElement.AddEventListener("touchcancel", domcore.NewEventListener(&CanvasListener{c, TOUCH_UP}), nil)

	c.Document.AddEventListener("keydown", domcore.NewEventListener(&CCListener{Animator: c.Animator}), nil)
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

type CanvasListener struct {
	c    *GameContext
	kind ListenerKind
}

func (l *CanvasListener) HandleEvent(e *domcore.Event) {
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
	e.PreventDefault()
	// lockScroll(c)
}

func dragCanvas(c *GameContext, e *domcore.Event) {
	if c.MouseDown {
		e.PreventDefault()
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
	// unlockScroll(c)
}

func handleDimension(c *GameContext, value js.Value) {
	i, _ := strconv.Atoi(value.Get("value").String())
	c.CubeDimension = i
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
	return 1.5 * float32(offsetX) / c.Width, 1.5 * float32(offsetY) / c.Height
}

type CCListener struct {
	Animator
}

func (l *CCListener) HandleEvent(e *domcore.Event) {
	shiftPressed := e.JSValue().Get("shiftKey").Bool()
	face := strings.ToLower(e.JSValue().Get("key").String())
	l.QueueEvent(face, shiftPressed)
}
