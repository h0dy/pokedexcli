package main

import (
	"time"

	"github.com/h0dy/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second, 5 * time.Minute ) // create a new client
	con := &config{
		pokeapiClient: pokeClient,
		pokemons: map[string]pokeapi.Pokemon{},
	}
	replCommand(con)
}

