package proteintranslation

import (
	"errors"
)

var ErrInvalidBase = errors.New("invalid base")
var ErrStop = errors.New("stop")

var proteinMap = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

func FromRNA(rna string) ([]string, error) {
	proteins := make([]string, 0, len(rna)/3)

	for i := 0; i < len(rna); i += 3 {
		if i+3 > len(rna) {
			return nil, ErrInvalidBase
		}

		protein, err := FromCodon(rna[i : i+3])

		if err != nil {
			if errors.Is(err, ErrStop) {
				return proteins, nil
			}
			return nil, err
		}

		proteins = append(proteins, protein)
	}

	return proteins, nil
}

func FromCodon(codon string) (string, error) {
	protein, ok := proteinMap[codon]

	if !ok {
		return "", ErrInvalidBase
	}

	if protein == "STOP" {
		return "", ErrStop
	}

	return protein, nil
}
