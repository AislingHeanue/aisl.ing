package main

import (
	"os"

	"github.com/gowebapi/webapi"
	"github.com/gowebapi/webapi/html/htmlcommon"

	"github.com/AislingHeanue/aisling-codes/wasm-demo/canvas"
	"github.com/AislingHeanue/aisling-codes/wasm-demo/games/common"
	"github.com/AislingHeanue/aisling-codes/wasm-demo/games/life"

	"github.com/AislingHeanue/aisling-codes/wasm-demo/games/rubiks"
)

var done chan struct{}

func main() {
	contexts := map[string]common.GameContext{
		"rubiks": {
			Square:   true,
			Window:   webapi.GetWindow(),
			Document: webapi.GetWindow().Document(),
			// RenderingCanvas: webapi.GetWindow().Document().GetElementById("wasm-canvas"),
			ResolutionScale: 1,
			Animator: rubiks.New(
				rubiks.CubeCubeOptions{
					TurnSeconds:     0.18,
					Dimension:       3,
					TotalSideLength: 0.5,
					GapProportion:   0.07,
				},
			),
		},
		"life": {
			Square:   true,
			Window:   webapi.GetWindow(),
			Document: webapi.GetWindow().Document(),
			// DisplayCanvas:   webapi.GetWindow().Document().GetElementById("wasm-canvas"),
			ResolutionScale: 0.5,
			Animator: &common.ShaderGame{
				GameInfo: life.New(
					life.LifeOptions{
						Tps:          30,
						Loop:         true,
						TrailLength:  25,
						ColourPeriod: 50,
					}),
			},
			PixelsHeight:    200,
			PixelsWidth:     100,
			Zoom:            1,
			RenderingCanvas: webapi.GetWindow().Document().CreateElement("canvas", &webapi.Union{}),
			ZoomCanvas:      webapi.GetWindow().Document().CreateElement("canvas", &webapi.Union{}),
			ZoomEnabled:     true,
			PanningEnabled:  true,
		},
	}
	// parser.ReadFile("oversized/41dots.lif")
	// fmt.Printf("%#v", os.Args)
	c := contexts[os.Args[0]]

	canvas.InitCanvas(&c)
	common.RegisterListeners(&c, nil, nil, canvas.CanvasActionHandler{})
	c.Animator.Init(&c)
	c.Animator.InitListeners(&c)
	c.Window.RequestAnimationFrame(htmlcommon.FrameRequestCallbackToJS(wrapAnimator(&c)))

	<-done
}

func wrapAnimator(c *common.GameContext) func(float64) {
	return func(time float64) {
		c.IntervalT = (float32(time) / 1000) - c.T // milliseconds to seconds
		c.T = float32(time) / 1000
		c.Animator.Render(c)
		c.Window.RequestAnimationFrame(htmlcommon.FrameRequestCallbackToJS(wrapAnimator(c)))
	}

}
