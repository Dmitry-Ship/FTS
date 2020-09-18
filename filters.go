package main

import (
	"strings"

	"github.com/kljensen/snowball"
)

func lowercaseFilter(tokens []string) []string {
	result := make([]string, 0, len(tokens))
	for _, token := range tokens {
		result = append(result, strings.ToLower(token))
	}
	return result
}

var stopwords = map[string]struct{}{ // I wish Go had built-in sets.
	"a": {}, "and": {}, "be": {}, "have": {}, "i": {},
	"in": {}, "of": {}, "that": {}, "the": {}, "to": {},
}

func stopwordFilter(tokens []string) []string {
	result := make([]string, 0, len(tokens))
	for _, token := range tokens {
		if _, ok := stopwords[token]; !ok {
			result = append(result, token)
		}
	}
	return result
}

func stemmerFilter(tokens []string) []string {
	result := make([]string, len(tokens))
	for i, token := range tokens {
		stemmed, err := snowball.Stem(token, "english", false)
		if err != nil {
			continue
		}
		result[i] = stemmed
	}
	return result
}
