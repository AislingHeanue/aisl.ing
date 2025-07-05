package canvas

import (
	"github.com/gowebapi/webapi/dom/domcore"
	"github.com/gowebapi/webapi/graphics/webgl"
	"github.com/gowebapi/webapi/html/canvas"
)

func InitCanvas(c *GameContext) {
	res := c.Window.JSValue().Call("resizeCanvas", float32(c.ResolutionScale), true)
	c.Height = float32(res.Get("height").Float())
	c.Width = float32(res.Get("width").Float())

	if c.ZoomCanvas != nil && c.RenderingCanvas != nil {
		cellWidth, cellHeight := c.Animator.Dimensions()
		res := c.Window.JSValue().Call("setupMultipleCanvases", c.RenderingCanvas.JSValue(), c.ZoomCanvas.JSValue(), cellHeight, cellWidth)

		c.GL = webgl.RenderingContextFromJS(res.Get("gl"))
		c.ZoomCtx = canvas.CanvasRenderingContext2DFromJS(res.Get("zoomContext"))
		c.DisplayCtx = canvas.CanvasRenderingContext2DFromJS(res.Get("displayContext"))
	} else {
		res := c.Window.JSValue().Call("setupCanvas", c.Height, c.Width)

		c.GL = webgl.RenderingContextFromJS(res.Get("gl"))
	}
}

func AddListener(c *GameContext, event string, listener domcore.EventListener) {
	c.Window.JSValue().Call("canvasEventListener", event, domcore.NewEventListener(listener).JSValue())
}
