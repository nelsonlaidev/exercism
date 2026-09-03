package rnatranscription

import "strings"

func ToRNA(dna string) string {
	result := make([]string, 0, len(dna))

	for _, r := range dna {
		str := string(r)

		switch str {
		case "C":
			result = append(result, "G")
		case "G":
			result = append(result, "C")
		case "T":
			result = append(result, "A")
		case "A":
			result = append(result, "U")
		default:
			return ""
		}
	}

	return strings.Join(result, "")
}
