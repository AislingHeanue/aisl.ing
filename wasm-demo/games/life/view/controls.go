package view

import (
	"fmt"

	"github.com/AislingHeanue/aisling-codes/wasm-demo/canvas"
	"github.com/AislingHeanue/aisling-codes/wasm-demo/games/life/parser"
)

func (lg *LifeGame) Reset(c *canvas.GameContext) {
	lg.setPixelsInTexture(c, emptyArray(lg.CellWidth, lg.CellHeight))
	// lg.swapTextures()
	lg.drawToCanvas(c)
}

func (lg *LifeGame) Random(c *canvas.GameContext) {
	lg.setPixelsInTexture(c, randomArray(lg.CellWidth, lg.CellHeight))
	lg.deathFrame(c)
	lg.swapTextures()
	lg.drawToCanvas(c)
}

func (lg *LifeGame) OpenFile(c *canvas.GameContext, path string) {
	newArray := parser.ReadFile(path)
	if newArray == nil {
		return
	}
	lg.CellHeight = len(newArray)
	lg.CellWidth = len(newArray[0])
	lg.Zoom = 3
	lg.DX = 0
	lg.DY = 0
	lg.ResizeBuffers(c)

	lg.setPixelsInTexture(c, newArray)
	fmt.Printf("width: %.2f, height: %.2f, cwidth: %d, cheight: %d, zoom: %.2f\n", c.Width, c.Height, lg.CellWidth, lg.CellHeight, lg.Zoom)
}

func (lg *LifeGame) OpenRandomFile(c *canvas.GameContext) {
	newArray, path := parser.ReadRandomFile() // TODO: should return ParsedStuff
	if newArray == nil {
		return
	}
	lg.OpenFileName = path
	lg.CellHeight = len(newArray)
	lg.CellWidth = len(newArray[0])
	lg.Zoom = 3
	lg.DX = 0
	lg.DY = 0
	lg.ResizeBuffers(c)

	lg.setPixelsInTexture(c, newArray)
	fmt.Printf("width: %.2f, height: %.2f, cwidth: %d, cheight: %d, zoom: %.2f\n", c.Width, c.Height, lg.CellWidth, lg.CellHeight, lg.Zoom)
}

func (lg *LifeGame) ResizeBuffers(c *canvas.GameContext) {
	fmt.Println("Resizing: width:", lg.CellWidth, "height:", lg.CellHeight)
	canvas.InitCanvas(c)
	lg.createBuffers(c)
}
