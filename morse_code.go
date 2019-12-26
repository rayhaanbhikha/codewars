package main

import (
	"strings"
)

const wordDelimitter = "   "
const letterDelimitter = " "

func DecodeMorse(morseCode string) string {
	mCode := make(map[string]string)

	words := make([]string, 0)
	for _, morseWord := range strings.Split(strings.Trim(morseCode, " "), wordDelimitter) {
		letters := make([]string, 0)
		for _, morseLetter := range strings.Split(morseWord, letterDelimitter) {
			letters = append(letters, mCode[morseLetter])
		}
		words = append(words, strings.Join(letters, ""))
	}
	return strings.Join(words, " ")
}
