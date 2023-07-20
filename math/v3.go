/*
V3

Same as V2, with the addition of a Z component.
*/
package math

import "math"

type V3 struct {
	X Scalar
	Y Scalar
	Z Scalar
}

func V3Zero() V3  { return V3{X: 0, Y: 0, Z: 0} }
func V3One() V3   { return V3{X: 1, Y: 1, Z: 1} }
func V3Up() V3    { return V3{X: 0, Y: -1, Z: 0} }
func V3Down() V3  { return V3{X: 0, Y: 1, Z: 0} }
func V3Left() V3  { return V3{X: -1, Y: 0, Z: 0} }
func V3Right() V3 { return V3{X: 1, Y: 0, Z: 0} }

// V3Forward, V3Back are intentionally undefined

func (a V3) Add(b V3) V3 {
	return V3{
		X: a.X + b.X,
		Y: a.Y + b.Y,
		Z: a.Z + b.Z,
	}
}

func (a V3) Sub(b V3) V3 {
	return V3{
		X: a.X - b.X,
		Y: a.Y - b.Y,
		Z: a.Z - b.Z,
	}
}

func (a V3) Mul(b V3) V3 {
	return V3{
		X: a.X * b.X,
		Y: a.Y * b.Y,
		Z: a.Z * b.Z,
	}
}

func (a V3) Div(b V3) (c V3) {
	return V3{
		X: a.X / b.X,
		Y: a.Y / b.Y,
		Z: a.Z / b.Z,
	}
}

func (a V3) Scale(amt Scalar) V3 {
	return V3{
		X: a.X * amt,
		Y: a.Y * amt,
		Z: a.Z * amt,
	}
}

func (a V3) Normalize() (res V3) {
	panic("Todo")
}

func (a V3) Length() Scalar {
	panic("Todo")
}

func (a V3) Dot(b V3) Scalar {
	panic("Todo")
}

func (a V3) DistanceTo(b V3) Scalar {
	panic("Todo")
}

func (a V3) AngleTo(b V3) Scalar {
	panic("Todo")
}

func (a V3) Lerp(b V3, t Scalar) (c V3) {
	c.X = a.X + t*(b.X-a.X)
	c.Y = a.Y + t*(b.Y-a.Y)
	c.Z = a.Z + t*(b.Z-a.Z)
	return
}

func (a V3) Clamp(min V3, max V3) V3 {
	return V3{
		X: Clamp[Scalar](a.X, min.X, max.X),
		Y: Clamp[Scalar](a.Y, min.Y, max.Y),
		Z: Clamp[Scalar](a.Z, min.Z, max.Z),
	}
}

func (a V3) Trunc() V3 {
	return V3{
		X: Scalar(math.Trunc(a.X)),
		Y: Scalar(math.Trunc(a.Y)),
		Z: Scalar(math.Trunc(a.Z)),
	}
}

func (a V3) Equals(b V3) bool {
	ax := float64(a.X)
	ay := float64(a.Y)
	az := float64(a.Z)
	bx := float64(b.X)
	by := float64(b.Y)
	bz := float64(b.Z)

	eps := float64(Epsilon)
	xdist := math.Abs(ax - bx)
	ydist := math.Abs(ay - by)
	zdist := math.Abs(az - bz)
	xeps := eps * math.Max(1, math.Max(math.Abs(ax), math.Abs(bx)))
	yeps := eps * math.Max(1, math.Max(math.Abs(ay), math.Abs(by)))
	zeps := eps * math.Max(1, math.Max(math.Abs(az), math.Abs(bz)))

	return xdist <= xeps &&
		ydist <= yeps &&
		zdist <= zeps
}
