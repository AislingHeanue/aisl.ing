package controller

import (
	"strconv"
	"syscall/js"

	"github.com/AislingHeanue/aisling-codes/wasm-demo/controller"
	"github.com/AislingHeanue/aisling-codes/wasm-demo/games/rubiks/model"
	"github.com/gowebapi/webapi/dom/domcore"
)

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

func InitListeners(c *controller.GameContext, ccc *model.CubeCubeContext) {
	c.CvsElement.AddEventListener("mousedown", domcore.NewEventListener(&CanvasListener{c, ccc, CLICK}), nil)
	c.CvsElement.AddEventListener("mousemove", domcore.NewEventListener(&CanvasListener{c, ccc, MOUSE_MOVE}), nil)
	c.CvsElement.AddEventListener("mouseup", domcore.NewEventListener(&CanvasListener{c, ccc, MOUSE_UP}), nil)
	c.CvsElement.AddEventListener("mouseleave", domcore.NewEventListener(&CanvasListener{c, ccc, MOUSE_UP}), nil)

	c.CvsElement.AddEventListener("touchstart", domcore.NewEventListener(&CanvasListener{c, ccc, TOUCH}), nil)
	c.CvsElement.AddEventListener("touchmove", domcore.NewEventListener(&CanvasListener{c, ccc, TOUCH_MOVE}), nil)
	c.CvsElement.AddEventListener("touchend", domcore.NewEventListener(&CanvasListener{c, ccc, TOUCH_UP}), nil)
	c.CvsElement.AddEventListener("touchcancel", domcore.NewEventListener(&CanvasListener{c, ccc, TOUCH_UP}), nil)

	c.Document.AddEventListener("keydown", domcore.NewEventListener(&CCListener{ccc}), nil)
	registerButtons(c, ccc)
}

type CanvasListener struct {
	c    *controller.GameContext
	ccc  *model.CubeCubeContext
	kind ListenerKind
}

func (l *CanvasListener) HandleEvent(e *domcore.Event) {
	switch l.kind {
	case CLICK:
		click(l.c, l.ccc, e)
	case MOUSE_MOVE:
		dragCanvas(l.c, l.ccc, e)
	case MOUSE_UP:
		mouseUp(l.ccc)
	case TOUCH:
		touch(l.c, l.ccc, e)
	case TOUCH_MOVE:
		dragCanvasTouch(l.c, l.ccc, e)
	case TOUCH_UP:
		touchUp(l.ccc)
	case DIMENSION_CHANGED:
		handleDimension(l.c, l.ccc, e.SrcElement().JSValue())
	}
}

func click(c *controller.GameContext, ccc *model.CubeCubeContext, e *domcore.Event) {
	ccc.AnchorX, ccc.AnchorY = getRelativeMousePosition(c, e.Value_JS)
	ccc.AnchorAngleX = ccc.AngleX
	ccc.AnchorAngleY = ccc.AngleY
	ccc.MouseDown = true
}

func touch(c *controller.GameContext, ccc *model.CubeCubeContext, e *domcore.Event) {
	ccc.AnchorX, ccc.AnchorY = getRelativeTouchPosition(c, e.Value_JS)
	ccc.AnchorAngleX = ccc.AngleX
	ccc.AnchorAngleY = ccc.AngleY
	ccc.MouseDown = true
	e.PreventDefault()
	// lockScroll(c)
}

func dragCanvas(c *controller.GameContext, ccc *model.CubeCubeContext, e *domcore.Event) {
	if ccc.MouseDown {
		e.PreventDefault()
		mouseX, mouseY := getRelativeMousePosition(c, e.JSValue())
		ccc.AngleX = (ccc.AnchorAngleX + 5*(ccc.AnchorY-mouseY)/c.ResolutionScale)
		ccc.AngleY = (ccc.AnchorAngleY + 5*(ccc.AnchorX-mouseX)/c.ResolutionScale)
	}
}

func dragCanvasTouch(c *controller.GameContext, ccc *model.CubeCubeContext, e *domcore.Event) {
	if ccc.MouseDown {
		mouseX, mouseY := getRelativeTouchPosition(c, e.JSValue())
		ccc.AngleX = (ccc.AnchorAngleX + 5*(ccc.AnchorY-mouseY)/c.ResolutionScale)
		ccc.AngleY = (ccc.AnchorAngleY + 5*(ccc.AnchorX-mouseX)/c.ResolutionScale)
	}
}

func mouseUp(ccc *model.CubeCubeContext) {
	ccc.MouseDown = false
}

func touchUp(ccc *model.CubeCubeContext) {
	ccc.MouseDown = false
	// unlockScroll(c)
}

func handleDimension(c *controller.GameContext, ccc *model.CubeCubeContext, value js.Value) {
	i, _ := strconv.Atoi(value.Get("value").String())
	ccc.CubeDimension = i
	c.Animator.Init(c)
}

func getRelativeMousePosition(c *controller.GameContext, click js.Value) (float32, float32) {
	relativeX := float32(click.Get("offsetX").Float()) / c.Width
	relativeY := float32(click.Get("offsetY").Float()) / c.Height
	return float32(relativeX), float32(relativeY)
}

func getRelativeTouchPosition(c *controller.GameContext, touch js.Value) (float32, float32) {
	rect := c.CvsElement.JSValue().Call("getBoundingClientRect")
	touchInfo := touch.Get("touches").Get("0") // only care about the first touch point
	offsetX := touchInfo.Get("clientX").Float() - rect.Get("left").Float()
	offsetY := touchInfo.Get("clientY").Float() - rect.Get("top").Float()
	return 1.5 * float32(offsetX) / c.Width, 1.5 * float32(offsetY) / c.Height
}
