package triangle

type Kind int

const (
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {

	if a <= 0 || b <= 0 || c <= 0 {
		return NaT
	}

	if a+b >= c == false || b+c >= a == false || a+c >= b == false {
		return NaT
	}

	if a == b && b == c {
		return Equ
	}

	if a == b || b == c || a == c {
		return Iso
	}

	if a != b && b != c && a != c {
		return Sca
	}
	var k Kind
	return k
}
