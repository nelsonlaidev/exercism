package scrabblescore

import (
	"strings"
)

func Score(word string) int {
	chars := strings.Split(word, "")
	total := 0

	for _, c := range chars {
		total += getScore(strings.ToLower(c))
	}

	return total
}

func getScore(character string) int {
	switch character {
	case "a", "e", "i", "o", "u", "l", "n", "r", "s", "t":
		return 1
	case "d", "g":
		return 2
	case "b", "c", "m", "p":
		return 3
	case "f", "h", "v", "w", "y":
		return 4
	case "k":
		return 5
	case "j", "x":
		return 8
	case "q", "z":
		return 10
	default:
		return 0
	}
}
