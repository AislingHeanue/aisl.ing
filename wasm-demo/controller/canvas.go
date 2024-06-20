package controller

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/AislingHeanue/aisling-codes/wasm-demo/model"

	"github.com/gowebapi/webapi"
	"github.com/gowebapi/webapi/graphics/webgl"
	"github.com/gowebapi/webapi/html/canvas"
	"github.com/gowebapi/webapi/html/htmlcommon"
)

func InitCanvas(c *model.GameContext) {
	c.Window = webapi.GetWindow()
	c.Document = c.Window.Document()
	c.CvsElement = c.Document.GetElementById("wasm-canvas")
	style := c.Window.GetComputedStyle(c.CvsElement, nil)
	// c.CvsElement = cvsElement.Value_JS

	widthWithBorder, _ := strconv.ParseFloat(strings.TrimSuffix(style.GetPropertyValue("width"), "px"), 64)
	heightWithBorder, _ := strconv.ParseFloat(strings.TrimSuffix(style.GetPropertyValue("height"), "px"), 64)
	borderUp, _ := strconv.ParseFloat(strings.TrimSuffix(style.GetPropertyValue("border-top-width"), "px"), 64)
	borderDown, _ := strconv.ParseFloat(strings.TrimSuffix(style.GetPropertyValue("border-bottom-width"), "px"), 64)
	borderLeft, _ := strconv.ParseFloat(strings.TrimSuffix(style.GetPropertyValue("border-left-width"), "px"), 64)
	borderRight, _ := strconv.ParseFloat(strings.TrimSuffix(style.GetPropertyValue("border-right-width"), "px"), 64)

	c.Height = (heightWithBorder - borderUp - borderDown) / c.ResolutionScale
	c.Width = (widthWithBorder - borderLeft - borderRight) / c.ResolutionScale

	c.CvsElement.SetAttribute("height", fmt.Sprint(c.Height))
	c.CvsElement.SetAttribute("width", fmt.Sprint(c.Width))

	cvsHTML := canvas.HTMLCanvasElementFromWrapper(c.CvsElement)
	glWrapper := cvsHTML.GetContext("webgl", nil)
	gl := webgl.RenderingContextFromWrapper(glWrapper)

	c.Animator.Init(c)
	// c.Animator.CreateBuffers(gl, c)
	program := c.Animator.CreateShaders(gl, c)

	c.Window.RequestAnimationFrame(htmlcommon.FrameRequestCallbackToJS(wrapAnimator(gl, program, c, c.Animator.Render)))
}

func wrapAnimator(gl *webgl.RenderingContext, p *webgl.Program, c *model.GameContext, f model.RenderFunc) func(float64) {
	return func(time float64) {
		c.T = time / 1000 // milliseconds to seconds
		f(gl, p, c)
		c.Window.RequestAnimationFrame(htmlcommon.FrameRequestCallbackToJS(wrapAnimator(gl, p, c, c.Animator.Render)))
	}

}
