package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

const (
	s = "this is the query string that needs to be normalized"
)

func main() {
	stopper, err := NewWordStopper()
	if err != nil {
		log.Fatalln(err)
		return
	}
	stemmer := NewStemmer()

	query := flag.String("s", "i'll follow you as long as you are following me", s)
	flag.Parse()

	if query == nil {
		log.Fatalln("query is not specified")
		return
	}

	words := strings.Split(strings.TrimSpace(*query), " ")
	fmt.Printf("current words: %v\n", words)

	lang, err := Detect(*query)

	if err != nil {
		fmt.Printf("%s\n", err.Error())
		fmt.Println("choosing english language...")
		lang = GetEnglish()
	} else {
		fmt.Printf("detected language: %s\n", lang.Name)
	}

	words, err = stopper.Run(*lang, words)

	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Println("removed stop words")
	fmt.Printf("current words: %v\n", words)

	words, err = stemmer.Run(*lang, words)

	if err != nil {
		log.Fatalln(err)
		return
	}

	fmt.Println("stemmed words")

	fmt.Println("resulting words")

	counter := countWords(words)
	for w, counted := range counter {
		fmt.Printf("word: %s, counted: %d\n", w, counted)
	}

}

func countWords(words []string) map[string]int {
	counter := make(map[string]int)

	for _, w := range words {
		counter[w] = 1
	}
	return counter
}
