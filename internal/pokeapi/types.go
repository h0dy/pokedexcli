package pokeapi

type PokeLocations struct {
	Count     int 	 `json:"count"`
	Next      *string `json:"next"`
	Previous  *string `json:"previous"`
	Results []struct{ 
		Name  string `json:"name"`
		Url   string `json:"url"`
	} `json:"results"`
}

type Pokemon map[string]any

type Location map[string]any
