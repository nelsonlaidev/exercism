package bafflingbirthdays

import (
	"fmt"
	"math/rand"
	"time"
)

func SharedBirthday(dates []time.Time) bool {
	seen := make(map[string]bool)

	for _, d := range dates {
		key := fmt.Sprintf("%v-%v", d.Month(), d.Day())

		if seen[key] {
			return true
		} else {
			seen[key] = true
		}
	}

	return false
}

func RandomBirthdates(size int) []time.Time {
	start := time.Date(2001, time.January, 1, 0, 0, 0, 0, time.UTC)
	dates := make([]time.Time, 0, size)

	for range size {
		randomNum := rand.Intn(365)
		dates = append(dates, start.AddDate(0, 0, randomNum))
	}

	return dates
}

func EstimatedProbability(size int) float64 {
	experiments := 10000
	sharedCount := 0

	for range experiments {
		dates := RandomBirthdates(size)

		if SharedBirthday(dates) {
			sharedCount += 1
		}
	}

	probability := float64(sharedCount) / float64(experiments) * 100

	return probability
}
