package main

import "strings"

func cleanInput(text string) []string {
	textLower := strings.ToLower(text)
	wordsSlice := strings.Fields(textLower)
	return wordsSlice
}
