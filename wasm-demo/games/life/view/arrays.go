package view

import (
	"math/rand"

	"github.com/gowebapi/webapi/core/jsconv"
	"github.com/gowebapi/webapi/graphics/webgl"

	"github.com/AislingHeanue/aisling-codes/wasm-demo/canvas"
)

func (lg *LifeGame) setPixelsInTexture(c *canvas.GameContext, in [][]bool) {
	c.GL.BindTexture(webgl.TEXTURE_2D, lg.writeTexture)
	c.GL.TexImage2D(webgl.TEXTURE_2D, 0, int(webgl.RGBA), lg.CellWidth, lg.CellHeight, 0, webgl.RGBA, webgl.UNSIGNED_BYTE, webgl.UnionFromJS(jsconv.UInt8ToJs(setupPixelArray(in))))
	c.GL.BindTexture(webgl.TEXTURE_2D, &webgl.Texture{})
}

func emptyArray(width int, height int) [][]bool {
	m := make([][]bool, height)
	for i := range m {
		m[i] = make([]bool, width)
	}

	return m
}

func randomArray(width int, height int) [][]bool {
	m := make([][]bool, height)
	for i := range m {
		m[i] = make([]bool, width)
		for j := range m[i] {
			if rand.Float32() > 0.8 {
				m[i][j] = true
			}
		}
	}

	return m
}

func setupPixelArray(m [][]bool) []uint8 {
	on := []uint8{255, 255, 255, 255}
	off := []uint8{0, 0, 0, 0}
	out := make([]uint8, 4*len(m)*len(m[0]))
	width := len(m[0])
	for i := range m {
		for j := range m[i] {
			if m[i][j] {
				out[4*(i*width+j)+0] = on[0]
				out[4*(i*width+j)+1] = on[1]
				out[4*(i*width+j)+2] = on[2]
				out[4*(i*width+j)+3] = on[3]
			} else {
				out[4*(i*width+j)+0] = off[0]
				out[4*(i*width+j)+1] = off[1]
				out[4*(i*width+j)+2] = off[2]
				out[4*(i*width+j)+3] = off[3]
			}
		}
	}

	return out
}
