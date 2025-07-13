package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

func commandExit(con *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(con *config, args ...string) error {
	fmt.Println(`
Welcome to the Pokedex!
Usage:
	`)
	for _, cmd := range commands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandMap(con *config, args ...string) error {
	locations, err := con.pokeapiClient.GetLocations(con.nextURL)
	if err != nil {
		return err
	}

	con.nextURL = locations.Next
	con.previousURL = locations.Previous

	for _, r := range locations.Results {
		fmt.Println(r.Name)
	}
	
	return nil
}

func commandMapb(con *config, args ...string) error {
	if con.previousURL == nil {
		return errors.New("you're on the first page, use map command")
	}

	locations, err := con.pokeapiClient.GetLocations(con.previousURL)
	if err != nil {
		return err
	}

	con.nextURL = locations.Next
	con.previousURL = locations.Previous


	for _, r := range locations.Results {
		fmt.Println(r.Name)
	}
	return nil
}

func commandExplore(con *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("please make sure to provide a location name")
	}
	area := args[0]
	fmt.Printf("Exploring %v...\n", area)
	locationInfo, err := con.pokeapiClient.ExploreLocation(area)
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

func commandCatch(con *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("please provide the name of the pokemon")
	}
	pokeName := args[0]
	fmt.Printf("Throwing a Pokeball at %v...\n", pokeName)

	pokemon, err := con.pokeapiClient.GetPokemon(pokeName)
	if err != nil {
		return err
	}

	pokeXP := int(pokemon["base_experience"].(float64))
	usersXP := 0

	r := rand.Float64()
	switch {
	case r < 0.5:
		usersXP = 60 + rand.Intn(61)
	case r < 0.70:
		usersXP = rand.Intn(60)
	default:
		usersXP = 121 + rand.Intn(480)
	}
	if usersXP >= pokeXP {
		fmt.Printf("%v was caught!\n", pokeName)
		con.pokemons[pokeName] = pokemon
		fmt.Println("You may now inspect it with the inspect command.")
	} else {
		fmt.Printf("%v escaped!\ntry again!\n", pokeName)
	}
	return nil
}

func commandInspect(con *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("please provide the name of the pokemon you want to inspect")
	}

	pokeName := args[0]
	foundPoke, ok := con.pokemons[pokeName]
	if !ok {
		return fmt.Errorf("you have not caught %v\ntry \"catch %v\" to try to catch it 😉", pokeName, pokeName)
	}
	fmt.Printf("Name: %v\nHeight: %v\nWeight: %v\nStats:\n", foundPoke["name"], foundPoke["height"], foundPoke["weight"])
	for _, s := range foundPoke["stats"].([]any) {
		stat := s.(map[string]any) 
		statNum := stat["base_stat"]
		statName := stat["stat"].(map[string]any)["name"]
		fmt.Printf(" - %v: %v\n", statName, statNum)
	}
	return nil
}

func commandPokedex(con *config, args ...string) error {
	if len(con.pokemons) < 1 {
		return fmt.Errorf("your Pokedex is empty ☹️\ntry to catch some pokemon first")
	}
	fmt.Println("Your Pokedex")
	for name := range con.pokemons {
		fmt.Printf(" - %v\n", name)
	}
	return nil
}