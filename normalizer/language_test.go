package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDetect(t *testing.T) {
	data := []struct {
		prompt string
		l      Language
	}{{prompt: "Every mistake is a step towards success.", l: Language{
		Code: "en",
		Name: "english",
	}}, {prompt: "Chaque erreur est un pas vers le succès.", l: Language{
		Code: "fr",
		Name: "french",
	}}, {prompt: "Cada error es un paso hacia el éxito.", l: Language{
		Code: "es",
		Name: "spanish",
	}}, {prompt: "Каждая ошибка — это шаг к успеху.", l: Language{
		Code: "ru",
		Name: "russian",
	}}, {prompt: "Varje misstag är ett steg mot framgång.", l: Language{
		Code: "sv",
		Name: "swedish",
	}}, {prompt: "Minden hiba egy lépés a siker felé.", l: Language{
		Code: "hu",
		Name: "hungarian",
	}}, {prompt: "Hver feil er et skritt mot suksess.", l: Language{
		Code: "nn",
		Name: "norwegian",
	}}}

	for _, v := range data {
		l, err := Detect(v.prompt)
		require.NoError(t, err)
		require.Equal(t, v.l.Name, l.Name)
		require.Equal(t, v.l.Code, l.Code)
	}
}
