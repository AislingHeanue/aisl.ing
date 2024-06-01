package animation

import (
	"image/color"
	"math"

	"github.com/AislingHeanue/aisling-codes/wasm-demo/model"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/draw2dkit"
)

func CirclingCircle(gc *draw2dimg.GraphicContext, c model.GameContext) bool {
	secs := c.T.Seconds()

	// fill is white
	gc.SetFillColor(color.RGBA{255, 255, 255, 255})
	// fill canvas with fill colour
	gc.Clear()

	// stroke and fill are magenta
	gc.SetFillColor(c.Colour)
	gc.SetStrokeColor(c.Colour)
	// draw a circle
	gc.BeginPath()
	draw2dkit.Circle(gc, c.Width/2*(1+math.Cos(secs)), c.Height/2*(1+math.Sin(secs)), min(c.Width/3, c.Height/3))
	gc.FillStroke()
	gc.Close()

	return true
}
