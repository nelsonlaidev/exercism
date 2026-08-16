package isbnverifier

import "strings"

func IsValidISBN(isbn string) bool {
	sum := 0
	cleaned := strings.ReplaceAll(isbn, "-", "")

	if len(cleaned) != 10 {
		return false
	}

	for i, r := range cleaned {
		if r == 'X' && i == 9 {
			sum += 10
			continue
		}

		if r < '0' || r > '9' {
			return false
		}

		sum += int(r-'0') * (10 - i)
	}

	return sum%11 == 0
}
