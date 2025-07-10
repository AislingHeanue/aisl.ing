package canvas

import (
	"github.com/gowebapi/webapi/graphics/webgl"
	"github.com/gowebapi/webapi/html/canvas"

	"github.com/AislingHeanue/aisling-codes/wasm-demo/common"
)

func InitCanvas(c *common.GameContext) {
	res := c.Window.JSValue().Call("resizeCanvas", float32(c.ResolutionScale), true)
	c.Height = float32(res.Get("height").Float())
	c.Width = float32(res.Get("width").Float())

	if c.AutoSizePixels {
		c.PixelsWidth = int(c.Width)
		c.PixelsHeight = int(c.Height)
	}

	if c.Zoom == 0 {
		c.Zoom = 1
	}
	res2 := c.Window.JSValue().Call("setupMultipleCanvases", c.RenderingCanvas.JSValue(), c.ZoomCanvas.JSValue(), c.PixelsHeight, c.PixelsWidth, c.SmoothImage)

	c.GL = webgl.RenderingContextFromJS(res2.Get("gl"))
	c.ZoomCtx = canvas.CanvasRenderingContext2DFromJS(res2.Get("zoomContext"))
	c.DisplayCtx = canvas.CanvasRenderingContext2DFromJS(res2.Get("displayContext"))
	// } else {
	// 	res := c.Window.JSValue().Call("setupCanvas", c.Height, c.Width)
	//
	// 	c.GL = webgl.RenderingContextFromJS(res.Get("gl"))
	// }
}
