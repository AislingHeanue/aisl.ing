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
	gc.SetFillColor(color.White)
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

func Cube(gc *draw2dimg.GraphicContext, c *model.GameContext) bool {
	// secs := c.T.Seconds()
	side := c.Width / 2

	// fill is white
	gc.SetFillColor(color.White)
	gc.SetStrokeColor(c.Colour)

	// fill canvas with white
	gc.Clear()

	// z is the direction going into the screen.
	// 0, 0 is top-left corner
	cubeOrigin := model.Vector{c.Width / 2, c.Width / 2, 0}
	cube := model.NewCube(cubeOrigin, side)
	for _, point := range cube.Points {
		c.Projector.DoFirstRotation(point, cubeOrigin)
	}

	cube.Rotate(*cube.GetCentre(), c.AngleY, model.Y)
	cube.Rotate(*cube.GetCentre(), c.AngleX, model.X)

	lines := cube.Lines
	gc.BeginPath()
	for _, line := range lines {
		gc.MoveTo(c.Projector.GetCoords(*line.Origin, c.Height, c.Width, *cube.GetCentre()))
		gc.LineTo(c.Projector.GetCoords(*line.End, c.Height, c.Width, *cube.GetCentre()))
		gc.FillStroke()
	}

	gc.Close()

	return true
}
