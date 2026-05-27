package main

import "strings"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	i := 1
	for i < len(strs) {
		if strings.HasPrefix(strs[i], prefix) {
			i++
		} else {
			prefix = prefix[:len(prefix)-1]

			if len(prefix) == 0 {
				return ""
			}
		}
	}

	return prefix
}
