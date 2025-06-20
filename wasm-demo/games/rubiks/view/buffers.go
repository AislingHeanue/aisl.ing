package view

import "github.com/AislingHeanue/aisling-codes/wasm-demo/games/rubiks/model"

type DrawShape struct {
	Points        []*model.Point
	VerticesArray []float32
	IndicesArray  []uint16
	ColourArray   []float32
	VCount        int
	ICount        int
	CCount        int
}

func GroupBuffers(s []DrawShape) DrawShape {
	var out DrawShape
	out.VerticesArray = []float32{}
	out.IndicesArray = []uint16{}
	out.ColourArray = []float32{}
	vCountOffset := 0
	iCountOffset := 0
	for _, shape := range s {
		out.VerticesArray = append(out.VerticesArray, shape.VerticesArray...)

		for _, i := range shape.IndicesArray {
			out.IndicesArray = append(out.IndicesArray, i+uint16(vCountOffset))
		}
		out.ColourArray = append(out.ColourArray, shape.ColourArray...)

		vCountOffset += shape.VCount
		iCountOffset += shape.ICount
	}
	out.VCount = vCountOffset
	out.ICount = iCountOffset

	return out
}

func GetBuffers(c model.Cube) DrawShape {
	var out DrawShape

	out.VerticesArray = make([]float32, 72)
	for i, index := range c.VertexArrayIndices {
		pointSlice := c.Points[index].ToSlice()
		out.VerticesArray[3*i] = pointSlice[0]
		out.VerticesArray[3*i+1] = pointSlice[1]
		out.VerticesArray[3*i+2] = pointSlice[2]
	}

	out.IndicesArray = make([]uint16, 36)
	for j := range 6 {
		// assume points are connected as 0->1->2->3
		// then we need 0,1,2,0,2,3
		out.IndicesArray[6*j] = uint16(4*j + 0)
		out.IndicesArray[6*j+1] = uint16(4*j + 1)
		out.IndicesArray[6*j+2] = uint16(4*j + 2)
		out.IndicesArray[6*j+3] = uint16(4*j + 0)
		out.IndicesArray[6*j+4] = uint16(4*j + 2)
		out.IndicesArray[6*j+5] = uint16(4*j + 3)

	}

	outColours := []float32{}
	for _, c := range c.Colours {
		outColours = append(outColours, float32(c.R)/256, float32(c.G)/256, float32(c.B)/256, float32(c.A)/256)
		outColours = append(outColours, float32(c.R)/256, float32(c.G)/256, float32(c.B)/256, float32(c.A)/256)
		outColours = append(outColours, float32(c.R)/256, float32(c.G)/256, float32(c.B)/256, float32(c.A)/256)
		outColours = append(outColours, float32(c.R)/256, float32(c.G)/256, float32(c.B)/256, float32(c.A)/256)
	}
	out.ColourArray = outColours

	out.VCount = 24
	out.ICount = 36
	out.CCount = 24

	return out
}

func GetBuffersForRubiksAnimator(a *model.RubiksAnimationHandler, origin *model.Point) DrawShape {
	a.CopyRubiksCube = a.RubiksCube.Copy()
	for _, event := range a.EventsWhichNeedToBeRotated {
		a.DoEvent(event, origin)
	}
	return GroupBuffersForRubiksCube(a.CopyRubiksCube)
}

func GroupBuffersForRubiksCube(r model.RubiksCube) DrawShape {
	d := []DrawShape{}
	for _, v := range r.Flatten() {
		d = append(d, GetBuffers(v))
	}

	return GroupBuffers(d)
}
