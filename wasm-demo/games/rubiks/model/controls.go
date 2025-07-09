package model

import (
	"fmt"
	"math/rand"
)

type CubeController struct {
	Context *CubeCubeContext
}

type Turn string

func (cc *CubeController) QueueEvent(turns ...Turn) {
	for _, t := range turns {
		if len(t) == 2 {
			switch string(t[1]) {
			case "'":
				cc.Context.AnimationHandler.AddEvent(string(t[0]), true)
			case "2":
				cc.Context.AnimationHandler.AddEvent(string(t[0]), false)
				cc.Context.AnimationHandler.AddEvent(string(t[0]), false)
			}
		} else if len(t) == 1 {
			cc.Context.AnimationHandler.AddEvent(string(t[0]), false)
		}
	}
}

func (cc *CubeController) ResetAngles() {
	cc.Context.AngleX = 0
	cc.Context.AngleY = 0
}

func (cc *CubeController) Shuffle() {
	fmt.Println("shuffling!")
	previousFace := ""
	face := ""
	faces := []string{"u", "d", "b", "f", "l", "r"}
	turnDirections := []string{"clockwise", "anticlockwise", "double"}
	for range 20 {
		for face == previousFace {
			face = faces[rand.Intn(len(faces))]
		}
		previousFace = face
		direction := turnDirections[rand.Intn(len(turnDirections))]
		switch direction {
		case "clockwise":
			cc.QueueEvent(Turn(face))
		case "anticlockwise":
			cc.QueueEvent(Turn(face + "'"))
		case "double":
			cc.QueueEvent(Turn(face))
			cc.QueueEvent(Turn(face))
		}
	}

}
