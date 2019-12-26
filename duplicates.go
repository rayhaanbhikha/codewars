package main

import "unicode"

func duplicateCount(s1 string) int {
	count := make(map[rune]int)
	repeatingRunes := 0

	for _, s := range s1 {
		if v, ok := count[unicode.ToLower(s)]; ok && v == 1 {
			repeatingRunes++
		}
		count[unicode.ToLower(s)]++
	}
	return repeatingRunes
}
