package model

type DrawShape struct {
	Points        []*Point
	VerticesArray []float32
	IndicesArray  []uint16
	ColourArray   []float32
	VCount        int
	ICount        int
	CCount        int
}

func GetBuffers(a *RubiksAnimationHandler, origin Point) DrawShape {
	var out DrawShape

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

	// for each cube, get the buffers for it and add them to the existing buffers.
	for _, shape := range a.CopyRubiksCube.FlattenBuffers() {
		offsetIndices := make([]uint16, len(shape.IndicesArray))
		for i, v := range shape.IndicesArray {
			offsetIndices[i] = v + uint16(vCountOffset)
		}

		out.VerticesArray = append(out.VerticesArray, shape.VerticesArray...)
		out.IndicesArray = append(out.IndicesArray, offsetIndices...)
		out.ColourArray = append(out.ColourArray, shape.ColourArray...)

		vCountOffset += shape.VCount
		iCountOffset += shape.ICount
	}
	out.VCount = vCountOffset
	out.ICount = iCountOffset

	return out
}
