package pokeapi

import (
	"net/http"
	"time"

	"github.com/wodorek/pokerepl/internal/pokecache"
)

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(timeout time.Duration, reapInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(reapInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

// debug method, to be deleted
func (c *Client) PrintCache() {
	c.cache.PrintCache()
}
