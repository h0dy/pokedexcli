package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/h0dy/pokedexcli/internal/pokeapi"
)
type configURL struct {
	next *string
	previous *string
	pokeapiClient pokeapi.Client
}

func replCommand(config *configURL) {
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
		err := cmd.callback(config)
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*configURL) error
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
		"map": {
			name:        "map",
			description: "Displays the names/next names of 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names/previous names of 20 location areas in the Pokemon world",
			callback:    commandMapb,
		},
	}
}

func CleanInput(text string) []string {
	textLower := strings.ToLower(text)
	wordsSlice := strings.Fields(textLower)
	return wordsSlice
}
