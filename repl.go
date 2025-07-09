package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
type cliCommand struct {
	name        string
	description string
	callback    func() error
}

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

		cmd, ok := commands()[sliceText[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := cmd.callback()
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
}

func commands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}

func CleanInput(text string) []string {
	textLower := strings.ToLower(text)
	wordsSlice := strings.Fields(textLower)
	return wordsSlice
}
