package bottlesong

import (
	"fmt"
	"strings"
)

func Recite(startBottles, takeDown int) []string {
	results := make([]string, 0)
	bottles := startBottles

	for i := range takeDown {
		results = append(results, fmt.Sprintf("%v green %v hanging on the wall,", numToString(bottles), getBottleString(bottles)))
		results = append(results, fmt.Sprintf("%v green %v hanging on the wall,", numToString(bottles), getBottleString(bottles)))
		results = append(results, "And if one green bottle should accidentally fall,")
		bottles -= 1
		results = append(results, fmt.Sprintf("There'll be %v green %v hanging on the wall.", strings.ToLower(numToString(bottles)), getBottleString(bottles)))
		if i != takeDown-1 {
			results = append(results, "")
		}
	}

	return results
}

func numToString(num int) string {
	switch num {
	case 10:
		return "Ten"
	case 9:
		return "Nine"
	case 8:
		return "Eight"
	case 7:
		return "Seven"
	case 6:
		return "Six"
	case 5:
		return "Five"
	case 4:
		return "Four"
	case 3:
		return "Three"
	case 2:
		return "Two"
	case 1:
		return "One"
	case 0:
		return "No"
	default:
		return "Invalid"
	}
}

func getBottleString(num int) string {
	switch num {
	case 1:
		return "bottle"
	default:
		return "bottles"
	}
}
