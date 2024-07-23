package graphics

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
	// c.CvsElement = cvsElement.Value_JS

	widthWithBorder, _ := strconv.ParseFloat(strings.TrimSuffix(style.GetPropertyValue("width"), "px"), 32)
	heightWithBorder, _ := strconv.ParseFloat(strings.TrimSuffix(style.GetPropertyValue("height"), "px"), 32)
	borderUp, _ := strconv.ParseFloat(strings.TrimSuffix(style.GetPropertyValue("border-top-width"), "px"), 32)
	borderDown, _ := strconv.ParseFloat(strings.TrimSuffix(style.GetPropertyValue("border-bottom-width"), "px"), 32)
	borderLeft, _ := strconv.ParseFloat(strings.TrimSuffix(style.GetPropertyValue("border-left-width"), "px"), 32)
	borderRight, _ := strconv.ParseFloat(strings.TrimSuffix(style.GetPropertyValue("border-right-width"), "px"), 32)

	c.Height = float32(heightWithBorder-borderUp-borderDown) / c.ResolutionScale
	c.Width = float32(widthWithBorder-borderLeft-borderRight) / c.ResolutionScale

	c.CvsElement.SetAttribute("height", fmt.Sprint(c.Height))
	c.CvsElement.SetAttribute("width", fmt.Sprint(c.Width))
}

func StartAnimation(c *GameContext) {
	c.Animator.Init(c)
	c.Animator.InitListeners(c)
	cvsHTML := canvas.HTMLCanvasElementFromWrapper(c.CvsElement)
	glWrapper := cvsHTML.GetContext("webgl", nil)
	gl := webgl.RenderingContextFromWrapper(glWrapper)
	program := c.Animator.CreateShaders(gl, c)
	c.Window.RequestAnimationFrame(htmlcommon.FrameRequestCallbackToJS(wrapAnimator(gl, program, c, c.Animator.Render)))
}

func wrapAnimator(gl *webgl.RenderingContext, p *webgl.Program, c *GameContext, f RenderFunc) func(float64) {
	return func(time float64) {
		c.T = float32(time) / 1000 // milliseconds to seconds
		f(gl, p, c)
		c.Window.RequestAnimationFrame(htmlcommon.FrameRequestCallbackToJS(wrapAnimator(gl, p, c, c.Animator.Render)))
	}

}
