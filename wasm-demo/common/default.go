package common

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
