package atbashcipher

import (
	"strings"
	"unicode"
)

func Atbash(s string) string {
	result := make([]rune, 0, len(s))
	count := 0

	for _, r := range strings.ToLower(s) {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			continue
		}

		if count > 0 && count%5 == 0 {
			result = append(result, ' ')
		}

		if r >= 'a' && r <= 'z' {
			r = 'z' - (r - 'a')
		}

		result = append(result, r)

		count++
	}

	return string(result)
}
