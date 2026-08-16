package luhn

import "slices"

func Valid(id string) bool {
	digits := make([]int, 0, len(id))
	for _, r := range id {
		if r == ' ' {
			continue
		}
		if r < '0' || r > '9' {
			return false
		}
		digits = append(digits, int(r-'0'))
	}
	if len(digits) < 2 {
		return false
	}

	sum := 0
	for i, d := range slices.Backward(digits) {

		if (len(digits)-1-i)%2 == 1 {
			d *= 2
			if d > 9 {
				d -= 9
			}
		}
		sum += d
	}
	return sum%10 == 0
}
