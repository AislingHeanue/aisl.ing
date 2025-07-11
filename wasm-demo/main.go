package main

import (
	"os"

	"github.com/gowebapi/webapi/html/htmlcommon"

	"github.com/AislingHeanue/aisling-codes/wasm-demo/canvas"
	"github.com/AislingHeanue/aisling-codes/wasm-demo/common"
	"github.com/AislingHeanue/aisling-codes/wasm-demo/games/life"
	"github.com/AislingHeanue/aisling-codes/wasm-demo/games/mandelbrot"

	"github.com/AislingHeanue/aisling-codes/wasm-demo/games/rubiks"
)

var done chan struct{}

func main() {
	contexts := map[string]common.GameContext{
		"rubiks": {
			ResolutionScale: 1.7,
			SmoothImage:     true,
			Is3D:            true,
			Game: rubiks.New(rubiks.CubeCubeOptions{
				TurnSeconds:     0.18,
				Dimension:       3,
				TotalSideLength: 0.5,
				GapProportion:   0.07,
				Tps:             60,
			}),
		},
		"life": {
			ResolutionScale: 1.4,
			Game: life.New(life.LifeOptions{
				Loop:         true,
				TrailLength:  25,
				ColourPeriod: 50,
				Tps:          30,
			}),
			PixelsWidth:    200,
			PixelsHeight:   200,
			ZoomEnabled:    true,
			PanningEnabled: true,
		},
		"mandelbrot": {
			ResolutionScale: 1.6,
			Game: mandelbrot.New(mandelbrot.MandelbrotOptions{
				CentreX:    -0.75,
				CentreY:    0,
				Zoom:       1,
				Iterations: 800,
				FpsTarget:  24,
			}),
			ZoomEnabled:    false,
			PanningEnabled: false,
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
