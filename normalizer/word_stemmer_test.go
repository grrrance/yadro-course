package main

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestWordStemmerEnglish(t *testing.T) {
	prompt := "follower brings bunch of questions"
	data := []string{"follow", "bring", "bunch", "of", "question"}
	s := NewStemmer()
	actual, err := s.Run(Language{
		Code: "en",
		Name: "english",
	}, strings.Split(prompt, " "))
	require.NoError(t, err)
	require.Equal(t, data, actual)

}
