package animation

import (
	"image/color"
	"math"
	"slices"

	"github.com/AislingHeanue/aisling-codes/wasm-demo/model"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/draw2dkit"
)

func CirclingCircle(gc *draw2dimg.GraphicContext, c model.GameContext) bool {
	secs := c.T.Seconds()

	// fill is white
	gc.SetFillColor(color.White)
	// fill canvas with fill colour
	gc.Clear()

	// stroke and fill are magenta
	gc.SetFillColor(c.Colour)
	gc.SetStrokeColor(c.Colour)
	// draw a circle
	gc.BeginPath()
	draw2dkit.Circle(gc, c.Width/2*(1+math.Cos(secs)), c.Height/2*(1+math.Sin(secs)), min(c.Width/3, c.Height/3))
	gc.FillStroke()
	gc.Close()

	return true
}

type CubeAnimator struct{}

func (cc *CubeAnimator) Init(c *model.GameContext) {
	side := c.Width / 2
	cubeOrigin := model.Vector{c.Width / 2, c.Width / 2, 0}
	c.Cube = model.NewCube(cubeOrigin, side)
}

func (cc *CubeAnimator) Render(gc *draw2dimg.GraphicContext, c *model.GameContext) bool {

	// fill is white
	gc.SetFillColor(color.White)
	gc.SetStrokeColor(c.Colour)

	// fill canvas with white
	gc.Clear()

	// calculate position of points
	for _, p := range c.Cube.Points {
		p.RenderedCoordinates = c.Projector.GetCoords(*p.Vector, c.Height, c.Width, c.AngleX, c.AngleY, *c.Cube.GetCentre())
	}

	// cull faces whose normal have a positive z component

	// render the points
	// gc.BeginPath()
	// rect := image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{int(c.Width), int(c.Height)}}
	// im := image.NewRGBA(rect)
	for _, f := range c.Cube.Faces {
		if f.GetRenderedNormal()[2] > 0 {
			continue
		}
		// controller.RasteriseSquare(f, im)
		gc.MoveTo(f.Vertices[3].RenderedCoordinates[0], f.Vertices[3].RenderedCoordinates[1])
		for _, p := range f.Vertices {
			gc.LineTo(p.RenderedCoordinates[0], p.RenderedCoordinates[1])
		}
	}
	// gc.DrawImage(im)
	// lines := c.Cube.Lines
	// for _, line := range lines {
	// 	gc.MoveTo(line.Origin.RenderedCoordinates[0], line.Origin.RenderedCoordinates[1])
	// gc.LineTo(line.End.RenderedCoordinates[0], line.End.RenderedCoordinates[1])
	// }

	gc.FillStroke()
	// gc.BeginPath()
	// a := 0
	// for x := 0.; x < c.Width; x++ {
	// 	for y := 0.; y < c.Height; y++ {
	// 		// if c.Cube.Faces[3].CheckContains(x, y) {
	// 		if c.Cube.Faces[3].CheckContains(x, y) {
	// 			a += 1
	// 		}
	// 	}
	// // }
	// gc.Stroke()

	gc.Close()

	return true
}

type CubeCube struct {
	Cubes  [][][]*model.Cube
	Sg     model.ShapeGroup
	Origin *model.Vector
	active bool
}

func (cc *CubeCube) IsActive() bool {
	return cc.active
}

