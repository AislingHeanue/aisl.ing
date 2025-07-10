package common

import (
	"github.com/gowebapi/webapi/dom/domcore"
)

type ListenerKind int

const (
	CLICK ListenerKind = iota
	MOUSE_MOVE
	MOUSE_UP
	TOUCH
	TOUCH_MOVE
	TOUCH_UP
	KEYBOARD
	RESIZE
	WHEEL
)

type ActionHandler[Context any, Controller any] interface {
	Click(c *GameContext, context *Context, controller Controller, e *domcore.Event)
	Drag(c *GameContext, context *Context, controller Controller, e *domcore.Event)
	MouseUp(c *GameContext, context *Context, controller Controller, e *domcore.Event)
	Touch(c *GameContext, context *Context, controller Controller, e *domcore.Event)
	DragTouch(c *GameContext, context *Context, controller Controller, e *domcore.Event)
	TouchUp(c *GameContext, context *Context, controller Controller, e *domcore.Event)
	Keyboard(c *GameContext, context *Context, controller Controller, e *domcore.Event)
	Resize(c *GameContext, context *Context, controller Controller, e *domcore.Event)
	Wheel(c *GameContext, context *Context, controller Controller, e *domcore.Event)
}

type Listener[Context any, Controller any] struct {
	c          *GameContext
	context    *Context
	controller Controller
	kind       ListenerKind
	impl       ActionHandler[Context, Controller]
}

func (l *Listener[Context, Controller]) HandleEvent(e *domcore.Event) {
	switch l.kind {
	case CLICK:
		// fmt.Println("click")
		l.impl.Click(l.c, l.context, l.controller, e)
	case MOUSE_MOVE:
		// fmt.Println("move")
		l.impl.Drag(l.c, l.context, l.controller, e)
	case MOUSE_UP:
		// fmt.Println("up")
		l.impl.MouseUp(l.c, l.context, l.controller, e)
	case TOUCH:
		// fmt.Println("touch")
		l.impl.Touch(l.c, l.context, l.controller, e)
	case TOUCH_MOVE:
		// fmt.Println("touch move")
		l.impl.DragTouch(l.c, l.context, l.controller, e)
	case TOUCH_UP:
		// fmt.Println("touch up")
		l.impl.TouchUp(l.c, l.context, l.controller, e)
	case RESIZE:
		// fmt.Println("resize")
		l.impl.Resize(l.c, l.context, l.controller, e)
	case KEYBOARD:
		// fmt.Println("keyboard")
		l.impl.Keyboard(l.c, l.context, l.controller, e)
	case WHEEL:
		// fmt.Println("wheel")
		l.impl.Wheel(l.c, l.context, l.controller, e)
	}
}

func RegisterListeners[Context, Controller any](c *GameContext, context *Context, controller Controller, d ActionHandler[Context, Controller]) {
	AddListenerToCanvas(c, "mousedown", domcore.NewEventListener(&Listener[Context, Controller]{c, context, controller, CLICK, d}))
	AddListenerToCanvas(c, "mousemove", domcore.NewEventListener(&Listener[Context, Controller]{c, context, controller, MOUSE_MOVE, d}))
	AddListenerToCanvas(c, "mouseup", domcore.NewEventListener(&Listener[Context, Controller]{c, context, controller, MOUSE_UP, d}))
	// disable mouseleave for the time being. Dragging while outside the bounds of the canvas is nicer because it doesn't interrupt the user.
	// AddListenerToCanvas(c, "mouseleave", domcore.NewEventListener(&Listener[Context, Controller]{c, context, controller, MOUSE_UP, d}))

	AddListenerToCanvas(c, "touchstart", domcore.NewEventListener(&Listener[Context, Controller]{c, context, controller, TOUCH, d}))
	AddListenerToCanvas(c, "touchmove", domcore.NewEventListener(&Listener[Context, Controller]{c, context, controller, TOUCH_MOVE, d}))
	AddListenerToCanvas(c, "touchend", domcore.NewEventListener(&Listener[Context, Controller]{c, context, controller, TOUCH_UP, d}))
	AddListenerToCanvas(c, "touchcancel", domcore.NewEventListener(&Listener[Context, Controller]{c, context, controller, TOUCH_UP, d}))

	AddListenerToCanvas(c, "wheel", domcore.NewEventListener(&Listener[Context, Controller]{c, context, controller, WHEEL, d}))

	c.Document.AddEventListener("keydown", domcore.NewEventListener(&Listener[Context, Controller]{c, context, controller, KEYBOARD, d}), nil)
	c.Window.AddEventListener("resize", domcore.NewEventListener(&Listener[Context, Controller]{c, context, controller, RESIZE, d}), nil)
}

func AddListenerToCanvas(c *GameContext, event string, listener domcore.EventListener) {
	c.Window.JSValue().Call("canvasEventListener", event, domcore.NewEventListener(listener).JSValue())
}
