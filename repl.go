package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/h0dy/pokedexcli/internal/pokeapi"
)
type config struct {
	nextURL *string
	previousURL *string
	pokeapiClient pokeapi.Client
	pokemons map[string]pokeapi.Pokemon
}

func replCommand(con *config) {
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
		args := []string{}
		if len(sliceText) > 1 {
			args = sliceText[1:]
		}
		err := cmd.callback(con, args...)
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
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
		"explore": {
			name:        "explore <location-name>",
			description: "list all the Pokémon located in an area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon-name>",
			description: "try to catch a pokemon!",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon-name>",
			description: "inspect your caught pokemon!",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "list all the the Pokemon that you have caught",
			callback:    commandPokedex,
		},
	}
}

func CleanInput(text string) []string {
	textLower := strings.ToLower(text)
	wordsSlice := strings.Fields(textLower)
	return wordsSlice
}
