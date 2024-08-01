package model

type CubeCubeContext struct {
	CubeDimension int
	TurnFrames    int
	AngleX        float32
	AngleY        float32
	AnchorX       float32
	AnchorY       float32
	AnchorAngleX  float32
	AnchorAngleY  float32
	MouseDown     bool

	AnimationHandler *RubiksAnimationHandler
}
