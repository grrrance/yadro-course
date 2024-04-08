package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type WordStopper struct {
	words map[string]map[string]struct{}
}

func NewWordStopper() (WordStopper, error) {
	jsonFile, err := os.Open("stopwords-iso.json")
	if err != nil {
		return WordStopper{}, err
	}
	defer jsonFile.Close()
	var words map[string][]string
	err = json.NewDecoder(jsonFile).Decode(&words)
	if err != nil {
		return WordStopper{}, err
	}

	return WordStopper{words: convertToWordsMap(words)}, nil
}

func (w *WordStopper) Run(language Language, words []string) ([]string, error) {
	stopWords, ok := w.words[language.Code]
	if !ok {
		return nil, fmt.Errorf("undefined language: %s", language.Name)
	}
	var cleanWords []string
	for _, word := range words {
		_, ok = stopWords[strings.ToLower(word)]
		if !ok && len([]rune(word)) > 2 {
			cleanWords = append(cleanWords, word)
		}
	}
	return cleanWords, nil
}

func convertToWordsMap(m map[string][]string) map[string]map[string]struct{} {
	stopWords := make(map[string]map[string]struct{}, len(m))
	for lang, words := range m {
		stopWords[lang] = make(map[string]struct{}, len(words))
		for _, w := range words {
			stopWords[lang][w] = struct{}{}
		}
	}
	return stopWords
}
