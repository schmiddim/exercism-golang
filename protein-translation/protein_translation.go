package protein

import (
	"errors"
)

var ErrInvalidBase = errors.New("ErrInvalidBase")
var ErrStop = errors.New("ErrStop")

func FromRNA(rna string) ([]string, error) {
	ret := []string{}
	for i := 0; i+2 < len(rna); i += 3 {
		codon := rna[i : i+3]
		protein, err := FromCodon(codon)

		if errors.Is(err, ErrStop) {
			break
		}
		if errors.Is(err, ErrInvalidBase) {
			return []string{}, err

		}
		ret = append(ret, protein)

	}
	return ret, nil
}

func FromCodon(codon string) (string, error) {

	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU":
		return "Phenylalanine", nil
	case "UUC":
		return "Phenylalanine", nil
	case "UUA":
		return "Leucine", nil
	case "UUG":
		return "Leucine", nil

	case "UCU":
		return "Serine", nil
	case "UCC":
		return "Serine", nil
	case "UCA":
		return "Serine", nil
	case "UCG":
		return "Serine", nil
	case "UAU":
		return "Tyrosine", nil
	case "UAC":
		return "Tyrosine", nil
	case "UGU":
		return "Cysteine", nil
	case "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil

	case "UAA":
		return "STOP", ErrStop
	case "UAG":
		return "STOP", ErrStop
	case "UGA":
		return "STOP", ErrStop
	default:
		return "invalid", ErrInvalidBase
	}

}
