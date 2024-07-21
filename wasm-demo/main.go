package main

import (
	"github.com/AislingHeanue/aisling-codes/wasm-demo/controller"
	"github.com/AislingHeanue/aisling-codes/wasm-demo/model"
	"github.com/AislingHeanue/aisling-codes/wasm-demo/rubiks"
)

var done chan struct{}

func main() {
	c := model.GameContext{}

	c.Animator = &rubiks.CubeCube{}
	c.ResolutionScale = 1

	controller.InitCanvas(&c)
	controller.StartAnimation(&c)
	controller.RegisterListeners(&c)

	<-done

}
