package controller

import (
	"github.com/AislingHeanue/aisling-codes/wasm-demo/canvas"
	"github.com/AislingHeanue/aisling-codes/wasm-demo/games/rubiks/model"
	"github.com/gowebapi/webapi/dom/domcore"
)

var buttonIDs = []string{
	"u", "r", "f", "x", "shuffle",
	"d", "l", "b", "y", "reset",
	"e", "m", "s", "z", "resetcam",
	"kilt", "s-flip", "cubecube", "checker", "snake",
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
	case "kilt":
		control := CubeController{l.ccc}
		control.QueueEvent("u'", "r2", "l2", "f2", "b2", "u'", "r", "l", "f", "b'", "u", "f2", "d2", "r2", "l2", "f2", "u2", "f2", "u'", "f2")
	case "superflip":
		control := CubeController{l.ccc}
		control.QueueEvent("u", "r2", "f", "b", "r", "b2", "r", "u2", "l", "b2", "r", "u'", "d'", "r2", "f", "r'", "l", "b2", "u2", "f2")
	case "cubecube":
		control := CubeController{l.ccc}
		control.QueueEvent("u'", "l'", "u'", "f'", "r2", "b'", "r", "f", "u", "b2", "u", "b'", "l", "u'", "f", "u", "r", "f'")
	case "checkerboard":
		control := CubeController{l.ccc}
		control.QueueEvent("m2", "e2", "s2")
	case "snake":
		control := CubeController{l.ccc}
		control.QueueEvent("l", "u", "b'", "u'", "r", "l'", "b", "r'", "f", "b'", "d", "r", "d'", "f'")
	default:
		control := CubeController{l.ccc}
		prime := ""
		if shiftPressed {
			prime = "'"
		}
		control.QueueEvent(Turn(face + prime))
	}
}
