package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)


func (client *Client)ExploreLocation(name string)(map[string]any, error) {
	url := baseURL + "/location-area/" + name
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error in making request in ExploreLocation func: %w", err)
	}
	
	if val, ok := client.cache.Get(name); ok {
		var locationInfo map[string]any
		if err := json.Unmarshal(val, &locationInfo); err != nil {
			return nil, fmt.Errorf("error unmarshaling cached data in ExploreLocation func: %w", err)
		}
		return locationInfo, nil
	}

	res, err := client.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request in ExploreLocation func: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err !=nil {
		return nil, fmt.Errorf("error reading the response in ExploreLocation func: %w", err)
	}

	var locationInfo map[string]any
	if err := json.Unmarshal(data, &locationInfo); err != nil {
		return nil, fmt.Errorf("error unmarshaling data in ExploreLocation func: %w", err)
	}

	client.cache.Add(name, data)
	return locationInfo, nil
}