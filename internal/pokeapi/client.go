package pokeapi

import (
	"github.com/langer-net/pokedexcli/internal/pokecache"
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func CreateNewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.CreateNewCache(cacheInterval),
	}
}
