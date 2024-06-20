package main

import (
	"github.com/AislingHeanue/aisling-codes/wasm-demo/animation"
	"github.com/AislingHeanue/aisling-codes/wasm-demo/controller"
	"github.com/AislingHeanue/aisling-codes/wasm-demo/model"
)

var done chan struct{}

func main() {
	c := model.GameContext{}

	c.Animator = &animation.CubeCube{}
	c.Dimension = 3
	c.ResolutionScale = 1

	controller.InitCanvas(&c)
	controller.RegisterListeners(&c)

	<-done

}
