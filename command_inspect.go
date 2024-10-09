package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("select pokemon to inspect")
	}

	pokemonName := args[0]

	pokemonData, ok := cfg.caughtPokemon[pokemonName]

	if !ok {
		return errors.New("you don't know that pokemon, try to catch it first")
	}

	fmt.Printf("Name: %s\n", pokemonData.Name)
	fmt.Printf("Height: %d\n", pokemonData.Height)
	fmt.Printf("Weight: %d\n", pokemonData.Height)
	fmt.Println("Stats:")
	for _, stat := range pokemonData.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")

	for _, pokemonType := range pokemonData.Types {
		fmt.Printf("  -%s\n", pokemonType.Type.Name)
	}
	return nil
}
