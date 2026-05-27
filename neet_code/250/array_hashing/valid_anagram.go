package main

import (
	"slices"
)

func isAnagram(s string, t string) bool {
	sb := []rune(s)
	tb := []rune(t)

	slices.Sort(sb)
	slices.Sort(tb)

	sortedSb := string(sb)
	sortedTb := string(tb)

	if sortedSb == sortedTb {
		return true
	}

	return false
}

func isAnagramOptimized(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var counts [26]int

	for i := 0; i < len(s); i++ {
		counts[s[i]-'a']++
		counts[t[i]-'a']--
	}

	for _, c := range counts {
		if c != 0 {
			return false
		}
	}

	return true
}
