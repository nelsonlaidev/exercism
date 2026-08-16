package darts

import "math"

func Score(x, y float64) int {
	score := 0
	distance := math.Sqrt(math.Pow((0-x), 2) + math.Pow((0-y), 2))

	if distance <= 1 {
		score = 10
	} else if distance <= 5 {
		score = 5
	} else if distance <= 10 {
		score = 1
	}

	return score
}
