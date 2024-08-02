package main

import (
	"github.com/AislingHeanue/aisling-codes/wasm-demo/canvas"
	"github.com/AislingHeanue/aisling-codes/wasm-demo/games/life"
	"github.com/AislingHeanue/aisling-codes/wasm-demo/games/rubiks"
	"github.com/gowebapi/webapi"
	"github.com/gowebapi/webapi/html/htmlcommon"
)

var done chan struct{}

func main() {
	c1 := canvas.GameContext{
		Window:          webapi.GetWindow(),
		Document:        webapi.GetWindow().Document(),
		CvsElement:      webapi.GetWindow().Document().GetElementById("wasm-canvas"),
		ResolutionScale: 1,
		Animator: rubiks.New(
			rubiks.CubeCubeOptions{
				TurnFrames: 12,
				Dimension:  3,
			},
		),
		// Animator:  &life.LifeGame{},
		// Height:    10,
		// Width:     10,
		// FixedSize: true,
	}

	c2 := canvas.GameContext{
		Window:          webapi.GetWindow(),
		Document:        webapi.GetWindow().Document(),
		CvsElement:      webapi.GetWindow().Document().GetElementById("wasm-canvas"),
		ResolutionScale: 1,
		Animator:        &life.LifeGame{},
		Height:          10,
		Width:           10,
		FixedSize:       true,
	}

	c := c1
	_ = c1
	_ = c2

	canvas.InitCanvas(&c)
	c.Animator.Init(&c)
	canvas.RegisterListeners(&c)
	c.Animator.InitListeners(&c)
	c.Window.RequestAnimationFrame(htmlcommon.FrameRequestCallbackToJS(wrapAnimator(&c)))

	<-done

}

func wrapAnimator(c *canvas.GameContext) func(float64) {
	return func(time float64) {
		c.T = float32(time) / 1000 // milliseconds to seconds
		c.Animator.Render(c)
		c.Window.RequestAnimationFrame(htmlcommon.FrameRequestCallbackToJS(wrapAnimator(c)))
	}

}
