package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (client *Client)GetPokemon(name string) (map[string]any, error) {
	url := baseURL + "/pokemon/" + name

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request in GetPokemon func: %w", err)
	}

	if val, ok := client.cache.Get(url); ok {
		var pokemon map[string]any
		if err := json.Unmarshal(val, &pokemon); err != nil {
			return nil, fmt.Errorf("error in unmarshaling cached data: %w", err)
		}
		return pokemon, nil
	}
	
	res, err := client.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request in GetPokemon func: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response in GetPokemon func: %w", err)
	}
	
	var pokemon map[string]any
	if err := json.Unmarshal(data, &pokemon); err != nil {
		return nil, fmt.Errorf("error unmarshaling the data in GetPokemon func: %w", err)
	}
	client.cache.Add(url, data)
	return pokemon, nil
}