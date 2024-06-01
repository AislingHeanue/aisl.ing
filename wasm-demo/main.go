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
var height float64 = 200
var width float64 = 200
var t float64 = 0

var renderDelay = 1000 * time.Millisecond

func main() {
	doc := js.Global().Get("document")
	cvsBody := doc.Call("getElementById", "wasm-canvas")
	cvs, _ = canvas.NewCanvas2d(false)
	cvs.Set(cvsBody, int(width), int(height))

	FrameRate := time.Second / renderDelay
	println("Hello Browser FPS:", FrameRate)

	cvs.Start(1, Render)

	<-done
}

func doEveryFrame(f func(time.Time)) {
	for x := range time.Tick(renderDelay) {
		f(x)
	}
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

// func updateHeightWidth(_ time.Time) {
// 	newHeight := int(js.Global().Get("innerHeight").Float())
// 	newWidth := int(js.Global().Get("innerWidth").Float())
// 	if newHeight != int(height) || newWidth != int(width) {
// 		width = float64(newWidth)
// 		height = float64(newHeight)
// 		cvs.Create(newWidth, newHeight)
// 	}

// }
