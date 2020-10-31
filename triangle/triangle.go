package triangle

import "math"

// Kind is a triangle type
type Kind string

const (
	// NaT is not a triangle
	NaT = "Nat"

	//Equ is equilateral
	Equ = "Equ"

	// Iso is isoceles
	Iso = "Iso"

	// Sca is scaleno
	Sca = "Sca"
)

// KindFromSides checks the type of a valid triangle
func KindFromSides(a, b, c float64) Kind {
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) || math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return NaT
	}

	if a <= 0 || b <= 0 || c <= 0 || a+b < c || a+c < b || c+b < a {
		return NaT
	}

	if a == b && a == c {
		return Equ
	}

	if (a == b && b != c) || (a == c && c != b) || (b == c && c != a) {
		return Iso
	}

	return Sca
}
