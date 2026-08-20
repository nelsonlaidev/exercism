package sumofmultiples

import (
	"slices"
)

func SumMultiples(limit int, divisors ...int) int {
	combined := []int{}
	sum := 0

	for _, d := range divisors {
		if d == 0 {
			combined = append(combined, 0)
			continue
		}

		count := 1
		done := false

		for done != true {
			if d*count < limit {
				combined = append(combined, d*count)
				count += 1
			} else {
				done = true
			}
		}
	}

	slices.Sort(combined)
	combined = slices.Compact(combined)

	for _, n := range combined {
		sum += n
	}

	return sum
}
