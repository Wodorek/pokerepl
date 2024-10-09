package main

import (
	"time"

	"github.com/wodorek/pokerepl/internal/pokeapi"
)

func main() {
	client := pokeapi.NewClient((5 * time.Second))

	cfg := &config{
		pokeapiClient: client,
	}

	startRepl(cfg)
}
