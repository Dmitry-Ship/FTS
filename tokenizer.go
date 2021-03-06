package main

import (
	"strings"
	"unicode"
)

func tokenize(text string) []string {
	return strings.FieldsFunc(text, func(symbol rune) bool {
		// Split on any character that is not a letter or a number.
		return !unicode.IsLetter(symbol) && !unicode.IsNumber(symbol)
	})
}

func analyze(text string) []string {
	tokens := tokenize(text)
	tokens = lowercaseFilter(tokens)
	tokens = stopwordFilter(tokens)
	tokens = stemmerFilter(tokens)
	return tokens
}
