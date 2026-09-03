package pangram

import "strings"

func IsPangram(input string) bool {
	alphabets := map[rune]int{
		'a': 0,
		'b': 0,
		'c': 0,
		'd': 0,
		'e': 0,
		'f': 0,
		'g': 0,
		'h': 0,
		'i': 0,
		'j': 0,
		'k': 0,
		'l': 0,
		'm': 0,
		'n': 0,
		'o': 0,
		'p': 0,
		'q': 0,
		'r': 0,
		's': 0,
		't': 0,
		'u': 0,
		'v': 0,
		'w': 0,
		'x': 0,
		'y': 0,
		'z': 0,
	}

	for _, r := range strings.ToLower(input) {
		if r >= 'a' && r <= 'z' {
			alphabets[r]++
		}
	}

	for _, a := range alphabets {
		if a == 0 {
			return false
		}
	}

	return true
}
