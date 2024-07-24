package main

import (
	"github.com/AislingHeanue/aisling-codes/wasm-demo/controller"
	"github.com/AislingHeanue/aisling-codes/wasm-demo/graphics"
)

var done chan struct{}

func main() {
	c := controller.GameContext{}

	c.Animator = &graphics.CubeRenderer{}
	c.ResolutionScale = 1
	c.TurnFrames = 12

	controller.InitCanvas(&c)
	controller.StartAnimation(&c)
	controller.RegisterListeners(&c)

	<-done

}
