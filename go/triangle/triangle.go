// Package triangle determines if a triangle is equilateral, isosceles, or scalene.
package triangle

import "math"

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

// KindFromSides verifies if the sides make a valid triangle and returns the kind of triangle
func KindFromSides(a, b, c float64) Kind {
	if !isValidTriangle(a, b, c) {
		return NaT
	}

	if a == b && b == c {
		return Equ
	}

	if a == b || b == c || a == c {
		return Iso
	}

	return Sca
}

func isValidTriangle(a, b, c float64) bool {
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return false
	}

	if math.IsInf(a, -1) || math.IsInf(b, -1) || math.IsInf(c, -1) {
		return false
	}

	if math.IsInf(a, 1) || math.IsInf(b, 1) || math.IsInf(c, 1) {
		return false
	}

	if a <= 0 || b <= 0 || c <= 0 {
		return false
	}

	if a+b < c || b+c < a || a+c < b {
		return false
	}

	return true
}
