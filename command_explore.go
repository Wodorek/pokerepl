package main

import "fmt"

func commandExplore(cfg *config, args ...string) error {

	location := args[0]

	explored, err := cfg.pokeapiClient.LocationExplore(location)

	if err != nil {
		return err
	}

	fmt.Printf("Exploring: %s...\n", explored.Name)
	fmt.Println("Found pokemon:")
	for _, enc := range explored.PokemonEncounters {
		fmt.Printf("- %s\n", enc.Pokemon.Name)
	}

	return nil
}
