package main

type WordProcessor interface {
	Run(language Language, words []string) ([]string, error)
}
