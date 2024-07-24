package controller

import (
	"strings"

	"github.com/gowebapi/webapi/dom/domcore"
)

type CCListener struct {
	Animator
}

func (l *CCListener) HandleEvent(e *domcore.Event) {
	shiftPressed := e.JSValue().Get("shiftKey").Bool()
	face := strings.ToLower(e.JSValue().Get("key").String())
	l.QueueEvent(face, shiftPressed)
}
