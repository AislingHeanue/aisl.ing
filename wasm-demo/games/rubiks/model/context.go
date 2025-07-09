package model

type CubeCubeContext struct {
	AngleX       float32
	AngleY       float32
	AnchorAngleX float32
	AnchorAngleY float32

	AnimationHandler *RubiksAnimationHandler

	TotalSideLength float32
	GapProportion   float32
	Origin          Point
	Dimension       int
	TurnSeconds     float32
	Tps             float32
}
