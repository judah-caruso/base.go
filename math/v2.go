/*
V2

Basic Vector2 type with Scalar components. Methods never modify their caller.
*/
package math

import "math"

type V2 struct {
	X Scalar
	Y Scalar
}

// Returns a Vector of [0, 0]
//
// Equivalent to
//
//	zero := V2{}
func V2Zero() V2 { return V2{X: 0, Y: 0} }

// Returns a Vector of [1, 1]
func V2One() V2 { return V2{X: 1, Y: 1} }

// Returns a Vector of [0, -1]
func V2Up() V2 { return V2{X: 0, Y: -1} }

// Returns a Vector of [0, 1]
func V2Down() V2 { return V2{X: 0, Y: 1} }

// Returns a Vector of [-1, 0]
func V2Left() V2 { return V2{X: -1, Y: 0} }

// Returns a Vector of [1, 0]
func V2Right() V2 { return V2{X: 1, Y: 0} }

// Adds 'b' to 'a'
func (a V2) Add(b V2) V2 {
	return V2{
		X: a.X + b.X,
		Y: a.Y + b.Y,
	}
}

// Subtracts 'b' from 'a'
func (a V2) Sub(b V2) V2 {
	return V2{
		X: a.X - b.X,
		Y: a.Y - b.Y,
	}
}

// Multiplies 'a' by 'b'
func (a V2) Mul(b V2) V2 {
	return V2{
		X: a.X * b.X,
		Y: a.Y * b.Y,
	}
}

// Divides 'a' by 'b'.
func (a V2) Div(b V2) (c V2) {
	return V2{
		X: a.X / b.X,
		Y: a.Y / b.Y,
	}
}

// Scale each component of 'a' by 'amt'
func (a V2) Scale(amt Scalar) V2 {
	return V2{
		X: a.X * amt,
		Y: a.Y * amt,
	}
}

// Normalize 'a'
func (a V2) Normalize() (res V2) {
	l := a.Length()
	if l > 0 {
		i := 1.0 / l
		res.X = a.X * i
		res.Y = a.Y * i
	}

	return
}

// Calculate length (or magnitude) of 'a'
func (a V2) Length() Scalar {
	return Scalar(math.Sqrt(float64((a.X * a.X) + (a.Y * a.Y))))
}

// Calculate dot product of 'a' and 'b'
func (a V2) Dot(b V2) Scalar {
	return Scalar((a.X * b.X) + (a.Y * b.Y))
}

// Calculate the distance between 'a' and 'b'
func (a V2) DistanceTo(b V2) Scalar {
	x := (a.X - b.X) * (a.X - b.X)
	y := (a.Y - b.Y) * (a.Y - b.Y)
	return Scalar(math.Sqrt(float64(x + y)))
}

// Calculate angle from 'a' to 'b'
func (a V2) AngleTo(b V2) Scalar {
	return Scalar(
		math.Atan2(float64(b.Y), float64(b.X)) -
			math.Atan2(float64(a.Y), float64(a.X)),
	)
}

// Linear interpolate 'a' to 'b' via 't'.
func (a V2) Lerp(b V2, t Scalar) (c V2) {
	c.X = a.X + t*(b.X-a.X)
	c.Y = a.Y + t*(b.Y-a.Y)
	return
}

// Clamp components of 'a' to be >= 'min' and <= 'max'.
func (a V2) Clamp(min V2, max V2) V2 {
	return V2{
		X: Clamp[Scalar](a.X, min.X, min.Y),
		Y: Clamp[Scalar](a.Y, min.Y, min.Y),
	}
}

// Truncate each component of 'a'.
func (a V2) Trunc() V2 {
	return V2{
		X: Scalar(math.Trunc(a.X)),
		Y: Scalar(math.Trunc(a.Y)),
	}
}

// Check loose equality between 'a' and 'b'.
func (a V2) Equals(b V2) bool {
	ax := float64(a.X)
	ay := float64(a.Y)
	bx := float64(b.X)
	by := float64(b.Y)

	eps := float64(Epsilon)
	xdist := math.Abs(ax - bx)
	ydist := math.Abs(ay - by)
	xeps := eps * math.Max(1, math.Max(math.Abs(ax), math.Abs(bx)))
	yeps := eps * math.Max(1, math.Max(math.Abs(ay), math.Abs(by)))

	return xdist <= xeps && ydist <= yeps
}
