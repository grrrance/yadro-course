package main

import (
	"fmt"
	"github.com/pemistahl/lingua-go"
	"strings"
)

type Language struct {
	Code string
	Name string
}

var languages = []lingua.Language{
	lingua.English,
	lingua.French,
	lingua.Spanish,
	lingua.Russian,
	lingua.Swedish,
	lingua.Hungarian,
	lingua.Nynorsk,
}

// Detect detects language from a sentence based on languages supported by snowball
func Detect(prompt string) (*Language, error) {

	detector := lingua.NewLanguageDetectorBuilder().
		FromLanguages(languages...).
		Build()

	if language, exists := detector.DetectLanguageOf(prompt); exists {
		return mapToModel(language), nil
	}

	return nil, fmt.Errorf("language not detected, available languages: %s", getAvailableLanguages())
}

func getAvailableLanguages() string {
	names := make([]string, len(languages))
	for i, l := range languages {
		names[i] = l.String()
	}
	return strings.Join(names, ", ")
}

func mapToModel(l lingua.Language) *Language {
	code := strings.ToLower(l.IsoCode639_1().String())
	name := strings.ToLower(l.String())

	if code == "nn" {
		// this is done because lingua only supports a certain form of Norwegian as nynorsk, and its iso code is different from the standard one
		name = "norwegian"
	}

	return &Language{
		Code: code,
		Name: name,
	}
}

func GetEnglish() *Language {
	return mapToModel(lingua.English)
}
