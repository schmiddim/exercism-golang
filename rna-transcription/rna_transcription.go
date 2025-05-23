package strand

func ToRNA(dna string) string {
	ret := ""
	for _, v := range dna {

		switch {
		case v == 'G':
			ret += "C"

		case v == 'C':
			ret += "G"
		case v == 'T':
			ret += "A"
		case v == 'A':
			ret += "U"

		}

	}
	return ret
}
