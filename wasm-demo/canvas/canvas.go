package canvas

import (
	"github.com/gowebapi/webapi"
	"github.com/gowebapi/webapi/graphics/webgl"
	"github.com/gowebapi/webapi/html/canvas"

	"github.com/AislingHeanue/aisling-codes/wasm-demo/common"
)

func InitCanvas(c *common.GameContext) {
	if !c.IsInitialised {
		c.Window = webapi.GetWindow()
		c.Document = c.Window.Document()
		c.Square = true // not configurable at the moment since I haven't tested it
		if c.ResolutionScale == 0 {
			c.ResolutionScale = 1
		}
		if c.Zoom == 0 {
			c.Zoom = 1
		}
		c.RenderingCanvas = c.Document.CreateElement("canvas", &webapi.Union{})
		c.ZoomCanvas = c.Document.CreateElement("canvas", &webapi.Union{})
		// create the animator which wraps the game and handles rendering and displaying the output
		c.Animator = &common.ShaderGame{
			GameInfo: c.Game,
		}
	}

	// set up heights and widths
	res := c.Window.JSValue().Call("resizeCanvas", float32(c.ResolutionScale), c.Square)
	c.Height = float32(res.Get("height").Float())
	c.Width = float32(res.Get("width").Float())
	if (c.PixelsWidth == 0 && c.PixelsHeight == 0) || c.AutoSizePixels {
		c.PixelsWidth = int(c.Width)
		c.PixelsHeight = int(c.Height)
		c.AutoSizePixels = true
	}

	// set up the canvases and rendering contexts
	res2 := c.Window.JSValue().Call("setupMultipleCanvases", c.RenderingCanvas.JSValue(), c.ZoomCanvas.JSValue(), c.PixelsHeight, c.PixelsWidth, c.SmoothImage)
	if !c.IsInitialised {
		c.GL = webgl.RenderingContextFromJS(res2.Get("gl"))
		c.ZoomCtx = canvas.CanvasRenderingContext2DFromJS(res2.Get("zoomContext"))
		c.DisplayCtx = canvas.CanvasRenderingContext2DFromJS(res2.Get("displayContext"))
	}

	c.IsInitialised = true
}
