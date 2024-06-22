package maths

type DrawShape struct {
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
	// fmt.Println("sup")
	for _, shape := range s {
		// fmt.Println("yup")
		// out.Shapes = append(out.Shapes, shape)
		// out.Points = append(out.Points, shape.Points...)

		out.VerticesArray = append(out.VerticesArray, shape.VerticesArray...)

		for _, i := range shape.IndicesArray {
			out.IndicesArray = append(out.IndicesArray, i+uint16(vCountOffset))
		}
		out.ColourArray = append(out.ColourArray, shape.ColourArray...)

		vCountOffset += shape.VCount
		iCountOffset += shape.ICount
	}
	// fmt.Println(len(out.VerticesArray) / 3 / 26)
	// fmt.Println(len(out.IndicesArray) / 26)
	// fmt.Println(len(out.ColourArray) / 4 / 26)
	out.VCount = vCountOffset
	out.ICount = iCountOffset

	return out
}
