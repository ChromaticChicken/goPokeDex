package pokeapi

import (
	"net/http"
	"time"

	"github.com/ChromaticChicken/goPokeDex/internal/pokecache"
)

// Client -
type Client struct {
	httpClient http.Client
	cache      pokecache.PokeCache
}

// NewClient -
func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(timeout),
	}
}
