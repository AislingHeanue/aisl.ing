package controller

import (
	"strings"

	"github.com/AislingHeanue/aisling-codes/wasm-demo/games/rubiks/model"
	"github.com/gowebapi/webapi/dom/domcore"
)

type CCListener struct {
	ccc *model.CubeCubeContext
}

func (l *CCListener) HandleEvent(e *domcore.Event) {
	controller := CubeController{l.ccc}

	shiftPressed := e.JSValue().Get("shiftKey").Bool()
	prime := ""
	if shiftPressed {
		prime = "'"
	}
	face := strings.ToLower(e.JSValue().Get("key").String())
	controller.QueueEvent(Turn(face + prime))
}
