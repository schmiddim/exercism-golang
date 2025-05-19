package kindergarten

import (
	"errors"
	"slices"
	"sort"
	"strings"
)

func hasMissMatchedRows(ss []string) bool {
	l1 := -1

	for _, v := range ss {
		if l1 == -1 {
			l1 = len(v)
			continue
		}
		if len(v) != l1 {
			return true
		}
	}

	return false
}

func hasDupes(s []string) bool {

	seen := make(map[string]struct{})
	for _, v := range s {
		if _, ok := seen[v]; ok {
			return true
		}
		seen[v] = struct{}{}
	}
	return false
}

func anyOdd(ss []string) bool {
	for _, s := range ss {
		if len(s)&1 == 1 {
			return true
		}
	}
	return false
}
func getPlantName(abbr string) string {
	switch abbr {
	case "G":
		return "grass"
	case "C":
		return "clover"
	case "R":
		return "radishes"
	case "V":
		return "violets"
	default:
		return "unknown"
	}
}

// Define the Garden type here.

type Garden struct {
	children  []string
	diagram   string
	gardenMap []string
}

// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//
//     diagram := `
//     VVCCGG
//     VVCCGG`

func NewGarden(diagram string, children []string) (*Garden, error) {
	//var diag []string

	rAll := strings.Split(diagram, "\n")
	r := rAll
	if len(rAll) > 0 && rAll[0] == "" {
		r = rAll[1:]
	}
	c := slices.Clone(children)
	sort.Strings(c)

	if hasDupes(children) {
		return nil, errors.New("duplicate Name")
	}

	if anyOdd(r) {
		return nil, errors.New("odd number of cups")
	}

	if hasMissMatchedRows(r) {
		return nil, errors.New("missmatched rows")
	}
	for _, row := range r {
		for _, char := range row {
			if getPlantName(string(char)) == "unknown" {
				return nil, errors.New("unknown Plant")
			}

		}
	}

	if len(rAll)-1 != len(r) {
		return nil, errors.New("invalid Diagram Format")
	}

	return &Garden{children: c, diagram: diagram, gardenMap: r}, nil
}

func (g *Garden) getIndicesForChild(name string) (int, int) {

	for i, candidate := range g.children {
		if candidate == name {
			countBefore := 0
			if i > 0 {
				countBefore = 2 * i
			}
			a := countBefore
			b := countBefore + 1
			return a, b
		}
	}
	return -1, -1

}

func (g *Garden) Plants(child string) ([]string, bool) {

	pos1, pos2 := g.getIndicesForChild(child)
	//kindergarten_garden_test.go:155: Lookup Patricia = ["clover" "radishes" "violets" "grass"], want: ["violets" "clover" "radishes" "violets"]
	//kindergarten_garden_test.go:155: Lookup Roger = ["clover" "radishes" "violets" "grass"], want: ["radishes" "radishes" "grass" "clover"]
	//kindergarten_garden_test.go:155: Lookup Patricia = ["clover" "radishes" "violets" "grass"], want: ["violets" "clover" "radishes" "violets"]

	if pos1 == -1 || pos2 == -1 {

		return nil, false
	}

	var result []string
	for _, row := range g.gardenMap {
		if len(row) < pos1 || len(row) < pos2 {
			return nil, false
		}

		result = append(result, getPlantName(string(row[pos1])))
		result = append(result, getPlantName(string(row[pos2])))
		//["Radish" "Clover" "Grass" "Grass"], want: ["radishes" "clover" "grass" "grass"]
	}

	return result, true
}