func (cc *CubeCube) Init(c *model.GameContext) {
	cc.active = true
	totalSide := c.Width / 2
	gap := 0.05
	side := totalSide / ((1+gap)*float64(c.Dimension) - gap)
	sideFull := side + gap*side
	cc.Origin = &model.Vector{c.Width / 2, c.Width / 2, 0}
	cc.Cubes = make([][][]*model.Cube, c.Dimension)

	shapes := []*model.Shape{}
	for x := 0; x < c.Dimension; x++ {
		cc.Cubes[x] = make([][]*model.Cube, c.Dimension)
		for y := 0; y < c.Dimension; y++ {
			cc.Cubes[x][y] = make([]*model.Cube, c.Dimension)
			for z := 0; z < c.Dimension; z++ {
				cubeOrigin := getCentre(x, y, z, side, sideFull, totalSide, cc.Origin)
				cc.Cubes[x][y][z] = model.NewCubeWithColours(*cubeOrigin, side, CubeColours(x, y, z, c.Dimension))
				shapes = append(shapes, (*model.Shape)(cc.Cubes[x][y][z]))
			}
		}
	}
	cc.Sg = model.NewShapeGroup(shapes)
}

func getCentre(x, y, z int, side, sideFull, totalSide float64, origin *model.Vector) *model.Vector {
	return origin.Subtract(model.Vector{totalSide / 2, totalSide / 2, totalSide / 2}).Add(model.Vector{sideFull * float64(x), sideFull * float64(y), sideFull * float64(z)}).Add(model.Vector{side / 2, side / 2, side / 2})
}

func CubeColours(x, y, z, dimension int) []color.Color {
	// White Orange Green Red Blue Yellow
	colours := []color.Color{model.BLACK, model.BLACK, model.BLACK, model.BLACK, model.BLACK, model.BLACK}
	if y == dimension-1 {
		colours[0] = model.DefaultColours[0]
	}
	if z == dimension-1 {
		colours[1] = model.DefaultColours[1]
	}
	if x == 0 {
		colours[2] = model.DefaultColours[2]
	}
	if z == 0 {
		colours[3] = model.DefaultColours[3]
	}
	if x == dimension-1 {
		colours[4] = model.DefaultColours[4]
	}
	if y == 0 {
		colours[5] = model.DefaultColours[5]
	}

	return colours
}

func (cc *CubeCube) Render(gc *draw2dimg.GraphicContext, c *model.GameContext) bool {

	// fill is white
	gc.SetFillColor(color.White)
	gc.SetStrokeColor(color.Black)
	gc.SetLineWidth(2)

	// fill canvas with white
	gc.Clear()

	normalOrigin := cc.Origin
	sg := cc.Sg

	// calculate position of points
	for _, s := range sg.Shapes {
		for _, p := range s.Points {
			p.RenderedCoordinates = c.Projector.GetCoords(*p.Vector, c.Height, c.Width, c.AngleX, c.AngleY, *normalOrigin)
		}
		for _, f := range s.Faces {
			f.Centre.RenderedCoordinates = c.Projector.GetCoords(*f.Centre.Vector, c.Height, c.Width, c.AngleX, c.AngleY, *normalOrigin)
		}
	}

	sortedFaces := []*model.Face{}
	for _, f := range sg.Faces {
		if f.GetRenderedNormal()[2] <= 0 {
			sortedFaces = append(sortedFaces, f)
		}
	}

	slices.SortFunc(sortedFaces, func(s1, s2 *model.Face) int {
		return int(s2.Centre.RenderedCoordinates[2] - s1.Centre.RenderedCoordinates[2])
	})

	gc.BeginPath()
	for _, f := range sortedFaces {
		gc.SetFillColor(f.Colour)
		gc.MoveTo((*f).Vertices[3].RenderedCoordinates[0], (*f).Vertices[3].RenderedCoordinates[1])
		for _, p := range f.Vertices {
			gc.LineTo(p.RenderedCoordinates[0], p.RenderedCoordinates[1])
		}
		gc.FillStroke()
	}

	// gc.BeginPath()
	// p := model.Point{Vector: cc.Origin}
	// centreCoords := c.Projector.GetCoords(*p.Vector, c.Height, c.Width, c.AngleX, c.AngleY, *normalOrigin)
	// gc.MoveTo(centreCoords[0], centreCoords[1])
	// draw2dkit.Circle(gc, centreCoords[0], centreCoords[1], 5)
	// gc.FillStroke()

	gc.Close()

	return true
}
