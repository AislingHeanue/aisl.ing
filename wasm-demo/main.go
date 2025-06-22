package main

import (
	"os"

	"github.com/gowebapi/webapi"
	"github.com/gowebapi/webapi/html/htmlcommon"

	"github.com/AislingHeanue/aisling-codes/wasm-demo/canvas"
	"github.com/AislingHeanue/aisling-codes/wasm-demo/games/life"
	"github.com/AislingHeanue/aisling-codes/wasm-demo/games/life/model"

	"github.com/AislingHeanue/aisling-codes/wasm-demo/games/rubiks"
)

var done chan struct{}

func main() {
	contexts := map[string]canvas.GameContext{
		"rubiks": {
			Square:   true,
			Window:   webapi.GetWindow(),
			Document: webapi.GetWindow().Document(),
			// RenderingCanvas: webapi.GetWindow().Document().GetElementById("wasm-canvas"),
			ResolutionScale: 1,
			Animator: rubiks.New(
				rubiks.CubeCubeOptions{
					TurnFrames: 12,
					Dimension:  3,
				},
			),
		},
		"life": {
			Square:   true,
			Window:   webapi.GetWindow(),
			Document: webapi.GetWindow().Document(),
			// DisplayCanvas:   webapi.GetWindow().Document().GetElementById("wasm-canvas"),
			ResolutionScale: 0.5,
			Animator: life.New(&model.LifeContext{
				CellHeight: 200,
				CellWidth:  200,
				Zoom:       2,
				Tps:        5,
				Loop:       true,
			}),
			RenderingCanvas: webapi.GetWindow().Document().CreateElement("canvas", &webapi.Union{}),
			ZoomCanvas:      webapi.GetWindow().Document().CreateElement("canvas", &webapi.Union{}),
		},
	}
	// parser.ReadFile("oversized/41dots.lif")
	// fmt.Printf("%#v", os.Args)
	c := contexts[os.Args[0]]

	canvas.InitCanvas(&c)
	canvas.RegisterListeners(&c)
	c.Animator.Init(&c)
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
