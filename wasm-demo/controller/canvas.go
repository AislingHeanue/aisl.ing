package controller

import (
	"strconv"
	"strings"
	"syscall/js"

	_ "github.com/AislingHeanue/aisling-codes/wasm-demo/animation"
	"github.com/AislingHeanue/aisling-codes/wasm-demo/model"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/markfarnan/go-canvas/canvas"
)

func InitCanvas(c *model.GameContext, _ js.Value, _ []js.Value) {
	c.CvsElement = c.Document.Call("getElementById", "wasm-canvas")
	computedStyle := js.Global().Get("window").Call("getComputedStyle", c.CvsElement)

	widthWithBorder, _ := strconv.ParseFloat(strings.TrimSuffix(computedStyle.Get("width").String(), "px"), 64)
	heightWithBorder, _ := strconv.ParseFloat(strings.TrimSuffix(computedStyle.Get("height").String(), "px"), 64)
	borderUp, _ := strconv.ParseFloat(strings.TrimSuffix(computedStyle.Get("border-top-width").String(), "px"), 64)
	borderDown, _ := strconv.ParseFloat(strings.TrimSuffix(computedStyle.Get("border-bottom-width").String(), "px"), 64)
	borderLeft, _ := strconv.ParseFloat(strings.TrimSuffix(computedStyle.Get("border-left-width").String(), "px"), 64)
	borderRight, _ := strconv.ParseFloat(strings.TrimSuffix(computedStyle.Get("border-right-width").String(), "px"), 64)

	c.Height = heightWithBorder - borderUp - borderDown
	c.Width = widthWithBorder - borderLeft - borderRight

	c.CvsElement.Set("width", int(c.Width))
	c.CvsElement.Set("height", int(c.Height))
	println(int(c.Height), int(c.Width))
	if c.Cvs != nil {
		go c.Cvs.Stop()
	}
	c.Cvs, _ = canvas.NewCanvas2d(false)
	c.Cvs.Set(c.CvsElement, int(c.Width), int(c.Height))
	c.Cvs.Start(c.Fps, render(c))
}

func render(c *model.GameContext) func(*draw2dimg.GraphicContext) bool {
	return func(gc *draw2dimg.GraphicContext) bool {
		val := c.Animation(gc, c)
		c.T += c.RenderDelay

		return val
	}
}
