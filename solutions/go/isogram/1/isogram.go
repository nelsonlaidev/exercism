package isogram

import "strings"

func IsIsogram(word string) bool {
	frequency := map[rune]int{}

	for _, r := range strings.ToLower(word) {
		if r >= 'a' && r <= 'z' {
			frequency[r]++
		}
	}

	for _, f := range frequency {
		if f > 1 {
			return false
		}
	}

	return true
}
