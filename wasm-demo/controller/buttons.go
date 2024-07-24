package controller

import (
	"github.com/gowebapi/webapi/dom/domcore"
)

var buttonIDs = []string{
	"u", "r", "f", "x", "shuffle",
	"d", "l", "b", "y", "reset",
	"e", "m", "s", "z", "reset-camera",
}

func registerButtons(c *GameContext) {
	for _, id := range buttonIDs {
		c.Document.GetElementById(id).AddEventListener("click", domcore.NewEventListener(&ButtonListener{id, c}), nil)
	}
}

type ButtonListener struct {
	id string
	c  *GameContext
}

func (l *ButtonListener) HandleEvent(e *domcore.Event) {
	shiftPressed := e.JSValue().Get("shiftKey").Bool()
	face := l.id
	switch face {
	case "shuffle":
		l.c.Animator.Shuffle()
	case "reset":
		l.c.Animator.Init(l.c)
	case "reset-camera":
		l.c.AngleX = 0
		l.c.AngleY = 0
	default:
		l.c.Animator.QueueEvent(face, shiftPressed)
	}
}
