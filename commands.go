package main

import (
	"errors"
	"fmt"
	"os"
)


func commandExit(config *configURL) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *configURL) error {
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

func commandMap(config *configURL) error {
	locations, err := config.pokeapiClient.GetLocations(config.next)
	if err != nil {
		return err
	}

	config.next = locations.Next
	config.previous = locations.Previous

	for _, r := range locations.Results {
		fmt.Println(r.Name)
	}
	
	return nil
}

func commandMapb(config *configURL) error {
	if config.previous == nil {
		return errors.New("you're on the first page, use map command")
	}

	locations, err := config.pokeapiClient.GetLocations(config.previous)
	if err != nil {
		return err
	}

	config.next = locations.Next
	config.previous = locations.Previous


	for _, r := range locations.Results {
		fmt.Println(r.Name)
	}
	return nil
}

