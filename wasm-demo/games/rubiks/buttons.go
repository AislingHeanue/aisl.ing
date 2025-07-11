package rubiks

import (
	"github.com/AislingHeanue/aisling-codes/wasm-demo/common"
	"github.com/AislingHeanue/aisling-codes/wasm-demo/games/rubiks/model"
	"github.com/gowebapi/webapi/dom/domcore"
)

var buttonIDs = []string{
	"u", "r", "f", "x", "shuffle",
	"d", "l", "b", "y", "reset",
	"e", "m", "s", "z", "resetcam",
	"kilt", "s-flip", "cubecube", "checker", "snake",
}

func RegisterButtons(c *common.GameContext, ccc *model.CubeCubeContext) {
	for _, id := range buttonIDs {
		c.Document.GetElementById(id).AddEventListener("click", domcore.NewEventListener(&ButtonListener{id, c, ccc}), nil)
	}
}

type ButtonListener struct {
	id  string
	c   *common.GameContext
	ccc *model.CubeCubeContext
}

func (l *ButtonListener) HandleEvent(e *domcore.Event) {
	shiftPressed := e.JSValue().Get("shiftKey").Bool()
	face := l.id
	switch face {
	case "shuffle":
		control := model.CubeController{Context: l.ccc}
		control.Shuffle()
	case "reset":
		l.c.Animator.Init(l.c)
	case "resetcam":
		l.ccc.AngleX = 0
		l.ccc.AngleY = 0
	case "kilt":
		control := model.CubeController{Context: l.ccc}
		control.QueueEvent("u'", "r2", "l2", "f2", "b2", "u'", "r", "l", "f", "b'", "u", "f2", "d2", "r2", "l2", "f2", "u2", "f2", "u'", "f2")
	case "s-flip":
		control := model.CubeController{Context: l.ccc}
		control.QueueEvent("u", "r2", "f", "b", "r", "b2", "r", "u2", "l", "b2", "r", "u'", "d'", "r2", "f", "r'", "l", "b2", "u2", "f2")
	case "cubecube":
		control := model.CubeController{Context: l.ccc}
		control.QueueEvent("u'", "l'", "u'", "f'", "r2", "b'", "r", "f", "u", "b2", "u", "b'", "l", "u'", "f", "u", "r", "f'")
	case "checker":
		control := model.CubeController{Context: l.ccc}
		control.QueueEvent("m2", "e2", "s2")
	case "snake":
		control := model.CubeController{Context: l.ccc}
		control.QueueEvent("l", "u", "b'", "u'", "r", "l'", "b", "r'", "f", "b'", "d", "r", "d'", "f'")
	default:
		control := model.CubeController{Context: l.ccc}
		prime := ""
		if shiftPressed {
			prime = "'"
		}
		control.QueueEvent(model.Turn(face + prime))
	}
}
