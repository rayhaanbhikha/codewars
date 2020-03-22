package main

import (
	"fmt"
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
	charMap := make(map[rune]int)
	for _, char := range strings.ToLower(word) {
		v, ok := charMap[char]
		if !ok {
			charMap[char] = 1
		} else {
			charMap[char] = v + 1
		}
	}

	newString := ""

	for _, char := range word {
		repetition := charMap[char]
		if repetition > 1 {
			newString += ")"
		} else {
			newString += "("
		}
	}
	fmt.Println(newString)
	return ""
}
