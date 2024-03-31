package main

import (
	"github.com/kljensen/snowball"
)

type stemmer struct {
}

func NewStemmer() WordProcessor {
	return &stemmer{}
}

func (n *stemmer) Run(language Language, words []string) ([]string, error) {
	stemmedWords := make([]string, len(words))

	for i, w := range words {
		stemmed, err := snowball.Stem(w, language.Name, true)
		if err != nil {
			return nil, err
		}
		stemmedWords[i] = stemmed
	}

	return stemmedWords, nil
}
