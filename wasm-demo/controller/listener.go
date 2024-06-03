package controller

import (
	"fmt"
	"strconv"
	"syscall/js"

	"github.com/AislingHeanue/aisling-codes/wasm-demo/model"
)

func Log(v js.Value) {
	js.Global().Get("console").Call("log", v)
}

func RegisterListeners(c *model.GameContext) {
	c.Window.Call("addEventListener", "resize", wrapListener(InitCanvas, c))

	// for _, k := range []string{"red", "green", "blue"} {
	// 	res := c.Document.Call("getElementById", k)
	// 	handleUint8(k, c)(res, nil)
	// 	res.Call("addEventListener", "input", js.FuncOf(handleUint8(k, c)))
	// }

	res := c.Document.Call("getElementById", "dimension")
	handleDimension(c)(res, nil)
	res.Call("addEventListener", "input", js.FuncOf(handleDimension(c)))

	c.CvsElement.Call("addEventListener", "mousedown", wrapListener(clickCanvas, c))
	c.CvsElement.Call("addEventListener", "mousemove", wrapListener(dragCanvas, c))
	c.CvsElement.Call("addEventListener", "mouseup", wrapListener(mouseUp, c))
	c.CvsElement.Call("addEventListener", "mouseleave", wrapListener(mouseUp, c))
}

func handleUint8(name string, c *model.GameContext) func(js.Value, []js.Value) any {
	return func(this js.Value, args []js.Value) any {
		i, _ := strconv.Atoi(this.Get("value").String())
		switch name {
		case "red":
			c.Colour.R = uint8(i)
		case "green":
			c.Colour.G = uint8(i)
		case "blue":
			c.Colour.B = uint8(i)
		}

		return js.Null()
	}
}

func handleDimension(c *model.GameContext) func(js.Value, []js.Value) any {
	return func(this js.Value, args []js.Value) any {
		i, _ := strconv.Atoi(this.Get("value").String())
		c.Dimension = i
		c.Animator.Init(c)

		return js.Null()
	}
}

func wrapListener(f func(*model.GameContext, js.Value, []js.Value), c *model.GameContext) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		f(c, this, args)

		return js.Null()
	})
}

func clickCanvas(c *model.GameContext, _ js.Value, args []js.Value) {
	c.AnchorX, c.AnchorY = getRelativeMousePosition(c, args[0])
	c.AnchorAngleX = c.AngleX
	c.AnchorAngleY = c.AngleY
	c.MouseDown = true
}

func dragCanvas(c *model.GameContext, _ js.Value, args []js.Value) {
	if c.MouseDown {
		// fmt.Println("drag")
		mouseX, mouseY := getRelativeMousePosition(c, args[0])
		c.AngleX = (c.AnchorAngleX + 5*(c.AnchorY-mouseY))
		c.AngleY = (c.AnchorAngleY + 5*(c.AnchorX-mouseX))
	}
	// x := args[0].Get("offsetX").Float()
	// y := c.Height - args[0].Get("offsetY").Float()
	// fmt.Println(c.Cube.Faces[3].Lines[0])
	// fmt.Println(c.Cube.Faces[3].Lines[1])
	// fmt.Println(c.Cube.Faces[3].Lines[2])
	// fmt.Println(c.Cube.Faces[3].Lines[3])

	// c.Cube.Faces[3].CheckContains(x, y)
	// fmt.Println(c.Cube.Faces[3].CheckContains(x, y), c.Cube.Faces[3].Lines[1].PointIsLeft(x, y), c.Cube.Faces[3].Lines[1].PointIsLeft(40000, 40000))
}

func mouseUp(c *model.GameContext, _ js.Value, args []js.Value) {
	fmt.Println("up")
	c.MouseDown = false
}

func getRelativeMousePosition(c *model.GameContext, click js.Value) (float64, float64) {
	relativeX := click.Get("offsetX").Float() / c.Width
	relativeY := click.Get("offsetY").Float() / c.Height
	return relativeX, relativeY
}
