package model

import (
	"github.com/AislingHeanue/aisling-codes/wasm-demo/common"
)

func GetBuffers(a *RubiksAnimationHandler, origin Point) common.DrawShape {
	var out common.DrawShape

	// handle any animations in progress
	a.CopyRubiksCube = a.RubiksCube.Copy()
	for _, event := range a.EventsWhichNeedToBeRotated {
		a.DoEvent(event, origin)
	}

	// get the buffers for all cubes and put them in an array
	out.VerticesArray = []float32{}
	out.IndicesArray = []uint16{}
	out.ColourArray = []float32{}
	vCountOffset := 0
	iCountOffset := 0
	eCountOffset := 0

	// for each cube, get the buffers for it and add them to the existing buffers.
	for _, shape := range a.CopyRubiksCube.FlattenBuffers() {
		offsetIndices := make([]uint16, len(shape.IndicesArray))
		offsetEdges := make([]uint16, len(shape.EdgeIndices))
		for i, v := range shape.IndicesArray {
			offsetIndices[i] = v + uint16(vCountOffset)
		}
		for i, v := range shape.EdgeIndices {
			offsetEdges[i] = v + uint16(vCountOffset)
		}

		out.VerticesArray = append(out.VerticesArray, shape.VerticesArray...)
		out.IndicesArray = append(out.IndicesArray, offsetIndices...)
		out.ColourArray = append(out.ColourArray, shape.ColourArray...)
		out.EdgeIndices = append(out.EdgeIndices, offsetEdges...)

		vCountOffset += shape.VCount
		iCountOffset += shape.ICount
		eCountOffset += shape.ECount
	}

	out.VCount = vCountOffset
	out.ICount = iCountOffset
	out.ECount = eCountOffset

	return out
}
