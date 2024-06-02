package model

// creates a new cube at the specified origin in the default rotation.
// The 'origin' is the cubes lowest x, y and z values.
/*
                        . 10.
                       11  9
                      . 8 .

                       6   7
  7 - 6    Y    Z
 /   /|    ^  /       4   5
4 - 5 |    | /
| 3 | 2    . --> X      . 2 .
|   |/                 3   1
0 - 1                 . 0 .
*/
type Cube Shape

func NewCube(origin Vector, side float64) *Cube {
	var out Cube
	// originCopy := origin
	out.Points = []*Vector{
		origin.Add(Vector{-side / 2, -side / 2, -side / 2}),
		origin.Add(Vector{side / 2, -side / 2, -side / 2}),
		origin.Add(Vector{side / 2, -side / 2, side / 2}),
		origin.Add(Vector{-side / 2, -side / 2, side / 2}),
		origin.Add(Vector{-side / 2, side / 2, -side / 2}),
		origin.Add(Vector{side / 2, side / 2, -side / 2}),
		origin.Add(Vector{side / 2, side / 2, side / 2}),
		origin.Add(Vector{-side / 2, side / 2, side / 2}),
	}
	out.Lines = []*Line{
		{out.Points[0], out.Points[1]},
		{out.Points[1], out.Points[2]},
		{out.Points[2], out.Points[3]},
		{out.Points[3], out.Points[0]},

		{out.Points[0], out.Points[4]},
		{out.Points[1], out.Points[5]},
		{out.Points[2], out.Points[6]},
		{out.Points[3], out.Points[7]},

		{out.Points[4], out.Points[5]},
		{out.Points[5], out.Points[6]},
		{out.Points[6], out.Points[7]},
		{out.Points[7], out.Points[4]},

		// {out.Points[0], out.Points[8]},
	}
	out.Faces = []*Face{
		FaceFromLines(out.Lines[0], out.Lines[1], out.Lines[2], out.Lines[3]),
		FaceFromLines(out.Lines[8], out.Lines[9], out.Lines[10], out.Lines[11]),
		FaceFromLines(out.Lines[3], out.Lines[4], out.Lines[11], out.Lines[6]),
		FaceFromLines(out.Lines[8], out.Lines[5], out.Lines[0], out.Lines[4]),
		FaceFromLines(out.Lines[9], out.Lines[7], out.Lines[1], out.Lines[5]),
		FaceFromLines(out.Lines[2], out.Lines[7], out.Lines[10], out.Lines[6]),
	}

	return &out
}

func (c Cube) GetCentre() *Vector {
	return c.Points[0].Add(*c.Points[6]).Scale(0.5)
}

func (c *Cube) Rotate(anchor Vector, angle float64, axis Axis) {
	shape := Shape(*c)
	(&shape).Rotate(anchor, angle, axis)
	*c = Cube(shape)
}
