package main

import (
	"image/color"
	"math"
	"syscall/js"
	"time"

	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/draw2dkit"
	"github.com/markfarnan/go-canvas/canvas"
)

var done chan struct{}

var cvs *canvas.Canvas2d
var height float64
var width float64
var t float64 = 0

var renderDelay = 30 * time.Millisecond

func main() {
	cvs, _ = canvas.NewCanvas2d(false)
	FrameRate := time.Second / renderDelay
	println("Hello Browser FPS:", FrameRate)
	cvs.Create(int(js.Global().Get("innerWidth").Float()*0.9), int(js.Global().Get("innerHeight").Float()*0.9))

	height = float64(cvs.Height())
	width = float64(cvs.Width())

	cvs.Start(60, Render)

	<-done
}

func Render(gc *draw2dimg.GraphicContext) bool {
	gc.SetFillColor(color.RGBA{255, 255, 255, 255})
	// fill canvas with fill colour (white)
	gc.Clear()

	// stroke and fill are magenta
	gc.SetFillColor(color.RGBA{255, 0, 255, 255})
	gc.SetStrokeColor(color.RGBA{255, 0, 255, 255})
	gc.BeginPath()
	draw2dkit.Circle(gc, width/2*(1+math.Cos(t/10)), height/2*(1+math.Sin(t/10)), min(width/3, height/3))
	gc.FillStroke()
	gc.Close()
	t++

	return true
}
