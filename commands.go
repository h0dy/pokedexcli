package main

import (
	"errors"
	"fmt"
	"os"
)


func commandExit(config *configURL, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *configURL, args ...string) error {
	fmt.Println(`
Welcome to the Pokedex!
Usage:
	`)
	fmt.Println()
	for _, cmd := range commands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	
	return nil
}

func commandMap(config *configURL, args ...string) error {
	locations, err := config.pokeapiClient.GetLocations(config.nextURL)
	if err != nil {
		return err
	}

	config.nextURL = locations.Next
	config.previousURL = locations.Previous

	for _, r := range locations.Results {
		fmt.Println(r.Name)
	}
	
	return nil
}

func commandMapb(config *configURL, args ...string) error {
	if config.previousURL == nil {
		return errors.New("you're on the first page, use map command")
	}

	locations, err := config.pokeapiClient.GetLocations(config.previousURL)
	if err != nil {
		return err
	}

	config.nextURL = locations.Next
	config.previousURL = locations.Previous


	for _, r := range locations.Results {
		fmt.Println(r.Name)
	}
	return nil
}

func commandExplore(config *configURL, args ...string) error {
	if len(args) < 1 {
		return errors.New("please make sure to provide a location name")
	}
	area := args[1]
	fmt.Printf("Exploring %v...\n", area)
	locationInfo, err := config.pokeapiClient.ExploreLocation(area)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon: ")
	for _, value := range locationInfo["pokemon_encounters"].([]any) {
		poke := (value.(map[string]any)["pokemon"]).(map[string]any)
		fmt.Printf(" - %v\n", poke["name"])
	}
	return nil
}
