package controller

import (
	"fmt"
	"math/rand"

	"github.com/AislingHeanue/aisling-codes/wasm-demo/games/rubiks/model"
)

type CubeController struct {
	ccc *model.CubeCubeContext
}

func (cc *CubeController) QueueEvent(face string, reverse bool) bool {
	return cc.ccc.AnimationHandler.AddEvent(face, reverse)
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
			cc.QueueEvent(face, false)
		case "anticlockwise":
			cc.QueueEvent(face, true)
		case "double":
			cc.QueueEvent(face, false)
			cc.QueueEvent(face, false)
		}
	}

}
