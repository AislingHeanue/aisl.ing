package main

import (
	"strconv"
	"strings"
	"syscall/js"
	"time"

	"github.com/AislingHeanue/aisling-codes/wasm-demo/animation"
	"github.com/AislingHeanue/aisling-codes/wasm-demo/model"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/markfarnan/go-canvas/canvas"
)

var done chan struct{}

func main() {
	c := model.GameContext{}
	c.Fps = 30
	c.RenderDelay = time.Second / time.Duration(c.Fps)
	c.Document = js.Global().Get("document")
	c.Window = js.Global().Get("window")

	c.Colour.A = 255
	for _, k := range []string{"red", "green", "blue"} {
		res := c.Document.Call("getElementById", k)
		handleUint8(k, &c)(res, nil)
		res.Call("addEventListener", "input", js.FuncOf(handleUint8(k, &c)))
	}

	println("Hello Browser FPS:", c.Fps)

	setCanvasSize(&c)
	c.Window.Call("addEventListener", "resize", wrapListener(setCanvasSize, &c))

	<-done

}

func handleUint8(name string, c *model.GameContext) func(js.Value, []js.Value) any {
	return func(this js.Value, args []js.Value) any {
		i, _ := strconv.Atoi(this.Get("value").String())
		switch name {
		case "red":
			c.Colour.R = uint8(i)
		case "green":
			c.Colour.G = uint8(i)
		case "blue":
			c.Colour.B = uint8(i)
		}

		return js.Null()
	}
}

func wrapListener(f func(*model.GameContext), c *model.GameContext) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		f(c)

		return js.Null()
	})
}

func setCanvasSize(c *model.GameContext) {
	cvsBody := c.Document.Call("getElementById", "wasm-canvas")
	computedStyle := js.Global().Get("window").Call("getComputedStyle", cvsBody)

	widthWithBorder, _ := strconv.ParseFloat(strings.TrimSuffix(computedStyle.Get("width").String(), "px"), 64)
	heightWithBorder, _ := strconv.ParseFloat(strings.TrimSuffix(computedStyle.Get("height").String(), "px"), 64)
	borderUp, _ := strconv.ParseFloat(strings.TrimSuffix(computedStyle.Get("border-top-width").String(), "px"), 64)
	borderDown, _ := strconv.ParseFloat(strings.TrimSuffix(computedStyle.Get("border-bottom-width").String(), "px"), 64)
	borderLeft, _ := strconv.ParseFloat(strings.TrimSuffix(computedStyle.Get("border-left-width").String(), "px"), 64)
	borderRight, _ := strconv.ParseFloat(strings.TrimSuffix(computedStyle.Get("border-right-width").String(), "px"), 64)

	c.Height = heightWithBorder - borderUp - borderDown
	c.Width = widthWithBorder - borderLeft - borderRight

	cvsBody.Set("width", int(c.Width))
	cvsBody.Set("height", int(c.Height))
	println(int(c.Height), int(c.Width))
	if c.Cvs != nil {
		go c.Cvs.Stop()
	}
	c.Cvs, _ = canvas.NewCanvas2d(false)
	c.Cvs.Set(cvsBody, int(c.Width), int(c.Height))
	c.Cvs.Start(c.Fps, render(animation.CirclingCircle, c))
}

func render(f model.AnimationFunc, c *model.GameContext) func(*draw2dimg.GraphicContext) bool {
	return func(gc *draw2dimg.GraphicContext) bool {
		val := f(gc, *c)
		c.T += c.RenderDelay

		return val
	}
}
