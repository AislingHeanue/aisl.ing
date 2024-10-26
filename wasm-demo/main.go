package main

import (
	"fmt"
	"os"

	"github.com/AislingHeanue/aisling-codes/wasm-demo/canvas"
	"github.com/AislingHeanue/aisling-codes/wasm-demo/games/life"
	"github.com/AislingHeanue/aisling-codes/wasm-demo/games/rubiks"
	"github.com/gowebapi/webapi"
	"github.com/gowebapi/webapi/html/htmlcommon"
)

var done chan struct{}

func main() {
	contexts := map[string]canvas.GameContext{
		"rubiks": {
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
		},
		"life": {
			Window:          webapi.GetWindow(),
			Document:        webapi.GetWindow().Document(),
			CvsElement:      webapi.GetWindow().Document().GetElementById("wasm-canvas"),
			ResolutionScale: 1,
			Animator:        life.New(),
			CellHeight:      200,
			CellWidth:       200,
			SecondaryCanvas: webapi.GetWindow().Document().CreateElement("canvas", &webapi.Union{}),
			ZoomCanvas:      webapi.GetWindow().Document().CreateElement("canvas", &webapi.Union{}),
		},
	}
	fmt.Printf("%#v", os.Args)
	c := contexts[os.Args[0]]

	canvas.InitCanvas(&c)
	c.Animator.Init(&c)
	canvas.RegisterListeners(&c)
	c.Animator.InitListeners(&c)
	c.Window.RequestAnimationFrame(htmlcommon.FrameRequestCallbackToJS(wrapAnimator(&c)))

	<-done

}

func wrapAnimator(c *canvas.GameContext) func(float64) {
	return func(time float64) {
		c.IntervalT = (float32(time) / 1000) - c.T // milliseconds to seconds
		c.T = float32(time) / 1000
		c.Animator.Render(c)
		c.Window.RequestAnimationFrame(htmlcommon.FrameRequestCallbackToJS(wrapAnimator(c)))
	}

}
