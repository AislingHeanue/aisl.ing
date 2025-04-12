package model

type LifeContext struct {
	DX        float32
	DY        float32
	AnchorX   float32
	AnchorY   float32
	AnchorDX  float32
	AnchorDY  float32
	MouseDown bool

	AnchorPinchDistance float32
	AnchorZoom          float32
	Zooming             bool

	CellHeight int
	CellWidth  int

	Zoom           float32
	Paused         bool
	SimulationSize int
	Tps            int
	NeedsRestart   bool // will probably remove this one
	Loop           bool

	OpenFileName string
}
