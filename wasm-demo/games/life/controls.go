package life

import (
	"github.com/AislingHeanue/aisling-codes/wasm-demo/canvas"
	"github.com/AislingHeanue/aisling-codes/wasm-demo/common"
	"github.com/AislingHeanue/aisling-codes/wasm-demo/games/life/parser"
)

func (lg *LifeGame) Reset(c *common.GameContext) {
	lg.Parent.SetPixelsInTexture(c, emptyArray(c.PixelsWidth, c.PixelsHeight))
}

func (lg *LifeGame) Random(c *common.GameContext) {
	lg.Parent.SetPixelsInTexture(c, randomArray(c.PixelsWidth, c.PixelsHeight))
}

func (lg *LifeGame) OpenFile(c *common.GameContext, path string) {
	newArray := parser.ReadFile(path)
	if newArray == nil {
		return
	}
	c.PixelsHeight = len(newArray)
	c.PixelsWidth = len(newArray[0])
	c.Zoom = 3
	c.DX = 0
	c.DY = 0
	lg.ResizeBuffers(c)

	lg.Parent.SetPixelsInTexture(c, newArray)
	// fmt.Printf("width: %.2f, height: %.2f, cwidth: %d, cheight: %d, zoom: %.2f\n", c.Width, c.Height, lg.CellWidth, lg.CellHeight, lg.Zoom)
}

func (lg *LifeGame) OpenRandomFile(c *common.GameContext) {
	newArray, path := parser.ReadRandomFile() // TODO: should return ParsedStuff
	if newArray == nil {
		return
	}
	lg.OpenFileName = path
	c.PixelsHeight = len(newArray)
	c.PixelsWidth = len(newArray[0])
	c.Zoom = 3
	c.DX = 0
	c.DY = 0
	lg.ResizeBuffers(c)

	lg.Parent.SetPixelsInTexture(c, newArray)
	// fmt.Printf("width: %.2f, height: %.2f, cwidth: %d, cheight: %d, zoom: %.2f\n", c.Width, c.Height, lg.CellWidth, lg.CellHeight, lg.Zoom)
}

func (lg *LifeGame) ResizeBuffers(c *common.GameContext) {
	// fmt.Println("Resizing: width:", lg.CellWidth, "height:", lg.CellHeight)
	canvas.InitCanvas(c)
	lg.Parent.CreateBuffers(c)
}
