package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %d.0", nb.Number())
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {

	switch fnb.(type) {
	case FancyNumber:
		val, _ := strconv.Atoi(fnb.Value())

		return val
	}

	return 0

}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {

	switch fnb.(type) {
	case FancyNumber:
		return fmt.Sprintf("This is a fancy box containing the number %s.0", fnb.Value())

	}
	return fmt.Sprintf("This is a fancy box containing the number %d.0", 0)
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {

	switch v := i.(type) {
	case int:
		return DescribeNumber(float64(v))
	case float64:
		return DescribeNumber(v)
	case FancyNumberBox:
		return DescribeFancyNumberBox(v)
	case NumberBox:
		return DescribeNumberBox(v)
	case string:
		return "Return to sender"
	default:
		return "Return to sender"
	}

}
