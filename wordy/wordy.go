package wordy

import (
	"fmt"
	"regexp"
	"strconv"
)

type Item struct {
	isNumber  bool
	number    int
	operation string
}

func NewItem(str string) *Item {
	isNumber := true
	operation := ""
	n, parseError := strconv.Atoi(str)

	if parseError != nil {
		isNumber = false
		// parse for an operation
		re := regexp.MustCompile(`plus|minus|multiplied by|divided by`)
		result := re.FindAllString(str, 1)
		if len(result) > 0 {
			isNumber = false
			operation = result[0]
		}
	}
	return &Item{
		isNumber:  isNumber,
		number:    n,
		operation: operation,
	}
}

func (i *Item) String() string {
	if i.isNumber {
		return fmt.Sprintf("%d", i.number)
	}
	return fmt.Sprintf("%s", i.operation)
}

type Solver struct {
	items []*Item
}

func NewSolver() *Solver {

	return &Solver{}
}

func (s *Solver) Append(str string) {
	s.items = append(s.items, NewItem(str))

}

func (s *Solver) String() string {
	return fmt.Sprintf("%v", s.items)
}

func (s *Solver) isValid() bool {
	if s.items == nil {
		return false
	}

	numOps := 0
	numNums := 0
	prevItem := s.items[0].isNumber
	for i, item := range s.items {

		//reject Postfix Notation
		if i > 0 {
			if prevItem == item.isNumber {
				return false
			}
			prevItem = item.isNumber
		}
		if item.isNumber {
			numNums++
		} else {
			numOps++
		}
	}
	if numOps != numNums-1 {
		return false
	}
	return true

}

func calc(a int, b int, op string) int {
	switch {
	case op == "plus":
		return a + b

	case op == "minus":
		return a - b
	case op == "divided by":
		return a / b
	case op == "multiplied by":
		return a * b
	default:
		return 0
	}
}

func (s *Solver) getFirstDashOp() int {
	for i, item := range s.items {
		if item.isNumber == false {
			if item.operation == "plus" || item.operation == "minus" {
				return i
			}
		}
	}
	return -1
}
func (s *Solver) getFirstPointOp() int {
	for i, item := range s.items {

		if item.isNumber == false {
			if item.operation == "multiplied by" || item.operation == "divided by" {
				return i
			}
		}
	}
	return -1
}

func (s *Solver) Calculate() (int, bool) {
	// just a number
	if len(s.items) == 1 && s.items[0].isNumber {
		return s.items[0].number, true
	}
	if s.isValid() == false {
		return 0, false
	}

	result := 0

	firstDashOp := s.getFirstDashOp()
	for firstDashOp > 0 {
		result = calc(s.items[firstDashOp-1].number, s.items[firstDashOp+1].number, s.items[firstDashOp].operation)
		s.items[firstDashOp-1].number = result
		s.items = append(s.items[:firstDashOp], s.items[firstDashOp+2:]...)
		firstDashOp = s.getFirstDashOp()
	}
	firstPointIndex := s.getFirstPointOp()
	for firstPointIndex > 0 {
		result = calc(s.items[firstPointIndex-1].number, s.items[firstPointIndex+1].number, s.items[firstPointIndex].operation)
		s.items[firstPointIndex-1].number = result
		s.items = append(s.items[:firstPointIndex], s.items[firstPointIndex+2:]...)
		firstPointIndex = s.getFirstPointOp()
	}

	return result, true
}

func Answer(question string) (int, bool) {
	re := regexp.MustCompile(`(?i)-?\d+|plus|minus|multiplied by|divided by`)

	cubeRe := regexp.MustCompile("cubed")
	if cubeRe.MatchString(question) {
		return 0, false
	}

	matchesNumbers := re.FindAllStringSubmatch(question, -1)
	solver := NewSolver()
	for _, n := range matchesNumbers {
		solver.Append(n[0])

	}
	fmt.Println(solver)
	return solver.Calculate()
}
