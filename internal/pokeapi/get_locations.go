package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (client *Client) GetLocations(pageURL *string) (PokeLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := client.cache.Get(url); ok {
		var locations PokeLocations
		if err := json.Unmarshal(val, &locations); err != nil {
			return PokeLocations{}, fmt.Errorf("error in getting cache: %w", err)
		}
		return locations, nil
	}
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokeLocations{}, fmt.Errorf("error creating request: %w", err)
	}

	res, err := client.httpClient.Do(req)
	if err != nil {
		return PokeLocations{}, fmt.Errorf("error sending request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return PokeLocations{}, fmt.Errorf("error reading the response: %w", err)
	}

	var locations PokeLocations
	if err := json.Unmarshal(data, &locations); err != nil {
		return PokeLocations{}, fmt.Errorf("error unmarshaling the locations: %w", err)
	}

	client.cache.Add(url, data)
	return locations, nil
}