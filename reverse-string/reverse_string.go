package reverse

func Reverse(input string) string {

	runes := []rune(input)
	var out []rune
	for i := len(runes) - 1; i >= 0; i-- {

		out = append(out, runes[i])
	}
	return string(out)

}
