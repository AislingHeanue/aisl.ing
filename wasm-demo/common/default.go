package common

import "github.com/gowebapi/webapi/dom/domcore"

type DefaultGame struct {
	Parent *ShaderGame
}

// CanRunBetweenFrames implements GameInfo.
func (d *DefaultGame) CanRunBetweenFrames() bool {
	return false
}

// GetDrawShape implements GameInfo.
func (d *DefaultGame) GetDrawShape(c *GameContext) DrawShape {
	return DrawShape{}
}

// GetTps implements GameInfo.
func (d *DefaultGame) GetTps() float32 {
	return 30
}

// InitListeners implements GameInfo.
func (d *DefaultGame) InitListeners(c *GameContext) {
}

// PostSetup implements GameInfo.
func (d *DefaultGame) PostSetup(c *GameContext) {
}

// PreSetup implements GameInfo.
func (d *DefaultGame) PreSetup(c *GameContext) {
}

// SetParent implements GameInfo.
func (d *DefaultGame) SetParent(parent *ShaderGame) {
	d.Parent = parent
}

// Tick implements GameInfo.
func (d *DefaultGame) Tick(c *GameContext) bool {
	return false
}

type DefaultActionHandler[Context, Controller any] struct{}

type a struct{}
type b struct{}

var _ ActionHandler[a, b] = DefaultActionHandler[a, b]{}

// Click implements ActionHandler.
func (d DefaultActionHandler[Context, Controller]) Click(c *GameContext, context *Context, controller Controller, e *domcore.Event) {
}

// Drag implements ActionHandler.
func (d DefaultActionHandler[Context, Controller]) Drag(c *GameContext, context *Context, controller Controller, e *domcore.Event) {
}

// DragTouch implements ActionHandler.
func (d DefaultActionHandler[Context, Controller]) DragTouch(c *GameContext, context *Context, controller Controller, e *domcore.Event) {
}

// Keyboard implements ActionHandler.
func (d DefaultActionHandler[Context, Controller]) Keyboard(c *GameContext, context *Context, controller Controller, e *domcore.Event) {
}

// MouseUp implements ActionHandler.
func (d DefaultActionHandler[Context, Controller]) MouseUp(c *GameContext, context *Context, controller Controller, e *domcore.Event) {
}

// Resize implements ActionHandler.
func (d DefaultActionHandler[Context, Controller]) Resize(c *GameContext, context *Context, controller Controller, e *domcore.Event) {
}

// Touch implements ActionHandler.
func (d DefaultActionHandler[Context, Controller]) Touch(c *GameContext, context *Context, controller Controller, e *domcore.Event) {
}

// TouchUp implements ActionHandler.
func (d DefaultActionHandler[Context, Controller]) TouchUp(c *GameContext, context *Context, controller Controller, e *domcore.Event) {
}

func (d DefaultActionHandler[Context, Controller]) Wheel(c *GameContext, context *Context, controller Controller, e *domcore.Event) {
}

