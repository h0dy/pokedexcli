package pokeapi

import (
	"net/http"
	"time"

	"github.com/h0dy/pokedexcli/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache pokecache.Cache
}

func NewClient(timeout time.Duration, cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(cacheInterval),
	}
}