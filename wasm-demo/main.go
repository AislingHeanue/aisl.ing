package main

import (
	"syscall/js"
	"time"

	"github.com/AislingHeanue/aisling-codes/wasm-demo/animation"
	"github.com/AislingHeanue/aisling-codes/wasm-demo/controller"
	"github.com/AislingHeanue/aisling-codes/wasm-demo/model"
)

var done chan struct{}

func main() {
	c := model.GameContext{}

	c.Fps = 30
	c.RenderDelay = time.Second / time.Duration(c.Fps)
	c.Document = js.Global().Get("document")
	c.Window = js.Global().Get("window")
	c.Animator = &animation.CubeCube{}
	c.Projector = animation.PerspectiveProjector{}
	c.Colour.A = 255
	c.Dimension = 2

	println("Hello Browser FPS:", c.Fps)

	controller.InitCanvas(&c, js.Null(), nil)
	controller.RegisterListeners(&c)

	<-done

}
