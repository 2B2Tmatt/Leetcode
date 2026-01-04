package main

import "strings"

func lengthOfLongestSubstring(s string) int {
	longest, cur := 0, ""

	for _, c := range s {
		i := strings.Index(cur, string(c))
		cur += string(c)
		if i != -1 {
			cur = cur[i+1:]
		}

		if longest < len(cur) {
			longest = len(cur)
		}
	}

	return longest
}
