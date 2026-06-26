package pokeapi

import (
	"net/http"
	"time"

	"github.com/JonasRH355/CLI-Pokedex-in-GO/internal"
)

// Client -
type Client struct {
	cache internal.Cache
	httpClient http.Client
}

// NewClient -
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: *internal.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}