package canvas

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gowebapi/webapi/graphics/webgl"
	"github.com/gowebapi/webapi/html/canvas"
)

func InitCanvas(c *GameContext) {
	if c.ResolutionScale != 0 {
		style := c.Window.GetComputedStyle(c.CvsElement, nil)

		pixelRatio := c.Window.DevicePixelRatio()
		fmt.Printf("device pixel ratio is %v\n", pixelRatio)

		widthWithBorder, _ := strconv.ParseFloat(strings.TrimSuffix(style.GetPropertyValue("width"), "px"), 32)
		heightWithBorder, _ := strconv.ParseFloat(strings.TrimSuffix(style.GetPropertyValue("height"), "px"), 32)
		borderUp, _ := strconv.ParseFloat(strings.TrimSuffix(style.GetPropertyValue("border-top-width"), "px"), 32)
		borderDown, _ := strconv.ParseFloat(strings.TrimSuffix(style.GetPropertyValue("border-bottom-width"), "px"), 32)
		borderLeft, _ := strconv.ParseFloat(strings.TrimSuffix(style.GetPropertyValue("border-left-width"), "px"), 32)
		borderRight, _ := strconv.ParseFloat(strings.TrimSuffix(style.GetPropertyValue("border-right-width"), "px"), 32)

		c.Height = float32(heightWithBorder-borderUp-borderDown) * float32(pixelRatio) / c.ResolutionScale
		c.Width = float32(widthWithBorder-borderLeft-borderRight) * float32(pixelRatio) / c.ResolutionScale
	}
	c.CvsElement.SetAttribute("height", fmt.Sprint(c.Width))
	c.CvsElement.SetAttribute("width", fmt.Sprint(c.Height))
	cvsHTML := canvas.HTMLCanvasElementFromWrapper(c.CvsElement)

	if c.SecondaryCanvas != nil {
		cellWidth, cellHeight := c.Animator.Dimensions()
		c.SecondaryCanvas.SetAttribute("height", fmt.Sprint(cellHeight))
		c.SecondaryCanvas.SetAttribute("width", fmt.Sprint(cellWidth))
		secondaryHTML := canvas.HTMLCanvasElementFromWrapper(c.SecondaryCanvas)

		c.ZoomCanvas.SetAttribute("height", fmt.Sprint(cellHeight)) //c.Height)) //cellHeight))
		c.ZoomCanvas.SetAttribute("width", fmt.Sprint(cellWidth))   //c.Width))   //cellWidth))
		zoomHTML := canvas.HTMLCanvasElementFromWrapper(c.ZoomCanvas)

		c.GL = webgl.RenderingContextFromWrapper(secondaryHTML.GetContext("webgl", map[string]any{"alpha": false}))
		c.ZoomCtx = canvas.CanvasRenderingContext2DFromWrapper(zoomHTML.GetContext("2d", map[string]any{"alpha": false}))
		c.DisplayCtx = canvas.CanvasRenderingContext2DFromWrapper(cvsHTML.GetContext("2d", map[string]any{"alpha": false}))
		c.DisplayCtx.SetImageSmoothingEnabled(false)
		c.GL.Viewport(0, 0, int(cellWidth), int(cellHeight))
	} else {
		c.GL = webgl.RenderingContextFromWrapper(cvsHTML.GetContext("webgl", map[string]any{"alpha": false}))
		c.GL.Viewport(0, 0, int(c.Width), int(c.Height))
	}
}
