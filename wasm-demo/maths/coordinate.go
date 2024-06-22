package maths

import (
	"fmt"
	"strconv"
	"strings"
)

type Coordinate []int

func (c Coordinate) String() string {
	vals := make([]string, len(c))
	for i, v := range c {
		vals[i] = fmt.Sprintf("%d", v)
	}

	return strings.Join(vals, ",")
}

func parseCoord(s string) Coordinate {
	vals := strings.Split(s, ",")
	out := make(Coordinate, len(vals))
	for i, v := range vals {
		var err error
		out[i], err = strconv.Atoi(v)
		if err != nil {
			panic(fmt.Sprintf("coordinate could not be parsed: %s", s))
		}
	}

	return out
}

func (c Coordinate) X() int {
	return c[0]
}

func (c Coordinate) Y() int {
	return c[1]
}

func (c Coordinate) Z() int {
	return c[2]
}
