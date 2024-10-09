package main

import (
	"time"

	"github.com/wodorek/pokerepl/internal/pokeapi"
)

func main() {
	client := pokeapi.NewClient(5*time.Second, time.Minute*5)

	cfg := &config{
		pokeapiClient: client,
		caughtPokemon: map[string]pokeapi.Pokemon{},
	}

	startRepl(cfg)
}
