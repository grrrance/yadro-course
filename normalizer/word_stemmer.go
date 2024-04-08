package main

import (
	"github.com/kljensen/snowball"
)

type Stemmer struct {
}

func NewStemmer() Stemmer {
	return Stemmer{}
}

func (n *Stemmer) Run(language Language, words []string) ([]string, error) {
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
