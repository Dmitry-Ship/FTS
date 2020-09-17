package main

import (
	"strings"

	"github.com/kljensen/snowball"
)

func lowercaseFilter(tokens []string) []string {
	r := make([]string, 0, len(tokens))
	for _, token := range tokens {
		r = append(r, strings.ToLower(token))
	}
	return r
}

var stopwords = map[string]struct{}{ // I wish Go had built-in sets.
	"a": {}, "and": {}, "be": {}, "have": {}, "i": {},
	"in": {}, "of": {}, "that": {}, "the": {}, "to": {},
}

func stopwordFilter(tokens []string) []string {
	r := make([]string, 0, len(tokens))
	for _, token := range tokens {
		if _, ok := stopwords[token]; !ok {
			r = append(r, token)
		}
	}
	return r
}

func stemmerFilter(tokens []string) []string {
	r := make([]string, len(tokens))
	for i, token := range tokens {
		stemmed, err := snowball.Stem(token, "english", false)
		if err != nil {
			continue
		}
		r[i] = stemmed
	}
	return r
}
