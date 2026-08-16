package isbnverifier

import "strings"

func IsValidISBN(isbn string) bool {
	sum := 0
	cleaned := strings.ReplaceAll(isbn, "-", "")

	if len(cleaned) != 10 {
		return false
	}

	for i, r := range cleaned {
		if (r < '0' || r > '9') && !(r == 'X' && i == 9) {
			return false
		}

		digit := int(r - '0')

		if digit == 40 {
			sum += 10 * (10 - i)
		} else {
			sum += digit * (10 - i)
		}
	}

	return sum%11 == 0
}
