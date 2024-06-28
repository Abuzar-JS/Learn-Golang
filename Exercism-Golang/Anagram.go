package main

import (
	"sort"
	"strings"
)

func SortString(s string) string {
	runes := []rune(s)

	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	return string(runes)
}

func Detect(subject string, candidates []string) []string {
	subject = strings.ToLower(subject)
	list := []string{}

	for i, candi := range candidates {
		candidates[i] = strings.ToLower(candidates[i])
		if SortString(subject) == SortString(candidates[i]) && subject != candidates[i] {
			list = append(list, candi)

		}
	}
	return list
}
