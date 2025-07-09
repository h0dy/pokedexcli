package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func replCommand() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		sliceText := CleanInput(text)
		if len(sliceText) == 0 {
			fmt.Println("please provide some text")
			continue
		}
		fmt.Printf("Your command was: %s\n", sliceText[0])
	}
}

func CleanInput(text string) []string {
	textLower := strings.ToLower(text)
	wordsSlice := strings.Fields(textLower)
	return wordsSlice
}
