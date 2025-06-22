package controller

import (
	"fmt"
	"math"

	"github.com/gowebapi/webapi/dom/domcore"

	"github.com/AislingHeanue/aisling-codes/wasm-demo/canvas"
	"github.com/AislingHeanue/aisling-codes/wasm-demo/games/life/model"
)

type ListenerKind int

const (
	CLICK ListenerKind = iota
	MOUSE_MOVE
	MOUSE_UP
	DIMENSION_CHANGED
	TOUCH
	TOUCH_MOVE
	TOUCH_UP
	KEYBOARD
)

func InitListeners(c *canvas.GameContext, lc *model.LifeContext, controller LifeController) {
	canvas.AddListener(c, "mousedown", domcore.NewEventListener(&LifeListener{c, lc, CLICK, controller}))
	canvas.AddListener(c, "mousemove", domcore.NewEventListener(&LifeListener{c, lc, MOUSE_MOVE, controller}))
	canvas.AddListener(c, "mouseup", domcore.NewEventListener(&LifeListener{c, lc, MOUSE_UP, controller}))
	canvas.AddListener(c, "mouseleave", domcore.NewEventListener(&LifeListener{c, lc, MOUSE_UP, controller}))

	canvas.AddListener(c, "touchstart", domcore.NewEventListener(&LifeListener{c, lc, TOUCH, controller}))
	canvas.AddListener(c, "touchmove", domcore.NewEventListener(&LifeListener{c, lc, TOUCH_MOVE, controller}))
	canvas.AddListener(c, "touchend", domcore.NewEventListener(&LifeListener{c, lc, TOUCH_UP, controller}))
	canvas.AddListener(c, "touchcancel", domcore.NewEventListener(&LifeListener{c, lc, TOUCH_UP, controller}))

	c.Document.AddEventListener("keydown", domcore.NewEventListener(&LifeListener{c, lc, KEYBOARD, controller}), nil)
}

type LifeListener struct {
	c          *canvas.GameContext
	lc         *model.LifeContext
	kind       ListenerKind
	controller LifeController
}

type LifeController interface {
	Reset(c *canvas.GameContext)
	Random(c *canvas.GameContext)
	ResizeBuffers(c *canvas.GameContext)
	OpenFile(c *canvas.GameContext, path string)
	OpenRandomFile(c *canvas.GameContext)
}

func (l *LifeListener) HandleEvent(e *domcore.Event) {
	switch l.kind {
	case CLICK:
		click(l.lc, e)
	case MOUSE_MOVE:
		dragCanvas(l.lc, e)
	case MOUSE_UP:
		mouseUp(l.lc)
	case TOUCH:
		touch(l.c, l.lc, e)
	case TOUCH_MOVE:
		dragCanvasTouch(l.c, l.lc, e)
	case TOUCH_UP:
		touchUp(l.lc)
	case KEYBOARD:
		handleKeyboard(l.c, l.lc, l.controller, e)
	}
}

func click(lc *model.LifeContext, e *domcore.Event) {
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
		lc.AnchorPinchDistance = getDistanceBetweenTouches(e)
		lc.AnchorZoom = lc.Zoom
		e.PreventDefault()
	} else {
		lc.AnchorX, lc.AnchorY = getRelativeTouchPosition(c, e)
		lc.AnchorDX = lc.DX
		lc.AnchorDY = lc.DY
		lc.MouseDown = true
		lc.Zooming = false
		e.PreventDefault()
	}
	// lockScroll(c)
}

func dragCanvas(lc *model.LifeContext, e *domcore.Event) {
	if lc.MouseDown {
		e.PreventDefault()
		mouseX, mouseY := getRelativeMousePosition(e)
		lc.DX = (lc.AnchorDX - (lc.AnchorX - mouseX))
		lc.DY = (lc.AnchorDY - (lc.AnchorY - mouseY))
	}
}

func dragCanvasTouch(c *canvas.GameContext, lc *model.LifeContext, e *domcore.Event) {
	if lc.MouseDown {
		mouseX, mouseY := getRelativeTouchPosition(c, e)
		lc.DX = (lc.AnchorDX - float32(c.Window.DevicePixelRatio())*(lc.AnchorX-mouseX))
		lc.DY = (lc.AnchorDY - float32(c.Window.DevicePixelRatio())*(lc.AnchorY-mouseY))
	}
	if lc.Zooming {
		distance := getDistanceBetweenTouches(e)
		setZoom(lc, lc.AnchorZoom*(distance/lc.AnchorPinchDistance))
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

func getRelativeTouchPosition(c *canvas.GameContext, e *domcore.Event) (float32, float32) {
	rect := c.RenderingCanvas.JSValue().Call("getBoundingClientRect")
	touch := e.JSValue().Get("touches").Get("0")
	offsetX := touch.Get("clientX").Float() - rect.Get("left").Float()
	offsetY := touch.Get("clientY").Float() - rect.Get("top").Float()
	return float32(offsetX), float32(offsetY)
}

func getDistanceBetweenTouches(e *domcore.Event) float32 {
	touches := e.JSValue().Get("touches")
	x1 := touches.Get("0").Get("clientX").Float()
	y1 := touches.Get("0").Get("clientY").Float()
	x2 := touches.Get("1").Get("clientX").Float()
	y2 := touches.Get("1").Get("clientY").Float()
	return float32(math.Hypot(float64(x1-x2), float64(y1-y2)))
}

func handleKeyboard(c *canvas.GameContext, lc *model.LifeContext, controller LifeController, e *domcore.Event) {
	switch e.JSValue().Get("key").String() {
	// pause simulation
	case " ":
		lc.Paused = !lc.Paused
		e.PreventDefault()
	// zoom out
	case "-":
		setZoom(lc, 9/10.*lc.Zoom)
	// zoom in (+)
	case "=":
		setZoom(lc, 10/9.*lc.Zoom)
	// reset zoom
	case "0":
		setZoom(lc, 1)
	// recentre
	case "o":
		lc.DX = 0
		lc.DY = 0
	// set simulation size to 210
	case "b":
		lc.CellWidth = 500
		lc.CellHeight = 500
		controller.ResizeBuffers(c)
		controller.Reset(c)
	case "c":
		controller.Reset(c)
	case "r":
		controller.Random(c)
	case "l":
		lc.Loop = !lc.Loop
	case "p":
		controller.OpenRandomFile(c)
	case "k":
		controller.OpenFile(c, lc.OpenFileName)
	}
}

func setZoom(lc *model.LifeContext, zoom float32) {
	oldZoom := lc.Zoom
	// cap max zoom in
	if zoom > oldZoom && zoom > 20 {
		return
	}
	// cap max zoom out
	if zoom < oldZoom && zoom < 0.2 {
		return
	}
	lc.Zoom = zoom
	// scale DX and DY so that the 'anchor' of the zoom is always in the centre of the screen.
	lc.DX *= zoom / oldZoom
	lc.DY *= zoom / oldZoom
}
