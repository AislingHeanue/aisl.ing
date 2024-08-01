package main

import (
	"github.com/AislingHeanue/aisling-codes/wasm-demo/controller"
	"github.com/AislingHeanue/aisling-codes/wasm-demo/games/rubiks"
)

var done chan struct{}

func main() {
	c := controller.GameContext{}

	c.Animator = rubiks.New(rubiks.CubeCubeOptions{
		TurnFrames: 12,
		Dimension:  3,
	})
	c.ResolutionScale = 1

	controller.InitCanvas(&c)
	controller.RegisterListeners(&c)

	<-done

}
