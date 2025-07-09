package canvas

import (
	"github.com/gowebapi/webapi/graphics/webgl"
	"github.com/gowebapi/webapi/html/canvas"

	"github.com/AislingHeanue/aisling-codes/wasm-demo/games/common"
)

func InitCanvas(c *common.GameContext) {
	res := c.Window.JSValue().Call("resizeCanvas", float32(c.ResolutionScale), true)
	c.Height = float32(res.Get("height").Float())
	c.Width = float32(res.Get("width").Float())

	if c.ZoomCanvas != nil && c.RenderingCanvas != nil {
		res := c.Window.JSValue().Call("setupMultipleCanvases", c.RenderingCanvas.JSValue(), c.ZoomCanvas.JSValue(), c.PixelsHeight, c.PixelsWidth)

		c.GL = webgl.RenderingContextFromJS(res.Get("gl"))
		c.ZoomCtx = canvas.CanvasRenderingContext2DFromJS(res.Get("zoomContext"))
		c.DisplayCtx = canvas.CanvasRenderingContext2DFromJS(res.Get("displayContext"))
	} else {
		res := c.Window.JSValue().Call("setupCanvas", c.Height, c.Width)

		c.GL = webgl.RenderingContextFromJS(res.Get("gl"))
	}
}
