package controller

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gowebapi/webapi"
	"github.com/gowebapi/webapi/graphics/webgl"
	"github.com/gowebapi/webapi/html/canvas"
	"github.com/gowebapi/webapi/html/htmlcommon"
)

func InitCanvas(c *GameContext) {
	c.Window = webapi.GetWindow()
	c.Document = c.Window.Document()
	c.CvsElement = c.Document.GetElementById("wasm-canvas")
	style := c.Window.GetComputedStyle(c.CvsElement, nil)

	pixelRatio := c.Window.DevicePixelRatio()
	fmt.Printf("device pixel ratio is %v\n", pixelRatio)
	// c.CvsElement = cvsElement.Value_JS

	widthWithBorder, _ := strconv.ParseFloat(strings.TrimSuffix(style.GetPropertyValue("width"), "px"), 32)
	heightWithBorder, _ := strconv.ParseFloat(strings.TrimSuffix(style.GetPropertyValue("height"), "px"), 32)
	borderUp, _ := strconv.ParseFloat(strings.TrimSuffix(style.GetPropertyValue("border-top-width"), "px"), 32)
	borderDown, _ := strconv.ParseFloat(strings.TrimSuffix(style.GetPropertyValue("border-bottom-width"), "px"), 32)
	borderLeft, _ := strconv.ParseFloat(strings.TrimSuffix(style.GetPropertyValue("border-left-width"), "px"), 32)
	borderRight, _ := strconv.ParseFloat(strings.TrimSuffix(style.GetPropertyValue("border-right-width"), "px"), 32)

	c.Height = float32(heightWithBorder-borderUp-borderDown) * float32(pixelRatio) / c.ResolutionScale
	c.Width = float32(widthWithBorder-borderLeft-borderRight) * float32(pixelRatio) / c.ResolutionScale

	c.CvsElement.SetAttribute("height", fmt.Sprint(c.Height))
	c.CvsElement.SetAttribute("width", fmt.Sprint(c.Width))

	cvsHTML := canvas.HTMLCanvasElementFromWrapper(c.CvsElement)
	glWrapper := cvsHTML.GetContext("webgl", nil)
	c.GL = webgl.RenderingContextFromWrapper(glWrapper)

	c.Animator.Init(c)

	c.Window.RequestAnimationFrame(htmlcommon.FrameRequestCallbackToJS(wrapAnimator(c)))
}

func wrapAnimator(c *GameContext) func(float64) {
	return func(time float64) {
		c.T = float32(time) / 1000 // milliseconds to seconds
		c.Animator.Render(c)
		c.Window.RequestAnimationFrame(htmlcommon.FrameRequestCallbackToJS(wrapAnimator(c)))
	}

}
