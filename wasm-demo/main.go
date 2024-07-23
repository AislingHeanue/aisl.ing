package main

import (
	"github.com/AislingHeanue/aisling-codes/wasm-demo/graphics"
)

var done chan struct{}

func main() {
	c := graphics.GameContext{}

	c.Animator = &graphics.CubeRenderer{}
	c.ResolutionScale = 1
	c.TurnFrames = 12

	graphics.InitCanvas(&c)
	graphics.StartAnimation(&c)
	graphics.RegisterListeners(&c)

	<-done

}
