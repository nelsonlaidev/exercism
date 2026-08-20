package anagram

import (
	"slices"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	result := []string{}

	for _, v := range candidates {
		if isAnagram(subject, v) {
			result = append(result, v)
		}
	}

	return result
}

func isAnagram(subject string, candidate string) bool {
	if strings.EqualFold(subject, candidate) {
		return false
	}

	sSlice := strings.Split(subject, "")
	cSlice := strings.SplitSeq(candidate, "")

	for v := range cSlice {
		index := slices.IndexFunc(sSlice, func(c string) bool {
			return strings.EqualFold(c, v)
		})

		if index == -1 {
			return false
		}

		sSlice = removeItem(sSlice, index)

	}

	if len(sSlice) > 0 {
		return false
	}

	return true
}

func removeItem(slice []string, index int) []string {
	result := []string{}

	s1 := slice[:index]
	s2 := slice[index+1:]

	result = append(result, s1...)
	result = append(result, s2...)

	return result
}
