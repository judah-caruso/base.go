/*
# Math

Math types and utility functions.
*/
package math

const (
	Epsilon Scalar = 0.000001
)

type (
	// Underlying type for all math data structures.
	// Allows this package to use float32 or integer types if wanted.
	Scalar = float64

	numeric interface {
		~int | ~float32 | ~float64
	}
)

// @note: temporary until Go's builtin versions are released
func Min[T numeric](v, min T) T {
	if v < min {
		return min
	}
	return v
}

func Max[T numeric](v, max T) T {
	if v > max {
		return max
	}
	return v
}

func Clamp[T numeric](v, min, max T) T {
	return Min(max, Max(min, v))
}

// Methods available for V2, V3, and V4. Only here for documentation.
type vec interface {
	Add(vec) vec
	Sub(vec) vec
	Mul(vec) vec
	Div(vec) vec
	Scale(Scalar) vec
	Normalize() vec
	Length() Scalar
	Dot(vec) Scalar
	DistanceTo(vec) Scalar
	AngleTo(vec) Scalar
	Lerp(vec, Scalar) vec
	Clamp(vec, vec) vec
	Trunc() vec
	Equals(vec) bool
}
