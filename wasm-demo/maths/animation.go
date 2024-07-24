package maths

import (
	"math"
)

type RubiksAnimationHandler struct {
	RubiksCube                 *RubiksCube
	CopyRubiksCube             RubiksCube
	currentEventIndices        []int
	EventsWhichNeedToBeRotated []RubiksEvent
	events                     []RubiksEvent
	done                       bool
	MaxTicks                   int
}

type RubiksEvent struct {
	face    Face
	reverse bool
	t       int
}

func (a *RubiksAnimationHandler) AddEvent(face string, reverse bool) bool {
	if _, ok := turnMap[Face(face)]; !ok {
		// face not recognised, do nothing
		return false
	}
	a.events = append(a.events, RubiksEvent{
		face:    Face(face),
		reverse: reverse,
		t:       0,
	})

	return true
}

func (a *RubiksAnimationHandler) Tick() bool {
	a.currentEventIndices = []int{}
	// for every event in order
	for i, event := range a.events {
		// if the event is not finished
		if event.t < a.MaxTicks {
			allowedToMove := true
			// if there are other moves currently moving first, check that this can also be done
			if len(a.currentEventIndices) != 0 {
				// check every move between the first currently moving move, and the
				// current move being looked at.
				for j := a.currentEventIndices[0]; j < i; j++ {
					if a.events[j].t == a.MaxTicks {
						continue
					}
					// Each move has a list of moves that can be done in parallel with them
					// If every move between the very first currently-moving move and this move
					// match this criteria, then this move can also be ticked.
					matchFoundInAllowedList := false
					for _, otherAllowedMove := range turnMap[a.events[j].face].allowedConcurrent {
						if otherAllowedMove == event.face {
							matchFoundInAllowedList = true
						}
					}
					if !matchFoundInAllowedList {
						allowedToMove = false
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

		// false: no events occurred
		return false
	}
	// count the events left in the list after the first currently-moving move. Zero if nothing is moving.
	eventsRemaining := len(a.events) - a.currentEventIndices[0]
	a.EventsWhichNeedToBeRotated = []RubiksEvent{}
	for _, j := range a.currentEventIndices {
		if a.events[j].t == a.MaxTicks {
			a.doTurn(a.events[j])
			// for each instance of doTurn run, the number of unfinished events decreases by one
			eventsRemaining--
		} else {
			a.EventsWhichNeedToBeRotated = append(a.EventsWhichNeedToBeRotated, a.events[j])
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
	a.RubiksCube.Turn(event.face, event.reverse)
}

func (a *RubiksAnimationHandler) DoEvent(event RubiksEvent, origin *Point) {
	info, ok := turnMap[event.face]
	rotationScale := 1.
	if info.reverse {
		rotationScale *= -1.
	}
	if event.reverse {
		rotationScale *= -1
	}
	if !ok {
		return // face not recognised, do nothing
	}
	coords := a.RubiksCube.getCubeSubset(info.xSelector, info.ySelector, info.zSelector)
	for _, coord := range coords {
		x := coord[0]
		y := coord[1]
		z := coord[2]
		a.CopyRubiksCube.Data[x][y][z] = a.RubiksCube.Data[x][y][z].Rotate(*origin, float32(float64(event.t)*rotationScale*math.Pi/float64(2*a.MaxTicks-2)), info.axis)
	}

}

func (a *RubiksAnimationHandler) FlushAll() {
	for !a.done {
		a.Tick()
	}
}
