package controller

import (
	"fmt"
	"math"
	"syscall/js"

	"github.com/AislingHeanue/aisling-codes/wasm-demo/canvas"
	"github.com/AislingHeanue/aisling-codes/wasm-demo/games/life/model"
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
	KEYBOARD
)

func InitListeners(c *canvas.GameContext, lc *model.LifeContext) {
	c.CvsElement.AddEventListener("mousedown", domcore.NewEventListener(&LifeListener{c, lc, CLICK}), nil)
	c.CvsElement.AddEventListener("mousemove", domcore.NewEventListener(&LifeListener{c, lc, MOUSE_MOVE}), nil)
	c.CvsElement.AddEventListener("mouseup", domcore.NewEventListener(&LifeListener{c, lc, MOUSE_UP}), nil)
	c.CvsElement.AddEventListener("mouseleave", domcore.NewEventListener(&LifeListener{c, lc, MOUSE_UP}), nil)

	c.CvsElement.AddEventListener("touchstart", domcore.NewEventListener(&LifeListener{c, lc, TOUCH}), nil)
	c.CvsElement.AddEventListener("touchmove", domcore.NewEventListener(&LifeListener{c, lc, TOUCH_MOVE}), nil)
	c.CvsElement.AddEventListener("touchend", domcore.NewEventListener(&LifeListener{c, lc, TOUCH_UP}), nil)
	c.CvsElement.AddEventListener("touchcancel", domcore.NewEventListener(&LifeListener{c, lc, TOUCH_UP}), nil)

	c.Document.AddEventListener("keydown", domcore.NewEventListener(&LifeListener{c, lc, KEYBOARD}), nil)
	// registerButtons(c, lc)
}

type LifeListener struct {
	c    *canvas.GameContext
	lc   *model.LifeContext
	kind ListenerKind
}

func (l *LifeListener) HandleEvent(e *domcore.Event) {
	switch l.kind {
	case CLICK:
		click(l.c, l.lc, e)
	case MOUSE_MOVE:
		dragCanvas(l.c, l.lc, e)
	case MOUSE_UP:
		mouseUp(l.lc)
	case TOUCH:
		touch(l.c, l.lc, e)
	case TOUCH_MOVE:
		dragCanvasTouch(l.c, l.lc, e)
	case TOUCH_UP:
		touchUp(l.lc)
	case KEYBOARD:
		handleKeyboard(l.c, l.lc, e)
	}
}

func click(c *canvas.GameContext, lc *model.LifeContext, e *domcore.Event) {
	lc.AnchorX, lc.AnchorY = getRelativeMousePosition(e)
	lc.AnchorDX = lc.DX
	lc.AnchorDY = lc.DY
	lc.MouseDown = true
}

func touch(c *canvas.GameContext, lc *model.LifeContext, e *domcore.Event) {
	fmt.Println("touch")
	if e.JSValue().Get("touches").Length() == 2 {
		// start zooming
		lc.Zooming = true
		// don't drag and zoom at the same time because it's probably complicated
		lc.MouseDown = false
		lc.AnchorPinchDistance = getDistanceBetweenTouches(c, e)
		lc.AnchorZoom = lc.Zoom
		e.PreventDefault()
	} else {
		lc.AnchorX, lc.AnchorY = getRelativeTouchPosition(c, e.JSValue().Get("touches").Get("0"))
		lc.AnchorDX = lc.DX
		lc.AnchorDY = lc.DY
		lc.MouseDown = true
		lc.Zooming = false
		e.PreventDefault()
	}
	// lockScroll(c)
}

func dragCanvas(c *canvas.GameContext, lc *model.LifeContext, e *domcore.Event) {
	if lc.MouseDown {
		e.PreventDefault()
		mouseX, mouseY := getRelativeMousePosition(e)
		lc.DX = (lc.AnchorDX - (lc.AnchorX-mouseX)/c.ResolutionScale)
		lc.DY = (lc.AnchorDY - (lc.AnchorY-mouseY)/c.ResolutionScale)
	}
}

func dragCanvasTouch(c *canvas.GameContext, lc *model.LifeContext, e *domcore.Event) {
	if lc.MouseDown {
		mouseX, mouseY := getRelativeTouchPosition(c, e.JSValue().Get("touches").Get("0"))
		lc.DX = (lc.AnchorDX - 3*(lc.AnchorX-mouseX)/c.ResolutionScale)
		lc.DY = (lc.AnchorDY - 3*(lc.AnchorY-mouseY)/c.ResolutionScale)
	}
	if lc.Zooming {
		distance := getDistanceBetweenTouches(c, e)
		lc.Zoom = lc.AnchorZoom * (distance / lc.AnchorPinchDistance) //(lc.AnchorPinchDistance - distance) / distance
		// zoom stuff
	}
}

func mouseUp(lc *model.LifeContext) {
	// fmt.Printf("%v %v\n", lc.DX, lc.DY)
	lc.MouseDown = false
}

func touchUp(lc *model.LifeContext) {
	lc.MouseDown = false
	lc.Zooming = false
}

func getRelativeMousePosition(e *domcore.Event) (float32, float32) {
	relativeX := float32(e.JSValue().Get("offsetX").Float())
	relativeY := float32(e.JSValue().Get("offsetY").Float())
	return float32(relativeX), float32(relativeY)
}

func getRelativeTouchPosition(c *canvas.GameContext, touch js.Value) (float32, float32) {
	rect := c.CvsElement.JSValue().Call("getBoundingClientRect")
	offsetX := touch.Get("clientX").Float() - rect.Get("left").Float()
	offsetY := touch.Get("clientY").Float() - rect.Get("top").Float()
	return float32(offsetX), float32(offsetY)
}

func getDistanceBetweenTouches(c *canvas.GameContext, e *domcore.Event) float32 {
	touches := e.JSValue().Get("touches")
	x1, y1 := getRelativeTouchPosition(c, touches.Get("0"))
	x2, y2 := getRelativeTouchPosition(c, touches.Get("1"))
	return float32(math.Hypot(float64(x1-x2), float64(y1-y2)))
}

func handleKeyboard(c *canvas.GameContext, lc *model.LifeContext, e *domcore.Event) {
	switch e.JSValue().Get("key").String() {
	case " ":
		e.PreventDefault()
		lc.Paused = !lc.Paused
	case "-":
		lc.Zoom *= 9. / 10.
		lc.DX = (lc.DX-c.Width/2)*9./10. + c.Width/2
		lc.DY = (lc.DY-c.Height/2)*9./10. + c.Height/2
		fmt.Println(lc.Zoom)
	case "=":
		lc.Zoom *= 10. / 9.
		lc.DX = (lc.DX-c.Width/2)*10./9. + c.Width/2
		lc.DY = (lc.DY-c.Height/2)*10./9. + c.Height/2
		fmt.Println(lc.Zoom)
	}
	// controller := CubeController{lc}
	// controller.QueueEvent(Turn(face + prime))
}
