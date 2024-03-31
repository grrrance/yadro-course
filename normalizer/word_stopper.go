package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

type wordStopper struct {
	words map[string][]string
}

func NewWordStopper() (WordProcessor, error) {
	jsonFile, err := os.Open("stopwords-iso.json")
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	bytes, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	var words map[string][]string
	err = json.Unmarshal(bytes, &words)
	if err != nil {
		return nil, err
	}
	return &wordStopper{words: words}, nil
}

func (w *wordStopper) Run(language Language, words []string) ([]string, error) {

	stopWords, ok := w.words[language.Code]
	if !ok {
		return nil, fmt.Errorf("undefined language: %s", language.Name)
	}
	var cleanWords []string
	for _, word := range words {
		if !slices.Contains(stopWords, strings.ToLower(word)) && len([]rune(word)) > 2 {
			cleanWords = append(cleanWords, word)
		}
	}
	return cleanWords, nil
}
