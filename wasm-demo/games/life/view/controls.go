package view

import (
	"github.com/AislingHeanue/aisling-codes/wasm-demo/canvas"
	"github.com/AislingHeanue/aisling-codes/wasm-demo/games/life/parser"
	"github.com/AislingHeanue/aisling-codes/wasm-demo/util"
)

func (lg *LifeGame) Reset(c *util.GameContext) {
	lg.setPixelsInTexture(c, emptyArray(c.PixelsWidth, c.PixelsHeight))
}

func (lg *LifeGame) Random(c *util.GameContext) {
	lg.setPixelsInTexture(c, randomArray(c.PixelsWidth, c.PixelsHeight))
}

func (lg *LifeGame) OpenFile(c *util.GameContext, path string) {
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

	lg.setPixelsInTexture(c, newArray)
	// fmt.Printf("width: %.2f, height: %.2f, cwidth: %d, cheight: %d, zoom: %.2f\n", c.Width, c.Height, lg.CellWidth, lg.CellHeight, lg.Zoom)
}

func (lg *LifeGame) OpenRandomFile(c *util.GameContext) {
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

	lg.setPixelsInTexture(c, newArray)
	// fmt.Printf("width: %.2f, height: %.2f, cwidth: %d, cheight: %d, zoom: %.2f\n", c.Width, c.Height, lg.CellWidth, lg.CellHeight, lg.Zoom)
}

func (lg *LifeGame) ResizeBuffers(c *util.GameContext) {
	// fmt.Println("Resizing: width:", lg.CellWidth, "height:", lg.CellHeight)
	canvas.InitCanvas(c)
	lg.createBuffers(c)
}
