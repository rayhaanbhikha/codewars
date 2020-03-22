package main

import (
	"strings"
)

/*
"din"      =>  "((("
"recede"   =>  "()()()"
"Success"  =>  ")())())"
"(( @"     =>  "))(("
*/
func main() {
	DuplicateEncode("Success")
	DuplicateEncode("(( @")
	DuplicateEncode("recede")
}

func DuplicateEncode(word string) string {
	charMap := make(map[rune]string)
	lString := strings.ToLower(word)
	for _, char := range lString {
		if v, ok := charMap[char]; !ok {
			charMap[char] = "("
		} else if v != ")" {
			charMap[char] = ")"
		}
	}

	newString := ""
	for _, char := range lString {
		val := charMap[char]
		newString += val
	}

	return newString
}
