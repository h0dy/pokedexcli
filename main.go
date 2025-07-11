package main

import (
	"time"

	"github.com/h0dy/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second) // create a new client
	config := &configURL{
		pokeapiClient: pokeClient,
	}
	replCommand(config)
}

