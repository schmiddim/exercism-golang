package strain

// Implement the "Keep" and "Discard" function in this file.

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1

func Keep[T any](cp []T, f func(T) bool) []T {
	var match []T
	for _, v := range cp {
		if f(v) {
			match = append(match, v)
		}
	}
	return match

}

func Discard[T any](cp []T, f func(T) bool) []T {
	var match []T
	for _, v := range cp {
		if f(v) == false {
			match = append(match, v)
		}
	}
	return match
}
