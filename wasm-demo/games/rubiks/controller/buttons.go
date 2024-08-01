package controller

import (
	"github.com/AislingHeanue/aisling-codes/wasm-demo/canvas"
	"github.com/AislingHeanue/aisling-codes/wasm-demo/games/rubiks/model"
	"github.com/gowebapi/webapi/dom/domcore"
)

var buttonIDs = []string{
	"u", "r", "f", "x", "shuffle",
	"d", "l", "b", "y", "reset",
	"e", "m", "s", "z", "reset-camera",
}

func registerButtons(c *canvas.GameContext, ccc *model.CubeCubeContext) {
	for _, id := range buttonIDs {
		c.Document.GetElementById(id).AddEventListener("click", domcore.NewEventListener(&ButtonListener{id, c, ccc}), nil)
	}
}

type ButtonListener struct {
	id  string
	c   *canvas.GameContext
	ccc *model.CubeCubeContext
}

func (l *ButtonListener) HandleEvent(e *domcore.Event) {
	shiftPressed := e.JSValue().Get("shiftKey").Bool()
	face := l.id
	switch face {
	case "shuffle":
		control := CubeController{l.ccc}
		control.Shuffle()
	case "reset":
		l.c.Animator.Init(l.c)
	case "reset-camera":
		l.ccc.AngleX = 0
		l.ccc.AngleY = 0
	default:
		control := CubeController{l.ccc}
		control.QueueEvent(face, shiftPressed)
	}
}
