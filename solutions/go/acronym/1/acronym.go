package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	letters := make([]string, 0)

	re := regexp.MustCompile(`[a-zA-Z]+('[a-zA-Z]+)?`)

	for _, w := range re.FindAllString(s, -1) {
		letters = append(letters, strings.ToUpper(string(w[0])))
	}

	return strings.Join(letters, "")
}
