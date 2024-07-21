package rubiks

import "github.com/AislingHeanue/aisling-codes/wasm-demo/maths"

type DrawShape struct {
	Points        []*maths.Point
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

func GetBuffers(c maths.Cube) DrawShape {
	var out DrawShape
	indexOfVerticesArray := []uint16{
		4, 5, 6, 7, // WHITE
		3, 2, 6, 7, // ORANGE
		0, 3, 7, 4, // GREEN
		1, 0, 4, 5, // RED
		1, 2, 6, 5, // BLUE
		0, 1, 2, 3, // YELLOW
	}

	points := []*maths.Point{
		c.CentrePoint.Add(maths.Point{-c.Side / 2, -c.Side / 2, -c.Side / 2}), // GREEN YELLOW RED
		c.CentrePoint.Add(maths.Point{c.Side / 2, -c.Side / 2, -c.Side / 2}),  // BLUE YELLOW RED
		c.CentrePoint.Add(maths.Point{c.Side / 2, -c.Side / 2, c.Side / 2}),   // BLUE YELLOW ORANGE
		c.CentrePoint.Add(maths.Point{-c.Side / 2, -c.Side / 2, c.Side / 2}),  // GREEN YELLOW ORANGE
		c.CentrePoint.Add(maths.Point{-c.Side / 2, c.Side / 2, -c.Side / 2}),  // GREEN WHITE RED
		c.CentrePoint.Add(maths.Point{c.Side / 2, c.Side / 2, -c.Side / 2}),   // BLUE WHITE RED
		c.CentrePoint.Add(maths.Point{c.Side / 2, c.Side / 2, c.Side / 2}),    // BLUE WHITE ORANGE
		c.CentrePoint.Add(maths.Point{-c.Side / 2, c.Side / 2, c.Side / 2}),   // GREEN WHITE ORANGE
	}
	// for i, p := range points {
	// 	points[i] = p.Rotate(c.CentrePoint, c.AngleY, Y)
	// }

	out.VerticesArray = make([]float32, 72)
	for i, index := range indexOfVerticesArray {
		pointSlice := points[index].ToSlice()
		out.VerticesArray[3*i] = pointSlice[0]
		out.VerticesArray[3*i+1] = pointSlice[1]
		out.VerticesArray[3*i+2] = pointSlice[2]
	}

	out.IndicesArray = make([]uint16, 36)
	for j := 0; j < 6; j++ {
		// assume points are connected as 0->1->2->3
		// then we need 0,1,2,0,2,3
		out.IndicesArray[6*j] = uint16(4*j + 0)
		out.IndicesArray[6*j+1] = uint16(4*j + 1)
		out.IndicesArray[6*j+2] = uint16(4*j + 2)
		out.IndicesArray[6*j+3] = uint16(4*j + 0)
		out.IndicesArray[6*j+4] = uint16(4*j + 2)
		out.IndicesArray[6*j+5] = uint16(4*j + 3)

	}
	// fmt.Println(out.IndicesArray)

	// out.IndicesArray = []uint16{
	// 	7., 6., 5., 7., 5., 4., // WHITE   +0
	// 	14, 15, 11, 14, 11, 10, // ORANGE  +8
	// 	23, 20, 16, 23, 16, 19, // GREEN   +16
	// 	12, 13, 9., 12, 9., 8., // RED     +8
	// 	17, 21, 22, 17, 22, 18, // BLUE    +16
	// 	0., 1., 2., 0., 2., 3., // YELLOW  +0
	// }

	outColours := []float32{}
	for _, c := range c.Colours {
		outColours = append(outColours, float32(c.R)/256, float32(c.G)/256, float32(c.B)/256, float32(c.A)/256)
		outColours = append(outColours, float32(c.R)/256, float32(c.G)/256, float32(c.B)/256, float32(c.A)/256)
		outColours = append(outColours, float32(c.R)/256, float32(c.G)/256, float32(c.B)/256, float32(c.A)/256)
		outColours = append(outColours, float32(c.R)/256, float32(c.G)/256, float32(c.B)/256, float32(c.A)/256)
	}
	out.ColourArray = outColours
	// fmt.Println(out.ColoursArray)

	// fmt.Println(len(c.Points))
	// fmt.Println(len(out.VerticesArray) / 3)
	// fmt.Println(len(out.IndicesArray))
	// fmt.Println(len(out.ColourArray) / 4)

	out.VCount = 24
	out.ICount = 36
	out.CCount = 24

	return out
}
