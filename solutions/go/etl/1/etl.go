package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	result := make(map[string]int)

	for k, v := range in {
		for _, c := range v {
			result[strings.ToLower((c))] = k
		}
	}

	return result
}
