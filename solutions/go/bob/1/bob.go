package bob

import (
	"strings"
	"unicode"
)

func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	switch {
	case remark == "":
		return "Fine. Be that way!"
	case isQuestion(remark) && isYelling(remark):
		return "Calm down, I know what I'm doing!"
	case isQuestion(remark):
		return "Sure."
	case isYelling(remark):
		return "Whoa, chill out!"
	default:
		return "Whatever."
	}
}

func isQuestion(remark string) bool {
	return strings.HasSuffix(remark, "?")
}

func isYelling(remark string) bool {
	hasLetter := strings.IndexFunc(remark, func(r rune) bool {
		return unicode.IsLetter(r)
	})

	isUppercase := strings.ToUpper(remark) == remark

	return hasLetter != -1 && isUppercase
}
