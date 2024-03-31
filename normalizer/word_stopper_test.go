package main

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestWordStopperEnglish(t *testing.T) {
	prompt := "i'll follow you as long as you are following me"
	data := []string{"follow"}
	s, err := NewWordStopper()
	require.NoError(t, err)
	actual, err := s.Run(Language{
		Code: "en",
		Name: "english",
	}, strings.Split(prompt, " "))
	require.NoError(t, err)
	require.Equal(t, data, actual)

}
