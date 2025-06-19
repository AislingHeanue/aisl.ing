package canvas

import (
	"github.com/gowebapi/webapi/dom/domcore"
)

func RegisterListeners(c *GameContext) {
	c.Window.AddEventListener("resize", domcore.NewEventListener(&CanvasListener{c, RESIZE}), nil)
}

type ListenerKind int

const (
	RESIZE ListenerKind = iota
)

type CanvasListener struct {
	c    *GameContext
	kind ListenerKind
}

func (l *CanvasListener) HandleEvent(e *domcore.Event) {
	switch l.kind {
	case RESIZE:
		println("confused")
		InitCanvas(l.c)
	}
}
