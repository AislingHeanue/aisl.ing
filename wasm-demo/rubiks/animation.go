package rubiks

import (
	"fmt"
	"math"

	"github.com/AislingHeanue/aisling-codes/wasm-demo/maths"
)

const maxTicks = 30

var concurrentTurnsAllowed map[string][]string = map[string][]string{
	"u": {"e", "d"},
	"e": {"d", "u"},
	"d": {"u", "e"},
	"y": {},
	"r": {"m", "l"},
	"m": {"l", "r"},
	"l": {"r", "m"},
	"x": {},
	"f": {"s", "b"},
	"s": {"b", "f"},
	"b": {"f", "s"},
	"z": {},
}

type animationHandler interface {
	AddEvent(face string, reverse bool)
	GetBuffers(origin *maths.Point) DrawShape
	FlushAll()
	Tick() bool
}

type RubiksAnimationHandler struct {
	controller                 CubeController
	rubiksCube                 RubiksCube
	copyRubiksCube             RubiksCube
	currentEventIndices        []int
	eventsWhichNeedToBeRotated []RubiksEvent
	events                     []RubiksEvent
	done                       bool
}

type RubiksEvent struct {
	face    string
	reverse bool
	t       int
}

var _ animationHandler = &RubiksAnimationHandler{}

func (a *RubiksAnimationHandler) AddEvent(face string, reverse bool) {
	// face not recognised, do nothing
	fmt.Println("Trying to make an event with the letter " + face)
	if _, ok := concurrentTurnsAllowed[face]; !ok {
		return
	}
	fmt.Println("event added")
	a.events = append(a.events, RubiksEvent{
		face:    face,
		reverse: reverse,
		t:       0,
	})
}

func (a *RubiksAnimationHandler) Tick() bool {
	a.currentEventIndices = []int{}
	// for every event in order
	for i, event := range a.events {
		fmt.Println(event)
		// if the event is not finished
		if event.t < maxTicks {
			allowedToMove := true
			// if there are other moves currently moving first, check that this can also be done
			if len(a.currentEventIndices) != 0 {
				// check every move between the first currently moving move, and the
				// current move being looked at.
				for j := a.currentEventIndices[0]; j <= i; j++ {
					// Each move has a list of moves that can be done in parallel with them
					// If every move between the very first currently-moving move and this move
					// match this criteria, then this move can also be ticked.
					for _, otherAllowedMove := range concurrentTurnsAllowed[a.events[j].face] {
						matchFoundInAllowedList := false
						if otherAllowedMove == event.face {
							matchFoundInAllowedList = true
						}
						if !matchFoundInAllowedList {
							allowedToMove = false
						}
					}
				}
			}
			if allowedToMove {
				a.currentEventIndices = append(a.currentEventIndices, i)
				a.events[i].t++
			}
		}
	}
	if len(a.currentEventIndices) == 0 {
		a.done = true

		// false: no events occured
		return false
	}
	// count the events left in the list after the first currently-moving move. Zero if nothing is moving.
	eventsRemaining := len(a.events) - a.currentEventIndices[0]
	a.eventsWhichNeedToBeRotated = []RubiksEvent{}
	for _, j := range a.currentEventIndices {
		if a.events[j].t == maxTicks {
			a.doTurn(a.events[j])
			// for each instance of doTurn run, the number of unfinished events decreases by one
			eventsRemaining--
		} else {
			a.eventsWhichNeedToBeRotated = append(a.eventsWhichNeedToBeRotated, a.events[j])
		}
	}

	// If there is at least one event left in the queue that isn't the maximum value,
	// we can assume that more ticks need to be done before
	// the queue is cleared.
	a.done = eventsRemaining != 0

	// true: at least one face was ticked
	return true
}

func (a *RubiksAnimationHandler) doTurn(event RubiksEvent) {
	switch event.face {
	case "u":
		a.controller.U(event.reverse)
	case "e":
		a.controller.E(event.reverse)
	case "d":
		a.controller.D(event.reverse)
	case "y":
		a.controller.Y(event.reverse)
	case "r":
		a.controller.R(event.reverse)
	case "m":
		a.controller.M(event.reverse)
	case "l":
		a.controller.L(event.reverse)
	case "x":
		a.controller.X(event.reverse)
	case "f":
		a.controller.F(event.reverse)
	case "s":
		a.controller.S(event.reverse)
	case "b":
		a.controller.B(event.reverse)
	case "z":
		a.controller.Z(event.reverse)
	}
}

func (a *RubiksAnimationHandler) doEvent(event RubiksEvent, origin *maths.Point) {
	info, ok := turnMap[event.face]
	rotationScale := 1.
	if info.reverse {
		rotationScale = -1.
	}
	if !ok {
		return // face not recognised, do nothing
	}
	coords := a.rubiksCube.getCubeSubset(info.xSelector, info.ySelector, info.zSelector)
	for _, coord := range coords {
		x := coord[0]
		y := coord[1]
		z := coord[2]
		a.copyRubiksCube.data[x][y][z] = a.rubiksCube.data[x][y][z].Rotate(*origin, float32(rotationScale*math.Pi/(2*maxTicks)), info.axis)
	}

}

func (a *RubiksAnimationHandler) FlushAll() {
	for !a.done {
		a.Tick()
	}
}

func (a *RubiksAnimationHandler) GetBuffers(origin *maths.Point) DrawShape {
	a.copyRubiksCube = a.rubiksCube.copy()
	for _, event := range a.eventsWhichNeedToBeRotated {
		a.doEvent(event, origin)
	}
	return a.copyRubiksCube.GroupBuffers()
}
