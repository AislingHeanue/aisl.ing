package controller

import (
	"fmt"
	"strconv"
	"syscall/js"

	"github.com/AislingHeanue/aisling-codes/wasm-demo/model"
	"github.com/gowebapi/webapi/dom/domcore"
)

func Log(v js.Value) {
	js.Global().Get("console").Call("log", v)
}

func RegisterListeners(c *model.GameContext) {
	c.Window.AddEventListener("resize", domcore.NewEventListener(&Listener{c, RESIZE}), nil)

	res := c.Document.GetElementById("dimension")
	handleDimension(c, res.JSValue())
	res.AddEventListener("input", domcore.NewEventListener(&Listener{c, DIMENSION_CHANGED}), nil)

	c.CvsElement.AddEventListener("mousedown", domcore.NewEventListener(&Listener{c, CLICK}), nil)
	c.CvsElement.AddEventListener("mousemove", domcore.NewEventListener(&Listener{c, MOUSE_MOVE}), nil)
	c.CvsElement.AddEventListener("mouseup", domcore.NewEventListener(&Listener{c, MOUSE_UP}), nil)
	c.CvsElement.AddEventListener("mouseleave", domcore.NewEventListener(&Listener{c, MOUSE_UP}), nil)
}

type ListenerKind int

const (
	CLICK ListenerKind = iota
	MOUSE_MOVE
	MOUSE_UP
	RESIZE
	DIMENSION_CHANGED
)

type Listener struct {
	c    *model.GameContext
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
	case RESIZE:
		InitCanvas(l.c)
	case DIMENSION_CHANGED:
		handleDimension(l.c, e.SrcElement().JSValue())
	}
}

func click(c *model.GameContext, e *domcore.Event) {
	c.AnchorX, c.AnchorY = getRelativeMousePosition(c, e.Value_JS)
	c.AnchorAngleX = c.AngleX
	c.AnchorAngleY = c.AngleY
	c.MouseDown = true
}

func dragCanvas(c *model.GameContext, e *domcore.Event) {
	if c.MouseDown {
		mouseX, mouseY := getRelativeMousePosition(c, e.JSValue())
		c.AngleX = (c.AnchorAngleX + 5*(c.AnchorY-mouseY)/c.ResolutionScale)
		c.AngleY = (c.AnchorAngleY + 5*(c.AnchorX-mouseX)/c.ResolutionScale)
	}
}

func mouseUp(c *model.GameContext) {
	fmt.Println("up")
	c.MouseDown = false
}

func handleDimension(c *model.GameContext, value js.Value) {
	i, _ := strconv.Atoi(value.Get("value").String())
	c.Dimension = i
	c.Animator.Init(c)
}

func getRelativeMousePosition(c *model.GameContext, click js.Value) (float32, float32) {
	relativeX := float32(click.Get("offsetX").Float()) / c.Width
	relativeY := float32(click.Get("offsetY").Float()) / c.Height
	return float32(relativeX), float32(relativeY)
}
